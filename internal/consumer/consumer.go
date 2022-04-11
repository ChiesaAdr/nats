package consumer

import (
	"encoding/json"
	"log"
	"nats/internal/config"
	"nats/internal/pkg"
	"runtime"

	"github.com/nats-io/nats.go"
)

func Run(address, subject string) {

	//Loading .env to get local variables
	config.LoadDotEnv()

	nats_addr := nats.DefaultURL
	if len(address) != 0 {
		nats_addr = address
	}
	su := config.GetSujectName()
	if len(subject) != 0 {
		su = config.GetTopicName() + "." + subject
		// Force to use all the string passed by CLI
		//su = subject
	}

	// Connect to NATS
	nc, _ := nats.Connect(nats_addr)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Create Pull based consumer with maximum 128 inflight.
	// PullMaxWaiting defines the max inflight pull requests.
	// sub, _ := js.PullSubscribe(su, "order-review", nats.PullMaxWaiting(128))
	// ctx, cancel := context.WithTimeout(context.Background(), d)

	// Create durable consumer monitor
	js.Subscribe(su, func(msg *nats.Msg) {
		msg.Ack()
		var order pkg.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
		reviewOrder(js, order, su)

	}, nats.ManualAck())

	runtime.Goexit()

}

// reviewOrder reviews the order and publishes ORDERS.approved event
func reviewOrder(js nats.JetStreamContext, order pkg.Order, sub string) {
	// Changing the Order status
	order.Status = "responded"
	orderJSON, _ := json.Marshal(order)
	_, err := js.Publish(sub+".resp", orderJSON)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Order with OrderID:%d has been %s\n", order.OrderID, order.Status)
}
