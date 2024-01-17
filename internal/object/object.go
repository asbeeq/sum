package object

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

const (
	// process type
	Concurrent = "concurrent"
	Sequential = "sequential"
)

type Object struct {
	A int `json: "a"`
	B int `json: "b"`
}

// read json file and map data to []Object
func Read(fileName string) ([]Object, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "read all from file")
	}

	var objects []Object
	err = json.Unmarshal(data, &objects)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal json to object struct")
	}

	return objects, nil
}
