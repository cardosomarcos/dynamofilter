package examples

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cardosomarcos/dynamofilter"
	"github.com/guregu/dynamo"
)

func exampleEquals() {
	var data []map[string]interface{}

	s, _ := session.NewSession()
	dynamoConnection := dynamo.New(s, &aws.Config{
		Region:     aws.String("sa-east-1"),
		MaxRetries: aws.Int(3),
	})

	filter := dynamofilter.NewFilter()
	filter.Equals("name", "Stor Gendibal")
	f, args := filter.Builder()

	err := dynamoConnection.Table("user").Scan().Filter(f, args...).All(&data)
	log.Println("result: ", data)
	log.Println("err: ", err)
}
