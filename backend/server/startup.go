package server

import "log"

func Start() {
	r := NewRouter()

	log.Println("server running on :8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
