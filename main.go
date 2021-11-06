package main

func main() {
	message, err := Hello("Master")
	if err != nil {
		print(err.Error())
		return
	}

	print(message)
}
