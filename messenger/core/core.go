package core

import "fmt"

var messengers map[Name]Messenger = make(map[Name]Messenger)

// Name !
type Name string

// Messenger !
type Messenger interface {
	StartMessageConsumer() error
	CloseConnection() error
}

// Register !
func Register(name Name, impl Messenger) {
	messengers[name] = impl
}

// Get !
func Get(name Name) (Messenger, error) {
	var broker, exists = messengers[name]
	if !exists {
		return nil, fmt.Errorf("Messager service %v not implemented", name)
	}
	return broker, nil
}
