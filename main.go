package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is the function invoked by AWS when the Lambda is triggered.
func HandleRequest(ctx context.Context, request Request) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(request)

	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}

// Request is the HTTP request body stucture with necessary data bindings and validations.
type Request struct {
	Zpl       string    `json:"zpl" binding:"required"`
	LabelSize LabelSize `json:"labelSize" binding:"required"`
	Dpi       string    `json:"dpi" binding:"required,oneof='152' '203' '300' '600'"`
}

// LabelSize is the nested struct for label size
type LabelSize struct {
	Width  float32 `json:"width" binding:"required,max=4"`
	Height float32 `json:"height" binding:"required,max=10"`
}

// SuccessResponse is the struct for successful response.
// It contains all the labels that are generated from the input request.
type SuccessResponse struct {
	Labels []string `json:"labels" binding:"required,min=1"`
}

// ErrorResponse is the struct for error response.
type ErrorResponse struct {
	Message     string `json:"message" binding:"required"`
	Details     string `json:"details" binding:"omitempty"`
	Suggestions string `json:"suggestions" binding:"omitempty"`
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
