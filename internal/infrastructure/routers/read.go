package routers

import (
	"fmt"
	"tweet-service/internal/domain"
	"tweet-service/internal/infrastructure/db"
	"tweet-service/utils"

	"github.com/aws/aws-lambda-go/events"
)


func ReadTweets(request events.APIGatewayProxyRequest) domain.RespAPI {
	var r domain.RespAPI
	r.Status = 400

	ID, ok := request.QueryStringParameters["id"]
	if !ok || len(ID) < 1 {
		r.Message = "ID es requerido en request"
		return r
	}
	fmt.Println("ID:", ID)

	cursor := request.QueryStringParameters["cursor"]

	tweets, hasMore, err := db.GetTweets(ID, cursor)
	if err != nil {
		r.Message = "Error al buscar los tweets: " + err.Error()
		return r
	}

	var nextCursor string
	if len(tweets) > 0 {
		nextCursor = tweets[len(tweets)-1].ID.Hex()
	}

	r.Status = 200
	r.Message = "OK"
	r.Data = tweets
	r.Meta = utils.Pagination(hasMore, nextCursor)

	return r
}