package main

func main() {
	chanMain := make(chan uint8)
	for i := 0; i < 10; i++ {
		chanMain <- uint8(i)
	}

}
