package messagebus

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type MyClient struct {
	client *pubsub.Client
}

func CreateClient(ProjectId string) *MyClient {

	client, err := pubsub.NewClient(context.Background(), ProjectId, option.WithCredentialsFile("secret/gcpserviceaccount.json"))
	if err != nil {
		panic(err)
	}
	return &MyClient{client: client}
}

// pass Client Struct as a receiver to use the Client:

func (c *MyClient) CreatePubSub(PubSubTopic string) {
	topic, err := c.client.CreateTopic(context.Background(), PubSubTopic)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pub Sub Topic Created with name: %v  \n", topic)
}

func (c *MyClient) CloseClient() error {
	return c.client.Close()
}
