package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func main() {
	ch := make(chan rxgo.Item)
	// Data producer
	go producer(ch)

	// Create an Observable
	observable := rxgo.FromChannel(ch)

	observable = observable.Filter(func(item interface{}) bool {
		// Filter operation
		customer := item.(Customer)
		return customer.Age > 18
	}).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			// Enrich operation
			customer := item.(Customer)
			taxNumber, err := getTaxNumber(customer)
			if err != nil {
				return nil, err
			}
			customer.TaxNumber = taxNumber
			return customer, nil
		},
			// Create multiple instances of the map operator
			// rxgo.WithPool(pool),
			// Serialize the items emitted by their Customer.ID
			rxgo.Serialize(func(item interface{}) int {
				customer := item.(Customer)
				return customer.ID
			}), rxgo.WithBufferedChannel(1))

	for customer := range observable.Observe() {
		if customer.Error() {
			panic(customer.E)
			//return err
		}
		fmt.Println(customer)
	}
	/*
		observable := rxgo.Just("Hello, World!")()
		rxgo.Of()
		rxgo.Error()
		ch := observable.Observe()
		item := <-ch
		fmt.Println(item.V)
	*/
}

func producer(ch chan rxgo.Item) {
	for i := 0; i < 3; i++ {
		c := Customer{
			ID:        i,
			Age:       i * 20,
			TaxNumber: "",
		}
		ch <- rxgo.Of(c)
	}
	close(ch)
}

func getTaxNumber(c Customer) (string, error) {
	return "123", nil
}
