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

// XdslChannelConfigurationProfilePart2ClassID is the 16-bit ID for the OMCI
// Managed entity xDSL channel configuration profile part 2
const XdslChannelConfigurationProfilePart2ClassID ClassID = ClassID(412)

var xdslchannelconfigurationprofilepart2BME *ManagedEntityDefinition

// XdslChannelConfigurationProfilePart2 (class ID #412)
//	This ME contains the channel configuration profile for an xDSL UNI. An instance of this ME is
//	created and deleted by the OLT.
//
//	NOTE - If [ITUT G.997.1] compatibility is required, bit rates should only be set to integer
//	multiples of 1000-bits/s. The ONU may reject attempts to set other values for bit rate
//	attributes.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the xDSL channel configuration
//			profile. (R, setbycreate) (mandatory) (2-bytes)
//
//		Minimum Expected Throughput For Retransmission Minetr_Rtx
//			Minimum expected throughput for retransmission (MINETR_RTX): If retransmission is used in a
//			given transmit direction, this attribute specifies the minimum expected throughput for the
//			bearer channel, in bits per second. See clause 7.3.2.1.8 of [ITU-T G.997.1]. (R,-W) (mandatory)
//			(4-bytes)
//
//		Maximum Expected Throughput For Retransmission Maxetr_Rtx
//			Maximum expected throughput for retransmission (MAXETR_RTX): If retransmission is used in a
//			given transmit direction, this parameter specifies the maximum expected throughput for the
//			bearer channel, in bits per second. See clause 7.3.2.1.9 of [ITU-T G.997.1]. (R,-W) (mandatory)
//			(4-bytes)
//
//		Maximum Net Data Rate For Retransmission Maxndr_Rtx
//			Maximum net data rate for retransmission (MAXNDR_RTX): If retransmission is used in a given
//			transmit direction, this parameter specifies the maximum net data rate for the bearer channel,
//			in bits per second. See clause 7.3.2.1.10 of [ITUT-G.997.1]. (R,-W) (mandatory) (4-bytes)
//
//		Maximum Delay For Retransmission Delaymax_Rtx
//			Maximum delay for retransmission (DELAYMAX_RTX): If retransmission is used in a given transmit
//			direction, this parameter specifies the maximum for the instantaneous delay due to the effect of
//			retransmission only. This delay is defined as the integer value of this attribute multiplied by
//			1-ms. The valid delay values are given in clause 7.3.2.11 of [ITU-T G.997.1]. (R,-W) (mandatory)
//			(1-bytes)
//
//		Minimum Delay For Retransmission Delaymin_Rtx
//			Minimum delay for retransmission (DELAYMIN_RTX): If retransmission is used in a given transmit
//			direction, this parameter specifies the minimum for the instantaneous delay due to the effect of
//			retransmission only. This delay is defined as the integer value of this attribute multiplied by
//			1 ms. The valid delay values are given in clause 7.3.2.12 of [ITU-T G.997.1]. (R,-W) (mandatory)
//			(1-bytes)
//
//		Minimum Impulse Noise Protection Against Single High Impulse Noise Event Shine For Retransmission Inpmin_Shine_Rtx
//			Minimum impulse noise protection against single high impulse noise event (SHINE) for
//			retransmission (INPMIN_SHINE_RTX): If retransmission is used in a given transmit direction, this
//			parameter specifies the minimum INP against a SHINE for the bearer channel if it is transported
//			over DMT symbols with a subcarrier spacing of 4.3125-kHz. The valid range of values is given in
//			clause-7.3.2.13 of [ITU-T G.997.1]. (R,-W) (mandatory) (1-bytes)
//
//		Minimum Impulse Noise Protection Against Shine For Retransmission For Systems Using 8.625 Khz Subcarrier Spacing Inpmin8_Shine_Rtx
//			Minimum impulse noise protection against SHINE for retransmission for systems using 8.625 kHz
//			subcarrier spacing (INPMIN8_SHINE_RTX): If retransmission is used in a given transmit direction,
//			this parameter specifies the minimum INP against SHINE for the bearer channel if it is
//			transported over DMT symbols with a subcarrier spacing of 8.625-kHz. The valid range of values
//			is given in clause 7.3.2.14 of [ITUT-G.997.1]. (R,-W) (mandatory) (1-bytes)
//
//		Shineratio_Rtx
//			SHINERATIO_RTX: If retransmission is used in a given transmit direction, this parameter
//			specifies the SHINE ratio. This ratio is defined as the integer value of this attribute
//			multiplied by 0.001. The valid range of values is given in clause-7.3.2.15 of [ITU-T G.997.1].
//			(R,-W) (mandatory) (1-bytes)
//
//		Minimum Impulse Noise Protection Against Rein For Retransmission Inpmin_Rein_Rtx
//			Minimum impulse noise protection against REIN for retransmission (INPMIN_REIN_RTX): If
//			retransmission is used in a given transmit direction, this parameter specifies the minimum INP
//			against REIN for the bearer channel if it is transported over DMT symbols with a subcarrier
//			spacing of 4.3125 kHz. The valid range of values is given in clause-7.3.2.16 of [ITU-T G.997.1].
//			(R,-W) (mandatory) (1-bytes)
//
//		Minimum Impulse Noise Protection Against Rein For Retransmission For Systems Using 8.625_Khz Subcarrier Spacing Inpmin8_Rein_Rtx
//			Minimum impulse noise protection against REIN for retransmission for systems using 8.625-kHz
//			subcarrier spacing (INPMIN8_REIN_RTX): If retransmission is used in a given transmit direction,
//			this parameter specifies the minimum INP against REIN for the bearer channel if it is
//			transported over DMT symbols with a subcarrier spacing of 8.625 kHz. The valid range of values
//			is given in clause 7.3.2.17 of [ITU-T G.997.1]. (R,-W) (mandatory) (1-bytes)
//
//		Rein Inter_Arrival Time For Retransmission Iat_Rein_Rtx
//			REIN inter-arrival time for retransmission (IAT_REIN_RTX): If retransmission is used in a given
//			transmit direction, this parameter specifies the IAT that shall be assumed for REIN protection.
//			The valid range of values is given in clause 7.3.2.18 of [ITU-T G.997.1]. (R,-W) (mandatory)
//			(1-bytes)
//
//		Target Net Data Rate Target_Ndr
//			Target net data rate (TARGET_NDR): If retransmission is not used in a given transmit direction,
//			this parameter specifies the target net data of the bearer channel, in bits per second. See
//			clause 7.3.2.19.1 of [ITU-T G.997.1]. (R,-W) (mandatory) (4-bytes)
//
//		Target Expected Throughput For Retransmission Target_Etr
//			Target expected throughput for retransmission (TARGET_ETR): If retransmission is used in a given
//			transmit direction, this parameter specifies the target expected throughput for the bearer
//			channel, in bits per second. See clause 7.3.2.19.2 of [ITUT-G.997.1]. (R,-W) (mandatory)
//			(4-bytes)
//
type XdslChannelConfigurationProfilePart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslchannelconfigurationprofilepart2BME = &ManagedEntityDefinition{
		Name:    "XdslChannelConfigurationProfilePart2",
		ClassID: 412,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  Uint32Field("MinimumExpectedThroughputForRetransmissionMinetrRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  Uint32Field("MaximumExpectedThroughputForRetransmissionMaxetrRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3:  Uint32Field("MaximumNetDataRateForRetransmissionMaxndrRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  ByteField("MaximumDelayForRetransmissionDelaymaxRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  ByteField("MinimumDelayForRetransmissionDelayminRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  ByteField("MinimumImpulseNoiseProtectionAgainstSingleHighImpulseNoiseEventShineForRetransmissionInpminShineRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  ByteField("MinimumImpulseNoiseProtectionAgainstShineForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ShineRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  ByteField("ShineratioRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionInpminReinRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 9),
			10: ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ReinRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 10),
			11: ByteField("ReinInterArrivalTimeForRetransmissionIatReinRtx", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 11),
			12: Uint32Field("TargetNetDataRateTargetNdr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 12),
			13: Uint32Field("TargetExpectedThroughputForRetransmissionTargetEtr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 13),
		},
	}
}

// NewXdslChannelConfigurationProfilePart2 (class ID 412 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelConfigurationProfilePart2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslchannelconfigurationprofilepart2BME, params...)
}
