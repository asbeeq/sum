package object

import "sync"

// iteration through objects slice sequentially
func CalculateSequentialSum(objects []Object) int {
	sum := 0
	for _, v := range objects {
		sum += v.A + v.B
	}
	return sum
}

// split objects into chunks and sum concurrently
func CalculateConcurrentSum(objects []Object, numGoroutines int) int {
	objectsSize := len(objects)
	if objectsSize == 0 {
		return 0
	}
	// number of goroutines should not be greater than number of objects
	if len(objects) < numGoroutines {
		numGoroutines = objectsSize
	}

	wg := sync.WaitGroup{}
	sumCh := make(chan int)

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start, end := calculateIndeces(i, numGoroutines, objectsSize)

		go calcConcurrentSum(sumCh, objects[start:end], &wg)
	}

	// do not block read from sumCh
	go func() {
		wg.Wait()
		close(sumCh)
	}()

	totalSum := 0
	for c := range sumCh {
		totalSum += c
	}
	return totalSum
}

// sum chunk of objects concurrently
func calcConcurrentSum(sumCh chan int, objects []Object, wg *sync.WaitGroup) {
	defer wg.Done()

	// log.Println("number of goroutine:", runtime.NumGoroutine())
	sum := 0
	for _, v := range objects {
		sum += v.A + v.B
	}
	sumCh <- sum
}

// calculate index in order to split objects into chucks using slice
func calculateIndeces(index, numGoroutines, objectsSize int) (int, int) {
	// objectSize: 10, numGoroutines: 2, groupNum: 5
	// every goroutine will handle 5 object
	groupNum := objectsSize / numGoroutines

	start := index * groupNum
	end := (index + 1) * groupNum
	// with end greater objectsSize lead to index out of range
	// or in case of odd number of objects
	// make end equal to last index
	if end > objectsSize || index == numGoroutines-1 {
		end = objectsSize
	}

	return start, end
}
