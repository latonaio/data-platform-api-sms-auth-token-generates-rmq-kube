package config

import "os"

type AWS struct {
	AWSPinpointSenderID string
}

func newAWS() *AWS {
	return &AWS{
		AWSPinpointSenderID: os.Getenv("AWS_PINPOINT_SENDER_ID"),
	}
}
