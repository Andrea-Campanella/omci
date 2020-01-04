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

// OltGClassID is the 16-bit ID for the OMCI
// Managed entity OLT-G
const OltGClassID ClassID = ClassID(131)

var oltgBME *ManagedEntityDefinition

// OltG (class ID #131)
//	This optional ME identifies the OLT to which an ONU is connected. This ME provides a way for the
//	ONU to configure itself for operability with a particular OLT. It also provides a way for the
//	OLT to communicate the time of day to the ONU.
//
//	An ONU that supports this ME automatically creates an instance of it. Immediately following the
//	start-up phase, the OLT should set the ONU to the desired configuration. Interpretation of the
//	OLT vendor ID, equipment ID and version attributes is a matter for negotiation between the two
//	vendors involved.
//
//	Relationships
//		The single instance of this ME is associated with the ONU ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only
//			one instance, number 0. (R) (mandatory) (2-bytes)
//
//		Olt Vendor Id
//			OLT vendor ID: This attribute identifies the OLT vendor. It is the same as the four most
//			significant bytes of an ONU serial number specified in the respective TC layer specification.
//			Upon instantiation, this attribute comprises all spaces. (R,-W) (mandatory) (4-bytes)
//
//		Equipment Id
//			Equipment ID: This attribute may be used to identify the specific type of OLT. The default value
//			of all spaces indicates that equipment ID information is not available or applicable to the OLT
//			being represented. (R,-W) (mandatory) (20-bytes)
//
//		Version
//			Version:	This attribute identifies the version of the OLT as defined by the vendor. The default
//			left-justified ASCII string "0" (padded with trailing nulls) indicates that version information
//			is not available or applicable to the OLT being represented. (R,-W) (mandatory) (14-bytes)
//
//		Time Of Day Information
//			NOTE - In ITU-T G.987/ITU-T G.989 systems, the superframe count field of the time of day
//			information attribute contains the 32 LSBs of the actual counter.
//
type OltG struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	oltgBME = &ManagedEntityDefinition{
		Name:    "OltG",
		ClassID: 131,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: MultiByteField("OltVendorId", StringAttributeType, 4, toOctets("ICAgIA=="), mapset.NewSetWith(Read, Write), false, false, false, 1),
			2: MultiByteField("EquipmentId", StringAttributeType, 20, toOctets("ICAgICAgICAgICAgICAgICAgICA="), mapset.NewSetWith(Read, Write), false, false, false, 2),
			3: MultiByteField("Version", StringAttributeType, 14, toOctets("MAAAAAAAAAAAAAAAAAA="), mapset.NewSetWith(Read, Write), false, false, false, 3),
			4: MultiByteField("TimeOfDayInformation", OctetsAttributeType, 14, toOctets("AAAAAAAAAAAAAAAAAAA="), mapset.NewSetWith(Read, Write), false, true, false, 4),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewOltG (class ID 131) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewOltG(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*oltgBME, params...)
}
