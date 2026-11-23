package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Go",
		117: "Python",
		346: "Java",
	}
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[346] = "Cloud-Lösungen"
	// TODO: replace one
	modules[104] = "Go Programming"

	fmt.Println(modules)
}