package monitor

import (
	"encoding/json"
	"log"
	"nats/internal/config"
	"nats/internal/pkg"
	"runtime"

	"github.com/nats-io/nats.go"
)

func Run(address, subject string) error {
	//Loading .env to get local variables
	config.LoadDotEnv()

	nats_addr := nats.DefaultURL
	if len(address) != 0 {
		nats_addr = address
	}
	su := config.GetSujects()
	if len(subject) != 0 {
		su = subject
		log.Println(su)
	}

	// Connect to NATS
	nc, _ := nats.Connect(nats_addr)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Create durable consumer monitor
	// js.Subscribe("incloud.>", func(msg *nats.Msg) {
	js.Subscribe(su, func(msg *nats.Msg) {
		msg.Ack()
		var order pkg.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("monitor service subscribes from subject:%s\n", msg.Subject)
		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
	}, nats.ManualAck())

	runtime.Goexit()

	return nil
}
