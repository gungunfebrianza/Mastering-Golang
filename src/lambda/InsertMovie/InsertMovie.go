package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"

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
		Name: "SAW",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       req.Body,
		}, nil
	}

	movies = append(movies, movie)
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "Application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(insert)
}
