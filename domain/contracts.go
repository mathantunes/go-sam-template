package domain

import (
	"github.com/mathantunes/go-sam-template/domain/entities"
	"github.com/mathantunes/go-sam-template/domain/valueobjects"
)

type ServiceInterface interface {
	ProcessRequest(valueobjects.Request) (entities.Data, error)
}
