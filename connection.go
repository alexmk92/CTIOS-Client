package main

import (
	"net"
	"fmt"
	"strconv"
	//"bufio"
)

type Connection struct {
	host string
	port uint16
	invokedID uint32
}

func (c *Connection) connect() {
	fmt.Println("Connecting")
	conn, err := net.Dial("tcp", (c.host + ":" + strconv.Itoa(int(c.port))))
	if err != nil {
		fmt.Println("Error connecting to: " + (c.host + ":" + strconv.Itoa(int(c.port))))
	}
	c.handleConnection(conn)
}

func (c *Connection) handleConnection(conn net.Conn) {
	go c.listen(conn)

	fmt.Println("Sending OPEN-REQ message")
	req := Request {
		Header: MessageHeader{
			66,
			3,
		},
		Body: OpenReq {
			// Fixed = 44 bytes
			InvokeID: 0,
			VersionNumber: 16,
			IdleTimeout: 30,
			PeripheralID: 5000,
			ServicesRequested: 0x000010,
			CallMsgMask: 0x000000,
			AgentStateMask: 0x000000,
			ConfigMsgMask: 0x000000,
			Reserved1: 0x000000,
			Reserved2: 0x000000,
			Reserved3: 0x000000,
			// Float = NamedVar A = 12 bytes NamedVar B = 12 bytes (56 bytes)
			ClientID_Tag: NamedVar {
				Tag: 1,
				FieldLength: 9,
				VariableName: "alexTest\x00",
			},
			ClientPassword_Tag: NamedVar {
				Tag: 2,
				FieldLength: 9,
				VariableName: "alexTest\x00",
			},
		},
	}

	b := req.serialize()
	conn.Write(b)
}

func (c *Connection) listen(conn net.Conn) {
	fmt.Println("Listening")
	for {
		var buf = make([]byte, 4329)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		} else if n > 0 {
			fmt.Println("Read " + strconv.Itoa(int(n)) + " bytes from the server")
			fmt.Println(buf[0:n])
		}
	}
}