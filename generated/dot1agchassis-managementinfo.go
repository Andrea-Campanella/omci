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

// Dot1AgChassisManagementInfoClassID is the 16-bit ID for the OMCI
// Managed entity Dot1ag chassis-management info
const Dot1AgChassisManagementInfoClassID ClassID = ClassID(306)

var dot1agchassismanagementinfoBME *ManagedEntityDefinition

// Dot1AgChassisManagementInfo (class ID #306)
//	This ME represents the system-level chassis ID or management address for [IEEE-802.1ag] CFM
//	messages, and potentially for other IEEE 802-based functions. Although [IEEE-802.1AB] allows for
//	several management addresses (synonyms in different formats or with granularity to the component
//	level), [IEEE-802.1ag] does not provide for more than one. Nor is it expected that an ONU would
//	require more than one format. Accordingly, this ME provides for only one.
//
//	According to sender ID permission attributes in several dot1ag MEs, transmitted IEEE-802.1ag CFM
//	messages may include either or both of the chassis ID or management address
//	fields.[IEEE-802.1ag] requires that CCMs do not exceed 128-bytes, of which 74 are separately
//	allocated to other purposes; the sender ID TLV, if present, must accommodate this requirement.
//	The chassis info and management info must fit, with a minimum of 4 additional overhead bytes,
//	into the remaining 54-bytes. This limit is exploited in defining the maximum size of the ME's
//	attributes.
//
//	Relationships
//		If an ONU supports [IEEE 802.1ag] functionality, it automatically creates an instance of this
//		ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies this ME. There is at most one instance,
//			whose value is 0. (R) (mandatory) (2-bytes)
//
//		Chassis Id Length
//			Chassis ID length: The length of the chassis ID attribute (not including the chassis ID subtype
//			attribute), default value 0. (R,-W) (mandatory) (1-byte)
//
//		Chassis Id Subtype
//			(R,-W) (mandatory) (1-byte)
//
//		Chassis Id Part 1 Chassis Id Part 2
//			Chassis ID part 1, Chassis ID part 2: These two attributes may be regarded as an octet string of
//			up to 50-bytes whose length is given by the chassis ID length attribute and whose value is the
//			left-justified chassis ID. (R,-W) (mandatory) (25-bytes-*-2 attributes)
//
//		Management Address Domain Length
//			Management address domain length: The length of the management address domain attribute, default
//			value 0. If this attribute has the value 0, all of the other management address attributes are
//			undefined. (R,-W) (mandatory) (1-byte)
//
//		Management Address Domain 1, Management Address Domain 2
//			Management address domain 1, Management address domain 2: These two attributes may be regarded
//			as an octet string of up to 50-bytes whose length is given by the management address domain
//			length attribute and whose value is the left-justified management address domain. The attribute
//			is coded as an object identifier (OID) as per [ITUT X.690], referring to a TDomain as defined in
//			[IETF RFC 2579]. Typical domain values include snmpUDPDomain (from SNMPv2-TM [IETF RFC 3417])
//			and snmpIeee802Domain (from SNMP-IEEE 802-TM-MIB [IETF RFC 4789]). (R,-W) (mandatory) (25-bytes
//			* 2 attributes)
//
//		Management Address Length
//			Management address length: The length of the management address attribute, default value 0.
//			(R,-W) (mandatory) (1-byte)
//
//		Management Address 1 Management Address 2
//			Management address 1, Management address 2: These two attributes may be regarded as an octet
//			string of up to 50-bytes whose length is given by the management address length attribute and
//			whose value is the left-justified management address. (R,-W) (mandatory) (25-bytes * 2
//			attributes)
//
type Dot1AgChassisManagementInfo struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agchassismanagementinfoBME = &ManagedEntityDefinition{
		Name:    "Dot1AgChassisManagementInfo",
		ClassID: 306,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfe00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: ByteField("ChassisIdLength", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 1),
			2: ByteField("ChassisIdSubtype", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 2),
			3: MultiByteField("ChassisIdPart1ChassisIdPart2", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, false, 3),
			4: ByteField("ManagementAddressDomainLength", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 4),
			5: MultiByteField("ManagementAddressDomain1,ManagementAddressDomain2", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, false, 5),
			6: ByteField("ManagementAddressLength", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, false, 6),
			7: MultiByteField("ManagementAddress1ManagementAddress2", OctetsAttributeType, 25, nil, mapset.NewSetWith(Read, Write), false, false, false, 7),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewDot1AgChassisManagementInfo (class ID 306) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewDot1AgChassisManagementInfo(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1agchassismanagementinfoBME, params...)
}
