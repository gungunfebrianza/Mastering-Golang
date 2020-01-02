package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "Serunya belajar serverless dengan golang!", nil
}

func main() {
	lambda.Start(handler)
}
