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

const MacBridgePortIcmpv6ProcessPreAssignTableClassId ClassID = ClassID(348)

var macbridgeporticmpv6processpreassigntableBME *ManagedEntityDefinition

// MacBridgePortIcmpv6ProcessPreAssignTable (class ID #348) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MacBridgePortIcmpv6ProcessPreAssignTable struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	macbridgeporticmpv6processpreassigntableBME = &ManagedEntityDefinition{
		Name:    "MacBridgePortIcmpv6ProcessPreAssignTable",
		ClassID: 348,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0XFF80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read, false, false, false, 0),
			1: ByteField("Icmpv6ErrorMessagesProcessing", 0, Read|Write, false, false, false, 1),
			2: ByteField("Icmpv6InformationalMessagesProcessing", 0, Read|Write, false, false, false, 2),
			3: ByteField("RouterSolicitationProcessing", 0, Read|Write, false, false, false, 3),
			4: ByteField("RouterAdvertisementProcessing", 0, Read|Write, false, false, false, 4),
			5: ByteField("NeighbourSolicitationProcessing", 0, Read|Write, false, false, false, 5),
			6: ByteField("NeighbourAdvertisementProcessing", 0, Read|Write, false, false, false, 6),
			7: ByteField("RedirectProcessing", 0, Read|Write, false, false, false, 7),
			8: ByteField("MulticastListenerQueryProcessing", 0, Read|Write, false, false, false, 8),
			9: ByteField("UnknownIcmpv6Processing", 0, Read|Write, false, false, false, 9),
		},
	}
}

// NewMacBridgePortIcmpv6ProcessPreAssignTable (class ID 348 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMacBridgePortIcmpv6ProcessPreAssignTable(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(macbridgeporticmpv6processpreassigntableBME, params...)
}
