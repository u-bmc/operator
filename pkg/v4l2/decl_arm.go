// SPDX-License-Identifier: BSD-3-Clause
//go:build linux && arm

package v4l2

const (
	VidiocDQBuf     = 0xc0445611
	VidiocQBuf      = 0xc044560f
	VidiocQueryBuf  = 0xc0445609
	VidiocGExtCtrls = 0xc0185647
	VidiocSExtCtrls = 0xc0185648
	VidiocSFmt      = 0xc0cc5605
)

const maxSizeBufferDotM = 4

type format struct {
	typ uint32
	fmt [maxSizeFormatDotFmt]byte
}

type requestbuffers struct {
	count    uint32
	typ      uint32
	memory   uint32
	reserved [2]uint32
}

type timeval struct {
	tvSec  uint32
	tvUsec uint32
}

type v4l2_buffer struct {
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
