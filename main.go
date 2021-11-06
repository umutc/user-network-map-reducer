package main

func main() {
	message, err := Hello("Umut")
	if err != nil {
		print(err.Error())
		return
	}

	print(message)
}
