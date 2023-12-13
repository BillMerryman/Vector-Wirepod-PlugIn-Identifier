package main

import (
	sdk_wrapper "github.com/fforchino/vector-go-sdk/pkg/sdk-wrapper"
	"context"
	"fmt"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"strings"
)

var Utterances = []string{"what is this thing"}
var Name = "Identifier Plugin"

func Action(transcribedText string, botSerial string) string {

	sdk_wrapper.InitSDK(botSerial)
	ctx := context.Background()
	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = sdk_wrapper.Robot.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
			case <-start:
				sdk_wrapper.SaveHiResCameraPicture("object.jpg")
				
				conn, err := grpc.Dial(
					"api.clarifai.com:443",
					grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
				)
				if err != nil {
					panic(err)
				}
				client := api.NewV2Client(conn)

				ctx_Clarifai := metadata.AppendToOutgoingContext(
					context.Background(),
					"Authorization", "Key {your Clarifai api key here}",
				)

				var GeneralModelId = "aaa03c23b3724a16a56b629203edc62c"

				fileBytes, err := ioutil.ReadFile("object.jpg")
				
				response, err := client.PostModelOutputs(
					ctx_Clarifai,
					&api.PostModelOutputsRequest{
					ModelId: GeneralModelId,
					Inputs: []*api.Input{
						{
							Data: &api.Data{
								Image: &api.Image{
									Base64: fileBytes,
								},
							},
						},
					},
					},
				)
				
				if err != nil {
					panic(err)
				}
				
				if response.Status.Code != status.StatusCode_SUCCESS {
					panic(fmt.Sprintf("Failed response: %s", response))
				}
				
				//Get the object name with response.Outputs[0].Data.Concepts[0].Name
				//Get the likelihood with response.Outputs[0].Data.Concepts[0].Value
				sdk_wrapper.SayText("That, my friend, could be any of the following things.")
				for index := 0; index < 5; index++ {
					var Article string
					Name := response.Outputs[0].Data.Concepts[index].Name
					if strings.Contains("AEIOUaeiou", Name[0:1]){
						Article = "an"
					}else{
						Article = "a"
					}
					sdk_wrapper.SayText(Article + " " + Name)
				}
				stop <- true
				return "intent_imperative_praise"
		}
	}

}
