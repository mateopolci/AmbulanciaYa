package utils

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "log"
)

func GenerateSecret() {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(b))
}
