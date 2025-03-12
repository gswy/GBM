package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

var (
	host string
	port int
)

func init() {
	flag.IntVar(&port, "port", 5060, "Please enter port numbers: 1~65535")
	flag.StringVar(&host, "host", "127.0.0.1", "Please enter an IPv4 address")
	flag.Parse()
	if !isIPv4(host) {
		log.Fatalf("Host: %s is not a valid IPv4 address", host)
	}
	if port < 1 || port > 65535 {
		log.Fatalf("Port: number must be between 1 and 65535")
	}
}

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(host), Port: port})
	if err != nil {
		log.Fatalf("UDP connection failed: %v", err)
	}
	log.Printf("Ready for sending to %s:%d\n", host, port)
	go shutdown(conn)
	num := 1
	for {
		for _, frame := range bytes {
			time.Sleep(time.Millisecond * 50)
			_, _ = conn.Write(frame)
		}
		log.Printf("Sent complete data in a loop: %d\n", num)
		num++
	}
}

// shutdown 优雅停机
func shutdown(conn *net.UDPConn) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	err := conn.Close()
	if err != nil {
		log.Fatalf("Failed to close UDP connection: %v", err)
	} else {
		log.Printf("UDP Connection closed")
	}
	os.Exit(0)
}

func isIPv4(ipv4 string) bool {
	// 正则表达式匹配 IPv4 地址
	pattern := `^(25[0-5]|2[0-4][0-9]|1?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|1?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|1?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|1?[0-9][0-9]?)$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(ipv4)
}
