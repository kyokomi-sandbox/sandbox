package main

import (
	msgpack "gopkg.in/vmihailenco/msgpack.v1"
	"log"
	"os"
	"io"
	"bytes"
	"github.com/ugorji/go/codec"
)

var (
	mh    = &codec.MsgpackHandle{RawToString: true}
	buf   = &bytes.Buffer{}
	bufR  = &bytes.Buffer{}
	me    = msgpack.NewEncoder(buf)
	md    = msgpack.NewDecoder(bufR)
)

var enc   = codec.NewEncoder(buf, mh)
var dec   = codec.NewDecoder(bufR, mh)

//go:generate msgp
type MsgPackSample struct {
	Name string
	Num int
	Message string
}

func msgPackExample() {

	var m = MsgPackSample{
		Name: "ユイにゃん",
		Num: 17,
		Message: "AngelBeats!",
	}

	packUnPackMsgPack(m)

	packUnPackCodec(m)
}

func packUnPackCodec(m MsgPackSample) {

	buf.Reset()

	if err := enc.Encode(&m); err != nil {
		log.Fatalln(err)
	}

//	fmt.Printf("size: [%d] data: %#v\n", buf.Len(), buf.Bytes())
	bufR.Reset()
	bufR.Write(buf.Bytes())

	var m2 MsgPackSample
	if err := dec.Decode(&m2); err != nil {
		log.Fatalln(err)
	}

//	pp.Println(m2)
}

func packUnPackMsgPack(m MsgPackSample) {

	buf.Reset()

	if err := me.Encode(m); err != nil {
		log.Fatalln(err)
	}

//	fmt.Printf("size: [%d] data: %#v\n", len(b), b)
	bufR.Reset()
	bufR.Write(buf.Bytes())

	var m2 MsgPackSample
	if err := md.Decode(&m2); err != nil {
		log.Fatalln(err)
	}

//	pp.Println(m2)
}

func writeFile(fileName string, b []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, bytes.NewReader(b)); err != nil {
		return err
	}

	return nil
}
