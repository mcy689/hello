package hello

import "github.com/mcy689/hello/world"
import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}

func CallWorld() string {
	return world.SayWorld()
}
