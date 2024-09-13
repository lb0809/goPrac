package main

import "fmt"

type avengers []string

func (d avengers) assemble() {
	for i, hero := range d {
		fmt.Println(i, hero)
	}
}
