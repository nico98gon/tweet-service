package routers

import (
	"fmt"
	"strconv"
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
	page := request.QueryStringParameters["page"]
	if len(page) < 1 {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Error al parsear el page: " + err.Error()
		return r
	}

	tweets, ok := db.GetTweets(ID, int64(pageInt))
	if !ok {
		r.Message = "Error al buscar los tweets: " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "OK"
	r.Data = tweets
	r.Meta = utils.Pagination{
		TotalPages:  0,
		CurrentPage: pageInt,
		NextPage:    0,
		HasNext:     false,
	}

	return r
}