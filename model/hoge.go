package model

import (
	"database/sql"
	"strconv"
)

type Hoge struct {
	ID   int64  `json:"id"`
	Col1 string `json:"col1"`
	Col2 string `json:"col2"`
}

func HogeAll(db *sql.DB) ([]*Hoge, error) {

	rows, err := db.Query(`select id from hoge`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []*Hoge
	for rows.Next() {
		m := &Hoge{}
		if err := rows.Scan(&m.ID); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

func HogeByID(db *sql.DB, id string) (*Hoge, error) {
	m := &Hoge{}

	// 1-1. ユーザー名を表示しよう
	if err := db.QueryRow(`select id from hoge where id = ?`, id).Scan(&m.ID); err != nil {
		return nil, err
	}

	return m, nil
}

func (h *Hoge) Insert(db *sql.DB) (*Hoge, error) {
	var c, _ = strconv.Atoi(h.Col1)
	res, err := db.Exec(`insert into hoge(id,col1,col2) values(null,?,?)`, c, h.Col2)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &Hoge{
		ID:   id,
		Col1: h.Col1,
		Col2: h.Col2,
	}, nil

}
