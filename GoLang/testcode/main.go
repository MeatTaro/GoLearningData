package main

import (
	"time"
	"io"
	"bytes"
	"net"
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage : host port %s",os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	coon, err := net.DialTimeout("tcp", service, 3*time.Second)
	checkerr(err)

	_, err = coon.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	err = coon.SetDeadline(time.Now().Add(5 * time.Second))
	checkerr(err)

	result, err := readfully(coon)
	checkerr(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkerr(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr,"fatol error is %s",err.Error())
		os.Exit(1)
	}
}

func readfully(coon net.Conn) ([]byte, error) {
	defer coon.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := coon.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}	
	}
	return result.Bytes(), nil
}