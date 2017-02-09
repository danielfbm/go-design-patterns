package main

import (
	"fmt"
	"time"
)

type Observable interface {
	Notify(data ...interface{})
}

type Subject interface {
	Subscribe(ob Observable)
	Unsubscribe(ob Observable)
	NotifyObservables(data ...interface{})
}

// Subject ----------

// EverySecond
type EverySecond struct {
	observables []Observable
}

func (e *EverySecond) Subscribe(ob Observable) {
	e.observables = append(e.observables, ob)
}

func (e *EverySecond) Unsubscribe(ob Observable) {
	for i, o := range e.observables {
		if o == ob {
			e.observables = append(e.observables[:i], e.observables[i+1:]...)
			break
		}
	}
}

func (e *EverySecond) NotifyObservables(data ...interface{}) {
	for _, r := range e.observables {
		if r != nil {
			r.Notify(data...)
		}
	}
}

func (e *EverySecond) Start() {
	for {
		time.Sleep(time.Second)
		e.NotifyObservables(time.Now())
	}
}

// Observable ----------

// Printer
type Printer struct{}

func (Printer) Notify(data ...interface{}) {
	fmt.Println(data...)
}

// FormattedPrinter
type FormatedPrinter struct{}

func (FormatedPrinter) Notify(data ...interface{}) {
	fmt.Printf("[{-_-}] ZZZzz zz z... %v \n", data...)
}

var secondNotifer *EverySecond

func init() {
	secondNotifer = &EverySecond{}
}

func main() {

	printer := Printer{}
	formated := FormatedPrinter{}

	secondNotifer.Subscribe(printer)
	secondNotifer.Subscribe(formated)

	secondNotifer.Start()
}
