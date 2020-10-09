# go-sam-template

## This repo can be used as a guideline for developing golang SAM projects with a DDD approach to serverless

### App

On app/* implement the lambda handler and initializa all layers used.

#### Domain

##### domain/entities

Define a file for each entity needed, and make sure any transformation to be done on the entity is a method provided by itself.

##### domain/valueobjects

define a file for each valueobject needed.

##### domain/***

Other than entities or valueobjects, implement a service layer that interacts with the domain entities and repositories from the infrastructure layer.

##### domain/contracts.go

Define all interfaces to be implemented by each service.

#### Infra

##### infra/contracts.go

Define all interfaces to be implemented by each service

##### infra/***

Implement infrasctructure layer based on the interfaces defines on contracts
