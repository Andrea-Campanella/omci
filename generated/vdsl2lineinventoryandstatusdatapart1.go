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

// Vdsl2LineInventoryAndStatusDataPart1ClassID is the 16-bit ID for the OMCI
// Managed entity VDSL2 line inventory and status data part 1
const Vdsl2LineInventoryAndStatusDataPart1ClassID ClassID = ClassID(168)

var vdsl2lineinventoryandstatusdatapart1BME *ManagedEntityDefinition

// Vdsl2LineInventoryAndStatusDataPart1 (class ID #168)
//	This ME extends the xDSL line configuration MEs. The ME name was chosen because its attributes
//	were initially unique to ITU-T G.993.2 VDSL2. Due to continuing standards development, some
//	attributes - and therefore this ME - have also become applicable to other Recommendations,
//	specifically [ITU-T G.992.3] and [ITU-T G.992.5].
//
//	This ME contains general and downstream attributes.
//
//	Relationships
//		This is one of the status data MEs associated with an xDSL UNI. It is meaningful if the PPTP
//		supports [ITU-T G.992.3], [ITU-T G.992.5] or [ITU-T G.993.2]. The ONU automatically creates or
//		deletes an instance of this ME upon creation and deletion of a PPTP xDSL UNI part 1 that
//		supports these attributes.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1 ME. (R)
//			(mandatory) (2-bytes)
//
//		Vdsl2 Transmission System Capability Xtu C
//			VDSL2 transmission system capability xTUC: This attribute extends the xTU-C transmission system
//			capability attribute of the xDSL line inventory and status data part 1 to include xTU-C VDSL2
//			capabilities. It is defined by bits 57..64 of Table 9.7.12-1. (R) (mandatory) (1-byte)
//
//		Vdsl2 Transmission System
//			VDSL2 transmission system: This attribute reports the transmission system in use. It extends the
//			xDSL transmission system attribute of the xDSL line inventory and status data part 2 ME with a
//			byte that includes VDSL2 capabilities currently in use. It is defined by bits 57..64 of Table
//			9.7.12-1. (R) (mandatory) (1-byte)
//
//		Vdsl2 Profile
//			(R) (mandatory) (1-byte)
//
//		Vdsl2 Limit Psd Mask And Bandplan
//			VDSL2 limit PSD mask and bandplan: This attribute defines the limit PSD mask and band plan in
//			use. It is a bit map as defined by Table 9.7.6-1. (R) (mandatory) (8-bytes)
//
//		Vdsl2 Us0 Psd Mask
//			VDSL2 US0 PSD mask: This attribute defines the US0 PSD mask in use. It is a bit map as defined
//			by Table 9.7.62. (R) (mandatory) (4-bytes)
//
//		Actsnrmodeds
//			(R) (mandatory) (1-byte)
//
//		Hlingds
//			HLINGds:	This attribute contains the number of subcarriers per group used to report HLINpsds.
//			(R) (mandatory) (1-byte)
//
//		Hloggds
//			HLOGGds:	This attribute contains the number of subcarriers per group used to report HLOGpsds.
//			(R) (mandatory) (1-byte)
//
//		Qlngds
//			QLNGds:	This attribute contains the number of subcarriers per group used to report QLNpsds. (R)
//			(mandatory) (1-byte)
//
//		Snrgds
//			SNRGds:	This attribute contains the number of subcarriers per group used to report SNRpsds. (R)
//			(mandatory) (1-byte)
//
//		Mrefpsdds Table
//			(R) (mandatory) (3 * N bytes, where N is the number of breakpoints)
//
//		Trellisds
//			(R) (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1-byte)
//
//		Actual Rate Adaptation Mode Downstream
//			(R) (optional) (1-byte)
//
//		Actual Impulse Noise Protection Robust Operations Channel Roc Downstream
//			Actual impulse noise protection robust operations channel (ROC) downstream: The ACTINP-ROC-ds
//			attribute reports the actual INP of the ROC in the downstream direction expressed in multiples
//			of T4k. The INP of this attribute is equal to the integer value multiplied by 0.1 symbols. Valid
//			values and usage are given in clause 7.5.1.34.1 of [ITUT-G.997.1]. (R) (optional) (1-byte)
//
//		Snr Margin Roc Downstream
//			SNR margin ROC downstream: The SNRM-ROC-ds attribute reports the actual signal-to-noise margin
//			of the ROC in the downstream direction. Its value ranges from 0 (-64.0-dB) to 1270 (+63.0-dB).
//			The special value 0xFFFF indicates that the attribute is out of range. (R) (optional) (2-bytes)
//
type Vdsl2LineInventoryAndStatusDataPart1 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vdsl2lineinventoryandstatusdatapart1BME = &ManagedEntityDefinition{
		Name:    "Vdsl2LineInventoryAndStatusDataPart1",
		ClassID: 168,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("Vdsl2TransmissionSystemCapabilityXtuC", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  ByteField("Vdsl2TransmissionSystem", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3:  ByteField("Vdsl2Profile", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  Uint64Field("Vdsl2LimitPsdMaskAndBandplan", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  Uint32Field("Vdsl2Us0PsdMask", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  ByteField("Actsnrmodeds", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  ByteField("Hlingds", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  ByteField("Hloggds", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
			9:  ByteField("Qlngds", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: ByteField("Snrgds", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: TableField("MrefpsddsTable", TableInfo{0, 3}, mapset.NewSetWith(Read), false, false, false, 11),
			12: ByteField("Trellisds", 0, mapset.NewSetWith(Read), false, false, false, false, 12),
			13: ByteField("ActualRateAdaptationModeDownstream", 0, mapset.NewSetWith(Read), false, false, true, false, 13),
			14: ByteField("ActualImpulseNoiseProtectionRobustOperationsChannelRocDownstream", 0, mapset.NewSetWith(Read), false, false, true, false, 14),
			15: Uint16Field("SnrMarginRocDownstream", 0, mapset.NewSetWith(Read), false, false, true, false, 15),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewVdsl2LineInventoryAndStatusDataPart1 (class ID 168) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVdsl2LineInventoryAndStatusDataPart1(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*vdsl2lineinventoryandstatusdatapart1BME, params...)
}
