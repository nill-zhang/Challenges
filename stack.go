package main

import (
	"fmt"
	"log"
)

type ArrayStack struct {
	datas []int
}

// Pop removes the top item from the stack
func (s *ArrayStack)Pop() int{
	defer func(){
		s.datas = s.datas[:(len(s.datas)-1)]
	}()
	return s.datas[len(s.datas)-1]

}

// Push adds an item to the top of the ArrayStack.
func (s *ArrayStack)Push(x int){
	s.datas = append(s.datas, x)
}

// Min returns the minimum element in the ArrayStack
func (s *ArrayStack)Min() int{
       temp := s
	if temp.IsEmpty(){
	       return -1
       }
	min := temp.Pop()
	for !temp.IsEmpty(){
                 v := temp.Pop()
		 if v < min{
			 min = v
		 }


	}
	return min



}

// Len returns the length of the ArrayStack
func (s *ArrayStack)Len() int{
 	return len(s.datas)



}

// Peek return the top of the ArrayStack
func (s *ArrayStack)Peek() int{
	return s[len(s)-1]



}

// isEmpty returns true if and only if the ArrayStack is empty
func (s *ArrayStack)IsEmpty() bool{
	return len(s.datas) == 0



}


type LinkedListStack struct {
	value int
	next *LinkedListStack
}

// Pop removes the top item from the stack
func (s *LinkedListStack)Pop() int{
	defer func(){
		s = s.next
	}()
	return s.value

}

// Push adds an item to the top of the LinkedListStack.
func (s *LinkedListStack)Push(x int){

	p := &LinkedListStack{x,s}
	s = p
}

// Min returns the minimum element in the LinkedListStack
func (s *LinkedListStack)Min() int{
       min := s.value
	for s!=nil{
		if s.value < min{
			min = s.value
		}
		s = s.next


	}
	return min



}

// Len returns the length of the LinkedListStack
func (s *LinkedListStack)Len() int{
 	length := 0
	for s != nil{
		length++
		s = s.next

	}
	return length



}

// Peek return the top of the LinkedListStack
func (s *LinkedListStack)Peek() int{
	return s.value



}

// isEmpty returns true if and only if the LinkedListStack is empty
func (s *LinkedListStack)IsEmpty() bool{
	return s == nil



}

//Three in One: Describe how you could use a single array to implement three stacks.
//Hints: #2, #72, #38, #58

func ThreeInOne(){





}

//Stack of Plates: Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
//Therefore, in real life, we would likely start a new stack when the previous stack exceeds some
//threshold. Implement a data structure SetOfStacks that mimics this. SetO-fStacks should be
//composed of several stacks and should create a new stack once the previous one exceeds capacity.
//SetOfStacks. push() and SetOfStacks. pop() should behave identically to a single stack
//(that is, pop () should return the same values as it would if there were just a single stack).
// SetOfStacks is a plates of stacks, once the previous substack exceeds its capacity
//  a new substack will be created and used as the current substack
// all stack ops will be preformed on the current substack
type SetOfStacks struct {
	stacks []*ArrayStack
	capacity int
}

// Push pushes x to the current substack in use
func (s *SetOfStacks)Push(x int){
	// if the current substack is full, created a new empty substack
	if s.stacks[len(s.stacks)-1].Len() == s.capacity{
		newstack := &ArrayStack{[]int{}}
		s.stacks = append(s.stacks, newstack)
	}
	// push data to the current substack
	s.stacks[len(s.stacks)-1].Push(x)

}

// Pop pops the last element on the current substack
func (s *SetOfStacks) Pop(){
	s.stacks[len(s.stacks)-1].Pop()
	// if current substack is empty then remove it
	if s.stacks[len(s.stacks)-1].Len() == 0{
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
}

//  PopAt performs a pop operation on a specific sub-stack identified by index
func (s *SetOfStacks)PopAt(idx int){
        if ! (idx >= 0 && idx <= len(s.stacks)){
		log.Fatalf("Invalid index %v for SetOfStacks\n",idx)
	}
	s.stacks[idx].Pop()
}


//Queue via Stacks: Implement a MyQueue data type which implements a queue using two stacks.




//Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use
//an additional temporary stack, but you may not copy the elements into any other data structure
//(such as an array). The stack supports the following operations: push, pop, peek, and is Empty.


//Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
//out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
//or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
//that type). They cannot select which specific animal they would like. Create the data structures to
//maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
//and dequeueCat. You may use the built-in Linked list data structure.

func main(){
	s := []int{2,2,43,4}
	s = s[:3]
	fmt.Println(len(s),cap(s),s)

}