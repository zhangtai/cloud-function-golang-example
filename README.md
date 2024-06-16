# Golang Example for Cloud Function

A example program to run on GCP Cloud Function, triggers by PubSub, retrieve Home Assistant entity state.

## References

- [Create and deploy an HTTP Cloud Function with Go](https://cloud.google.com/functions/docs/writing/specifying-dependencies-go)
- [Call local functions](https://cloud.google.com/functions/docs/running/calling)

## Quick start

```bash
# Start
export HA_HOST="https://ha.example.com"
export HA_TOKEN="token"
export ENTITY_ID="entity_id"
export FUNCTION_TARGET=HelloPubSub
go run cmd/main.go

# Test
curl localhost:8080 \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2020-01-02T12:34:56.789Z" \
  -H "ce-type: google.cloud.pubsub.topic.v1.messagePublished" \
  -H "ce-source: //pubsub.googleapis.com/projects/MY-PROJECT/topics/MY-TOPIC" \
  -d '{
        "message": {
          "data": "d29ybGQ=",
          "attributes": {
             "attr1":"attr1-value"
          }
        },
        "subscription": "projects/MY-PROJECT/subscriptions/MY-SUB"
      }'

# Deploy
# TBC
```

