package topic

import (
	"context"
	"log"

	"github.com/iamgauravpande/gcp-infra/pkg/gcp/gcpclient"
)

func CreatePubSub(gcpclient *gcpclient.MyClient, PubSubTopic string) {
	topic, err := gcpclient.Client.CreateTopic(context.Background(), PubSubTopic)
	if err != nil {
		log.Fatalf("Failed to Create Topic Error is: %v \n", err)
	}
	log.Printf("Pub Sub Topic Created with name: %v  \n", topic)
}
