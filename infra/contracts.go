package infra

import (
	"github.com/mathantunes/go-sam-template/domain/entities"
	"github.com/mathantunes/go-sam-template/domain/valueobjects"
)

// DB repository contract in order to
type DB interface {
	Get(msg valueobjects.Request) (entities.Data, error)
}
