package repo

import (
	"github.com/mathantunes/go-sam-template/domain/entities"
	"github.com/mathantunes/go-sam-template/domain/valueobjects"
)

type RepoClient struct{}

func NewRepoClient() *RepoClient {
	return &RepoClient{}
}

func (c *RepoClient) Get(msg valueobjects.Request) (entities.Data, error) {
	return entities.Data{
		Id: "123",
	}, nil
}
