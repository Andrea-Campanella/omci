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

const VoipApplicationServiceProfileClassId uint16 = 146

// VoipApplicationServiceProfile (class ID #146) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type VoipApplicationServiceProfile struct {
	BaseManagedEntityDefinition
}

// NewVoipApplicationServiceProfile (class ID 146 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVoipApplicationServiceProfile(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "VoipApplicationServiceProfile",
		ClassID:  146,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			1: ByteField("CidFeatures", 0, Read|SetByCreate|Write),
			2: ByteField("CallWaitingFeatures", 0, Read|SetByCreate|Write),
			3: Uint16Field("CallProgressOrTransferFeatures", 0, Read|SetByCreate|Write),
			4: Uint16Field("CallPresentationFeatures", 0, Read|SetByCreate|Write),
			5: ByteField("DirectConnectFeature", 0, Read|SetByCreate|Write),
			6: Uint16Field("DirectConnectUriPointer", 0, Read|SetByCreate|Write),
			7: Uint16Field("BridgedLineAgentUriPointer", 0, Read|SetByCreate|Write),
			8: Uint16Field("ConferenceFactoryUriPointer", 0, Read|SetByCreate|Write),
			9: Uint16Field("DialToneFeatureDelayWArmlineTimerNew", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &VoipApplicationServiceProfile{entity}, nil
}
