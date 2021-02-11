## PortChecker Go Library

# Using

func PortChecker() (p int, e error) 


Example   
```go
package portChecker

import (
	"fmt"
	"testing"
)

func TestPortChecker(t *testing.T) {
	port, err := PortChecker()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(port)
	}
}
```