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

const MgcConfigDataClassId ClassID = ClassID(155)

var mgcconfigdataBME *ManagedEntityDefinition

// MgcConfigData (class ID #155) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MgcConfigData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	mgcconfigdataBME = &ManagedEntityDefinition{
		Name:    "MgcConfigData",
		ClassID: 155,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFE0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, 0),
			1:  Uint16Field("PrimaryMgc", 0, Read|SetByCreate|Write, false, false, false, 1),
			2:  Uint16Field("SecondaryMgc", 0, Read|SetByCreate|Write, false, false, false, 2),
			3:  Uint16Field("TcpUdpPointer", 0, Read|SetByCreate|Write, false, false, false, 3),
			4:  ByteField("Version", 0, Read|SetByCreate|Write, false, false, false, 4),
			5:  ByteField("MessageFormat", 0, Read|SetByCreate|Write, false, false, false, 5),
			6:  Uint16Field("MaximumRetryTime", 0, Read|Write, false, false, true, 6),
			7:  Uint16Field("MaximumRetryAttempts", 0, Read|SetByCreate|Write, false, false, true, 7),
			8:  Uint16Field("ServiceChangeDelay", 0, Read|Write, false, false, true, 8),
			9:  MultiByteField("TerminationIdBase", 25, nil, Read|Write, false, false, true, 9),
			10: Uint32Field("Softswitch", 0, Read|SetByCreate|Write, false, false, false, 10),
			11: Uint16Field("MessageIdPointer", 0, Read|SetByCreate|Write, false, false, true, 11),
		},
	}
}

// NewMgcConfigData (class ID 155 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMgcConfigData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(mgcconfigdataBME, params...)
}
