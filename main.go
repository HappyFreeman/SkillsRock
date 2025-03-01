package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/HappyFreeman/SkillsRock/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	//_ "github.com/lib/pq"
)

func main() {
	/* cd sql\schema goose postgres postgres://username:password@localhost:5432/database_name up */
	// go run main.go
	// go build ; ./SkillsRock
	// go get github.com/joho/godotenv
	// go get github.com/jackc/pgx/v5
	// go get github.com/gofiber/fiber/v2
	// go install github.com/pressly/goose/v3/cmd/goose@latest
	// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	// go mod vendor
	// go mod tidy
	// sqlc generate

	godotenv.Load(".env") // загружаем данные из .env

	// postgres://username:password@localhost:5432/database_name
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
	}

	db, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удается подключиться к базе данных: %v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	defer db.Close(context.Background())

	app := fiber.New()
	v1 := fiber.New()
	app.Mount("/api/v1", v1)

	TaskHandlers(v1, dbQueries)

	log.Fatal(app.Listen(":3000"))
}