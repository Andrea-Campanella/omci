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

const AniGClassId uint16 = 263

var anigBME *BaseManagedEntityDefinition

// AniG (class ID #263) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type AniG struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	anigBME = &BaseManagedEntityDefinition{
		Name:    "AniG",
		ClassID: 263,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
			Test,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false),
			1:  ByteField("SrIndication", 0, Read, false, false, false),
			2:  Uint16Field("TotalTcontNumber", 0, Read, false, false, false),
			3:  Uint16Field("GemBlockLength", 0, Read|Write, false, false, false),
			4:  ByteField("PiggybackDbaReporting", 0, Read, false, false, false),
			5:  ByteField("Deprecated", 0, Read, false, false, false),
			6:  ByteField("SignalFailThreshold", 0, Read|Write, false, false, false),
			7:  ByteField("SignalDegradeSdThreshold", 0, Read|Write, false, false, false),
			8:  ByteField("Arc", 0, Read|Write, true, false, true),
			9:  ByteField("ArcInterval", 0, Read|Write, false, false, true),
			10: Uint16Field("OpticalSignalLevel", 0, Read, false, false, true),
			11: ByteField("LowerOpticalThreshold", 0, Read|Write, false, false, true),
			12: ByteField("UpperOpticalThreshold", 0, Read|Write, false, false, true),
			13: Uint16Field("OnuResponseTime", 0, Read, false, false, true),
			14: Uint16Field("TransmitOpticalLevel", 0, Read, false, false, true),
			15: ByteField("LowerTransmitPowerThreshold", 0, Read|Write, false, false, true),
			16: ByteField("UpperTransmitPowerThreshold", 0, Read|Write, false, false, true),
		},
	}
}

// NewAniG (class ID 263 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewAniG(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: anigBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
