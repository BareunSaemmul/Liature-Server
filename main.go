package main

import "Liature-Server/server"

func main() {
	s, err := server.New()
	if err != nil {
		panic(err)
	}
	s.Run(":3000")
}
