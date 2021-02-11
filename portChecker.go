package portChecker

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// PortChecker
func PortChecker() (int64, error) {

	var port int64 = 8080
	for ; port <= 9999; port++ {
		fmt.Println(strconv.FormatInt(int64(port), 10))
		port, err := portCheck(port)
		if err == nil {
			return port, nil
		}
	}
	return -1, errors.New("port error")
}

func portCheck(port int64) (int64, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic")
		}
	}()

	var sb strings.Builder
	sb.WriteString(":")
	sb.WriteString(strconv.FormatInt(int64(port), 10))

	newSocket, err := net.Listen("tcp", sb.String())

	if err == nil && newSocket != nil {
		return port, nil
		newSocket.Close()
	}

	return -1, err
}
