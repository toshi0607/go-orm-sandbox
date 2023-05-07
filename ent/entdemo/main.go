package main

import (
	"context"
	"fmt"
	"log"

	"github.com/toshi0607/go-orm-sandbox/ent/entdemo/ent"
	"github.com/toshi0607/go-orm-sandbox/ent/entdemo/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	user, err := CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}
	fmt.Printf("user name: %s\n", user.Name)

	gotUser, err := QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed query user: %v", err)
	}
	fmt.Printf("user name: %s\n", gotUser.Name)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
