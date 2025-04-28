//go:build tools
// +build tools

//go:generate go run cmd/emptyFile/main.go output.yaml
//go:generate go tool yq eval-all --inplace "select(fileIndex == 0) * select(fileIndex == 1) * select(fileIndex == 2) * select(fileIndex == 3)" output.yaml reduct-api-docs/bucket.yaml reduct-api-docs/entry-read.yaml reduct-api-docs/entry-write.yaml
//go:generate go tool swagger generate client -A reduct -f output.yaml -t ../reduct/integration --template-dir tmpl --allow-template-override

package main

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/mikefarah/yq/v4"
)
