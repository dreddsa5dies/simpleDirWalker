package main

import (
	"fmt"
	"log"
	"os"

	s "github.com/dreddsa5dies/simpleDirWalker"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(os.Args[0] + " catalog")
	} else {
		lstFiles, err := s.SDW(os.Args[1])
		if err != nil {
			log.Println(err)
		}

		for _, v := range lstFiles {
			fmt.Println(v.FullPath, v.FileInfo.Size())
		}
	}
}
