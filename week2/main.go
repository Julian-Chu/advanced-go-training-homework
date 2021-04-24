package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Julian-Chu/advanced-go-training-homework/user"
)

func main() {
	db, _ := sql.Open("postgres", "pg_conn_str")
	userStore:= user.NewStore(db)
	ctx:= context.Background()
	_, err := userStore.QueryUserByEmail(ctx, "testuser@testmail.com")
	if err != nil{
		if errors.Is(err, user.ErrNotFound){
			fmt.Println("bad request")
			return
		}
		// handle err
	}

}
