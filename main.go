package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//model.InitDb()
	//routes.InitRouter()
	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")
}
