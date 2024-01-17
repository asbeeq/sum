package main

import (
	"flag"

	"github.com/asbeeq/sum/internal/object"

	"github.com/sirupsen/logrus"
)

func main() {
	// create and set logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// example: -g=2 -f=objects.json -p=concurrent
	numGoroutines := flag.Int("g", 1, "number of goroutines")
	fileName := flag.String("f", "objects.json", "objects file")
	processType := flag.String("p", object.Concurrent, "process type: { concurrent | sequencial }")
	flag.Parse()

	// validation
	if *numGoroutines <= 0 {
		logger.Warn("number of goroutines set to default 1")
		*numGoroutines = 1
	}
	if *processType != object.Concurrent && *processType != object.Sequential {
		logger.Warnf("provided processType is invalid: `%s`, set to default `concurrent`", *processType)
		*processType = object.Concurrent
	}

	// read json file and map data to []Object
	objects, err := object.Read(*fileName)
	if err != nil {
		// logs error and exits
		logger.Fatal(err)
	}

	sum := 0
	switch *processType {
	case object.Concurrent:
		sum = object.CalculateConcurrentSum(objects, *numGoroutines)
	case object.Sequential:
		sum = object.CalculateSequentialSum(objects)
	default:
		logger.Fatal("sum calculation is invalid, processType: ", *processType)
	}

	logger.Info("sum of all objects in json file is: ", sum)
}
