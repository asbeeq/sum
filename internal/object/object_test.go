package object

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	// test data
	testObjects := []Object{
		{A: 5, B: 3},
		{A: 7, B: 2},
		{A: -10, B: -4},
		{A: -1, B: 0},
	}

	tempFile, err := ioutil.TempFile("", "testfile.json")
	if err != nil {
		t.Fatalf("create temp file err: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// json encode
	jsonData, err := json.Marshal(testObjects)
	if err != nil {
		t.Fatalf("marshal json data err: %v", err)
	}
	if _, err := tempFile.Write(jsonData); err != nil {
		t.Fatalf("write json data to file err: %v", err)
	}

	// test Read function
	objects, err := Read(tempFile.Name())
	if err != nil {
		t.Fatalf("read json file err: %v", err)
	}

	// assert objects size
	if len(objects) != len(testObjects) {
		t.Fatalf("expected %d objects, got %d", len(testObjects), len(objects))
	}

	// assert values
	for i := range testObjects {
		if objects[i].A != testObjects[i].A || objects[i].B != testObjects[i].B {
			t.Errorf("mismatch at index %d: expected %+v, got %+v", i, testObjects[i], objects[i])
		}
	}
}
