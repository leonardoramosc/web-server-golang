package main

type P struct {
	name string
}

func main() {

	server := NewServer(":3000")

	server.Listen()

}
