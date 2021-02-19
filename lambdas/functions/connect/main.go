package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	servicewebsocket "github.com/estebanbacl/show_tr/service/websocket"

	"github.com/estebanbacl/show_tr/functions"
	"github.com/estebanbacl/show_tr/log"
	websocketmodel "github.com/estebanbacl/show_tr/model/websocket"
	"github.com/estebanbacl/show_tr/util"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda createDynamoTrigger websocket connection")
	log.Logger.Debug().Msgf("Start lambda createDynamoTrigger claims %v", event.RequestContext.Authorizer)
	websocket := &websocketmodel.WebSocket{
		ID: event.RequestContext.ConnectionID,
	}
	log.Logger.Debug().Msgf("id %v", websocket)
	err := functions.Instances["websocketService"].(servicewebsocket.WebsocketService).Create(cxt, websocket)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the connect %v", websocket)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
