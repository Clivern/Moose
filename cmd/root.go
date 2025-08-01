// Copyright 2025 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var baseURL string
var sseMode bool
var logOutput string
var logLevel string
var logFormat string

var rootCmd = &cobra.Command{
	Use: "moose",
	Short: `🐺 A Dynamic MCP Server.

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://github.com/clivern/moose/issues>`,
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
