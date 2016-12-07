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
    password := []byte("        ")
    cnt := 0

    for i:=0;; i++ {
        data := make([]byte,0)
        data = append(input, []byte(strconv.Itoa(i))...)
        //fmt.Println(string(data))

        md5sum := md5.Sum(data)
        md5sumStr := hex.EncodeToString(md5sum[:])
        if(strings.HasPrefix(md5sumStr,"00000")) {
            if pos, err := strconv.Atoi(string(md5sumStr[5])); err == nil {
                if pos < 8 {
                    if password[pos] == ' ' {
                        password[pos] = md5sumStr[6]
                        fmt.Println(md5sumStr)
                        cnt++
                        if cnt == 8 { break }
                    }
                }
            }
        }
    }
    fmt.Println("Password: ", string(password))
}
