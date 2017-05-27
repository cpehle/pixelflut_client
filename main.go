package main

import "net"
import "fmt"
import "flag"
import "bytes"
import "runtime"

var ip = flag.String("ip", "127.0.0.1:8080", "The IP and port of the server to pixel flood")
var x = flag.Int("x", 0, "x coordinate of rectangle to send")
var y = flag.Int("y", 0, "x coordinate of rectangle to send")
var width = flag.Int("w", 400, "width of rectangle to send")
var height = flag.Int("h", 400, "height of rectangle to send")


func pixel(x int, y int, r int, g int, b int) string {
	return (fmt.Sprintf("PX %d %d %02x%02x%02x\n", x, y, r, g, b))
}

func rectangle(x int, y int, width int, height int, r int, g int, b int) bytes.Buffer {
     var buffer bytes.Buffer
     for i := x; i < x + width; i++ {
     	 for j := y; j < y + height; j++ {
		 buffer.WriteString(pixel(i,j,r,g,b))
	 }
     }
     return buffer
}

func pixel_client(server string, x int, y int, r int, g int, b int, stride int,  c chan int) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("pixel server %d %d not created\n", x, y)
		return
	}
	
	buffer := rectangle(x, y, stride, stride, r, g, b)
	bytes := buffer.Bytes()
	
	for {
	    conn.Write(bytes)
	}
}

func main() {
	runtime.GOMAXPROCS(16)
//	server := *ip //"127.0.0.1:8080"
	x_base := *x
	y_base := *y
	
        server := "94.45.231.39:1234"
   	// server := "94.45.234.31:1234"
	flag.Parse()

	c := make(chan int)
	
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			stride := 50
//						go pixel_client(server, x_base+stride*i, y_base+stride*j, 0,0,0, stride, c)
		 	go pixel_client(server, x_base+stride*i, y_base+stride*j, 255,105,180, stride, c)
		}
	}

	x := <- c
	fmt.Println(x)
}
