package _package

import "fmt"

func (this *MinStack) PrintStack()  {
	fmt.Print("printf normal stack:")
	for _, v := range this.stack{
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Print("printf min stack:")
	for _, v := range this.min{
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
