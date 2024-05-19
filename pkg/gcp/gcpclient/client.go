package gcpclient

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type MyClient struct {
	Client *pubsub.Client
}

func CreateClient(ProjectId string) *MyClient {

	client, err := pubsub.NewClient(context.Background(), ProjectId, option.WithCredentialsFile("secret/gcpserviceaccount.json"))
	if err != nil {
		log.Fatalf("Connection failed: %v \n", err)
	}
	return &MyClient{Client: client}
}

func (c *MyClient) CloseClient() error {
	return c.Client.Close()
}
