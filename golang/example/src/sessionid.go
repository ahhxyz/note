package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	sessionId := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
	fmt.Println(sessionId)
}

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}
