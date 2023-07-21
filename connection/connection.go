package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


var Conn *pgx.Conn

func DatabaseConnect() {

	var err error
	databaseURL := "postgres://postgres:200799@localhost:5432/b48-s1"

	Conn, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully connected to database.")
}