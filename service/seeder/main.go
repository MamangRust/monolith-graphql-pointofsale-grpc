package main

import (
	"context"
	"log"

	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/database"
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/seeder"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/dotenv"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/hash"
	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
)

func main() {
	if err := dotenv.Viper(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	l, err := logger.NewLogger("seeder", nil)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	pool, err := database.NewClient(l)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()

	ctx := context.Background()
	queries := db.New(pool)

	s := seeder.NewSeeder(seeder.Deps{
		Db:     queries,
		Ctx:    ctx,
		Logger: l,
		Hash:   hash.NewHashingPassword(),
	})

	if err := s.Run(); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	l.Info("Seeding completed successfully.")
}
