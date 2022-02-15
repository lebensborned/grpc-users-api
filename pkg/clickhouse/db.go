package clickhouse

import (
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type ClickHouse struct {
	*sqlx.DB
}

func Connect(host, dbname, password, user string) (*ClickHouse, error) {
	connStr := fmt.Sprintf("tcp://%s:9000?debug=true&database=%s&password=%s&user=%s", host, dbname, password, user)
	db, err := sqlx.Open("clickhouse", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	ch := &ClickHouse{db}
	ch.db_init()
	return ch, nil
}

func (ch *ClickHouse) db_init() error {
	// log.Println("clickhouse_init")
	_, err := ch.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			description   String,
			day   Date,
			time  DateTime
		) ENGINE = MergeTree()
	`)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
