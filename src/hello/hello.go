/*

package main

import "fmt"

func main() {
	fmt.Println("hello World")
}

*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// MIN min value of rand
const MIN = 1

// MAX max value of rand
const MAX = 100

var shutdown = false

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		if shutdown {
			break
		}
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		} else if temp == "SHUTDOWN" {
			shutdown = true
			break
		} else {
			fmt.Println("RCV:" + temp)
		}

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(result)))
		fmt.Println("SND:" + result)
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		if shutdown {
			l.Close()
			break
		}

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
