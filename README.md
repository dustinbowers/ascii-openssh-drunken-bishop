# Drunken Bishop Fingerprint Generator
Create an OpenSSH ASCII fingerprint from a Hex string

## Usage
```
import "fmt"
import "github.com/dustinbowers/ascii-openssh-drunken-bishop/drunkenbishop"

func main() {

  fingerprint := "a1d0c6e83f027327d8461063f4ac58a6"
  
  db := drunkenbishop.NewDrunkenBishop()
  db.SetTopLabel("MD5")
  db.SetBottomLabel("Test Hash")
  ascii, _ := db.ToAscii(fingerprint)
  
  fmt.Print(ascii)
}
```
will produce this output:
```
+------[MD5]------+
| .*.             |
| . = +           |
|  o * + .        |
| = * o . .       |
|E = * o S        |
|   = +           |
|    . o          |
|     . .         |
|                 |
+---[Test Hash]---+
```

## More info
- http://www.dirk-loss.de/sshvis/drunken_bishop.pdf
- https://www.jfurness.uk/the-drunken-bishop-algorithm/
- https://dev.to/krofdrakula/improving-security-by-drawing-identicons-for-ssh-keys-24mc
