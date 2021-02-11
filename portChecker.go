package portChecker

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var port int = 80

func PortChecker() (p int, e error) {

	fmt.Println(port)
	port += 1
	if port >= 9999 {
		panic(errors.New("No Ip ..."))
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic")
			PortChecker()
		}
	}()

	fmt.Println("Port Checking Start ...")
	var sb strings.Builder
	for true {
		sb.WriteString(":")
		sb.WriteString(strconv.FormatInt(int64(port), 10))
		fmt.Println(sb.String())
		newSocket, err := net.Listen("tcp", sb.String())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			return port, nil
		}
		newSocket.Close()
		sb.Reset()
	}

	err := errors.New("All port disable ...")
	return -1, err
}
