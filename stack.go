package main

import (
	"fmt"
)

// ArrayStack is a stack implemented using an array
type ArrayStack struct {
	elems []int
}

// Pop removes the top item from the stack
func (s *ArrayStack)Pop() (int,error){
	if s.IsEmpty(){
		return -1,fmt.Errorf("Stack is empty!")
	}
	// use a temp variable to hold the last element
	temp := s.elems[len(s.elems)-1]
	s.elems = s.elems[:(len(s.elems)-1)]
	return temp,nil

}

// Push adds an item to the top of the ArrayStack.
func (s *ArrayStack)Push(x int){
	s.elems = append(s.elems, x)
}

// Min returns the minimum element in the ArrayStack
func (s *ArrayStack)Min() (int,error){
       temp := s
	min,err := temp.Pop()
	if err != nil{
		return -1,err
	}
	for v,err := temp.Pop();err == nil;{
		 if v < min{
			 min = v
		 }


	}
	return min,nil



}

// Len returns the length of the ArrayStack
func (s *ArrayStack)Len() int{
 	return len(s.elems)



}

// Peek returns the top element of the ArrayStack
func (s *ArrayStack)Peek() (int,error){
	if s.IsEmpty(){
		return -1, fmt.Errorf("Stack is Empty!")
	}
	return s.elems[len(s.elems)-1],nil



}

// isEmpty returns true if and only if the ArrayStack is empty
func (s *ArrayStack)IsEmpty() bool{
	return len(s.elems) == 0



}

// LinkedListStack is a stack implemented using a linked list
// its head node is not a stack element, its second node is the
// top element of the stack, all push and pop operations are performed
// on the second node in order to satisfy O(1) complexity
type LinkedListStack struct {
	value int
	next *LinkedListStack
}

// Pop removes the top item from the stack
// it removes the head node of the underlying linked list
func (l *LinkedListStack)Pop() (int,error){
	if !l.IsEmpty(){
		l.value = l.next.value
		l.next = l.next.next
		return l.value,nil

	}
	return -1, fmt.Errorf("Stack is empty!")


}

// Push adds an item to the top of the LinkedListStack.
// it adds the new element as the head node of the underlying linked list
func (l *LinkedListStack)Push(x int){
        // because assignment to method receiver only propagates to callees but not to callers
	// the following won't work
	// p := &LinkedListStack{x,s}
	// l = p
	// we can not change the value of l, because the caller will not see it
	// the workaround here is change l.value and l.next, as what is done in Pop()

	temp := l.next
	newNode := &LinkedListStack{x,temp}
	l.next = newNode
}

// Min returns the minimum element in the LinkedListStack
func (l *LinkedListStack)Min() (int,error){
	temp := l
	if temp.IsEmpty(){
		return -1, fmt.Errorf("Stack is empty!")
	}
	// if temp is not empty, Pop is guaranteed to succeed
	// ignore the error
        min,_ := temp.Pop()
	for {
		v,err := temp.Pop()
		if err != nil{
			break
		}
		if v < min{
			min = v
		}


	}
	return min,nil



}

// Len returns the length of the LinkedListStack
func (l *LinkedListStack)Len() int{
 	length := 0
	for l.next != nil{
		length++
		l = l.next

	}
	return length



}

// Peek returns the top element of the LinkedListStack
func (l *LinkedListStack)Peek() (int,error){
	if !l.IsEmpty(){
		return l.next.value,nil
	}
	return -1,fmt.Errorf("Stack is empty!")




}

// isEmpty returns true if LinkedListStack is empty
func (l *LinkedListStack)IsEmpty() bool{
	return l.next == nil



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
		newSubStack := &ArrayStack{[]int{}}
		s.stacks = append(s.stacks, newSubStack)
	}
	// push data to the current substack
	s.stacks[len(s.stacks)-1].Push(x)

}

// Pop pops the last element on the current substack
func (s *SetOfStacks) Pop() (int,error){
	output,err := s.stacks[len(s.stacks)-1].Pop()
	if err != nil{
		return -1, err
	}
	// if current substack is empty then remove it
	if s.stacks[len(s.stacks)-1].Len() == 0{
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return output,nil
}

//  PopAt performs a pop operation on a specific sub-stack identified by index
func (s *SetOfStacks)PopAt(idx int) (int,error){
        if ! (idx >= 0 && idx <= len(s.stacks)){
		return -1, fmt.Errorf("Invalid index %v for SetOfStacks\n",idx)
	}
	return s.stacks[idx].Pop()
}

func (s *SetOfStacks) IsEmpty() bool{

	for i:=0;i<len(s.stacks);i++{
		if !s.stacks[i].IsEmpty(){
			return false
		}
	}
	return true
}

// todo
//Queue via Stacks: Implement a MyQueue data type which implements a queue using two stacks.



// todo
//Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use
//an additional temporary stack, but you may not copy the elements into any other data structure
//(such as an array). The stack supports the following operations: push, pop, peek, and is Empty.

// todo
// Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
// out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
// or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
// that type). They cannot select which specific animal they would like. Create the data structures to
// maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
// and dequeueCat. You may use the built-in Linked list data structure.



//type StackPrinter interface {
//	Pop()(int,error)
//	IsEmpty() bool
//}

func (l *LinkedListStack) String() string{
        tempArray := []int{}
	temp := l
	item,err := temp.Pop()

	for err==nil{
		tempArray = append(tempArray, item)
		item,err = temp.Pop()
	}

	s := ""
	for _,j :=range tempArray {
	     s += fmt.Sprintf("|__%-2d__|\n",j)

	}
	return s
}


func (a *ArrayStack) String() string{
        tempArray := []int{}
	item,err := a.Pop()
	if err != nil {
		tempArray = append(tempArray, item)
	}

	s := ""
	for _,j :=range tempArray {
	     s += fmt.Sprintf("|__%-2d__|\n",j)

	}
	return s
}

func (s *SetOfStacks) String() string{
        tempArray := []int{}
	item,err := s.Pop()
	if err != nil {
		tempArray = append(tempArray, item)
	}

	output := ""
	for _,j :=range tempArray {
	     output += fmt.Sprintf("|__%-2d__|\n",j)

	}
	return output
}

func LinkedListStack_test(){
	p := fmt.Printf
	l := &LinkedListStack{}
	//l := new(LinkedListStack)
	//l.value = 3
	//l.next = nil
	p("Type: %T\tValue: %p\n",l,l.next)
	p("Empty? %t\n",l.IsEmpty())
	l.Push(3)
	l.Push(-4)
	l.Push(5)
	l.Push(10)
	p("Stack: \n%s",l)
	// todo: use deep copy instead of shallow copy
	// current implementations of  Min(), String() use
	// Pop() to have 0(1) complexity, but This method introduces
	// a bug, the local variable temp inside Min(),String() is
	// a shallow copy(reference) of the receiver l
	// when I do temp.Pop(), receiver l will be changed as well
	// I certainly can implement those two using linked list traversal, but
	// they will have O(n) complexity.
	// the solution is to write deepcopy function to get independent local
	// variables
	p("Length : %v\n", l.Len())
	top,_ := l.Peek()
	p("Peek : %v\n", top)
	min,_ := l.Min()
	p("Minimum : %v\n", min)
	l.Pop()
	l.Pop()
	p("Stack: %s\n",l)
	top,_ = l.Peek()
	p("Peek : %v\n", top)
	p("Length : %v\n", l.Len())
}


func main(){
	LinkedListStack_test()

}