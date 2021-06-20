package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func nPerson(name string, age int) *person {

    p := person{name: name,age: age}
    return &p
}


func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))
    fmt.Println(nPerson("Thomas",23))
    s1 := nPerson("Tom",26)
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
    fmt.Println(s1.name,s1.age)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}
