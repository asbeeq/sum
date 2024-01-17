package object

import (
	"sync"
	"testing"
)

func TestCalcConcurrentSum(t *testing.T) {
	// test data
	objects := []Object{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: 5, B: 6},
	}
	expectedSum := 21
	numGoroutines := len(objects)

	sumCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for _, v := range objects {
		go calcConcurrentSum(sumCh, []Object{v}, &wg)
	}

	go func() {
		wg.Wait()
		close(sumCh)
	}()

	totalSum := 0
	for c := range sumCh {
		totalSum += c
	}

	// assert actual sum with expected sum
	if totalSum != expectedSum {
		t.Errorf("expected %d, got %d", expectedSum, totalSum)
	}
}

func TestCalculateIndeces(t *testing.T) {
	tests := []struct {
		index         int
		numGoroutines int
		objectsSize   int
		expectedStart int
		expectedEnd   int
	}{
		{0, 5, 20, 0, 4}, // 5 goroutines each of 4 objects
		{1, 5, 20, 4, 8},
		{3, 5, 20, 12, 16},
		{4, 5, 20, 16, 20}, // end should be equal to objectsSize
		{3, 7, 25, 9, 12},
		{0, 7, 25, 0, 3},
		{6, 7, 25, 18, 25}, // last goroutine, end should be equal to objectsSize
	}

	for _, test := range tests {
		start, end := calculateIndeces(test.index, test.numGoroutines, test.objectsSize)

		if start != test.expectedStart || end != test.expectedEnd {
			t.Errorf("for input (%d, %d, %d), expected (%d, %d), but got (%d, %d)",
				test.index,
				test.numGoroutines,
				test.objectsSize,
				test.expectedStart,
				test.expectedEnd,
				start,
				end,
			)
		}
	}
}

func TestCalculateConcurrentSum(t *testing.T) {
	// test data
	objects := []Object{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: 5, B: 6},
		{A: 7, B: 8},
		{A: 9, B: 10},
	}
	numGoroutines := 2
	expectedSum := 55

	actualSum := CalculateConcurrentSum(objects, numGoroutines)

	// assert actual sum with expected sum
	if actualSum != expectedSum {
		t.Errorf("expected %d, got %d", expectedSum, actualSum)
	}
}

func TestCalculateSequentialSum(t *testing.T) {
	// test data
	objects := []Object{
		{A: 1, B: 2},
		{A: 3, B: 4},
		{A: 5, B: 6},
	}
	expectedSum := 21

	actualSum := CalculateSequentialSum(objects)

	// assert actual sum with expected sum
	if actualSum != expectedSum {
		t.Errorf("expected %d, got %d", expectedSum, actualSum)
	}
}
