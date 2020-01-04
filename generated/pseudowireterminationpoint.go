/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// PseudowireTerminationPointClassID is the 16-bit ID for the OMCI
// Managed entity Pseudowire termination point
const PseudowireTerminationPointClassID ClassID = ClassID(282)

var pseudowireterminationpointBME *ManagedEntityDefinition

// PseudowireTerminationPoint (class ID #282)
//	The pseudowire TP supports packetized (rather than TDM) transport of TDM services, transported
//	either directly over Ethernet, over UDP/IP or over MPLS. Instances of this ME are created and
//	deleted by the OLT.
//
//	Relationships
//		One pseudowire TP ME exists for each distinct TDM service that is mapped to a pseudowire.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Underlying Transport
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Service Type
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Signalling
//			(R,-W, setbycreate) (mandatory for structured service type) (1-byte)
//
//		Tdm Uni Pointer
//			TDM UNI pointer: If service type-= structured, this attribute points to a logical N-* 64-kbit/s
//			subport CTP. Otherwise, this attribute points to a PPTP CES UNI. (R,-W, setbycreate) (mandatory)
//			(2-bytes)
//
//		North_Side Pointer
//			North-side pointer: When the pseudowire service is transported via IP, as indicated by the
//			underlying transport attribute, the northside pointer attribute points to an instance of the
//			TCP/UDP config data ME. When the pseudowire service is transported directly over Ethernet, the
//			north-side pointer attribute is not used - the linkage to the Ethernet flow TP is implicit in
//			the ME IDs. When the pseudowire service is transported over MPLS, the northside pointer
//			attribute points to an instance of the MPLS PW TP. (R, W, setbycreate) (mandatory) (2 bytes)
//
//		Far_End Ip Info
//			A null pointer is appropriate if the pseudowire is not transported via IP. (R,-W, setbycreate)
//			(mandatory for IP transport) (2-bytes)
//
//		Payload Size
//			(R,-W, setbycreate) (mandatory for unstructured service) (2-bytes)
//
//		Payload Encapsulation Delay
//			(R,-W, setbycreate) (mandatory for structured service) (1-byte)
//
//		Timing Mode
//			(R,-W) (mandatory) (1-byte)
//
//		Transmit Circuit Id
//			(R,-W) (mandatory for MEF 8 transport) (8-bytes)
//
//		Expected Circuit Id
//			(R,-W) (optional for MEF 8 transport) (8-bytes)
//
//		Received Circuit Id
//			Received circuit ID: This attribute indicates the actual ECID(s) received on the payload and
//			signalling channels, respectively. It may be used for diagnostic purposes. (R) (optional for MEF
//			8 transport) (8-bytes)
//
//		Exception Policy
//			Exception policy: This attribute points to an instance of the pseudowire maintenance profile ME.
//			If the pointer has its default value 0, the ONU's internal defaults apply. (R,-W) (optional)
//			(2-bytes)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
type PseudowireTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	pseudowireterminationpointBME = &ManagedEntityDefinition{
		Name:    "PseudowireTerminationPoint",
		ClassID: 282,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  ByteField("UnderlyingTransport", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2:  ByteField("ServiceType", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3:  ByteField("Signalling", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 3),
			4:  Uint16Field("TdmUniPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 4),
			5:  Uint16Field("NorthSidePointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 5),
			6:  Uint16Field("FarEndIpInfo", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 6),
			7:  Uint16Field("PayloadSize", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 7),
			8:  ByteField("PayloadEncapsulationDelay", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 8),
			9:  ByteField("TimingMode", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 9),
			10: Uint64Field("TransmitCircuitId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 10),
			11: Uint64Field("ExpectedCircuitId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 11),
			12: Uint64Field("ReceivedCircuitId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 12),
			13: Uint16Field("ExceptionPolicy", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 13),
			14: ByteField("Arc", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), true, true, false, 14),
			15: ByteField("ArcInterval", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 15),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewPseudowireTerminationPoint (class ID 282) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewPseudowireTerminationPoint(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*pseudowireterminationpointBME, params...)
}
