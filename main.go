package main

import (
	log "Landsat-Extractor/logger"
	"Landsat-Extractor/mnist"
)

func main() {

	defer log.Destroy()

	mnist.GetTrainingSet()
	mnist.GetTestSet()
}
