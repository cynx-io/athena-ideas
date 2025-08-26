package exampleservice

import (
	"github.com/cynx-io/athena-ideas/internal/repository/database"
)

type Service struct {
	ExampleRepo *database.ExampleRepo
}

func New(exampleRepo *database.ExampleRepo) *Service {
	return &Service{
		ExampleRepo: exampleRepo,
	}
}
