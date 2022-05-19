package main

import (
    "context"
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
    fmt.Println("Hello! This will show up in your CloudWatch Log Group.")
    return "hello - this will only show up as the response, not in any Log Groups", nil
}

func main() {
    lambda.Start(handler)
}
