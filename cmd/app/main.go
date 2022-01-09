package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/config"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/adapter"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/repository/instance"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/internal/routes"
	"github.com/Delaram-Gholampoor-Sagha/DynamoDB-Bulletproof-CRUD-API/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	config := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("waiting for the service to start .... ", nil)

	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migration : .... ", err)
		}
	}

	logger.PANIC("", checkTables(connection))

	port := fmt.Sprint(":%v", config.Port)
	router := routes.NewRouter().SetRoutes(repository)
	logger.INFO("service is running on port ", port)
	server := http.ListenAndServe(port, router)

	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndappendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndappendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}

}

func checkTables(connection dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables Not Found : ", nil)
		}

		for _, tableName := range response.TableNames {
			logger.INFO("Table Found :", *tableName)
		}
	}

	return err

}
