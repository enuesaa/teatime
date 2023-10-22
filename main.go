package main

func main() {
	app := setupRouter()
	app.Run(":3000")
}
