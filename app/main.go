package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

type Resp struct {
	Body string
}

func handleRequest(event events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	resp := Resp{
		Body: fmt.Sprintf("%s request to %v is successful!", event.RequestContext.HTTP.Method, event.RequestContext.RouteKey),
	}
	msg, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(msg),
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
