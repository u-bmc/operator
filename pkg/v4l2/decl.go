// SPDX-License-Identifier: BSD-3-Clause
//go:build linux

package v4l2

import "unsafe"

const (
	BufTypeVideoCapture = 1

	FieldAny  = 0
	FieldNone = 1

	PixFmtJPEG = 'J' | 'P'<<8 | 'E'<<16 | 'G'<<24
	PixFmtH264 = 'H' | '2'<<8 | '6'<<16 | '4'<<24

	MemoryMmap = 1

	VidiocQueryCap  = 0x80685600
	VidiocReqBufs   = 0xc0145608
	VidiocStreamOn  = 0x40045612
	VidiocStreamOff = 0x40045613
	VidiocSCtrl     = 0xc008561c
)

// Controls (from linux/v4l2-controls.h).
const (
	CtrlClassUser = 0x00980000
	CtrlClassMpeg = 0x00990000
)

// User-class control IDs.
const (
	CidBase             = CtrlClassUser | 0x900
	CidUserBase         = CidBase
	CidUserClass        = CtrlClassUser | 1
	CidBrightness       = CidBase + 0
	CidContrast         = CidBase + 1
	CidSaturation       = CidBase + 2
	CidHue              = CidBase + 3
	CidAutoWhiteBalance = CidBase + 12
	CidDoWhiteBalance   = CidBase + 13
	CidRedBalance       = CidBase + 14
	CidBlueBalance      = CidBase + 15
	CidGamma            = CidBase + 16
	CidExposure         = CidBase + 17
	CidAutogain         = CidBase + 18
	CidGain             = CidBase + 19
	CidHflip            = CidBase + 20
	CidVflip            = CidBase + 21
)

// MPEG-class control IDs.
const (
	CidMpegBase                 = CtrlClassMpeg | 0x900
	CidMpegClass                = CtrlClassMpeg | 1
	CidMpegVideoBFrames         = CidMpegBase + 202
	CidMpegVideoGopSize         = CidMpegBase + 203
	CidMpegVideoBitrate         = CidMpegBase + 207
	CidMpegVideoRepeatSeqHeader = CidMpegBase + 226
	CidMpegVideoH264IPeriod     = CidMpegBase + 358
	CidMpegVideoH264Level       = CidMpegBase + 359
	CidMpegVideoH264Profile     = CidMpegBase + 363
)

// H.264 Levels.
const (
	MpegVideoH264Level10 = iota
	MpegVideoH264Level1B
	MpegVideoH264Level11
	MpegVideoH264Level12
	MpegVideoH264Level13
	MpegVideoH264Level20
	MpegVideoH264Level21
	MpegVideoH264Level22
	MpegVideoH264Level30
	MpegVideoH264Level31
	MpegVideoH264Level32
	MpegVideoH264Level40
	MpegVideoH264Level41
	MpegVideoH264Level42
	MpegVideoH264Level50
	MpegVideoH264Level51
)

// H.264 Profiles.
const (
	MpegVideoH264ProfileBaseline = iota
	MpegVideoH264ProfileConstrainedBaseline
	MpegVideoH264ProfileMain
	MpegVideoH264ProfileExtended
	MpegVideoH264ProfileHigh
	MpegVideoH264ProfileHigh10
	MpegVideoH264ProfileHigh422
	MpegVideoH264ProfileHigh444Predictive
	MpegVideoH264ProfileHigh10Intra
	MpegVideoH264ProfileHigh422Intra
	MpegVideoH264ProfileHigh444Intra
	MpegVideoH264ProfileCavlc444Intra
	MpegVideoH264ProfileScalableBaseline
	MpegVideoH264ProfileScalableHigh
	MpegVideoH264ProfileScalableHighIntra
)

const (
	maxSizeExtControlDotValue = 8
	maxSizeFormatDotFmt       = 200
	sizePixFormat             = 48
)

type capability struct {
	driver       [16]uint8
	card         [32]uint8
	busInfo      [32]uint8
	version      uint32
	capabilities uint32
	deviceCaps   uint32
	reserved     [3]uint32
}

type pixFormat struct {
	width        uint32
	height       uint32
	pixelformat  uint32
	field        uint32
	bytesperline uint32
	sizeimage    uint32
	colorspace   uint32
	priv         uint32
	flags        uint32
	xxEnc        uint32
	quantization uint32
	xferFunc     uint32
}

type control struct {
	id    uint32
	value int32
}

type timecode struct {
	typ      uint32
	flags    uint32
	frames   uint8
	seconds  uint8
	minutes  uint8
	hours    uint8
	userbits [4]uint8
}

type extControl struct {
	id        uint32
	size      uint32
	reserved2 [1]uint32
	value     [maxSizeExtControlDotValue]byte
}

type extControls struct {
	ctrlClass uint32
	count     uint32
	errorIdx  uint32
	reserved  [2]uint32
	controls  unsafe.Pointer
}

func (pfmt *pixFormat) marshal() [maxSizeFormatDotFmt]byte {
	var b [maxSizeFormatDotFmt]byte

	copy(b[0:sizePixFormat], (*[sizePixFormat]byte)(unsafe.Pointer(pfmt))[:])

	return b
}
