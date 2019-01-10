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

const Ieee8021PMapperServiceProfileClassId uint16 = 130

// Ieee8021PMapperServiceProfile (class ID #130) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Ieee8021PMapperServiceProfile struct {
	BaseManagedEntityDefinition
}

// NewIeee8021PMapperServiceProfile (class ID 130 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewIeee8021PMapperServiceProfile(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "Ieee8021PMapperServiceProfile",
		ClassID:  130,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			1:  Uint16Field("TpPointer", 0, Read|SetByCreate|Write),
			2:  Uint16Field("InterworkTpPointerForPBitPriority0", 0, Read|SetByCreate|Write),
			3:  Uint16Field("InterworkTpPointerForPBitPriority1", 0, Read|SetByCreate|Write),
			4:  Uint16Field("InterworkTpPointerForPBitPriority2", 0, Read|SetByCreate|Write),
			5:  Uint16Field("InterworkTpPointerForPBitPriority3", 0, Read|SetByCreate|Write),
			6:  Uint16Field("InterworkTpPointerForPBitPriority4", 0, Read|SetByCreate|Write),
			7:  Uint16Field("InterworkTpPointerForPBitPriority5", 0, Read|SetByCreate|Write),
			8:  Uint16Field("InterworkTpPointerForPBitPriority6", 0, Read|SetByCreate|Write),
			9:  Uint16Field("InterworkTpPointerForPBitPriority7", 0, Read|SetByCreate|Write),
			10: ByteField("UnmarkedFrameOption", 0, Read|SetByCreate|Write),
			11: MultiByteField("DscpToPBitMapping", 24, nil, Read|Write),
			12: ByteField("DefaultPBitAssumption", 0, Read|SetByCreate|Write),
			13: ByteField("TpType", 0, Read|SetByCreate|Write),
		},
	}
	entity.computeAttributeMask()
	return &Ieee8021PMapperServiceProfile{entity}, nil
}
