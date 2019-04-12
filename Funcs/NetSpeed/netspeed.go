package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const sizeToTest = 200000
const bufSize = 1000

func main() {
	listen := flag.String("listen", "", "")
	connect := flag.String("connect", "", "")
	flag.Parse()
	if len(*listen) > 0 {
		doListen(*listen)
	} else if len(*connect) > 0 {
		doConnect(*connect)
	} else {
		print("Nethier listen nor connect not specified")
	}
}
func doConnect(addr string) {
	conn, err := net.Dial("tcp", addr)
	Must(err, "net.Dial")

	conn.Close()
}
func formatSpeed(op string, d time.Duration, size int) {
	kBytes := float64(size) / float64(8) / float64(1000)
	str := fmt.Sprintf("%s Duration: %s, Speed : %f kb/s", op, d, kBytes/d.Seconds())
	fmt.Println(str)
}
func writeData(conn net.Conn) {
	writtenBytes := 0
	data := randBytes(bufSize)
	for writtenBytes < sizeToTest {
		toWriteSize := min(bufSize, sizeToTest-writtenBytes)
		n, err := conn.Write(data[:toWriteSize])
		Must(err, "conn write")
		if n < toWriteSize {
			log.Panicf("Expected to write %d, but written %d", toWriteSize, n)
		}
		writtenBytes += toWriteSize
	}

}
func readData(conn net.Conn) {
	readBytes := 0
	data := make([]byte, bufSize)
	for readBytes < sizeToTest {
		toreadsize := min(bufSize, sizeToTest-readBytes)
		n, err := conn.Read(data)
		Must(err, "conn read")
		if n < toreadsize {
			log.Panicf("Expected to read %d,but read %d", toreadsize, n)
		}
		readBytes += toreadsize
	}
}
func randBytes(len int) []byte {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func stopWatch(f func()) time.Duration {
	s := time.Now()
	f()
	return time.Now().Sub(s)
}
func doListen(addr string) {
	l, err := net.Listen("tcp", addr)
	Must(err, "net Listen")
	for {
		conn, err := l.Accept()
		Must(err, "L accept")
		formatSpeed("Server read", stopWatch(func() { readData(conn) }), sizeToTest)
		formatSpeed("Server write", stopWatch(func() { writeData(conn) }), sizeToTest)
		conn.Close()
	}
	l.Close()
}
func Must(err error, msg ...string) {
	if err != nil {
		log.Panic(err, msg)
	}
}
