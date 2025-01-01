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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var client *dynamodb.Client

func init() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("error fetching default configuration: %v", err)
	}

	client = dynamodb.NewFromConfig(cfg)
}

func subCategories(ctx context.Context, event events.APIGatewayV2HTTPRequest) ([]dtypes.CategorySummary, error) {

	iPCatID := event.QueryStringParameters["iPCatID"]

	pk, _ := attributevalue.Marshal("Category")
	sk, _ := attributevalue.Marshal(fmt.Sprintf("CAT#%s", iPCatID))
	key := map[string]types.AttributeValue{
		"PK": pk,
		"SK": sk,
	}

	qryIn := &dynamodb.GetItemInput{
		TableName:       aws.String("MarioGallery"),
		Key:             key,
		AttributesToGet: []string{"LChildren"},
		ConsistentRead:  aws.Bool(false),
	}

	qryOut, err := client.GetItem(ctx, qryIn)
	if err != nil {
		return nil, err
	}

	var children []dtypes.CategorySummary
	attributevalue.Unmarshal(qryOut.Item["LChildren"], &children)

	return children, nil
}

func main() {
	lambda.Start(subCategories)
}
