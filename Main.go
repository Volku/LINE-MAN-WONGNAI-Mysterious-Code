package main

import "fmt"
import b64 "encoding/base64"
import "strings"

func reverse(words []string) []string {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return words
}

func main() {
    secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := b64.StdEncoding.DecodeString(secret)
	decodeString := string(sd)
	fmt.Println(decodeString)
	array := strings.Split(decodeString,"")
	array = reverse(array)
	decodeString = strings.ReplaceAll(strings.Join(array[:],""),":"," ")
	fmt.Println(decodeString)
}