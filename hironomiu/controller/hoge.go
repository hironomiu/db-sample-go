package controller

import (
	"database/sql"
	//"errors"
	"net/http"

	"github.com/VG-Tech-Dojo/vg-1day-2017-06-17/hironomiu/httputil"
	"github.com/VG-Tech-Dojo/vg-1day-2017-06-17/hironomiu/model"
	"github.com/gin-gonic/gin"
)

// Message is controller for requests to messages
type Hoge struct {
	DB     *sql.DB
	Stream chan *model.Message
}

// All は全てのメッセージを取得してJSONで返します
func (h *Hoge) All(c *gin.Context) {
	msgs, err := model.HogeAll(h.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(msgs) == 0 {
		c.JSON(http.StatusOK, make([]*model.Message, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": msgs,
		"error":  nil,
	})
}

// GetByID はパラメーターで受け取ったidのメッセージを取得してJSONで返します
func (m *Hoge) GetByID(c *gin.Context) {
	msg, err := model.HogeByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusNotFound, resp)
		return
	case err != nil:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": msg,
		"error":  nil,
	})
}

