// Copyright 2025 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package tool

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	log "github.com/sirupsen/logrus"
)

// HelloHandler handles the hello_world tool requests
func HelloHandler(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to run hello tool")

		return mcp.NewToolResultError(err.Error()), nil
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Run hello tool")

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
