package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func handler() error {

	clientTwilio := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%s", os.Getenv("TO_NUMBER")))
	params.SetFrom(fmt.Sprintf("whatsapp:%s", os.Getenv("FROM_NUMBER")))
	params.SetBody("Message sent from Go API AWS")

	resp, err := clientTwilio.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}

	log.Println("Message has been sent")

	return nil
}

func main() {
	lambda.Start(handler)
}
