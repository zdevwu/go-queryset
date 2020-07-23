// +build tools

package tools

// This files is here to force go module to download the following command tools

import (
	//_ "github.com/golang/mock/mockgen"
	//_ "github.com/google/wire/cmd/wire"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
