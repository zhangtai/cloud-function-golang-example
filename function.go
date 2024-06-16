package roomtemperatue

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("HelloPubSub", helloPubSub)
}

type EntityState struct {
	State       string `json:"state"`
	LastChanged string `json:"last_changed"`
}

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func helloPubSub(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	haHost, exists := os.LookupEnv("HA_HOST")
	if !exists {
		haHost = "https://example.com"
	}
	haToken, exists := os.LookupEnv("HA_TOKEN")
	if !exists {
		haToken = "Invalid Token"
	}
	entityId, exists := os.LookupEnv("ENTITY_ID")
	if !exists {
		entityId = "Invalid Token"
	}
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	haEndpoint := haHost + "/api/states/" + entityId

	client := &http.Client{}
	req, err := http.NewRequest("GET", haEndpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+haToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var entityState EntityState
	if err := json.Unmarshal(body, &entityState); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	fmt.Printf("Entity State: %+v\n", entityState)
	return nil
}
