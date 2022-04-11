package publish

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"nats/internal/config"
	"nats/internal/pkg"
	"strconv"

	"github.com/hako/durafmt"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func Run(address, subject string, qty uint, delta time.Duration, data []byte) error {

	//Loading .env to get local variables
	config.LoadDotEnv()

	nats_addr := nats.DefaultURL
	if len(address) != 0 {
		nats_addr = address
	}
	sub := config.GetSujectName()
	if len(subject) != 0 {
		sub = config.GetTopicName() + "." + subject
		// Force to use all the string passed by CLI
		//sub = subject
	}
	q := config.GetNumberMsg()
	if qty != 0 {
		q = qty
	}

	d := delta
	if delta == 0 {
		if config.GetTime2Wait() != "0" {
			duration, err := durafmt.ParseString(config.GetTime2Wait())
			if err != nil {
				fmt.Println(err)
			}
			d = duration.Duration()
		}
	}

	// Connect to NATS
	nc, _ := nats.Connect(nats_addr)
	// Creates JetStreamContext
	js, err := nc.JetStream()
	if err != nil {
		logrus.Error(err)
		return err
	}
	// Creates stream
	err = createStream(js, config.GetTopicName(), sub)
	if err != nil {
		logrus.Error(err)
		return err
	}
	// Create orders by publishing messages
	err = createOrder(js, int(q), sub, d)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// createOrder publishes stream of events
func createOrder(js nats.JetStreamContext, msgs int, sub string, t time.Duration) error {
	var order pkg.Order
	for i := 1; i <= msgs; i++ {
		order = pkg.Order{
			OrderID:    i,
			CustomerID: "Cust-" + strconv.Itoa(i),
			Status:     "publish",
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(sub, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("%v:Order with OrderID:%d has been published\n", sub, i)
		time.Sleep(t)
	}
	return nil
}

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext, topicName string, subject string) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(config.GetTopicName())
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", topicName, subject)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     topicName,
			Subjects: []string{topicName + ".>"},
			// Subjects: []string{subject, subject + ".resp"},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
