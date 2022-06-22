package main

import (
	"context"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"io"
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://dev.cip.sysco.no/oauthservice/operation", nil)
	//request, err := http.NewRequest("GET", "http://localhost:9001/oauthservice/operation", nil)
	if err != nil {
		log.Fatalf("http.NewRequest: %v", err)
	}
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")

	audience := "63785118804-tkq80964jqq266r4mufm6kl9u40an90c.apps.googleusercontent.com"
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, audience, option.WithCredentialsFile("/Users/prakhar/workspace/sysco/oauth-example/cip-dev.json"))
	if err != nil {
		log.Fatalf("idtoken.NewClient: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("client.Do: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)
	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("io.Copy: %v", err)
	}
	log.Println("mmm", string(b))
	log.Println("mmm", response.Status)
	for k, h := range response.Header {
		log.Println(k, "-->", h)
	}

}
