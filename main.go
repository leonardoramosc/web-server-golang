package main

type P struct {
	name string
}

func main() {

	server := NewServer(":3000")

	server.Handle("/", HandleRoot)
	server.Handle("/products", HandleProducts)

	server.Listen()

}
