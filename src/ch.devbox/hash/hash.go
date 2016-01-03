package hash

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
)

type FileHash struct {
	FileName string
	Path     string
	Hashes   []string
}

var (
	encoder hash.Hash = sha256.New()
)

func File(f *os.File, blockSize int64) (hashes []string, err error) {
	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		return
	}

	// calculate number of parts
	size := stat.Size()
	parts := (size / blockSize) + 1
	hashes = make([]string, parts)
	reader := bufio.NewReaderSize(f, 32*1000*1000)

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
	iterations := blockSize / 32
	bufferRest := blockSize % 32

	if bufferRest > 0 {
		iterations++
	}

	for i := 0; i < int(iterations); i++ {
		buffer := make([]byte, 32)

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
	}

	hash = hex.EncodeToString(encoder.Sum(nil))
	log.Printf("%s - %s", hash, err)

	return
}

func ListFiles(path string) (files []FileHash, err error) {
	handleFile := func(path string, info os.FileInfo, err error) error {
		log.Printf("Handle file %s with name: %s", path, info.Name())
		if info.IsDir() {
			return err
		}

		f, err := os.Open(path)
		defer f.Close()

		hs, err := File(f, 64*1000*1000)
		if err != nil {
			log.Fatal(err)
			return err
		}

		files = append(files, FileHash{
			FileName: info.Name(),
			Path:     path,
			Hashes:   hs,
		})

		return err
	}

	filepath.Walk(path, handleFile)
	return
}
