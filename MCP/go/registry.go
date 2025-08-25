package main

import (
	"github.com/spotify-web-api-node-api/mcp-server/config"
	"github.com/spotify-web-api-node-api/mcp-server/models"
	tools_callback "github.com/spotify-web-api-node-api/mcp-server/tools/callback"
	tools_login "github.com/spotify-web-api-node-api/mcp-server/tools/login"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_callback.CreateGet_callbackTool(cfg),
		tools_login.CreateGet_loginTool(cfg),
	}
}
