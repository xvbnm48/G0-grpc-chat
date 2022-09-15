package main

import (
	"context"

	chat "github.com/xvbnm48/Go-grpc-chat/schema"
)

const tokenHeader = "X-chat-token"

type client struct {
	chat.ChatClient
	Name, token string
}

func (c *client) login() (string, error) {
	ctx := context.Background()

	res, err := c.ChatClient.Login(ctx, &chat.LoginRequest{
		Username: c.Name,
	})
	if err != nil {
		return "", nil
	}

	return res.Token, nil
}
