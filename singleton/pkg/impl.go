package pkg

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Count int
}

type SingletonItf interface {
	AddOne() int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(
		func() {
			fmt.Println("Created new object")
			instance = new(Singleton)
		},
	)
	return instance
}

func (s *Singleton) AddOne() int {
	s.Count++
	return s.Count
}
