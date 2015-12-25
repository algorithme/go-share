package main

import (
	"fmt"
	"os"

	"ch.devbox/hash"
)

func main() {
	f, err := os.Open("/Users/omorel/Dropbox/video/clojure/Rich Hickey - Deconstructing the Database-Cym4TZwTCNU.mp4")
	defer f.Close()

	if err != nil {
		fmt.Printf("An error has occured when opening file %s", err)
		os.Exit(1)
	}

	_, err = hash.File(f, 32*1000*1000)

	if err != nil {
		fmt.Printf("An error has occured when opening file %s", err)
		os.Exit(1)
	}

}
