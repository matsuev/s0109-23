package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	api.Post("/connect", processConnect)

	if err := api.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}

// ConnectRequest ...
type ConnectRequest struct {
	Client    string             `json:"client"`
	Transport string             `json:"transport"`
	Protocol  string             `json:"protocol"`
	Encoding  string             `json:"encoding"`
	Name      string             `json:"name,omitempty"`
	Version   string             `json:"version,omitempty"`
	Data      ConnectRequestData `json:"data,omitempty"`
	B64Data   string             `json:"b64data,omitempty"`
	Channels  []string           `json:"channels,omitempty"`
}

// ConnectRequestData ...
type ConnectRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectResponseResult ...
type ConnectResponseResult struct {
	User string `json:"user"`
}

// ConnectResponseError ...
type ConnectResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// ConnectResponseDisconnect ...
type ConnectResponseDisconnect struct {
	Code   int    `json:"code"`
	Reason string `json:"reason,omitempty"`
}

// ConnectResponse ...
type ConnectResponse struct {
	Result     *ConnectResponseResult     `json:"result,omitempty"`
	Error      *ConnectResponseError      `json:"error,omitempty"`
	Disconnect *ConnectResponseDisconnect `json:"disconnect,omitempty"`
}

func processConnect(ctx *fiber.Ctx) error {
	var request ConnectRequest

	if err := json.Unmarshal(ctx.Body(), &request); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if request.Data.Username == "alex" && request.Data.Password == "qwerty" {
		return ctx.JSON(ConnectResponse{
			Result: &ConnectResponseResult{
				User: request.Data.Username,
			},
		})
	}

	return ctx.SendStatus(http.StatusUnauthorized)
}
