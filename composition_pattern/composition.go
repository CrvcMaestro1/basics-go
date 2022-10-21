package main

import "fmt"

type liquid struct {
	color string
}

type brand struct {
	country string
}

func (l *liquid) isColor() {
	fmt.Println(l.color)
}

type beer struct {
	name    string
	alcohol float32
	liquid
	brand
}

func (b *beer) isBeer() {
	fmt.Println("My beer is", b.name, b.color, b.alcohol, b.country)
}

func main() {
	l := liquid{color: "lila"}
	br := brand{country: "Germany"}
	b := beer{name: "GG", alcohol: 7.5, liquid: l, brand: br}
	l.isColor()
	b.isBeer()
}
