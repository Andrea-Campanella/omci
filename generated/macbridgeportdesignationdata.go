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

// MacBridgePortDesignationDataClassID is the 16-bit ID for the OMCI
// Managed entity MAC bridge port designation data
const MacBridgePortDesignationDataClassID ClassID = ClassID(48)

var macbridgeportdesignationdataBME *ManagedEntityDefinition

// MacBridgePortDesignationData (class ID #48)
//	This ME records data associated with a bridge port. The ONU automatically creates or deletes an
//	instance of this managed entity upon the creation or deletion of a MAC bridge port configuration
//	data ME.
//
//	Relationships
//		An instance of this managed entity is associated with one MAC bridge port configuration data ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the MAC bridge port configuration
//			data. (R) (mandatory) (2-bytes)
//
//		Designated Bridge Root Cost Port
//			Upon ME instantiation, the ONU sets this attribute to 0. (R) (mandatory) (24-bytes)
//
//		Port State
//			The value (R)stp_off is introduced to denote the port status where the (rapid) spanning tree
//			protocol has been disabled by setting the port spanning tree ind attribute of the MAC bridge
//			port configuration data to false, and the Ethernet link state is up. This value distinguishes
//			whether frame forwarding is under the control of (R)STP.
//
type MacBridgePortDesignationData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	macbridgeportdesignationdataBME = &ManagedEntityDefinition{
		Name:    "MacBridgePortDesignationData",
		ClassID: 48,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0xC000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1: MultiByteField("DesignatedBridgeRootCostPort", 24, nil, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: ByteField("PortState", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewMacBridgePortDesignationData (class ID 48 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMacBridgePortDesignationData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*macbridgeportdesignationdataBME, params...)
}
