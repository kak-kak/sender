package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func GetHeader() []byte {
	return []byte{0xBB}
}

func GetFooter() []byte {
	return []byte{0xAA}
}

func main() {
	c := &serial.Config{Name: "/dev/tty98", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		go newFunction1(s)

		go newFunction(s)

		time.Sleep(1 * time.Second)
	}
}

func newFunction1(s *serial.Port) {
	now := time.Now().Format(time.ANSIC)
	buf := GetHeader()
	buf = append(buf, now...)
	buf = append(buf, GetFooter()...)

	fmt.Printf("%q\n", buf)
	_, err := s.Write(buf)

	if err != nil {
		log.Fatal(err)
	}
}

func newFunction(s *serial.Port) {
	readbuf := make([]byte, 1000)
	n, err := s.Read(readbuf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("recieved %o", readbuf[:n])
}
