/*
 * PetSitter API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	swagger "go-server/go"
	"log"
	"net/http"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
