package main

// https://medium.com/capital-one-tech/generics-are-the-generics-of-go-3e0ef0cb9e04
import (
	"fmt"
)

type LinkedList[type T] struct {
value T
next  *LinkedList[T]
}

func (ll *LinkedList[T]) Len() int {
	count := 0
	for node := ll; node != nil; node = node.next {
		count++
	}
	return count
}

func (ll *LinkedList[T]) Append(value T) *LinkedList[T] {
	return ll.InsertAt(ll.Len(), value)
}

func (ll *LinkedList[T]) InsertAt(pos int, value T) *LinkedList[T] {
	if ll == nil || pos <= 0 {
		return &LinkedList[T]{
		value: value,
			next:  ll,
		}
	}
	ll.next = ll.next.InsertAt(pos-1, value)
	return ll
}

func (ll *LinkedList[T]) String() string {
	if ll == nil {
		return "nil"
	}
	return fmt.Sprintf("%v->%v", ll.value, ll.next.String())
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var head *LinkedList[string]
	head = head.Append("Hello")
	fmt.Println(head.String())
	fmt.Println(head.Len())
	head = head.Append("Hola")
	head = head.Append("हैलो")
	head = head.Append("こんにちは")
	head = head.Append("你好")
	fmt.Println(head.String())
	fmt.Println(head.Len())

	var peopleList *LinkedList[Person]
	peopleList = peopleList.Append(Person{"Fred", 23})
	peopleList = peopleList.Append(Person{"Joan", 30})
	fmt.Println(peopleList)
}
