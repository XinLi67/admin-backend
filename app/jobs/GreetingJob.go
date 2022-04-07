package jobs

import "fmt"

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("hello", g.Name)
}
