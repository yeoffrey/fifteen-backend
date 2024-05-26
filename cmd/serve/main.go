package main

import (
	"log"
	"os"

	// Blank-import the function package so the init() runs
	_ "github.com/yeoffrey/fifteen"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := getPort()

	log.Printf("Listenning on port %v...", port)

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return envPort
	}

	return "8080"
}
