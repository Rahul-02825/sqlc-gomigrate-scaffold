package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Rahul-02825/user-service/internal/db"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()
	dbURL, ok := os.LookupEnv("DATABASE_URL");
	if !ok || dbURL == ""{
		fmt.Printf("error in gettting dburl from env")
	}
	fmt.Println(dbURL)
	// os.Getenv("DATABASE_URL")
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(ctx,dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	users,err := queries.Getusers(ctx)

	fmt.Println(users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

}