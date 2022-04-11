package main

import (
	"github.com/sirupsen/logrus"

	"nats/cmd"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	cmd.Execute()
}
