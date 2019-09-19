package ws_flogo

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Server       string                 `md:"server,required"` // The broker URL
}

type Input struct {
	Destination     string       `md:"destination"`     // The message to send
	Message     string       `md:"message"`     // The message to send
}

type Output struct {
	Data interface{} `md:"data"` // The data recieved
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"destination":     i.Destination,
		"message":     i.Message,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.Destination, _ = coerce.ToString(values["destination"])
	i.Message, _ = coerce.ToString(values["message"])
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	o.Data = values["data"]
	return nil
}
