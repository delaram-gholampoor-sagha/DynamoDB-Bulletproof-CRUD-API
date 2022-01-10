package controllers

import (
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/handlers/product"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/adapter"
	"github.com/golangci/golangci-lint/pkg/report"
	"github.com/google/uuid"
)

type Controller struct {
	repository adapter.Interface

}

type Interface interface {
	ListOne(id uuid.UUID) (entity product.Product , err error)
	ListAll() (entities []product.Product , err error)
	Create(entity *product.Product ) (uuid.UUID , error)
	Update(id uuid.UUID , entity  *product.Product )
	Remove(id uuid.UUID) error 

}

func NewController() Interface {
	

	return &Controller{repository: repository}
}


func (c *Controller) ListOne(id uuid.UUID) (entity product.Product , err error){
     entity.ID = id 
	 response , err := c.repository.FindOne(entity.GetFilterId() , entity.TableName())
	 if err != nil {
		 return entity , err 
	 }
	 return product.ParseDynamoAttributeToStruct(response.Item) 
}





func (c *Controller) ListAll() (entities []product.Product , err error){
	
}



func (c *Controller) Create(entity *product.Product ) (uuid.UUID , error) {
	
}



func (c *Controller) Update(id uuid.UUID , entity  *product.Product ) {
	
}

func (c *Controller) Remove(id uuid.UUID) error{
	
}