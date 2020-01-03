package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	Movie{
		ID:   1,
		Name: "Avengers",
	},
	Movie{
		ID:   2,
		Name: "Gun Gun Febrianza",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid Payload",
		}, nil
	}

	movies = append(movies, movie)
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:500,
			Body: err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode:200,
		Headers: map[string]string{
			"Content-Type": "Application/json",
		},
		Body: string(response),
	}, nil
}

func main()  {
	lambda.Start(insert)
}
