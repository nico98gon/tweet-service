package routers

import (
	"fmt"
	"tweet-service/internal/domain"
	"tweet-service/internal/infrastructure/db"

	"github.com/aws/aws-lambda-go/events"
)

func FollowingTweets(request events.APIGatewayProxyRequest, claim *domain.Claim) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	userID := claim.ID.Hex()
	if userID == "" || userID == "000000000000000000000000" {
		r.Message = "Error: userID no válido en el token"
		fmt.Println(r.Message)
		return r
	}

	cursor := request.QueryStringParameters["cursor"]

	tweets, nextCursor, ok := db.GetFollowingTweets(userID, cursor)
	if !ok {
		r.Message = "Error al obtener los tweets de los seguidos"
		return r
	}

	r.Status = 200
	r.Message = "OK"
	r.Data = tweets
	r.Meta = map[string]interface{}{"nextCursor": nextCursor}

	return r
}