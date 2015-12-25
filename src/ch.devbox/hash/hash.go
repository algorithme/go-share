package hash

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"log"
	"os"
)

type Hash string

var encoder hash.Hash = sha256.New()

func File(f *os.File, blockSize int64) (hashes []string, err error) {
	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return
	}

	// calculate number of parts
	size := stat.Size()
	parts := (size / blockSize) + 1
	hashes = make([]string, size)
	reader := bufio.NewReaderSize(f, int(size))

	log.Printf("number of parts %d with a size of %d", parts, size)

	var i int64

	for i = 0; i < parts; i++ {
		offset := blockSize * i
		_, err = f.Seek(offset, 0)
		if err != nil {
			log.Fatal(err)
			return
		}

		hash, err := FilePart(reader, blockSize, offset)

		if err != nil {
			log.Fatal(err)
		} else {
			hashes[i] = hash
		}

	}
	log.Printf("Finished")

	return
}

func FilePart(b *bufio.Reader, blockSize int64, offset int64) (hash string, err error) {
	buffer := make([]byte, blockSize)

	_, err = b.Read(buffer)

	if err != nil && err != io.EOF {
		log.Fatal(err)
		return
	}

	_, err = encoder.Write(buffer)

	if err != nil {
		log.Fatal(err)
		return
	}
	hash = hex.EncodeToString(encoder.Sum(nil))
	log.Printf("%s - %s", hash, err)

	return
}
