package mnist

import (
	log "Landsat-Extractor/logger"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

// GetTrainingSet provides images and associated labels for training
func GetTrainingSet() ([]*mat.Dense, []byte) {

	images := loadMNISTImages("mnist/data/train-images-idx3-ubyte")
	labels := loadMNISTLabels("mnist/data/train-labels-idx1-ubyte")

	return images, labels
}

// GetTestSet provides images and associated labels for evaluation
func GetTestSet() ([]*mat.Dense, []byte) {

	images := loadMNISTImages("mnist/data/t10k-images-idx3-ubyte")
	labels := loadMNISTLabels("mnist/data/t10k-labels-idx1-ubyte")

	return images, labels

}

func loadMNISTLabels(path string) []byte {
	log.Info("Start loading MNIST labels " + path)

	file, error := os.Open(path)
	if error != nil {
		panic("Cannot open " + path)
	}
	defer file.Close()

	// Reading the data
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	bytes := make([]byte, size)
	file.Read(bytes)

	// Analysing the 'Magic Number'
	// 1st and 2nd byte are always 0x00 (ignored in the code)
	// 3rd byte (bytes[2]) defines the data type of the payload
	// 4th byte (bytes[3]) defines the number of dimensions of the payload

	// data type has to be unsigned byte (value 0x08)
	if bytes[2] != byte(0x08) {
		panic("Label format not correct")
	}

	// dimension has to be 1 (number of labels)
	if bytes[3] != 1 {
		panic("Number of dimensions not correct")
	}

	// 1st dimension is number of labels
	numLabels := uint32(bytes[7])
	numLabels |= uint32(bytes[6]) << 8
	numLabels |= uint32(bytes[5]) << 16
	numLabels |= uint32(bytes[4]) << 24

	labels := make([]byte, numLabels)

	// Read every label and collect them in an array
	for i := uint32(0); i < numLabels; i++ {
		var label byte = bytes[8+i]
		labels[i] = label
	}

	log.Info(strconv.Itoa(len(labels)) + " labels loaded")

	return labels
}

func loadMNISTImages(path string) []*mat.Dense {

	log.Info("Load MNIST images " + path)

	file, error := os.Open(path)
	if error != nil {
		panic("Cannot open " + path)
	}
	defer file.Close()

	// Reading the data
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	bytes := make([]byte, size)
	file.Read(bytes)

	// Analysing the 'Magic Number'
	// 1st and 2nd byte are always 0x00 (ignored in the code)
	// 3rd byte (bytes[2]) defines the data type of the payload
	// 4th byte (bytes[3]) defines the number of dimensions of the payload

	// data type has to be unsigned byte (value 0x08)
	if bytes[2] != byte(0x08) {
		panic("Image format not correct")
	}

	// dimension has to be 3 (number of images, number of rows, number of columns)
	if bytes[3] != 3 {
		panic("Number of dimensions not correct")
	}

	// 1st dimension is number of images
	numImages := uint32(bytes[7])
	numImages |= uint32(bytes[6]) << 8
	numImages |= uint32(bytes[5]) << 16
	numImages |= uint32(bytes[4]) << 24

	// 2nd dimension is number of rows
	numRows := uint32(bytes[11])
	numRows |= uint32(bytes[10]) << 8
	numRows |= uint32(bytes[9]) << 16
	numRows |= uint32(bytes[8]) << 24

	// 3rd dimension ist number of columns
	numCols := uint32(bytes[15])
	numCols |= uint32(bytes[14]) << 8
	numCols |= uint32(bytes[13]) << 16
	numCols |= uint32(bytes[12]) << 24

	pixelPerImage := numRows * numCols

	images := make([]*mat.Dense, numImages)

	// Read every image, build a pixel matrix
	// and collect them in an array
	for i := uint32(0); i < numImages; i++ {

		pixels := make([]float64, pixelPerImage)

		for j := uint32(0); j < pixelPerImage; j++ {
			offset := i*pixelPerImage + j + 16
			pixels[j] = float64(bytes[offset])
		}

		image := mat.NewDense(int(numRows), int(numCols), pixels)
		images[i] = image
	}

	log.Info(strconv.Itoa(len(images)) + " images loaded")

	return images
}
