package main

import ("fmt"
        "time")

type Order struct {
	TableNumber int
	PrepTime time.Duration
}  

func processOrder(order Order) {
  //simulate cooking time
  fmt.Printf("Preparing order for table %d...\n", order.TableNumber)

  time.Sleep(order.PrepTime)

  fmt.Printf("Order ready for table %d!\n\n", order.TableNumber)
}

func main(){
	orders:=[]Order{
		{TableNumber:1, PrepTime:2*time.Second},
		{TableNumber:2, PrepTime:3*time.Second},
		{TableNumber:3, PrepTime:1*time.Second},
		{TableNumber:4, PrepTime:4*time.Second},
		{TableNumber:5, PrepTime:2*time.Second},
	}
	for _,order:=range orders{
		go processOrder(order) // Process each order concurrently, what happening there? In there is a loop that iterates over the orders and for each order, it starts a new goroutine to process the order concurrently. 
		// This means that all orders will be processed at the same time, rather than waiting for one order to finish before starting the next one. 
		// The main function will continue executing and will wait for user input (via fmt.Scanln()) to prevent the program from exiting immediately, allowing all goroutines to complete their tasks.
	}
	// Wait for all orders to be processed
	fmt.Scanln()
}