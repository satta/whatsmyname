package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func mmd5(w string) string {
	h := md5.New()
	h.Write([]byte(w))
	return hex.EncodeToString(h.Sum(nil))
}

func msha1(w string) string {
	h := sha1.New()
	h.Write([]byte(w))
	return hex.EncodeToString(h.Sum(nil))
}

func hostnameToRange(word string, max int64) int64 {
	bi := big.NewInt(0)
	hexstr := mmd5(word)
	bi.SetString(hexstr, 16)
	out := big.NewInt(0)
	return out.Mod(bi, big.NewInt(max)).Int64()
}

func main() {
	data, err := Asset("names.txt")
	if err != nil {
		log.Fatal(err)
	}

	words := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		word := scanner.Text()
		if string([]rune(word)[0]) != "#" {
			words = append(words, word)
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	v := hostnameToRange(hostname, int64(len(words)))
	w := hostnameToRange(msha1(hostname), int64(len(words)))
	fmt.Printf("%s%s\n", words[v], words[w])
}
