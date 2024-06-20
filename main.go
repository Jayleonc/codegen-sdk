package main

import (
	"github.com/Jayleonc/codegen-sdk/codegen"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

func main() {
	// 创建 etcd 客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}

	serviceName := "example-repo"
	baseURL := "http://example.com"
	outputPath := "gen"

	err = codegen.GenerateClientCode(serviceName, baseURL, outputPath, client)
	if err != nil {
		log.Fatalf("Failed to generate client code: %v", err)
	}

	log.Printf("Client code generated and written to %s", outputPath)
}
