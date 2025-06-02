package pkg

import (
	"fmt"
	"sync"
	"time"
)

func output(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func ex1() {
	// Declare a waitgroup
	var wg sync.WaitGroup

	// Add(3) because of 3 functions of concurrent
	wg.Add(3)

	// Concurrent function list
	go output(&wg, "Playing Dota2")
	go output(&wg, "Reading")
	go output(&wg, "GGEZ")

	// Waiting to finish
	wg.Wait()
}

func ex2() {
	// Declare a waitgroup
	var wg sync.WaitGroup

	// Add(3) because of 3 functions of concurrent
	wg.Add(3)

	// Concurrent function list
	go func() {
		defer wg.Done()
		fmt.Println("Playing Dota2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Reading")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("GGEZ")
	}()

	// Waiting to finish
	wg.Wait()
}

func ex3() {
	// Declare a waitgroup
	var wg sync.WaitGroup

	// Add(1) because of 1 functions of concurrent
	wg.Add(1)

	// Concurrent function list
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%d \n", i)
		}
	}()

	// Waiting to finish
	wg.Wait()
}

func ex4() {
	// Declare a waitgroup
	var wg sync.WaitGroup

	// Add(1) because of 1 functions of concurrent

	// Concurrent function list
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d \n", i)
		}()

		wg.Wait()
	}

}

func ex5() {
	arr1 := []int{
		1, 2, 3, 4, 5, 6,
	}
	arr2 := []int{
		10, 30, 50, 70,
	}

	sumArr1 := make(chan int)
	sumArr2 := make(chan int)

	go func(arr1 []int, sum chan int) {
		var result int
		for i := range arr1 {
			result += arr1[i]
		}

		sum <- result
		close(sum)
	}(arr1, sumArr1)

	go func(arr2 []int, sum chan int) {
		var result int
		for i := range arr2 {
			result += arr2[i]
		}

		sum <- result
		close(sum)
	}(arr2, sumArr2)

	sum1 := <-sumArr1
	sum2 := <-sumArr2

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func ex6() {
	result1 := make(chan string)
	result2 := make(chan string)

	go func() {
		for {
			result1 <- "100 ms"
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for {
			result2 <- "200 ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	for {
		select {
		case <-result1:
			fmt.Println(<-result1)
		case <-result2:
			fmt.Println(<-result2)
		}
	}
}

func RunGoRoutine() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	ex6()
}
