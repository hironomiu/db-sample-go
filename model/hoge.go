package model

import (
	"database/sql"
)

type Hoge struct {
	ID       int64  `json:"id"`
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
