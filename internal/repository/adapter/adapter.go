package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tablename string) (respoonse *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tablename string) (response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tablename string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

func NewAdapter(con *dynamodb.DynamoDB) interface{} {
	return &Database{
		connection: con,
		logMode:    false,
	}
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll(condition expression.Expression, tablename string) (respoonse *dynamodb.ScanInput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tablename),
	}
	return input, nil

}

// what do you want to match ?
func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}

	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate(entity interface{}, tablename string) (response *dynamodb.PutItemOutput, err error) {

	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tablename),
	}

	return db.connection.PutItem(input)

}

// which collection you want to delete something from and you have to find the exact one item you want to delete
func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {

	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}

	return db.connection.DeleteItem(input)
}
