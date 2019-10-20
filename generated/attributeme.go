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

// AttributeMeClassID is the 16-bit ID for the OMCI
// Managed entity Attribute ME
const AttributeMeClassID ClassID = ClassID(289)

var attributemeBME *ManagedEntityDefinition

// AttributeMe (class ID #289)
//	This ME describes a particular attribute type that is supported by the ONU. This ME is not
//	included in an MIB upload.
//
//	Relationships
//		One or more attribute entities are related to each ME entity. More than one ME entity can refer
//		to a given attribute entity.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. This number is
//			the same as the one that appears in the attributes table in the ME. Only one instance of each
//			unique attribute need be created. The ONU can assign attribute numbering as it pleases, out of
//			the pool of 64K IDs; however, it is suggested that the numbering follow a rational scheme to aid
//			human readability. (R) (mandatory) (2-bytes)
//
//		Name
//			Name:	This attribute contains a 25-byte mnemonic tag for the attribute. Strings shorter than
//			25-bytes are padded with null characters. (R) (mandatory) (25-bytes)
//
//		Size
//			Size:	This attribute contains the size of the attribute, in bytes. The value 0 indicates that
//			the attribute can have a variable/unknown size. (R) (mandatory) (2-bytes)
//
//		Access
//			(R) (mandatory) (1-byte)
//
//		Format
//			(R) (mandatory) (1-byte)
//
//		Lower Limit
//			Lower limit:	This attribute provides the lowest value for the attribute. Valid for numeric types
//			(pointer, signed integer, unsigned integer) only. For attributes smaller than 4-bytes, the
//			desired numeric value is expressed in 4-byte representation (for example, the 2s complement
//			1-byte integer 0xFE is expressed as 0xFFFF-FFFE; the unsigned 1-byte integer 0xFE is expressed
//			as 0x0000-00FE). (R) (mandatory) (4-bytes)
//
//		Upper Limit
//			Upper limit:	This attribute provides the highest value for the attribute. It has the same
//			validity and format as the lower limit attribute. (R) (mandatory) (4-bytes)
//
//		Bit Field
//			Bit field:	This attribute is a mask of the supported bits in a bit field attribute, valid for
//			bit field type only. A 1 in any position signifies that its code point is supported, while 0
//			indicates that it is not supported. For bit fields smaller than 4-bytes, the attribute is
//			aligned at the least significant end of the mask. (R) (mandatory) (4-bytes)
//
//		Code Points Table
//			Code points table: This attribute lists the code points supported by an enumerated attribute.
//			(R) (mandatory) (2 * Q bytes, where Q is the number of entries in the table.)
//
//		Support
//			(R) (mandatory) (1-byte)
//
type AttributeMe struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	attributemeBME = &ManagedEntityDefinition{
		Name:    "AttributeMe",
		ClassID: 289,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0XFF80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1: MultiByteField("Name", 25, nil, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("Size", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3: ByteField("Access", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: ByteField("Format", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("LowerLimit", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6: Uint32Field("UpperLimit", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7: Uint32Field("BitField", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8: TableField("CodePointsTable", TableInfo{0, 2}, mapset.NewSetWith(Read), false, false, false, 8),
			9: ByteField("Support", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
		},
	}
}

// NewAttributeMe (class ID 289 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewAttributeMe(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*attributemeBME, params...)
}
