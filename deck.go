package main

import "fmt"

// 	Create A New Type Of 'deck'
//	Which Is A Slice Of Strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
