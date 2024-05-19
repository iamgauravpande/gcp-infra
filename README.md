## GCP Infra Using GCP Go Client

This repo is a POC for Creating infra on GCP using google cloud Client library: [google-cloud-go](https://github.com/googleapis/google-cloud-go)

**NOTE: Presently the New Client is created using `option.WithCredentialsFile` a GCP service account so make sure you have the path locally along with json key for your GCP Service account , See below:**

```golang
client, err := pubsub.NewClient(context.Background(), ProjectId, option.WithCredentialsFile("secret/gcpserviceaccount.json"))
```