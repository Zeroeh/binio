package main

import (
	"fmt"
	"github.com/Zeroeh/binio"
)

//Example struct for demonstration of the "binio" library
type Player struct {
	PlayerID uint32
	PlayerName string
	PlayerLocation PlayerPosition
	PlayerStats []int16
}

type PlayerPosition struct {
	X float32
	Y float32
	Z float32
}

func main() {
	//Create a new player struct
	newPlayer := Player{}
	newPlayer.PlayerID = 1337
	newPlayer.PlayerName = "Zeroeh"
	newPlayer.PlayerLocation.X = 103.71
	newPlayer.PlayerLocation.Y = 2123.71
	newPlayer.PlayerLocation.Z = 14.71
	newPlayer.PlayerStats = make([]int16, 4)
	for i := 0; i < len(newPlayer.PlayerStats); i++ {
		newPlayer.PlayerStats[i] = int16(i)
	}
	fmt.Println("Player:", newPlayer)

	//Now that we have our struct, we can start dumping it into bytes
	p := new(binio.Packet)

	//need to allocate p.Data before we can write anything to it
	//it will be left to you to decide how much will need to be allocated per buffer
	p.Data = make([]byte, 16+len(newPlayer.PlayerName)+len(newPlayer.PlayerStats)*4)
	newPlayer.WriteToBytes(p)
	fmt.Println("Player struct as bytes:", p.Data)
	
	/*
		at this point you can send your data over the wire
		i.e. net.Conn.Write(p.Data)
	*/

	//now we can read our bytes back to a struct
	p.Index = 0 //we're reading from the same packet so reset the index
	//if this were a remote app/server we'd make a new packet here
	fmt.Println("Player re-read to a struct:", ReadPlayerToStruct(p))

}


//Example of writing struct to bytes buffer
func (p *Player)WriteToBytes(x *binio.Packet) {
	x.WriteUInt32(p.PlayerID)
	x.WriteString(p.PlayerName)
	x.WriteFloat(p.PlayerLocation.X)
	x.WriteFloat(p.PlayerLocation.Y)
	x.WriteFloat(p.PlayerLocation.Z)
	statSize := len(p.PlayerStats)
	x.WriteInt16(int16(statSize))
	//if writing an array of types it's best to write the len() of the slice as an int16
	for i := 0; i < statSize; i++ {
		x.WriteInt16(p.PlayerStats[i])
	}
}

//be aware of the order in which you write the bytes as they will need to be read back in that same order
func ReadPlayerToStruct(x *binio.Packet) Player {
	tmp := Player{}
	tmp.PlayerID = x.ReadUInt32()
	tmp.PlayerName = x.ReadString()
	tmp.PlayerLocation.X = x.ReadFloat()
	tmp.PlayerLocation.Y = x.ReadFloat()
	tmp.PlayerLocation.Z = x.ReadFloat()
	sliceSize := int(x.ReadInt16()) //read the size of our PlayerStats slice
	tmp.PlayerStats = make([]int16, sliceSize) //make sure to allocate for the structs slice
	for i := 0; i < sliceSize; i++ { //now we can loop and read each player stat
		tmp.PlayerStats[i] = x.ReadInt16()
	}
	return tmp
}

