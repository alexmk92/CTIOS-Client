package main

type Body interface {}

type Request struct {
	Header MessageHeader
	Body Body
}

// See page 224 for tag implementation
type NamedVar struct {
	Tag uint8
	FieldLength uint8
	VariableName string
}

type MessageHeader struct {
	MessageLength uint32
	MessageType uint32
}
