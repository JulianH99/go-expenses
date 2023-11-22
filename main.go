package main

func main() {

	server := NewApiServer("localhost", 8080)

	server.Run()

}
