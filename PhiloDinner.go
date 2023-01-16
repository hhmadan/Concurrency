package main

import (
	"fmt"
	"sync"
)

var host = make(chan bool, 2)
var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	philoNum        int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for i := 1; i <= 3; i++ {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %d is Eating...\n", p.philoNum+1)
		fmt.Printf("Philosopher %d is Finishes Eating...!\n", p.philoNum+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
