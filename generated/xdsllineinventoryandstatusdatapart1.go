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

const XdslLineInventoryAndStatusDataPart1ClassId uint16 = 100

// XdslLineInventoryAndStatusDataPart1 (class ID #100) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslLineInventoryAndStatusDataPart1 struct {
	BaseManagedEntityDefinition
}

// NewXdslLineInventoryAndStatusDataPart1 (class ID 100 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslLineInventoryAndStatusDataPart1(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslLineInventoryAndStatusDataPart1",
		ClassID:  100,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, false),
			1:  Uint64Field("XtuCG9941VendorId", 0, Read, false, false, false, false),
			2:  Uint64Field("XtuRG9941VendorId", 0, Read, false, false, false, false),
			3:  Uint64Field("XtuCSystemVendorId", 0, Read, false, false, false, false),
			4:  Uint64Field("XtuRSystemVendorId", 0, Read, false, false, false, false),
			5:  MultiByteField("XtuCVersionNumber", 16, nil, Read, false, false, false, false),
			6:  MultiByteField("XtuRVersionNumber", 16, nil, Read, false, false, false, false),
			7:  MultiByteField("XtuCSerialNumberPart1", 16, nil, Read, false, false, false, false),
			8:  MultiByteField("XtuCSerialNumberPart2", 16, nil, Read, false, false, false, false),
			9:  MultiByteField("XtuRSerialNumberPart1", 16, nil, Read, false, false, false, false),
			10: MultiByteField("XtuRSerialNumberPart2", 16, nil, Read, false, false, false, false),
			11: Uint32Field("XtuCSelfTestResults", 0, Read, false, false, false, false),
			12: Uint32Field("XtuRSelfTestResults", 0, Read, false, false, false, false),
			13: MultiByteField("XtuCTransmissionSystemCapability", 7, nil, Read, false, false, false, false),
			14: MultiByteField("XtuRTransmissionSystemCapability", 7, nil, Read, false, false, false, false),
			15: ByteField("InitializationSuccessFailureCause", 0, Read, false, false, false, false),
		},
	}
	entity.computeAttributeMask()
	return &XdslLineInventoryAndStatusDataPart1{entity}, nil
}
