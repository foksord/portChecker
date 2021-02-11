package portChecker

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestPortChecker(t *testing.T) {
	port, err := PortChecker()
	if err != nil {
		t.Error(err.Error())
	} else {
		var sb strings.Builder
		sb.WriteString(":")
		sb.WriteString(strconv.FormatInt(int64(port), 10))
		fmt.Println(sb.String())
		newSocket, err := net.Dial("tcp", sb.String())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Port Oki :D")
			newSocket.Close()
			sb.Reset()
		}
	}
}
