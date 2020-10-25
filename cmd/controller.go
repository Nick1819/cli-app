package cmd

import (
	"log"
	"time"
)

func loopController(result string, loop string) {
	if loop == "yes" {
		playMP3("la.mp3")
		log.Println("done 1")
		time.Sleep(time.Second * 3)
		playMP3(result)
		log.Println("done 2")
	}
	if loop == "no" {
		playMP3(result)
	}
}
