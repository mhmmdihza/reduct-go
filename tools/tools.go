//go:build tools
// +build tools

//go:generate go run cmd/emptyFile/main.go output.yaml
//go:generate go tool yq eval-all --inplace "select(fileIndex == 0) * select(fileIndex == 1) * select(fileIndex == 2) * select(fileIndex == 3)" output.yaml reduct-api-docs/bucket.yaml reduct-api-docs/entry-read.yaml reduct-api-docs/entry-write.yaml
//go:generate go tool swagger generate client -f output.yaml --target ../reduct/integration

package main

import (
	_ "github.com/mikefarah/yq/v4"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
