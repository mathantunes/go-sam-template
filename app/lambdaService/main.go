package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mathantunes/go-sam-template/domain"
	"github.com/mathantunes/go-sam-template/domain/entities"
	"github.com/mathantunes/go-sam-template/domain/service"
	"github.com/mathantunes/go-sam-template/domain/valueobjects"
	"github.com/mathantunes/go-sam-template/infra/repo"
)

type h struct {
	svc domain.ServiceInterface
}

func (h *h) handle(msg valueobjects.Request) (entities.Data, error) {
	return h.svc.ProcessRequest(msg)
}

func main() {
	h := &h{
		svc: service.NewService(repo.NewRepoClient()),
	}
	lambda.Start(h.handle)
}
