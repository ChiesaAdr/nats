package config

import (
	"os"
	"strconv"
)

func GetScalable() bool {

	scalable, exists := os.LookupEnv("SCALABLE")
	if !exists {
		return false
	}

	return scalable == "true"
}

func GetSujectName() string {

	url, exists := os.LookupEnv("SUBJECTS_NAME")
	if !exists {
		return "incloud.default"
	}
	return url
}

func GetTopicName() string {

	url, exists := os.LookupEnv("TOPIC_NAME")
	if !exists {
		return "incloud"
	}
	return url
}

func GetSujects() string {

	url, exists := os.LookupEnv("SUBJECTS")
	if !exists {
		return "incloud.>"
	}
	return url
}

func GetNumberMsg() uint {

	url, exists := os.LookupEnv("NUMBER_MSG")
	if !exists {
		return 1
	}
	v, e := strconv.Atoi(url)
	if e != nil {
		//Default Limit
		return 1
	}
	return uint(v)
}

func GetAddressServer() string {

	addr, exists := os.LookupEnv("ADDRESS_SERVER")
	if !exists {
		return "nats://127.0.0.1:4222"
	}
	return addr
}

func GetTime2Wait() string {

	url, exists := os.LookupEnv("TIME2WAIT")
	if !exists {
		return "2s"
	}
	return url
}
