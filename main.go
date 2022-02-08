package main

import (
	"flag"

	"github.com/nao11aihara/golang-euler/euler1"
	"github.com/nao11aihara/golang-euler/euler2"
	"github.com/nao11aihara/golang-euler/euler3"
)

func main(){
	flag.Parse()

	switch flag.Arg(0){
		case "1":
			euler1.Execute()
		case "2":
			euler2.Execute()
		case "3":
			euler3.Execute()
	}
}
