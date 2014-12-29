package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/ugorji/go/codec"
	msgpack "gopkg.in/vmihailenco/msgpack.v1"
)

type MsgPackExample struct {
	mh   *codec.MsgpackHandle
	buf1 *bytes.Buffer
	buf2 *bytes.Buffer
	enc  *codec.Encoder
	dec  *codec.Decoder
	menc *msgpack.Encoder
	mdec *msgpack.Decoder
}

//go:generate msgp
type MsgPackSample struct {
	Name    string
	Num     int
	Message string
}

func main() {

	var m = MsgPackSample{
		Name:    "ユイにゃん",
		Num:     17,
		Message: "AngelBeats!",
	}

	mh := &codec.MsgpackHandle{RawToString: true}
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	var msgpackExample = MsgPackExample{
		mh:   mh,
		buf1: buf1,
		buf2: buf2,
		enc:  codec.NewEncoder(buf1, mh),
		dec:  codec.NewDecoder(buf2, mh),
		menc: msgpack.NewEncoder(buf1),
		mdec: msgpack.NewDecoder(buf2),
	}

	msgpackExample.packUnPackMsgPack(m)

	msgpackExample.packUnPackCodec(m)
}

func (e MsgPackExample) packUnPackCodec(m MsgPackSample) {

	e.buf1.Reset()

	if err := e.enc.Encode(&m); err != nil {
		log.Fatalln(err)
	}

	e.buf2.Reset()
	e.buf2.Write(e.buf1.Bytes())

	var m2 MsgPackSample
	if err := e.dec.Decode(&m2); err != nil {
		log.Fatalln(err)
	}

	pp.Println(m2)
}

func (e MsgPackExample) packUnPackMsgPack(m MsgPackSample) {

	e.buf1.Reset()

	if err := e.menc.Encode(m); err != nil {
		log.Fatalln(err)
	}

	e.buf2.Reset()
	e.buf2.Write(e.buf1.Bytes())

	var m2 MsgPackSample
	if err := e.mdec.Decode(&m2); err != nil {
		log.Fatalln(err)
	}

	pp.Println(m2)
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
