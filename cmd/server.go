// Copyright 2025 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/clivern/moose/core/tool"
	"github.com/clivern/moose/core/util"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start moose server",
	Run: func(_ *cobra.Command, _ []string) {

		fs := util.NewFileSystem()

		if logOutput != "stdout" {
			dir, _ := filepath.Split(logOutput)

			if !fs.DirExists(dir) {
				if err := fs.EnsureDir(dir, 775); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}
		}

		if logOutput == "stdout" {
			log.SetOutput(os.Stdout)
		} else {
			f, err := os.OpenFile(logOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

			if err != nil {
				panic(fmt.Sprintf(
					"Error while opening log file [%s]: %s",
					logOutput,
					err.Error(),
				))
			}
			log.SetOutput(f)
		}

		lvl := strings.ToLower(logLevel)
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if logFormat == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		s := server.NewMCPServer(
			"Moose",
			"0.1.1",
			server.WithToolCapabilities(false),
		)

		// Add tool
		helloTool := mcp.NewTool("hello_world",
			mcp.WithDescription("Say hello to someone"),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("Name of the person to greet"),
			),
		)

		// Add tool handler
		s.AddTool(helloTool, tool.HelloHandler)

		// Run server in appropriate mode
		if sseMode {
			// Create and start SSE server
			sseServer := server.NewSSEServer(s, server.WithBaseURL(baseURL))
			log.Printf("Starting SSE server on %s", baseURL)
			// Extract port from baseURL for the Start method
			port := ":8080" // default port
			if strings.Contains(baseURL, ":") {
				parts := strings.Split(baseURL, ":")
				if len(parts) >= 3 {
					port = ":" + parts[2]
				}
			}
			if err := sseServer.Start(port); err != nil {
				log.Fatalf("Server error: %v", err)
			}
		} else {
			// Run as stdio server
			if err := server.ServeStdio(s); err != nil {
				log.Fatalf("Server error: %v", err)
			}
		}
	},
}

func init() {
	serverCmd.Flags().BoolVarP(
		&sseMode,
		"sse",
		"s",
		false,
		"Run in SSE mode instead of stdio mode",
	)

	serverCmd.Flags().StringVarP(
		&baseURL,
		"url",
		"u",
		"http://localhost:8080",
		"Base URL for the server",
	)

	serverCmd.Flags().StringVarP(
		&logOutput,
		"log-output",
		"o",
		"stdout",
		"Log output file (stdout or file path)",
	)

	serverCmd.Flags().StringVarP(
		&logLevel,
		"log-level",
		"l",
		"info",
		"Log level (debug, info, warn, error, fatal, panic)",
	)

	serverCmd.Flags().StringVarP(
		&logFormat,
		"log-format",
		"f",
		"text",
		"Log format (text or json)",
	)

	rootCmd.AddCommand(serverCmd)
}
