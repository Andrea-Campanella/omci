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

// PwAtmConfigurationDataClassID is the 16-bit ID for the OMCI
// Managed entity PW ATM configuration data
const PwAtmConfigurationDataClassID ClassID = ClassID(337)

var pwatmconfigurationdataBME *ManagedEntityDefinition

// PwAtmConfigurationData (class ID #337)
//	This ME contains generic configuration data for an ATM pseudowire. Definitions of attributes are
//	from PW-ATM-MIB [IETF RFC 5605]. Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME is associated with an instance of the MPLS pseudowire TP ME with a
//		pseudowire type attribute equal to one of the following.////		2	ATM AAL5 SDU VCC transport////		3	ATM transparent cell transport////		9	ATM n-to-one VCC cell transport////		10	ATM n-to-one VPC cell transport////		12	ATM one-to-one VCC cell mode////		13	ATM one-to-one VPC cell mode////		14	ATM AAL5 PDU VCC transport////		Alternatively, an instance of this ME may be associated with an Ethernet flow TP or a TCP/UDP
//		config data ME, depending on the transport layer of the pseudowire.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R,
//			setbycreate)-(mandatory) (2 bytes)
//
//		Tp Type
//			2	TCP/UDP config data
//
//		Transport Tp Pointer
//			Transport TP pointer: This attribute points to an associated instance of the transport layer TP,
//			whose type is specified by the TP type attribute. (R, W, setbycreate) (mandatory) (2 bytes)
//
//		Pptp Atm Uni Pointer
//			PPTP ATM UNI pointer: This attribute points to an associated instance of the ITU-T G.983.2 PPTP
//			ATM UNI. Refer to [ITUT G.983.2] for the definition of the target ME. (R, W, setbycreate)
//			(mandatory) (2 bytes)
//
//		Max C Ell C Oncatenation
//			Max cell concatenation: This attribute specifies the maximum number of ATM cells that can be
//			concatenated into one PW packet in the upstream direction. (R, W, setbycreate) (mandatory) (2
//			bytes)
//
//		Far End M Ax C Ell C Oncatenation
//			Far-end max cell concatenation: This attribute specifies the maximum number of ATM cells that
//			can be concatenated into one PW packet as provisioned at the far end. This attribute may be used
//			for error checking of downstream traffic. The value 0 specifies that the ONU uses its internal
//			default. (R, W, set-by-create) (optional) (2 bytes)
//
//		Atm Cell Loss Priority Clp Qos Mapping
//			The value 0 specifies that the ONU uses its internal default. (R, W, setbycreate) (optional) (1
//			byte)
//
//		Timeout Mode
//			The value 0 specifies that the ONU uses its internal default. (R, W, setbycreate) (optional) (1
//			byte)
//
//		Pw Atm Mapping Table
//			(R,-W) (mandatory) (21N bytes, where N is the number of entries in the list)
//
type PwAtmConfigurationData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	pwatmconfigurationdataBME = &ManagedEntityDefinition{
		Name:    "PwAtmConfigurationData",
		ClassID: 337,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: ByteField("TpType", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2: Uint16Field("TransportTpPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3: Uint16Field("PptpAtmUniPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 3),
			4: Uint16Field("MaxCEllCOncatenation", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 4),
			5: Uint16Field("FarEndMAxCEllCOncatenation", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 5),
			6: ByteField("AtmCellLossPriorityClpQosMapping", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 6),
			7: ByteField("TimeoutMode", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 7),
			8: TableField("PwAtmMappingTable", TableAttributeType, TableInfo{nil, 21}, mapset.NewSetWith(Read, Write), false, false, false, 8),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewPwAtmConfigurationData (class ID 337) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewPwAtmConfigurationData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*pwatmconfigurationdataBME, params...)
}
