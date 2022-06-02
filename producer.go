package main

import (
	"log"
	"time"
)

func Produce(ticker *time.Ticker) {
	log.Println(">>>> Produce")

	defer log.Println("<<<< Produce")
	defer ticker.Stop()

	for {
		select {
		case <-Shutdown:
			log.Println("Done!")
			return
		case <-ticker.C:
			switch {
			case ItemCount <= MinThreshold:
				ProductionRate++
				ItemCount += ProductionRate

			case ItemCount >= MaxThreshold:
				ProductionRate--
				continue
			}
			log.Println("Item count:", ItemCount)
		}

	}

}
