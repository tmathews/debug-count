// This is "debug count", a package for simply spitting out numbers in your code
// it is purely for debugging purposes. If anyone adds a singleton to this
// package there will be some serious repercussions...

package dc

import "fmt"

type Counter struct {
	Name  string
	Count int
}

func (c *Counter) Print(xs... interface{}) {
	c.Count++
	fmt.Printf("%s: %d", c.Name, c.Count)
	if len(xs) > 0 {
		fmt.Printf(": ")
		fmt.Print(xs...)
	}
	fmt.Printf("\n")
}

func (c *Counter) Oprint(xs... interface{}) {
	fmt.Printf("%s: %d: ", c.Name, c.Count)
	fmt.Print(xs...)
	fmt.Print("\n")
}

func New(name string) *Counter {
	d := &Counter{
		Count: 0,
		Name: name,
	}
	d.Print()
	return d
}
