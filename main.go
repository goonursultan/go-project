package main

import "fmt"

func main()  {
	var s string = Hi("Nursultan")
	fmt.Println(s)
}
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}