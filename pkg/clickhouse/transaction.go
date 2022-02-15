package clickhouse

import (
	"time"
)

// insert transaction
func (ch *ClickHouse) InsertData(data interface{}) error {
	var (
		tx, _   = ch.Begin()
		stmt, _ = tx.Prepare("INSERT INTO logs (description , day, time) VALUES (?, ?, ?)")
	)
	now := time.Now()
	if _, err := stmt.Exec(data, now, now); err != nil {
		stmt.Close()
		return err
	}
	defer stmt.Close()
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
