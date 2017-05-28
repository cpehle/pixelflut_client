package main

import "net"
import "fmt"
import "flag"
import "bytes"
import "runtime"
import "os"
import "image"
import _ "image/png"

var ip = flag.String("ip", "127.0.0.1:8080", "The IP and port of the server to pixel flood")
var x = flag.Int("x", 0, "x coordinate of rectangle to send")
var y = flag.Int("y", 0, "x coordinate of rectangle to send")
var width = flag.Int("w", 400, "width of rectangle to send")
var r = flag.Int("r", 0, "red")
var g = flag.Int("g", 0, "green")
var b = flag.Int("b", 0, "blue")
var height = flag.Int("h", 400, "height of rectangle to send")


func pixel(x int, y int, r uint32, g uint32, b uint32) string {
	return (fmt.Sprintf("PX %d %d %02x%02x%02x\n", x, y, r, g, b))
}

func image_to_pixel(img image.Image, x_offset int, y_offset int, x int, y int, width int, height int) bytes.Buffer {
     var buffer bytes.Buffer
     for i := x; i < x + width; i++ {
     	 for j := y; j < y + height; j++ {
	       	 r,g,b,_ := img.At(i,j).RGBA()
		 buffer.WriteString(pixel(x_offset + i, y_offset + j,r,g,b))
	 }
     }
     return buffer
}

func rectangle(x int, y int, width int, height int, r uint32, g uint32, b uint32) bytes.Buffer {
     var buffer bytes.Buffer
     for i := x; i < x + width; i++ {
     	 for j := y; j < y + height; j++ {
		 buffer.WriteString(pixel(i,j,r,g,b))
	 }
     }
     return buffer
}

func image_client(server string, img image.Image, x_offset int, y_offset int, x_index int, y_index int, width int, height int, c chan int) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("pixel server %d %d not created\n", x, y)
		return
	}
	
	buffer := image_to_pixel(img, x_offset, y_offset, width*x_index, height*y_index, width, height)
	bytes := buffer.Bytes()
	
	for {
	    conn.Write(bytes)
	}
}


func pixel_client(server string, x int, y int, r uint32, g uint32, b uint32, stride int,  c chan int) {
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
//	x_base := *x
//	y_base := *y	
	infile, _ := os.Open("cat.png")
	defer infile.Close()
	img,_,_ := image.Decode(infile)
	bounds := img.Bounds()
	fmt.Println(bounds)
	
        // server := "94.45.231.39:1234"
   	server := "94.45.234.31:1234"
	flag.Parse()
	c := make(chan int)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
//			stride := 50
//			go pixel_client(server, x_base+stride*i, y_base+stride*j, 0,0,0, stride, c)
//		 	go pixel_client(server, x_base+stride*i, y_base+stride*j, *r,*g,*b, stride, c)
			go image_client(server, img, 200, 200, i, j, 10, 10, c)
		}
	}
	
	x := <- c
	fmt.Println(x)
}
