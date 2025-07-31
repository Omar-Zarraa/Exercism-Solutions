package airportrobot

import "fmt"

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(visitorsName string) string
}

func SayHello(visitorsName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorsName))
}

type Italian struct {
}

func (italian Italian) LanguageName() string {
	return "Italian"
}
func (italian Italian) Greet(visitorsName string) string {
	return fmt.Sprintf("Ciao %s!", visitorsName)
}

type Portuguese struct {
}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (portuguese Portuguese) Greet(visitorsName string) string {
	return fmt.Sprintf("Olá %s!", visitorsName)
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
