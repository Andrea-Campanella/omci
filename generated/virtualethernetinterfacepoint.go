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

const VirtualEthernetInterfacePointClassId uint16 = 329

// VirtualEthernetInterfacePoint (class ID #329) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type VirtualEthernetInterfacePoint struct {
	BaseManagedEntityDefinition
}

// NewVirtualEthernetInterfacePoint (class ID 329 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVirtualEthernetInterfacePoint(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "VirtualEthernetInterfacePoint",
		ClassID:  329,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read, false, false, false, false),
			1: ByteField("AdministrativeState", 0, Read|Write, false, false, false, false),
			2: ByteField("OperationalState", 0, Read, false, false, false, true),
			3: MultiByteField("InterdomainName", 25, nil, Read|Write, false, false, false, true),
			4: Uint16Field("TcpUdpPointer", 0, Read|Write, false, false, false, true),
			5: Uint16Field("IanaAssignedPort", 0, Read, false, false, false, false),
		},
	}
	entity.computeAttributeMask()
	return &VirtualEthernetInterfacePoint{entity}, nil
}
