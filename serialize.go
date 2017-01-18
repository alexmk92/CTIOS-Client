package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"reflect"
)

// uses reflection to
func (r *Request) serialize() []byte {

	buf := new(bytes.Buffer)

	serializeStruct(buf, reflect.ValueOf(r.Header).Interface());
	serializeStruct(buf, reflect.ValueOf(r.Body).Interface());

	fmt.Println("Serialized the buffer to:")
	fmt.Println(buf.Bytes())

	return buf.Bytes()
}

// Rework this in future to recursively serialize struct
func serializeStruct(b *bytes.Buffer, s interface{}) {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			for i := 0; i < field.NumField(); i++ {
				k := field.Field(i).Kind()
				val := field.Field(i).Interface()

				writeToBuffer(b, k, val)
			}
		} else {
			k := v.Field(i).Kind()
			val := v.Field(i).Interface()

			writeToBuffer(b, k, val)
		}
	}
}

// Byte order to host is Big Endian
func writeToBuffer(b *bytes.Buffer, k reflect.Kind, val interface{}) {

	switch k {
	case reflect.String:
		out := []byte(val.(string))
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Uint:
		out := val.(uint)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Uint8:
		out := val.(uint8)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Uint16:
		out := val.(uint16)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Uint32:
		out := val.(uint32)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Uint64:
		out := val.(uint64)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Int:
		out := val.(int)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Int8:
		out := val.(int8)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Int16:
		out := val.(int16)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Int32:
		out := val.(int32)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	case reflect.Int64:
		out := val.(int64)
		err := binary.Write(b, binary.BigEndian, out)
		checkErr(err)
	default:
		fmt.Println("Uncaught type: ", k)
	}
}

func checkErr(err error) {
	if err != nil {
		panic("Fatal error, binary write failed: ")
	}
}