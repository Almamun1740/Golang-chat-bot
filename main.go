package main

import (
	"fmt"
	"net"
	"os"
) 

func main() {

	//net = network
	//ring ring ring
	nl, err := net.Listen("tcp", "192.168.31.33:8888") //ip: port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1 ==0
	}

	//call receive == green button press
	conn, err := nl.Accept() //Layer-5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	remoteAddr := conn.RemoteAddr().String() //caller
	fmt.Println(remoteAddr) 				//[::1]:52451
	//localhost

	//byte
	/*8 bit = 1 byte
	1024 byte = 1 kilobyte
	map, slice*/
	bs := make([]byte, 1024) //1500 bytes
	conn.Read(bs) 			//listen

	//90 10 20 40 50 70 90
	// H e l l o
	fmt.Println(string(bs)) //human readable formated message

	//byte
	conn.Write([]byte("welcome we have received your message")) //speak

	conn.Close() // green button press
	nl.Close()	//red button press

}