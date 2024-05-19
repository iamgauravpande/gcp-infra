package main

import (
	"github.com/iamgauravpande/gcp-infra/pkg/gcp/gcpclient"
	"github.com/iamgauravpande/gcp-infra/pkg/gcp/topic"
)

func main() {
	ProjectId := "bitlost"                         // ProjectID of GCP Project
	PubSubTopic := "lostinopensrc"                 // PubSub Topic Name
	gcpclient := gcpclient.CreateClient(ProjectId) // Create gcp go Client
	defer gcpclient.CloseClient()                  // Close gcp go Client
	topic.CreatePubSub(gcpclient, PubSubTopic)     // Create PubSubTopic

}
