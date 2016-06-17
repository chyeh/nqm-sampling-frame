package main

func main() {
	go pull()
	go update()
	select {}
}
