package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"tweet-service/internal/domain"
	"tweet-service/internal/domain/tweets"
	"tweet-service/internal/infrastructure/db"
)

func CreateTweet(ctx context.Context, claim *domain.Claim) domain.RespAPI {
	var t tweets.Tweet
	var r domain.RespAPI
	r.Status = 400

	IDUser := claim.ID.Hex()

	body, ok := ctx.Value(domain.Key("body")).(string)
	if !ok {
		r.Message = "Error: No se pudo obtener el body de la solicitud"
		fmt.Println(r.Message)
		return r
	}

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Error al parsear el body: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	register := tweets.CreateTweet{
		UserID:  IDUser,
		Content: t.Content,
		Date:    time.Now(),
	}

	err = tweets.CreateTweetValidations(t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		r.Message = "Error al insertar el tweet: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "Error al insertar el tweet"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Tweet creado correctamente"
	fmt.Println(r.Message)

	return r
}