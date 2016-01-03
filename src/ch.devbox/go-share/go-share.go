package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"ch.devbox/hash"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var memprofile = flag.String("memprofile", "", "write memory profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		cf, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cf)
		defer pprof.StopCPUProfile()
	}

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

	hash.ListFiles("/Users/omorel/Dropbox/video/")

	if *memprofile != "" {
		mf, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(mf)
		defer mf.Close()
	}
}
