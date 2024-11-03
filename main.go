package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	n := 10                       // Set the upper limit
	printed := make(map[int]bool) // Create a map to track printed numbers

	printNumbers(n, printed) // Call the function to print numbers
}

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Animal noises"
}

func (a Dog) Speak() string {
	return "Bark"
}

type Dog struct {
	Animal
	Breed string
}

type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func (p *Person) CelebrateBirthday() {
	p.Age++
}

type Greeter interface {
	Greet() string
}

func add(x int, y int) int {
	return x + y
}

func divide(divident, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return divident / divisor, nil
}

func multiply(x, y int) (result int) {
	result = x * y
	return
}

func fizzBuzz() {
	for i := 0; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func channelFunc() {
	ch := make(chan int, 2)
	defer close(ch)

	// go func() {
	// 	ch <- 42
	// }()
	ch <- 1
	ch <- 2

	value := <-ch
	fmt.Println(value)
	value = <-ch
	ch <- 3
	fmt.Println(value)
	value = <-ch
	fmt.Println(value)
}

func doSomething() error {
	return &CustomError{Msg: "Something went wrong"}
}

// Fan-out: Distributing tasks to multiple goroutines.
func fanOut() {
	tasks := make(chan int, 10)

	for i := 0; i < 5; i++ {
		go func() {
			for task := range tasks {
				fmt.Println(task)
			}
		}()
	}
}

// Fan-out: Distributing tasks to multiple goroutines.
func fanIn() {
	results := make(chan int)

	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()
}

// This pattern uses a fixed number of goroutines to process jobs from a job queue.
func workerPool() {
	const numWorkers = 5
	jobs := make(chan int)

	for w := 0; w < numWorkers; w++ {
		go func() {
			for job := range jobs {
				fmt.Println(job)
			}
		}()
	}
}

func selectChan() {
	results := make(chan string)
	errors := make(chan error)

	select {
	case result := <-results:
		fmt.Println("Received result:", result)
	case err := <-errors:
		fmt.Println("Received error:", err)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout!")
	}
}

func waitGroup() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Completed work")
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
		}
	}()
	wg.Wait()
}

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an int:", v)
	case string:
		fmt.Println("This is a string:", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

// rand function that simulates generating a random number
func randNum(n int) int {
	return rand.Intn(n + 1)
}

// Function to print all numbers from 0 to n
func printNumbers(n int, printed map[int]bool) {
	// Base case: if all numbers are printed, return
	if len(printed) == n+1 {
		return
	}

	// Generate a random number
	num := randNum(n)

	// If the number hasn't been printed, print it
	if !printed[num] {
		fmt.Println(num)
		printed[num] = true // Mark it as printed
	}

	// Recur to print more numbers
	printNumbers(n, printed)
}
