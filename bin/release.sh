#!/bin/bash

# Fetch latest version
export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clivern/moose/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/')

# Update go checksum database (sum.golang.org) immediately after release
curl --silent https://sum.golang.org/lookup/github.com/clivern/moose@{$LATEST_VERSION}
