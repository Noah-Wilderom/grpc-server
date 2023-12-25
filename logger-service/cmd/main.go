package main

func main() {

	app := NewServer()

	go app.Run()
}
