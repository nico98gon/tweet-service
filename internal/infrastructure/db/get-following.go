package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"tweet-service/internal/domain/users"
)

func GetFollowingFromUserService(userID string) ([]string, bool) {
	userServiceURL := os.Getenv("USER_SERV_URL")
	if userServiceURL == "" {
		fmt.Println("Error: USER_SERV_URL no está configurada")
		return nil, false
	}

	strReq := fmt.Sprintf("%s/get-following?id=%s", userServiceURL, userID)
	fmt.Println("Solicitud HTTP:", strReq)
	req, err := http.NewRequest("GET", strReq, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud HTTP:", err)
		return nil, false
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al hacer la solicitud HTTP:", err)
		return nil, false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta HTTP:", err)
		return nil, false
	}

	var followingResponse users.FollowingResponse
	if err := json.Unmarshal(body, &followingResponse); err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return nil, false
	}

	if followingResponse.Status != 200 {
		fmt.Println("Error en la respuesta del microservicio de usuarios:", followingResponse.Message)
		return nil, false
	}

	var followingIDs []string
	for _, user := range followingResponse.Data {
		followingIDs = append(followingIDs, user.ID)
	}

	fmt.Println("Usuarios seguidos obtenidos:", len(followingIDs))

	return followingIDs, true
}