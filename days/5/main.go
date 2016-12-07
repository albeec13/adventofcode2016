package main

import (
    "crypto/md5"
    "fmt"
    "strconv"
    "strings"
    "encoding/hex"
)

func main() {
    input := []byte("ffykfhsq")
    password := ""

    for i:=0;; i++ {
        data := make([]byte,0)
        data = append(input, []byte(strconv.Itoa(i))...)
        //fmt.Println(string(data))

        md5sum := md5.Sum(data)
        md5sumStr := hex.EncodeToString(md5sum[:])
        if(strings.HasPrefix(md5sumStr,"00000")) {
            fmt.Println(md5sumStr)
            password += string(md5sumStr[5])

            if len(password) == 8 { break }
        }
    }
    fmt.Println("Password: ", password)
}
