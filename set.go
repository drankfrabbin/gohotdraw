package gohotdraw

import (
	"container/vector"
)

type Set struct {
	*vector.Vector
}

func NewSet() *Set {
	return &Set{new(vector.Vector)} 
}

//Convenience method for Push()
func (this *Set) Push(element interface{}) {
	this.Add(element)
}

//Adds an element only if it is not already in the set
func (this *Set) Add(element interface{}) {
	if !this.Contains(element) {
		this.Vector.Push(element)
	}
}

//Checks if element is conatined in set
func (this *Set) Contains(element interface{}) bool {
	for i := 0; i < this.Vector.Len(); i++ {
		currentElement := this.Vector.At(i)
		if currentElement == element {
			return true
		}
	}
	return false
}

//Removes an element from the set
func (this *Set) Remove(element interface{}) {
	for i := 0; i < this.Vector.Len(); i++ {
		currentElement := this.Vector.At(i)
		if currentElement == element {
			this.Vector.Delete(i)
			return
		}
	}
}

func (this *Set) Replace(toBeReplaced, replacement interface{}) {
	for i := 0; i < this.Vector.Len(); i++ {
		currentElement := this.Vector.At(i)
		if currentElement == toBeReplaced {
			this.Vector.Set(i, replacement)
		}
	}
}

//Returns a new Set with the elements of the original Set
func (this *Set) Clone() *Set {
	set := NewSet()
	for element := range this.Iter() {
		set.Push(element)
	}
	return set
}

func (this *Set) Iter() <-chan interface{} {
	c := make(chan interface{})
	go this.iterate(c)
	return c
}

func (this *Set) iterate(c chan<- interface{}) {
	for _,v := range *this.Vector {
		c <- v
	}
	close(c)
}
