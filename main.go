package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/c2h5oh/datasize"
)

func generateRandomChar(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	var args struct {
		File string            `arg:"required,positional" help:"File to write"`
		Size datasize.ByteSize `arg:"-s,--size" help:"Size to write" default:"1M"`
	}
	arg.MustParse(&args)

	file, err := os.Create(args.File)
	if err != nil {
		_ = fmt.Errorf("Failed to open file.")
		os.Exit(1)
	}

	var i float64 = 0
	var mbytes = args.Size.MBytes()

	for ; i < mbytes; i++ {
		file.Write([]byte(generateRandomChar(int(datasize.MB))))
	}

	file.Close()
}
