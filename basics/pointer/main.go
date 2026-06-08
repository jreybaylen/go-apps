package pointer

type Book struct {
	title  string
	author string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func (b *Book) setAuthor(author string) {
	b.author = author
}

func PointerMain() {
	x := 10
	y := &x

	println("Value of x:", x)
	println("Address of x:", &x)
	println("Value of y (address of x):", y)
	println("Dereferenced value of y (value of x):", *y)

	x = 15
	println("\nAfter changing x to 15:")
	println("Value of x:", x)
	println("Dereferenced value of y (value of x):", *y)

	b := &Book{}
	b.setTitle("The Go Programming Language")
	b.setAuthor("Alan A. A. Donovan and Brian W. Kernighan")

	println("Title:", b.title)
	println("Author:", b.author)
}
