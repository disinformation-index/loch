package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/nacl/secretbox"
)

func loadKey(keyArg string) [blake2b.Size256]byte {
	var key string
	if keyArg == "" {
		var rawKey string
		fmt.Printf("passphrase:")
		reader := bufio.NewReader(os.Stdin)
		rawKey, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		key = strings.TrimSpace(rawKey)
	} else {
		keyData, err := ioutil.ReadFile(keyArg)
		if err != nil {
			log.Fatal(err)
		}
		key = strings.TrimSpace(string(keyData))
	}
	h := blake2b.Sum256([]byte(key))
	return h
}

func doEncrypt(inputArg string, key [blake2b.Size256]byte) []byte {
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		log.Fatal(err)
	}
	plainData, err := ioutil.ReadFile(inputArg)
	if err != nil {
		log.Fatal(err)
	}
	encrypted := secretbox.Seal(nonce[:], plainData, &nonce, &key)
	return encrypted
}

func doDecrypt(inputArg string, key [blake2b.Size256]byte) []byte {
	encrypted, err := ioutil.ReadFile(inputArg)
	if err != nil {
		log.Fatal(err)
	}
	var nonce [24]byte
	copy(nonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open(nil, encrypted[24:], &nonce, &key)
	if !ok {
		log.Fatal("Decryption error")
	}
	return decrypted
}

func outputFile(content []byte, outputArg string) {
	if outputArg == "" {
		os.Stdout.Write(content)
	} else {
		err := ioutil.WriteFile(outputArg, content, 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	keyArg := flag.String("key", "", "Key file")
	outArg := flag.String("out", "", "Output file")
	flag.Parse()

	switch flag.Arg(0) {
	case "encrypt":
		privKey := loadKey(*keyArg)
		content := doEncrypt(flag.Arg(1), privKey)
		outputFile(content, *outArg)
	case "decrypt":
		privKey := loadKey(*keyArg)
		content := doDecrypt(flag.Arg(1), privKey)
		outputFile(content, *outArg)
	default:
		fmt.Println("encrypt or decrypt subcommand is required")
		os.Exit(1)
	}
}
