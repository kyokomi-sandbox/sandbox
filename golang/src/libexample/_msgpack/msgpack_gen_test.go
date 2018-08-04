package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"bytes"
	"testing"

	"github.com/philhofer/msgp/msgp"
)

func TestMsgPackSampleMarshalUnmarshal(t *testing.T) {
	v := new(MsgPackSample)
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func BenchmarkMsgPackSampleMarshalMs(b *testing.B) {
	v := new(MsgPackSample)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkMsgPackSampleAppendMsg(b *testing.B) {
	v := new(MsgPackSample)
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkMsgPackSampleUnmarshal(b *testing.B) {
	v := new(MsgPackSample)
	bts, _ := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMsgPackSampleEncodeDecode(t *testing.T) {
	v := new(MsgPackSample)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Maxsize() for %v is inaccurate", v)
	}

	vn := new(MsgPackSample)
	err := msgp.Decode(&buf, vn)
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkMsgPackSampleEncode(b *testing.B) {
	v := new(MsgPackSample)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	b.SetBytes(int64(buf.Len()))
	en := msgp.NewWriter(msgp.Nowhere)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.EncodeMsg(en)
	}
	en.Flush()
}

func BenchmarkMsgPackSampleDecode(b *testing.B) {
	v := new(MsgPackSample)
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	b.SetBytes(int64(buf.Len()))
	rd := msgp.NewEndlessReader(buf.Bytes())
	dc := msgp.NewReader(rd)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := v.DecodeMsg(dc)
		if err != nil {
			b.Fatal(err)
		}
	}
}
