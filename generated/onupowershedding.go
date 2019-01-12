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

const OnuPowerSheddingClassId uint16 = 133

// OnuPowerShedding (class ID #133) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type OnuPowerShedding struct {
	BaseManagedEntityDefinition
}

// NewOnuPowerShedding (class ID 133 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewOnuPowerShedding(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "OnuPowerShedding",
		ClassID:  133,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, false),
			1:  Uint16Field("RestorePowerTimerResetInterval", 0, Read|Write, false, false, false, false),
			2:  Uint16Field("DataClassSheddingInterval", 0, Read|Write, false, false, false, false),
			3:  Uint16Field("VoiceClassSheddingInterval", 0, Read|Write, false, false, false, false),
			4:  Uint16Field("VideoOverlayClassSheddingInterval", 0, Read|Write, false, false, false, false),
			5:  Uint16Field("VideoReturnClassSheddingInterval", 0, Read|Write, false, false, false, false),
			6:  Uint16Field("DigitalSubscriberLineClassSheddingInterval", 0, Read|Write, false, false, false, false),
			7:  Uint16Field("AtmClassSheddingInterval", 0, Read|Write, false, false, false, false),
			8:  Uint16Field("CesClassSheddingInterval", 0, Read|Write, false, false, false, false),
			9:  Uint16Field("FrameClassSheddingInterval", 0, Read|Write, false, false, false, false),
			10: Uint16Field("SdhSonetClassSheddingInterval", 0, Read|Write, false, false, false, false),
			11: Uint16Field("SheddingStatus", 0, Read, false, false, false, true),
		},
	}
	entity.computeAttributeMask()
	return &OnuPowerShedding{entity}, nil
}
