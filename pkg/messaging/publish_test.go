package messaging

import (
	"encoding/json"
//	"os"
	"log"
	"testing"
)

type DeploymentContent1 struct {
	Content1 string `json:"content1"`
	Content2 int	`json:"content2"`
}

type DeploymentContent2 struct {
	Content3 string `json:"content1"`
	Content4 int	`json:"content2"`
	Nested DeploymentContentNest
}

type DeploymentContentNest struct {
	Content5 string `json:"content5"`
}

func TestMarshallMessage(t *testing.T) {
	log.Printf("Inside TestPublish")

	message := DeploymentMessage {
		Header: DeploymentMessageHeader {
			Topic: "TestTopic",
		},
		Body: "TestContent",
	}

	jsonContent, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("TestPublish() failed to marshal message: %v", err)
	}
	jsonString := string(jsonContent)
	log.Printf("TestPublish result - %s", jsonString)

	message2 := DeploymentMessage {
		Header: DeploymentMessageHeader {
			Topic: "TestTopic",
		},
		Body: DeploymentContent1{
				Content1: "TestContent1",
				Content2: 2,
		},
	}

	jsonContent, err = json.Marshal(message2)
	if err != nil {
		t.Fatalf("TestPublish() failed to marshal message: %v", err)
	}
	jsonString = string(jsonContent)
	log.Printf("TestPublish result2 - %s", jsonString)

	message3 := DeploymentMessage {
		Header: DeploymentMessageHeader {
			Topic: "TestTopic",
		},
		Body: DeploymentContent2{
				Content3: "TestContent3",
				Content4: 3,
				Nested: DeploymentContentNest{
					Content5: "TestContent5",
				},
		},
	}

	jsonContent, err = json.Marshal(message3)
	if err != nil {
		t.Fatalf("TestPublish() failed to marshal message: %v", err)
	}
	jsonString = string(jsonContent)
	log.Printf("TestPublish result2 - %s", jsonString)
}