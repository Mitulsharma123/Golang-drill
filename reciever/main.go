// receiver setup method on variable we create

package main

func main() {
	fruits := basket{"apple", "banana", "oranages", "kiwi"}
	fruits = append(fruits, "mango")

	fruits.show()
}
