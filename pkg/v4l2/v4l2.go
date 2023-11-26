// SPDX-License-Identifier: BSD-3-Clause
//go:build linux && (arm || arm64)

package v4l2

import (
	"encoding/binary"
	"io"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const defaultNumBuffers = 2

type Buffer struct {
	Data []byte

	fd    int
	index int
}

func (b *Buffer) Release() error {
	return enqueue(b.fd, b.index)
}

type Device struct {
	C       chan Buffer
	buffers [][]byte
	fd      int
}

func Open(path string) (*Device, error) {
	fd, err := unix.Open(path, unix.O_RDWR, 0o666)
	if nil != err {
		return nil, err
	}

	return &Device{
		C:       make(chan Buffer, defaultNumBuffers),
		buffers: make([][]byte, defaultNumBuffers),
		fd:      fd,
	}, nil
}

func (dev *Device) Close() error {
	return unix.Close(dev.fd)
}

func (dev *Device) SetBitrate(bitrate int32) error {
	return setCodecControl(dev.fd, CidMpegVideoBitrate, bitrate)
}

// SetPixelFormat configures frame geometry and pixel format. The pixel
// format may be a compressor supported by the device, such as MJPEG or
// H.264.
func (dev *Device) SetPixelFormat(width, height, format int) error {
	pfmt := pixFormat{
		width:       uint32(width),
		height:      uint32(height),
		pixelformat: uint32(format),
		field:       FieldAny,
	}

	fmt := format{
		typ: BufTypeVideoCapture,
		fmt: pfmt.marshal(),
	}

	return ioctl(dev.fd, VidiocSFmt, unsafe.Pointer(&fmt))
}

// SetRepeatSequenceHeader configures the device to output sequence
// parameter sets (SPS) and picture parameter sets (PPS) before each
// group-of-pictures (GoP). This is H.264 specific and not supported by
// all devices.
func (dev *Device) SetRepeatSequenceHeader(on bool) error {
	var value int32
	if on {
		value = 1
	}

	return setCodecControl(dev.fd, CidMpegVideoRepeatSeqHeader, value)
}

func (dev *Device) Start() error {
	// Request specified number of kernel-space buffers from device.
	if err := requestBuffers(dev.fd, len(dev.buffers)); nil != err {
		return err
	}

	// For each buffer...
	for i := 0; i < len(dev.buffers); i++ {
		// Get length and offset of i-th buffer.
		length, offset, err := queryBuffer(dev.fd, uint32(i))
		if nil != err {
			return err
		}

		// Memory map i-th buffer to user-space.
		if dev.buffers[i], err = unix.Mmap(
			dev.fd,
			int64(offset),
			int(length),
			unix.PROT_READ|unix.PROT_WRITE,
			unix.MAP_SHARED,
		); nil != err {
			return err
		}

		// Enqueue to device for population.
		if err := enqueue(dev.fd, i); nil != err {
			return err
		}
	}

	go func(dev *Device) {
		for {
			// Dequeue buffer.
			i, n, err := dequeue(dev.fd)
			if nil != err {
				if err == syscall.EINVAL {
					err = io.EOF
				}
				return
			}

			// Write buffer to channel.
			// Note: Zero-copy. Slice bounds are written, not contents.
			dev.C <- Buffer{
				Data: dev.buffers[i][:n],

				fd:    dev.fd,
				index: i,
			}
		}
	}(dev)

	// Enable stream.
	typ := BufTypeVideoCapture

	return ioctl(dev.fd, VidiocStreamOn, unsafe.Pointer(&typ))
}

func (dev *Device) Stop() error {
	// Disable stream (dequeues any outstanding buffers as well).
	typ := BufTypeVideoCapture
	if err := ioctl(dev.fd, VidiocStreamOff, unsafe.Pointer(&typ)); nil != err {
		return nil
	}

	// Unmap each buffer from user space.
	for i := 0; i < len(dev.buffers); i++ {
		if dev.buffers[i] != nil {
			if err := unix.Munmap(dev.buffers[i]); err != nil {
				return err
			}
			dev.buffers[i] = nil
		}
	}

	// Count of zero frees all buffers, after aborting or finishing any
	// DMA in progress.
	return requestBuffers(dev.fd, 0)
}

func setCodecControl(fd int, id uint32, value int32) error {
	return setControl(fd, CtrlClassMpeg, id, value)
}

func setControl(fd int, class, id uint32, value int32) error {
	const numControls = 1

	ctrls := [numControls]extControl{
		extControl{
			id:   id,
			size: 0,
		},
	}
	binary.NativeEndian.PutUint32(ctrls[0].value[:], uint32(value))

	extctrls := extControls{
		ctrl_class: class,
		count:      numControls,
		controls:   unsafe.Pointer(&ctrls),
	}

	return ioctl(fd, VidiocSExtCtrls, unsafe.Pointer(&extctrls))
}

func ioctl(fd int, req uint, arg unsafe.Pointer) error {
	if _, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(req),
		uintptr(arg),
	); errno != 0 {
		return errno
	}

	return nil
}

func queryBuffer(fd int, n uint32) (length, offset uint32, err error) {
	qb := buffer{
		index:  n,
		typ:    BufTypeVideoCapture,
		memory: MemoryMmap,
	}
	if err = ioctl(fd, VidiocQueryBuf, unsafe.Pointer(&qb)); err != nil {
		return
	}

	length = qb.length
	offset = binary.NativeEndian.Uint32(qb.m[0:4])

	return
}

func requestBuffers(fd int, n int) error {
	rb := requestbuffers{
		count:  uint32(n),
		typ:    BufTypeVideoCapture,
		memory: MemoryMmap,
	}

	return ioctl(fd, VidiocReqBufs, unsafe.Pointer(&rb))
}

func enqueue(fd int, index int) error {
	qbuf := buffer{
		typ:    BufTypeVideoCapture,
		memory: MemoryMmap,
		index:  uint32(index),
	}

	return ioctl(fd, VidiocQBuf, unsafe.Pointer(&qbuf))
}

func dequeue(fd int) (int, int, error) {
	dqbuf := buffer{
		typ: BufTypeVideoCapture,
	}

	err := ioctl(fd, VidiocDQBuf, unsafe.Pointer(&dqbuf))

	return int(dqbuf.index), int(dqbuf.bytesused), err
}
