package main

import "github.com/iamgauravpande/gcp-infra/pkg/gcp/messagebus"

func main() {
	ProjectId := "bitlost"                          // ProjectID of GCP Project
	PubSubTopic := "lostinopensrc"                  // PubSub Topic Name
	gcpclient := messagebus.CreateClient(ProjectId) // Create gcp go Client
	defer gcpclient.CloseClient()                   // Close gcp go Client
	gcpclient.CreatePubSub(PubSubTopic)             // Create PubSubTopic

}
