package ws_flogo

import (
	"net"
	//"os"
	//"os/exec"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
)

func Pour(port string) {
	for {
		conn, _ := net.Dial("tcp", net.JoinHostPort("", port))
		if conn != nil {
			conn.Close()
			break
		}
	}
}

func TestMain(m *testing.M) {
	/*
	command := exec.Command("docker", "start", "mqtt")
	err := command.Run()
	if err != nil {
		command := exec.Command("docker", "run", "-p", "1883:1883", "-p", "9001:9001", "--name", "mqtt", "-d", "eclipse-mosquitto")
		err := command.Run()
		if err != nil {
			panic(err)
		}
	}
	Pour("1883")
	os.Exit(m.Run())
	*/
}

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.SetClientID("TestAbc123")
	client := mqtt.NewClient(options)
	token := client.Connect()
	token.Wait()
	assert.Nil(t, token.Error())
	fini := make(chan bool, 1)
	token = client.Subscribe("/x/+/y/#", 0, func(client mqtt.Client, message mqtt.Message) {
		topic, payload := message.Topic(), string(message.Payload())
		assert.Equal(t, `{"message": "hello world"}`, payload)
		assert.Equal(t, "/x/test/y/j/k", topic)
		fini <- true
	})
	token.Wait()
	assert.Nil(t, token.Error())

	settings := Settings{
		Server: "localhost:5001",
	}
	init := test.NewActivityInitContext(settings, nil)
	act, err := New(init)
	assert.Nil(t, err)
	context := test.NewActivityContext(activityMd)
	context.SetInput("message", `{"message": "hello world"}`)
	done, err := act.Eval(context)
	assert.True(t, done)
	assert.Nil(t, err)

	select {
	case <-fini:
	case <-time.Tick(time.Second):
		t.Fatal("didn't get message in time")
	}
}
