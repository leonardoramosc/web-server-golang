package main

func main() {

	server := NewServer(":3000")

	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/products", server.AddMiddleware(HandleProducts, CheckAuth(), Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/users", UserPostRequest)

	server.Listen()

}
