package routers

import (
	"fmt"
	"tweet-service/internal/domain"
	"tweet-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	userID := claim.ID.Hex()
	if len(userID) == 0 || userID == "000000000000000000000000" {
		r.Message = "Error: userID no válido en el token"
		fmt.Println(r.Message)
		return r
	}

	ID, ok := request.QueryStringParameters["id"]
	if !ok || len(ID) < 1 {
		r.Message = "ID es requerido en request"
		return r
	}

	err := db.DeleteTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al eliminar el tweet: " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Se elimino el tweet correctamente"

	return r
}