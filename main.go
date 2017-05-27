package main

import "net"
import "fmt"
import "runtime"

func pixel(x int, y int, r int, g int, b int) string {
	return (fmt.Sprintf("PX %d %d %02x%02x%02x\n", x, y, r, g, b))
}

func rectangle(x int, y int, width int, height int, r int, g int, b int) []byte {
     var buffer []byte
     for i := x; i < x + width; i++ {
     	 for j := y; j < y + height; j++ {
	     buffer = append(buffer, []byte(pixel(i, j, r, g, b)))
	 }
     }
}

func pixel_client(server string, x int, y int, r int, g int, b int, stride int) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("pixel server %d %d not created\n", x, y)
		return
	}
	
	bytes := rectangle(x, y, stride, stride, r, g, b)

	for {
	    conn.Write(bytes)
	}
}

func main() {
	runtime.GOMAXPROCS(16)
	server := "127.0.0.1:8080"
// server := "94.45.231.39:1234"
// server := "94.45.234.31:1234"
	
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			stride := 50
			go pixel_client(server, 500+stride*i, 500+stride*j, 255,105,180, stride)
		}
	}
	for {
	}
}
