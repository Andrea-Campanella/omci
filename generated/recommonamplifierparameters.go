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

// ReCommonAmplifierParametersClassID is the 16-bit ID for the OMCI
// Managed entity RE common amplifier parameters
const ReCommonAmplifierParametersClassID ClassID = ClassID(328)

var recommonamplifierparametersBME *ManagedEntityDefinition

// ReCommonAmplifierParameters (class ID #328)
//	This ME organizes data associated with each OA supported by the RE. The management ONU
//	automatically creates one instance of this ME for each upstream or downstream OA.
//
//	Relationships
//		An instance of this ME is associated with an instance of the RE downstream amplifier or RE
//		upstream amplifier ME.
//
//	Attributes
//		Managed Entity Id
//			NOTE - The type of the linked ME can be determined by uniqueness of slot and port.
//
//		Gain
//			Gain:	This attribute reports the current measurement of the OA's gain, in decibels. Its value is
//			a 2s complement integer with 0.25-dB granularity, and with a range from -32-dB to 31.5-dB. The
//			value 0x7F indicates that the current measured gain is 0, i.e., negative infinity in decibels
//			terms. (R) (optional) (1-byte)
//
//		Lower Gain Threshold
//			Lower gain threshold: This attribute specifies the gain the RE uses to declare the low gain
//			alarm. Valid values are 0-dB (coded as 0x00) to 63.5-dB (coded as 0xFE). The default value 0xFF
//			selects the RE's internal policy. (R,-W) (optional) (1-byte)
//
//		Upper Gain Threshold
//			Upper gain threshold: This attribute specifies the gain the RE uses to declare the high gain
//			alarm. Valid values are 0-dB (coded as 0x00) to 63.5-dB (coded as 0xFE). The default value 0xFF
//			selects the RE's internal policy. (R,-W) (optional) (1-byte)
//
//		Target Gain
//			Target gain:	This attribute specifies the target gain, when the operational mode of the parent
//			RE downstream or upstream amplifier is set to constant gain mode. Valid values are 0-dB (coded
//			as 0x00) to 63.5-dB (coded as 0xFE). The default value 0xFF selects the RE's internal policy.
//			(R,-W) (optional) (1-byte)
//
//		Device Temperature
//			Device temperature: This attribute reports the temperature in degrees Celcius of the active
//			device (SOA or pump) in the OA. Its value is a 2s complement integer with granularity
//			1/256-degree-C. (R) (optional) (2-bytes)
//
//		Lower Device Temperature Threshold
//			Lower device temperature threshold: This attribute is a 2s complement integer that specifies the
//			temperature the RE uses to declare the low temperature alarm. Valid values are -64 to
//			+63-degree-C in 0.5-degree-C increments. The default value 0x7F selects the RE's internal
//			policy. (R,-W) (optional) (1-byte)
//
//		Upper Device Temperature Threshold
//			Upper device temperature threshold: This attribute is a 2s complement integer that specifies the
//			temperature the RE uses to declare the high temperature alarm. Valid values are -64 to
//			+63-degree-C in 0.5-degree-C increments. The default value 0x7F selects the RE's internal
//			policy. (R,-W) (optional) (1-byte)
//
//		Device Bias Current
//			Device bias current: This attribute contains the measured bias current applied to the SOA or
//			pump laser. Its value is an unsigned integer with granularity 2-mA. Valid values are 0 to
//			512-mA. (R) (optional) (1-byte)
//
//		Amplifier Saturation Output Power
//			Amplifier saturation output power: This attribute reports the saturation output power of the
//			amplifier as specified by the manufacturer. Its value is an unsigned integer referred to 1-mW
//			(i.e., dBm), with 0.1-dB granularity. (R) (optional) (2-bytes)
//
//		Amplifier Noise Figure
//			Amplifier noise figure: This attribute reports the intrinsic noise figure of the amplifier, as
//			specified by the manufacturer. Its value is an unsigned integer with 0.1-dB granularity (R)
//			(optional) (1-byte)
//
//		Amplifier Saturation Gain
//			Amplifier saturation gain: This attribute reports the gain of the amplifier at saturation, as
//			specified by the manufacturer. Its value is an unsigned integer with 0.25-dB granularity, and
//			with a range from 0 to 63.75-dB. (R) (optional) (1-byte)
//
type ReCommonAmplifierParameters struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	recommonamplifierparametersBME = &ManagedEntityDefinition{
		Name:    "ReCommonAmplifierParameters",
		ClassID: 328,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xffe0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1:  ByteField("Gain", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 1),
			2:  ByteField("LowerGainThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 2),
			3:  ByteField("UpperGainThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 3),
			4:  ByteField("TargetGain", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 4),
			5:  Uint16Field("DeviceTemperature", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 5),
			6:  ByteField("LowerDeviceTemperatureThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 6),
			7:  ByteField("UpperDeviceTemperatureThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 7),
			8:  ByteField("DeviceBiasCurrent", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 8),
			9:  Uint16Field("AmplifierSaturationOutputPower", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 9),
			10: ByteField("AmplifierNoiseFigure", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 10),
			11: ByteField("AmplifierSaturationGain", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, false, 11),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewReCommonAmplifierParameters (class ID 328) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewReCommonAmplifierParameters(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*recommonamplifierparametersBME, params...)
}
