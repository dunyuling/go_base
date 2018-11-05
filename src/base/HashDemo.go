package main

import (
    "crypto/sha512"
    "encoding/hex"
    "log"
)

func calculateHash(toBeHashed string) string {
    hashInBytes := sha512.Sum512([]byte(toBeHashed)) 
    log.Printf("%v,%d",hashInBytes,len(hashInBytes)) 
    hashInString := hex.EncodeToString(hashInBytes[:]) 
    log.Printf("%s,%s,%d", toBeHashed,hashInString,len(hashInString))
    log.Println()
    return hashInString
}

func main() {
    calculateHash("test1")
    calculateHash("test1")
    calculateHash("test2")
}


