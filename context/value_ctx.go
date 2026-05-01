package main

import (
	"context"
	"fmt"
)

type keyType string

const reqIDKey keyType = "requestID"

func main() {
	ctx := context.WithValue(context.Background(), reqIDKey, "req-12345")
	process(ctx)
}

func process(ctx context.Context) {
	if id, ok := ctx.Value(reqIDKey).(string); ok {
		fmt.Println("request ID:", id)
	} else {
		fmt.Println("no request ID")
	}
}
