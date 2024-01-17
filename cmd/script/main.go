package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

// TODO make arg
const objectsNum = 1_000_000 // 2

type Object struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	objects := make([]Object, objectsNum)

	for i := 0; i < objectsNum; i++ {
		objects[i] = Object{
			A: rand.Intn(21) - 10, // rand num in [-10, 10]
			B: rand.Intn(21) - 10,
		}
	}

	jsonData, err := json.Marshal(objects)
	if err != nil {
		log.Fatal("json marshal err:", err)
	}

	file, err := os.Create("objects.json")
	if err != nil {
		log.Fatal("file create err:", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("write file err:", err)
	}

	log.Printf("success, %d objects has been stored in objects.json", objectsNum)
}
