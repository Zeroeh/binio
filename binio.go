package binio

import (
	"math"
	"encoding/binary"
)

//Packet is a struct containing an underlying bytes buffer
type Packet struct {
	Index  int
	Length int
	Data   []byte
}

//Advance the buffer index while returning the amount increased
func (p *Packet) advance(amount int) int {
	p.Index += amount
	return amount
}

//ResizeBuffer resizes our packet to include any extra data portion
//Feel free to customize this function, as it can be tailored to many individual needs
func (p *Packet) resizeBuffer(size, keep, index int) {
	if index > len(p.Data) { //could also substitute p.Length
		return //out of bounds
	}
	tmp := p.Data
	p.Data = make([]byte, size)
	//keep the specified amount of bytes in the new buffer, starting from index
	for index < keep {
		p.Data[index] = tmp[index]
		index++
	}
}

//ReadString reads the expected string size (n) and read until n
func (p *Packet)readString() string {
	n := int(p.readUInt16()) //absolute
	if n == 0 {
		return ""
	}
	var str []byte
	str = p.Data[p.Index:p.Index+n]
	p.advance(n)
	return string(str)
}

//WriteString writes int16 (len of string) and then the contents of s as bytes
func (p *Packet)writeString(s string) {
	if s == "" {
		p.writeUInt16(uint16(0))
		return
	}
	p.writeUInt16(uint16(len(s)))
	for i := range s {
		p.writeByte(s[i])
	}
}

func (p *Packet)readUTFString() string {
	n := int(p.readUInt32())
	if n == 0 {
		return ""
	}
	var str []byte
	str = p.Data[p.Index:p.Index+n]
	p.advance(n)
	return string(str)
}

func (p *Packet)writeUTFString(s string) {
	if s == "" {
		p.writeUInt32(0)
		return
	}
	p.writeUInt32(uint32(len(s)))
	for i := range s {
		p.writeByte(s[i])
	}
}

func (p *Packet)readBool() bool {
	if p.readByte() == 1 {
		return true
	}
	return false //assume anything else is false
}

func (p *Packet)writeBool(b bool) {
	if b == true {
		p.writeByte(1)
	} else {
		p.writeByte(0)
	}
}

//ReadFloat reads 4 bytes representing a float
func (p *Packet)readFloat() float32 {
	return math.Float32frombits(p.readUInt32())
}

//WriteFloat writes 4 bytes representing a float
func (p *Packet)writeFloat(f float32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.advance(4)], math.Float32bits(f))
}

func (p *Packet)readInt16() int16 {
	return int16(binary.BigEndian.Uint16(p.Data[p.Index:p.Index+p.advance(2)]))
}

func (p *Packet)writeInt16(i int16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.advance(2)], uint16(i))
}

func (p *Packet)readUInt16() uint16 {
	return binary.BigEndian.Uint16(p.Data[p.Index:p.Index+p.advance(2)])
}

func (p *Packet)writeUInt16(i uint16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.advance(2)], i)
}

func (p *Packet)readInt32() int32 {
	return int32(binary.BigEndian.Uint32(p.Data[p.Index:p.Index+p.advance(4)]))
}

func (p *Packet)writeInt32(i int32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.advance(4)], uint32(i))
}

func (p *Packet)readUInt32() uint32 {
	return binary.BigEndian.Uint32(p.Data[p.Index:p.Index+p.advance(4)])
}

func (p *Packet)writeUInt32(i uint32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.advance(4)], i)
}

//ReadByte reads and returns a singular byte
func (p *Packet)readByte() byte {
	return p.Data[p.Index:p.Index+p.advance(1)][0]
}

//WriteByte writes a singular byte to the packet buffer
func (p *Packet)writeByte(d byte) {
	p.Data[p.Index] = d
	p.advance(1)
}

//ReadBytes is experimental and has not been tested
func (p *Packet)readBytes(amount int) []byte {
	return p.Data[p.Index:p.Index+p.advance(amount)]
}


