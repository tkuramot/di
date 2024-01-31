package main

import (
  "fmt"
)

// Message
type Message string

// Event
func NewEvent(g Greeter) Event {
    return Event{Greeter: g}
}

type Event struct {
    Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
    msg := e.Greeter.Greet()
    fmt.Println(msg)
}

func NewMessage() Message {
    return Message("Hi there!")
}

// Gteeter
func NewGreeter(m Message) Greeter {
    return Greeter{Message: m}
}

type Greeter struct {
    Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
    return g.Message
}

func main() {
    message := NewMessage()
    greeter := NewGreeter(message)
    event := NewEvent(greeter)

    event.Start()
}
