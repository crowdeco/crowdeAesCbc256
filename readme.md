Crowde Aes-Cbc-256 golang SDK

## How to install
```$ go get github.com/crowdeco/crowdeAesCbc256```

## How to use
```
   package main
   
   import (
   	"github.com/crowdeco/crowdeAesCbc256"
   	"fmt"
   )
   
   func main() {
   	t := new(crowdeAesCbc256.AesEncrypt)

        // initial credentials: initial vector & secret key
   	t.Init("ijzh84t1w9xa56s9", "4bd393e7a457f9023d9ba95fffb5a2e1")

        // encrypt payload
   	encrypted := t.AESEncrypt("tes_enkripsi")
   	fmt.Printf("encrypt: %s \n", encrypted)
   
        // decrypt payload
   	decrypt := t.AESDecrypt(encrypted)
   	fmt.Printf("decrypt: %s \n", decrypt)
   }
```

### TODO
##### - Unit test
