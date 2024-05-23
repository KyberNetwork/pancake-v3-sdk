package entities

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/daoleno/uniswap-sdk-core/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Tick) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.Index, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "LiquidityGross")
			return
		}
		z.LiquidityGross = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.LiquidityGross))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityGross")
				return
			}
			z.LiquidityGross = msgpencode.DecodeInt(zb0002)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "LiquidityNet")
			return
		}
		z.LiquidityNet = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.LiquidityNet))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityNet")
				return
			}
			z.LiquidityNet = msgpencode.DecodeInt(zb0003)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Tick) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Index)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	if z.LiquidityGross == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.LiquidityGross))
		if err != nil {
			err = msgp.WrapError(err, "LiquidityGross")
			return
		}
	}
	if z.LiquidityNet == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.LiquidityNet))
		if err != nil {
			err = msgp.WrapError(err, "LiquidityNet")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tick) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o = msgp.AppendInt(o, z.Index)
	if z.LiquidityGross == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.LiquidityGross))
	}
	if z.LiquidityNet == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.LiquidityNet))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tick) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.Index, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.LiquidityGross = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.LiquidityGross))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityGross")
				return
			}
			z.LiquidityGross = msgpencode.DecodeInt(zb0002)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.LiquidityNet = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.LiquidityNet))
			if err != nil {
				err = msgp.WrapError(err, "LiquidityNet")
				return
			}
			z.LiquidityNet = msgpencode.DecodeInt(zb0003)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tick) Msgsize() (s int) {
	s = 1 + msgp.IntSize
	if z.LiquidityGross == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.LiquidityGross))
	}
	if z.LiquidityNet == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.LiquidityNet))
	}
	return
}
