package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/msgp"
)


// MarshalMsg implements the msgp.Marshaler interface
func (z *MsgPackSample) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 3)

	o = msgp.AppendString(o, "Name")

	o = msgp.AppendString(o, z.Name)

	o = msgp.AppendString(o, "Num")

	o = msgp.AppendInt(o, z.Num)

	o = msgp.AppendString(o, "Message")

	o = msgp.AppendString(o, z.Message)

	return
}

// UnmarshalMsg unmarshals a MsgPackSample from MessagePack, returning any extra bytes
// and any errors encountered
func (z *MsgPackSample) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Name":

			z.Name, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Num":

			z.Num, bts, err = msgp.ReadIntBytes(bts)

			if err != nil {
				return
			}

		case "Message":

			z.Message, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Msgsize implements the msgp.Sizer interface
func (z *MsgPackSample) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 4

	s += msgp.StringPrefixSize + len(z.Name)

	s += msgp.StringPrefixSize + 3

	s += msgp.IntSize
	s += msgp.StringPrefixSize + 7

	s += msgp.StringPrefixSize + len(z.Message)

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *MsgPackSample) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Name":

			z.Name, err = dc.ReadString()

			if err != nil {
				return
			}

		case "Num":

			z.Num, err = dc.ReadInt()

			if err != nil {
				return
			}

		case "Message":

			z.Message, err = dc.ReadString()

			if err != nil {
				return
			}

		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}

	return
}

// EncodeMsg implements the msgp.Encodable interface
func (z *MsgPackSample) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}

	err = en.WriteString("Name")
	if err != nil {
		return
	}

	err = en.WriteString(z.Name)

	if err != nil {
		return
	}

	err = en.WriteString("Num")
	if err != nil {
		return
	}

	err = en.WriteInt(z.Num)

	if err != nil {
		return
	}

	err = en.WriteString("Message")
	if err != nil {
		return
	}

	err = en.WriteString(z.Message)

	if err != nil {
		return
	}

	return
}
