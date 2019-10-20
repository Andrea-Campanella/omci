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

// PhysicalPathTerminationPointXdslUniPart1ClassID is the 16-bit ID for the OMCI
// Managed entity Physical path termination point xDSL UNI part 1
const PhysicalPathTerminationPointXdslUniPart1ClassID ClassID = ClassID(98)

var physicalpathterminationpointxdslunipart1BME *ManagedEntityDefinition

// PhysicalPathTerminationPointXdslUniPart1 (class ID #98)
//	This ME represents the point where physical paths terminate on an xDSL CO modem (xTU-C). The
//	xDSL ME family is used for ADSL VDSL2 and FAST services. A legacy family of VDSL MEs remains
//	valid for ITUT G.993.1 VDSL, if needed. It is documented in [ITUT G.983.2].
//
//	The ONU automatically creates an instance of this ME per port:
//
//	o	when the ONU has xDSL ports built into its factory configuration;
//
//	o	when a cardholder is provisioned to expect a circuit pack of the xDSL type;
//
//	o	when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the xDSL
//	type. Note that the installation of a plug-and-play card may indicate the presence of xDSL ports
//	via equipment ID as well as its type, and indeed may cause the ONU to instantiate a port-mapping
//	package that specifies xDSL ports.
//
//	The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to
//	expect an xDSL circuit pack, nor is it equipped with an xDSL circuit pack.
//
//	Relationships
//		An instance of this ME is associated with each instance of a real or pre-provisioned xDSL port.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID:	This attribute uniquely identifies each instance of this ME. This 2 byte
//			number indicates the physical position of the UNI. The six LSBs of the first byte are the slot
//			ID, defined in clause 9.1.5. The two MSBs indicate the channel number in some of the implicitly
//			linked MEs, and must be 0 in the PPTP itself. This reduces the possible number of physical slots
//			to 64. The second byte is the port ID, with the range 1..255. (R) (mandatory) (2-bytes)
//
//		Loopback Configuration
//			Upon ME instantiation, the ONU sets this attribute to 0. (R,-W) (mandatory) (1-byte)
//
//		Administrative State
//			Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this
//			ME. Administrative state is further described in clause A.1.6. (R,-W) (mandatory) (1-byte)
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1-byte)
//
//		Xdsl Line Configuration Profile
//			xDSL line configuration profile: This attribute points to an instance of the xDSL line
//			configuration profiles (part 1, 2 and 3) MEs, and if necessary, also to VDSL2 line configuration
//			extensions (1 and 2) MEs, also to vectoring line configuration extension MEs. Upon ME
//			instantiation, the ONU sets this attribute to 0, a null pointer. (R,-W) (mandatory) (2-bytes)
//
//		Xdsl Subcarrier Masking Downstream Profile
//			xDSL subcarrier masking downstream profile: This attribute points to an instance of the xDSL
//			subcarrier masking downstream profile ME. Upon ME instantiation, the ONU sets this attribute to
//			0, a null pointer. (R,-W) (mandatory) (2-bytes)
//
//		Xdsl Subcarrier Masking Upstream Profile
//			xDSL subcarrier masking upstream profile: This attribute points to an instance of the xDSL
//			subcarrier masking upstream profile ME. Upon ME instantiation, the ONU sets this attribute to 0,
//			a null pointer. (R,-W) (mandatory) (2-bytes)
//
//		Xdsl Downstream Power Spectral Density Psd Mask Profile
//			xDSL downstream power spectral density (PSD) mask profile: This attribute points to an instance
//			of the xDSL PSD mask profile ME that defines downstream parameters. Upon ME instantiation, the
//			ONU sets this attribute to 0, a null pointer. (R,-W) (mandatory) (2-bytes)
//
//		Xdsl Downstream Rfi Bands Profile
//			xDSL downstream RFI bands profile: This attribute points to an instance of the xDSL downstream
//			RFI bands profile ME. Upon ME instantiation, the ONU sets this attribute to 0, a null pointer.
//			(R,-W) (mandatory) (2-bytes)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Modem Type
//			NOTE - Many newer VDSL2 chip sets support only PTM. The ATM default is retained for backward
//			compatibility, but implementers should be aware that the default may need to be overridden by
//			provisioning before the xDSL UNI can be brought into service.
//
//		Upstream Psd Mask Profile
//			Upstream PSD mask profile: This attribute points to an instance of the xDSL PSD mask profile
//			that defines upstream parameters. Upon ME instantiation, the ONU sets this attribute to 0, a
//			null pointer. (R,-W) (optional) (2-bytes)
//
//		Network Specific Extensions Pointer
//			Network specific extensions pointer: This attribute points to a network address ME that contains
//			the path and name of a file containing network specific parameters for the associated UNI. Upon
//			ME instantiation, the ONU sets this attribute to 0xFFFF, a null pointer. (R,-W) (optional)
//			(2-bytes)
//
type PhysicalPathTerminationPointXdslUniPart1 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointxdslunipart1BME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointXdslUniPart1",
		ClassID: 98,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("LoopbackConfiguration", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  ByteField("AdministrativeState", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3:  ByteField("OperationalState", 0, mapset.NewSetWith(Read), true, false, true, false, 3),
			4:  Uint16Field("XdslLineConfigurationProfile", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  Uint16Field("XdslSubcarrierMaskingDownstreamProfile", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  Uint16Field("XdslSubcarrierMaskingUpstreamProfile", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  Uint16Field("XdslDownstreamPowerSpectralDensityPsdMaskProfile", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  Uint16Field("XdslDownstreamRfiBandsProfile", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  ByteField("Arc", 0, mapset.NewSetWith(Read, Write), true, false, true, false, 9),
			10: ByteField("ArcInterval", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 10),
			11: ByteField("ModemType", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 11),
			12: Uint16Field("UpstreamPsdMaskProfile", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 12),
			13: Uint16Field("NetworkSpecificExtensionsPointer", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 13),
		},
	}
}

// NewPhysicalPathTerminationPointXdslUniPart1 (class ID 98 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointXdslUniPart1(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*physicalpathterminationpointxdslunipart1BME, params...)
}
