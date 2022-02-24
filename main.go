package main

import (
	"flag"
	"fmt"
	"lesson6/mycrypt"
	"log"
)

func main() {
	//action := "enc"
	var fileSource, hashFile, outFile string
	flag.StringVar(&fileSource, "source-file", "source.txt", "File Source")
	flag.StringVar(&hashFile, "hash-file", "hash.txt", "File hash")
	flag.StringVar(&outFile, "out-file", "sign.txt", "File output")
	flag.Parse()
	action := flag.Args()[0]
	switch action {
	case "enc":
		encoder, err := mycrypt.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(encoder)
		err = encoder.EncryptSha256()
		if err != nil {
			panic(err)
		}
		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}
	case "dec":
		decoder, err := mycrypt.NewDecryptor(hashFile, fileSource, outFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(decoder)
		err = decoder.Validate()
		if err != nil {
			panic(err)
		}
		err = decoder.SaveToFile("decoded.txt")
		if err != nil {
			panic(err)
		}
	default:
		log.Fatal("use enc or dec")
	}
	return
}
