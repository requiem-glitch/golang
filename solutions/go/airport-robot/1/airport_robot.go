package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

type Italian struct {}

func (a Italian) LanguageName() string {
    return " Italian"
}

func (a Italian) Greet(b string) string {
    return " Ciao " + b + "!"
}

type Portuguese struct {}

func (a Portuguese) LanguageName() string {
    return " Portuguese"
}

func (a Portuguese) Greet(b string) string {
    return " Olá " + b + "!"
}

func SayHello(visitor string, a Greeter) string {
    return "I can speak" + a.LanguageName() + ":" + a.Greet(visitor) 
}
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
