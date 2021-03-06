// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_freebsd.go

package disk

const (
	sizeofPtr        = 0x4
	sizeofShort      = 0x2
	sizeofInt        = 0x4
	sizeofLong       = 0x4
	sizeofLongLong   = 0x8
	sizeofLongDouble = 0x8

	DEVSTAT_NO_DATA = 0x00
	DEVSTAT_READ    = 0x01
	DEVSTAT_WRITE   = 0x02
	DEVSTAT_FREE    = 0x03
)

const (
	sizeOfDevstat = 0xf0
)

type (
	_C_short       int16
	_C_int         int32
	_C_long        int32
	_C_long_long   int64
	_C_long_double int64
)

type Statfs struct {
	Version     uint32
	Type        uint32
	Flags       uint64
	Bsize       uint64
	Iosize      uint64
	Blocks      uint64
	Bfree       uint64
	Bavail      int64
	Files       uint64
	Ffree       int64
	Syncwrites  uint64
	Asyncwrites uint64
	Syncreads   uint64
	Asyncreads  uint64
	Spare       [10]uint64
	Namemax     uint32
	Owner       uint32
	Fsid        Fsid
	Charspare   [80]int8
	Fstypename  [16]int8
	Mntfromname [88]int8
	Mntonname   [88]int8
}
type Fsid struct {
	Val [2]int32
}

type Devstat struct {
	Sequence0     uint32
	Allocated     int32
	Start_count   uint32
	End_count     uint32
	Busy_from     Bintime
	Dev_links     _Ctype_struct___0
	Device_number uint32
	Device_name   [16]int8
	Unit_number   int32
	Bytes         [4]uint64
	Operations    [4]uint64
	Duration      [4]Bintime
	Busy_time     Bintime
	Creation_time Bintime
	Block_size    uint32
	Tag_types     [3]uint64
	Flags         uint32
	Device_type   uint32
	Priority      uint32
	Id            *byte
	Sequence1     uint32
}
type Bintime struct {
	Sec  int32
	Frac uint64
}

type _Ctype_struct___0 struct {
	Empty uint32
}
