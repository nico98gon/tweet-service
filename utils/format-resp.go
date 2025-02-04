package utils

import (
	"encoding/json"
	"tweet-service/internal/domain"

	"github.com/aws/aws-lambda-go/events"
)

func FormatResponse(status int, message string, data interface{}, meta interface{}) *events.APIGatewayProxyResponse {
	body, _ := json.Marshal(domain.RespAPI{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}