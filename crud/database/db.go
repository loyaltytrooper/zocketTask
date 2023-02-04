package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

var CRUD_DB *pgx.Conn

func InitDB() {
	var err error
	if CRUD_DB, err = pgx.Connect(context.Background(), "postgresql://postgres:U0JxW4N6PzPkeCaff0RW@containers-us-west-45.railway.app:7316/railway"); err != nil {
		log.Fatalf("Error fetching DB, %s", err.Error())
	}
	log.Println("Connected to DB now")
}

func CloseDB() {
	defer func(crudDB *pgx.Conn, ctx context.Context) {
		err := crudDB.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing DB, %s", err.Error())
		}
	}(CRUD_DB, context.Background())
	log.Println("Closed DB now")
}
