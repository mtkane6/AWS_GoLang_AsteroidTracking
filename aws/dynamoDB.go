package aws

import (
	"fmt"
	"log"
	"os"

	d "Asteroids/datastructs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoInit creates the table and returns a pointer
func DynamoInit() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		log.Fatal("Err:", err)
	}
	svc := dynamodb.New(sess)

	// create dynamodb table
	tablename := dynamodb.DescribeTableInput{
		TableName: aws.String("NearEarthObjects"),
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: tablename.TableName,
	}

	_, err = svc.CreateTable(input)
	if err != nil {
		if err.(awserr.Error).Code() == dynamodb.ErrCodeResourceInUseException {
			fmt.Println("Err, table exist, but skipping error")
		} else {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err)
			os.Exit(1)
		}

	}
	svc.WaitUntilTableExists(&tablename)
	return svc
}

// InputItem inputs items into the dynamodb
func InputItem(asteroidCollection []d.Asteroid, dyna *dynamodb.DynamoDB) {
	for i := range asteroidCollection {
		av, err := dynamodbattribute.MarshalMap(asteroidCollection[i])
		if err != nil {
			fmt.Println("Got error marshalling new asteroid item:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		itemInput := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("NearEarthObjects"),
		}

		_, err = dyna.PutItem(itemInput)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
