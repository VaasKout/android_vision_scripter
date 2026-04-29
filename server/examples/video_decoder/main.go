// Package main ...
package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"image"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/asticode/go-astiav"
	"gocv.io/x/gocv"
)

const (
	BufSize    = 1 * 1024 * 1024 //1MB
	HeaderSize = 12
)

const (
	SCPacketFlagConfig   uint64 = 1 << 63
	SCPacketFlagKeyFrame uint64 = 1 << 62

	SCPacketPtsMask uint64 = SCPacketFlagKeyFrame - 1
)

type DecoderData struct {
	CodecContext   *astiav.CodecContext
	Decoder        *astiav.Codec
	Pkt            *astiav.Packet
	Frame          *astiav.Frame
	DrawFrame      *astiav.Frame
	Buf            []byte
	HeaderBuf      []byte
	ConfigFrameBuf []byte
	Img            *image.YCbCr
	mu             sync.Mutex
}

func (d *DecoderData) Free() {
	if d == nil {
		return
	}
	d.DrawFrame.Free()
	d.Frame.Free()
	d.Pkt.Free()
	d.CodecContext.Free()
	d.Buf = []byte{}
	d.HeaderBuf = []byte{}
	d.ConfigFrameBuf = []byte{}
}

func (d *DecoderData) Allocate(width, height int) {
	if d == nil {
		return
	}
	codecID := astiav.CodecIDH264
	d.Decoder = astiav.FindDecoder(codecID)
	d.CodecContext = astiav.AllocCodecContext(d.Decoder)

	var newFlags = d.CodecContext.Flags().Add(astiav.CodecContextFlagLowDelay)
	d.CodecContext.SetFlags(newFlags)
	d.CodecContext.SetWidth(width)
	d.CodecContext.SetHeight(height)
	d.CodecContext.SetPixelFormat(astiav.PixelFormatYuv420P)
	d.Pkt = astiav.AllocPacket()
	d.Frame = astiav.AllocFrame()
	d.DrawFrame = astiav.AllocFrame()
	d.Img = &image.YCbCr{}
	d.HeaderBuf = make([]byte, HeaderSize)
	d.Buf = make([]byte, BufSize)
}

func main() {
	run()
}

func run() {
	astiav.SetLogLevel(astiav.LogLevelDebug)
	astiav.SetLogCallback(func(c astiav.Classer, l astiav.LogLevel, _, msg string) {
		var cs string
		if c != nil {
			if cl := c.Class(); cl != nil {
				cs = " - class: " + cl.String()
			}
		}
		log.Printf("ffmpeg log: %s%s - level: %d\n", strings.TrimSpace(msg), cs, l)
	})

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("Failed to connect to scrcpy socket: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to scrcpy server")

	width, height := readMetaData(conn)
	fmt.Printf("Width: %d\n", width)
	fmt.Printf("Height: %d\n", height)

	var data = &DecoderData{}
	data.Allocate(width, height)
	defer data.Free()

	// Open codec context
	if err := data.CodecContext.Open(data.Decoder, nil); err != nil {
		log.Println(fmt.Errorf("main: opening codec context failed: %w", err))
		return
	}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			err := handlePackets(conn, data)
			if err != nil {
				cancel()
				break
			}
		}
	}()
	showVideo(ctx, data)
}

func readMetaData(conn net.Conn) (width, height int) {
	var dummyByte = make([]byte, 1)
	_, err := conn.Read(dummyByte)
	if err != nil {
		fmt.Println(err)
		return width, height
	}
	fmt.Println(dummyByte)

	var deviceName = make([]byte, 64)
	_, err = conn.Read(deviceName)
	if err != nil {
		fmt.Println(err)
		return width, height
	}
	fmt.Println(deviceName)

	var codecID = make([]byte, 4)
	_, err = conn.Read(codecID)
	if err != nil {
		fmt.Println(err)
		return width, height
	}
	fmt.Println(codecID)

	var widthBuf = make([]byte, 4)
	_, err = conn.Read(widthBuf)
	if err != nil {
		fmt.Println(err)
		return width, height
	}
	fmt.Println(widthBuf)

	var heightBuf = make([]byte, 4)
	_, err = conn.Read(heightBuf)
	if err != nil {
		fmt.Println(err)
		return width, height
	}
	fmt.Println(heightBuf)

	width = int(binary.BigEndian.Uint32(widthBuf))
	height = int(binary.BigEndian.Uint32(heightBuf))
	return width, height
}

func handlePackets(
	conn net.Conn,
	data *DecoderData,
) error {
	if data == nil {
		return fmt.Errorf("decoder data is nil")
	}

	headerSize, err := io.ReadFull(conn, data.HeaderBuf)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if headerSize != 12 {
		fmt.Println("header size is not compatible")
		return nil
	}

	pts := binary.BigEndian.Uint64(data.HeaderBuf[0:8])
	fmt.Println(pts)
	packetSize := binary.BigEndian.Uint32(data.HeaderBuf[8:12])
	fmt.Println(packetSize)

	if pts <= 0 || int(packetSize) <= 0 {
		return nil
	}

	if BufSize < int(packetSize) {
		fmt.Printf("SIZE: %d not compatible\n", packetSize)
		return nil
	}

	size, err := io.ReadFull(conn, data.Buf[:packetSize])
	if err != nil {
		fmt.Printf("Connection closed or read error: %v\n", err)
		return err
	}

	if size < 0 || uint32(size) < packetSize {
		fmt.Printf("SIZE: %d not compatible\n", packetSize)
		return nil
	}

	var configSize = len(data.ConfigFrameBuf)
	var neededSpace = configSize + size
	if configSize > 0 {
		if neededSpace > len(data.Buf) {
			data.Buf = append(data.ConfigFrameBuf, data.Buf...)
		} else {
			copy(data.Buf[configSize:], data.Buf[:len(data.Buf)])
			copy(data.Buf[:configSize], data.ConfigFrameBuf)
		}
		data.ConfigFrameBuf = []byte{}
	}

	err = data.Pkt.FromData(data.Buf[:neededSpace])
	if err != nil {
		log.Println(fmt.Errorf("data set failed: %w", err))
		return err
	}
	defer data.Pkt.Unref()

	if pts&SCPacketFlagKeyFrame != 0 {
		var updatedFlags = data.Pkt.Flags().Add(astiav.PacketFlagKey)
		data.Pkt.SetFlags(updatedFlags)
	}

	var initConfigBuf = false
	if pts&SCPacketFlagConfig != 0 {
		data.Pkt.SetPts(astiav.NoPtsValue)
		initConfigBuf = true
	} else {
		data.Pkt.SetPts(int64(pts & SCPacketPtsMask))
	}

	data.Pkt.SetDts(data.Pkt.Pts())

	if initConfigBuf {
		data.ConfigFrameBuf = data.Pkt.Data()
		return nil
	}

	if err := data.CodecContext.SendPacket(data.Pkt); err != nil {
		log.Println(fmt.Errorf("sending packet failed: %w", err))
		return err
	}

	receiveFrames(data)
	return nil
}

func receiveFrames(data *DecoderData) {
	for {
		if err := data.CodecContext.ReceiveFrame(data.Frame); err != nil {
			data.Frame.Unref()
			if err == astiav.ErrEagain || err == astiav.ErrEof {
				return
			}
		}

		data.mu.Lock()
		data.DrawFrame.Unref()
		data.DrawFrame.MoveRef(data.Frame)
		data.mu.Unlock()
	}
}

func frameToMat(data *DecoderData) (gocv.Mat, error) {
	data.mu.Lock()
	if err := data.DrawFrame.Data().ToImage(data.Img); err != nil {
		data.mu.Unlock()
		return gocv.NewMat(), fmt.Errorf("failed to convert frame to image: %w", err)
	}
	data.mu.Unlock()

	// Convert the Go image to a gocv.Mat
	mat, err := gocv.ImageToMatRGB(data.Img)
	if err != nil {
		mat.Close()
		return gocv.NewMat(), fmt.Errorf("failed to convert image to Mat: %w", err)
	}
	return mat, nil
}

func showVideo(
	ctx context.Context,
	data *DecoderData,
) {
	window := gocv.NewWindow("H264 Stream")
	defer window.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			img, err := frameToMat(data)
			if err != nil {
				img.Close()
				continue
			}

			if img.Empty() {
				log.Println("image is empty")
				img.Close()
				continue
			}

			window.IMShow(img)
			img.Close()
			window.WaitKey(1)
		}
	}
}
