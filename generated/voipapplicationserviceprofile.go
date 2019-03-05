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

const VoipApplicationServiceProfileClassId ClassID = ClassID(146)

var voipapplicationserviceprofileBME *ManagedEntityDefinition

// VoipApplicationServiceProfile (class ID #146) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type VoipApplicationServiceProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	voipapplicationserviceprofileBME = &ManagedEntityDefinition{
		Name:    "VoipApplicationServiceProfile",
		ClassID: 146,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, 0),
			1: ByteField("CidFeatures", 0, Read|SetByCreate|Write, false, false, false, 1),
			2: ByteField("CallWaitingFeatures", 0, Read|SetByCreate|Write, false, false, false, 2),
			3: Uint16Field("CallProgressOrTransferFeatures", 0, Read|SetByCreate|Write, false, false, false, 3),
			4: Uint16Field("CallPresentationFeatures", 0, Read|SetByCreate|Write, false, false, false, 4),
			5: ByteField("DirectConnectFeature", 0, Read|SetByCreate|Write, false, false, false, 5),
			6: Uint16Field("DirectConnectUriPointer", 0, Read|SetByCreate|Write, false, false, false, 6),
			7: Uint16Field("BridgedLineAgentUriPointer", 0, Read|SetByCreate|Write, false, false, false, 7),
			8: Uint16Field("ConferenceFactoryUriPointer", 0, Read|SetByCreate|Write, false, false, false, 8),
			9: Uint16Field("DialToneFeatureDelayWArmlineTimerNew", 0, Read|Write, false, false, true, 9),
		},
	}
}

// NewVoipApplicationServiceProfile (class ID 146 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVoipApplicationServiceProfile(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(voipapplicationserviceprofileBME, params...)
}
