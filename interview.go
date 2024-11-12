package main

import "sync"

type result struct {
	num int
	idx int
}

func sliceSum(arr []int) []int {
	var wg sync.WaitGroup
	output := make([]int, len(arr), len(arr))
	resChan := make(chan result)

	totalSum := 0

	for _, v := range arr {
		totalSum += v
	}

	for i, v := range arr {
		wg.Add(1)

		go func() {
			defer wg.Done()
			sum := totalSum - v

			resChan <- result{num: sum, idx: i}
		}()
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()

	for res := range resChan {
		output[res.idx] = res.num
	}

	return output
}
