/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package omci

import (
	"encoding/binary"
	"errors"
	"fmt"
	me "github.com/cboling/omci/generated"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"hash/crc32"
)

type DeviceIdent byte

var (
	LayerTypeOMCI gopacket.LayerType
)

func init() {
	LayerTypeOMCI = gopacket.RegisterLayerType(1000,
		gopacket.LayerTypeMetadata{
			Name:    "OMCI",
			Decoder: gopacket.DecodeFunc(decodeOMCI),
		})
}

const (
	// Device Identifiers
	_                         = iota
	BaselineIdent DeviceIdent = 0x0A // All G-PON OLTs and ONUs support the baseline message set
	ExtendedIdent             = 0x0B
)

var OmciIK = []byte{0x18, 0x4b, 0x8a, 0xd4, 0xd1, 0xac, 0x4a, 0xf4,
	0xdd, 0x4b, 0x33, 0x9e, 0xcc, 0x0d, 0x33, 0x70}

func (di DeviceIdent) String() string {
	switch di {
	default:
		return "Unknown"

	case BaselineIdent:
		return "Baseline"

	case ExtendedIdent:
		return "Extended"
	}
}

// MaxBaselineLength is the maximum number of octets allowed in an OMCI Baseline
// message.  Depending on the adapter, it may or may not include the
const MaxBaselineLength = 48

// MaxExtendedLength is the maximum number of octets allowed in an OMCI Extended
// message (including header).
const MaxExtendedLength = 1980

// OMCI defines the common protocol. Extended will be added once
// I can get basic working (and layered properly).  See ITU-T G.988 11/2017 section
// A.3 for more information
type OMCI struct {
	layers.BaseLayer
	TransactionID    uint16
	MessageType      uint8
	DeviceIdentifier DeviceIdent
	Payload          []byte
	padding          []byte
	Length           uint16
	MIC              uint32
}

func (omci *OMCI) String() string {
	msgType := me.MsgType(omci.MessageType & me.MsgTypeMask)
	if me.IsAutonomousNotification(msgType) {
		return fmt.Sprintf("OMCI: Type: %v:", msgType)
	} else if omci.MessageType&me.AK == me.AK {
		return fmt.Sprintf("OMCI: Type: %v Response", msgType)
	}
	return fmt.Sprintf("OMCI: Type: %v Request", msgType)
}

// LayerType returns LayerTypeOMCI
func (omci *OMCI) LayerType() gopacket.LayerType {
	return LayerTypeOMCI
}

func (omci *OMCI) LayerContents() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint16(b, omci.TransactionID)
	b[2] = omci.MessageType
	b[3] = byte(omci.DeviceIdentifier)
	return b
}

func (omci *OMCI) CanDecode() gopacket.LayerClass {
	return LayerTypeOMCI
}

// NextLayerType returns the layer type contained by this DecodingLayer.
func (omci *OMCI) NextLayerType() gopacket.LayerType {
	return gopacket.LayerTypeZero
}

func decodeOMCI(data []byte, p gopacket.PacketBuilder) error {
	// Allow baseline messages without Length & MIC, but no less
	if len(data) < MaxBaselineLength-8 {
		return errors.New("frame header too small")
	}
	switch DeviceIdent(data[3]) {
	default:
		return errors.New("unsupported message type")

	case BaselineIdent:
		//omci := &BaselineMessage{}
		omci := &OMCI{}
		return omci.DecodeFromBytes(data, p)

	case ExtendedIdent:
		//omci := &ExtendedMessage{}
		omci := &OMCI{}
		return omci.DecodeFromBytes(data, p)
	}
}

func calculateMicCrc32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

//func calculateAes128(upstream bool, key uint16, data []byte) uint32 {
//	block, err := aes.NewCipher(key)
//	return crc32.ChecksumIEEE(data)
//}

/////////////////////////////////////////////////////////////////////////////
//   Baseline Message encode / decode
//
// TODO: For OMCI and MessageType decode, look into what the proper values for Padding and Payload should be
//       and what they are generally used for.  Currently often being set to 'nil'

func (omci *OMCI) DecodeFromBytes(data []byte, p gopacket.PacketBuilder) error {
	if len(data) < 10 {
		p.SetTruncated()
		return errors.New("frame too small")
	}
	omci.TransactionID = binary.BigEndian.Uint16(data[0:])
	omci.MessageType = data[2]
	omci.DeviceIdentifier = DeviceIdent(data[3])

	isNotification := (int(omci.MessageType) & ^me.MsgTypeMask) == 0
	if omci.TransactionID == 0 && !isNotification {
		return errors.New("omci Transaction ID is zero for non-Notification type message")
	}
	// Decode length
	var payloadOffset int
	var micOffset int
	if omci.DeviceIdentifier == BaselineIdent {
		omci.Length = MaxBaselineLength - 8
		payloadOffset = 8
		micOffset = MaxBaselineLength - 4

		if len(data) >= micOffset {
			length := binary.BigEndian.Uint32(data[micOffset-4:])
			if uint16(length) != omci.Length {
				return errors.New("invalid baseline message length")
			}
		}
	} else {
		payloadOffset = 10
		omci.Length = binary.BigEndian.Uint16(data[8:10])
		micOffset = int(omci.Length) + payloadOffset

		if omci.Length > MaxExtendedLength {
			return errors.New("extended frame exceeds maximum allowed")
		}
		if int(omci.Length) != micOffset {
			if int(omci.Length) < micOffset {
				p.SetTruncated()
			}
			return errors.New("extended frame too small")
		}
	}
	// Extract MIC if present in the data
	//if len(data) >= micOffset+4 {
	//	omci.MIC = binary.BigEndian.Uint32(data[micOffset:])
	//	actual := calculateMicCrc32(data[:micOffset])
	//	if omci.MIC != actual {
	//		msg := fmt.Sprintf("invalid MIC, expected %#x, got %#x",
	//			omci.MIC, actual)
	//		return errors.New(msg)
	//	}
	//}
	omci.BaseLayer = layers.BaseLayer{data[:4], data[4:]}
	p.AddLayer(omci)
	nextLayer, err := MsgTypeToNextLayer(omci.MessageType)
	if err != nil {
		return err
	}
	return p.NextDecoder(nextLayer)
}

// SerializeTo writes the serialized form of this layer into the
// SerializationBuffer, implementing gopacket.SerializableLayer.
// See the docs for gopacket.SerializableLayer for more info.
func (omci *OMCI) SerializeTo(b gopacket.SerializeBuffer, opts gopacket.SerializeOptions) error {
	// TODO: Hardcoded for baseline message format for now. Will eventually need to support
	//       the extended message format.
	bytes, err := b.PrependBytes(4)
	if err != nil {
		return err
	}
	// OMCI layer error checks
	isNotification := (int(omci.MessageType) & ^me.MsgTypeMask) == 0
	if omci.TransactionID == 0 && !isNotification {
		return errors.New("omci Transaction ID is zero for non-Notification type message")
	}
	if omci.DeviceIdentifier == BaselineIdent {
		if omci.Length != MaxBaselineLength-8 {
			msg := fmt.Sprintf("invalid Baseline message length: %v", omci.Length)
			return errors.New(msg)
		}
	} else if omci.DeviceIdentifier == ExtendedIdent {
		if omci.Length > MaxExtendedLength {
			msg := fmt.Sprintf("invalid Baseline message length: %v", omci.Length)
			return errors.New(msg)
		}
	} else {
		msg := fmt.Sprintf("invalid device identifier: %#x, Baseline or Extended expected",
			omci.DeviceIdentifier)
		return errors.New(msg)
	}
	binary.BigEndian.PutUint16(bytes, omci.TransactionID)
	bytes[2] = byte(omci.MessageType)
	bytes[3] = byte(omci.DeviceIdentifier)
	b.PushLayer(LayerTypeOMCI)

	bufLen := len(b.Bytes())
	padSize := int(omci.Length) - bufLen + 4
	padding, err := b.AppendBytes(padSize)
	copy(padding, lotsOfZeros[:])

	// For baseline, always provide the length
	binary.BigEndian.PutUint32(b.Bytes()[MaxBaselineLength-8:], 40)

	if opts.ComputeChecksums {
		micBytes, err := b.AppendBytes(4)
		if err != nil {
			return err
		}
		// TODO: Look up MIC definition and see if it includes the length
		omci.MIC = calculateMicCrc32(bytes[:MaxBaselineLength-4])
		binary.BigEndian.PutUint32(micBytes, omci.MIC)
	}
	return nil
}

// hacky way to zero out memory... there must be a better way?
var lotsOfZeros [MaxExtendedLength]byte // Extended OMCI messages may be up to 1980 bytes long, including headers
