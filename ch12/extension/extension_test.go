package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet
	Pet // Anonymouis nested type 匿名嵌套类型
}

/*
func (d *Dog) Speak() {
	// d.p.Speak()
	fmt.Println("Wang!") // override
}

func (d *Dog) SpeakTo(host string) {
	// d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ", host)
}
*/

func (d *Dog) Speak() {
	fmt.Println("Wang!") // override
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao!") // output ... Chao! not Wang! Chao!
}
