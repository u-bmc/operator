// SPDX-License-Identifier: BSD-3-Clause
//go:build linux && arm64

package v4l2

const (
	VidiocDQBuf     = 0xc0585611
	VidiocQBuf      = 0xc058560f
	VidiocQueryBuf  = 0xc0585609
	VidiocGExtCtrls = 0xc0205647
	VidiocSExtCtrls = 0xc0205648
	VidiocSFmt      = 0xc0d05605
)

const maxSizeBufferDotM = 8

type format struct {
	typ uint64
	fmt [maxSizeFormatDotFmt]byte
}

type requestbuffers struct {
	count        uint32
	typ          uint32
	memory       uint32
	capabilities uint32
	flags        uint8
	reserved     [3]uint8
}

type timeval struct {
	tvSec  uint64
	tvUsec uint64
}

type buffer struct {
	index     uint32
	typ       uint32
	bytesused uint32
	flags     uint32
	field     uint32
	timestamp timeval
	timecode  timecode
	sequence  uint32
	memory    uint32
	m         [maxSizeBufferDotM]byte
	length    uint32
	reserved2 uint32
	reserved  uint32
}
