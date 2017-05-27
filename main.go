package main

import "net"
import "fmt"
import "runtime"

func pixel(x int, y int, r int, g int, b int) string {
	return (fmt.Sprintf("PX %d %d %02x%02x%02x\n", x, y, r, g, b))
}

func pixel_client(server string, x int, y int, r int, g int, b int, stride int) {
	conn, err := net.Dial("tcp",)
	if err != nil {
		fmt.Printf("pixel server %d %d not created\n", x, y)
		return
	}

	for {
		for px := x; px < x+stride; px++ {
			for py := y; py < y+stride; py++ {
				bytes := []byte(pixel(x, y, r, g, b))
				conn.Write(bytes)
			}
		}
	}
}

func main() {
        runtime.GOMAXPROCS(16)
	server := "94.45.231.39:1234"
	fmt.Printf("PX %d %d %02x%02x%02x\n", 500, 500, 10, 20, 30)
	
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			go pixel_client(server, 500+5*i, 500+5*j, 0, 0, 0, 5)
		}
	}
	for {
	}
}
