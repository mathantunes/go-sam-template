package service

import (
	"github.com/mathantunes/go-sam-template/domain/entities"
	"github.com/mathantunes/go-sam-template/domain/valueobjects"
	"github.com/mathantunes/go-sam-template/infra"
)

type Service struct {
	db infra.DB
}

func NewService(db infra.DB) *Service {
	return &Service{
		db: db,
	}
}

func (sv *Service) ProcessRequest(req valueobjects.Request) (entities.Data, error) {
	return sv.db.Get(req)
}
