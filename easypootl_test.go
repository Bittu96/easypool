package easypool

import (
	"fmt"
	"log"
	"testing"
	"time"

	"math/rand"

	"go.uber.org/goleak"
)

// go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

func TestCase1(t *testing.T) {
	defer goleak.VerifyNone(t)

	// case 1 mock
	log.Println("case 1 mock test start")
	defer log.Println("case 1 mock test end")

	ep := New(mockTaskFunc)
	ep.Deploy(10)
}

func mockTaskFunc(msg interface{}) interface{} {
	// log.Println("receiving..", msg)
	mockFunctionDelay()
	// log.Println("received", msg)
	return msg
}

func mockFunctionDelay() {
	delay := time.Duration(rand.Intn(10)) * time.Millisecond
	time.Sleep(delay)
}

func TestCase2(t *testing.T) {
	defer goleak.VerifyNone(t)

	// case 2 mock
	log.Println("case 2 mock start")
	defer log.Println("case 2 mock end")

	// mock params
	var (
		mockChnlSize = 100
		inflow       = make(chan interface{}, mockChnlSize)
	)

	for i := range mockChnlSize {
		inflow <- fmt.Sprintf("msg %v", i+1)
	}

	ep := New(mockTaskFunc).AddInflow(&inflow)
	ep.Deploy(10)
}

func TestCase3(t *testing.T) {
	defer goleak.VerifyNone(t)

	// case 3 mock
	log.Println("case 3 mock start")
	defer log.Println("case 3 mock end")

	// mock params
	var (
		mockChnlSize = 100
		inflow       = make(chan interface{}, mockChnlSize)
		outflow      = make(chan interface{}, mockChnlSize)
	)

	for i := range mockChnlSize {
		inflow <- fmt.Sprintf("msg in %v", i+1)
	}

	ep := New(mockTaskFunc).AddInflow(&inflow).AddOutflow(&outflow)
	ep.Deploy(10)

	for range mockChnlSize {
		<-outflow
		// res := <-outflow
		// fmt.Println("msg out:", res)
	}
}
