package main

import (
	"aws-lambda/dtypes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *dynamodb.Client
var tableName = "MarioGallery"
var indexName = "VCategoryURLName-CTypeStatus-Index"

func init() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("error fetching default configuration: %v", err)
	}

	client = dynamodb.NewFromConfig(cfg)
}

func categoryProducts(ctx context.Context, event events.APIGatewayV2HTTPRequest) (
	[]dtypes.ProductValue, error) {

	catUrl, ok := event.QueryStringParameters["vCategoryUrlName"]
	if !ok {
		return nil, fmt.Errorf("required parameter category url is required")
	}

	condCat := expression.Key("VCategoryURLName").Equal(expression.Value(catUrl))
	condStatus := expression.Key("CTypeStatus").Equal(expression.Value("PA"))
	cond := expression.KeyAnd(condCat, condStatus)
	expr, err := expression.NewBuilder().WithKeyCondition(cond).Build()
	if err != nil {
		return nil, fmt.Errorf("error generating expression: %s", err)
	}

	qryIn := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String(indexName),
		ConsistentRead:            aws.Bool(false),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	qryOut, err := client.Query(context.TODO(), qryIn)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %s", err)
	}

	var products []dtypes.ProductValue
	if err := attributevalue.UnmarshalListOfMaps(qryOut.Items, &products); err != nil {
		return nil, fmt.Errorf("error marshaling query output: %s", err)
	}

	return products, nil
}

func main() {
	lambda.Start(categoryProducts)
}
