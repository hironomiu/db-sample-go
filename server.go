package main

import (
	//"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/VG-Tech-Dojo/vg-1day-2017-06-17/hironomiu/controller"
	"github.com/VG-Tech-Dojo/vg-1day-2017-06-17/hironomiu/model"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
        _ "github.com/go-sql-driver/mysql"
)

// Server はAPIサーバーが実装された構造体です
type Server struct {
	db          *sql.DB
	Engine      *gin.Engine
}

// NewServer は新しいServerの構造体のポインタを返します
func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

// Init はサーバーを初期化します
func (s *Server) Init(dbconf, env string) error {
        db, err := sql.Open("mysql", "root:redash@tcp(172.18.0.2:3306)/redash")
	if err != nil {
		return err
	}
        //defer db.Close()
	s.db = db

	// routing
	s.Engine.LoadHTMLGlob("./templates/*")

	s.Engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	s.Engine.Static("/assets", "./assets")

	// api
	api := s.Engine.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	msgStream := make(chan *model.Hoge)

	hctr := &controller.Hoge{DB: db, Stream: msgStream}
	api.GET("/hoge", hctr.All)
	api.GET("/hoge/:id", hctr.GetByID)

	return nil
}

// Close はDBとの接続を閉じてサーバーを終了します
func (s *Server) Close() error {
	return s.db.Close()
}

// Run はサーバーを起動します
func (s *Server) Run(port string) {
	s.Engine.Run(fmt.Sprintf(":%s", port))
}

func main() {
	var (
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
		port   = flag.String("port", "8080", "listening port.")
	)
	flag.Parse()

	s := NewServer()
	if err := s.Init(*dbconf, *env); err != nil {
		log.Fatalf("fail to init server: %s", err)
	}
	defer s.Close()

	s.Run(*port)
}
