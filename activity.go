package ws_flogo

import (
	"github.com/gorilla/websocket"
	"net/url"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func init() {
	_ = activity.Register(&Activity{}, New)
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{
		settings: settings,
	}
	return act, nil
}

type Activity struct {
	settings *Settings
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)

	if err != nil {
		return true, err
	}

	wsDestination := input.Destination
	wsMessage := input.Message
	wsHost := a.settings.Server

	wsURL := url.URL{Scheme: "ws", Host: wsHost}
	ctx.Logger().Infof("connecting to %s", wsURL.String())

	wsConn, _, err2 := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err2 != nil {
		ctx.Logger().Infof("Error while dialing to wsHost: ", err)
	}

	textMessage := `{"body": {"_dest":"` + wsDestination + `", "text":"` + wsMessage + `"}, "seq": 1}`

	ctx.Logger().Infof("Preparing to send message: [%s]", textMessage)

	err = wsConn.WriteMessage(websocket.TextMessage, []byte(textMessage))
	if err != nil {
		ctx.Logger().Infof("Error while sending message to wsHost: [%s]", err)
		return
	}
	wsConn.Close()

	return true, nil
}
