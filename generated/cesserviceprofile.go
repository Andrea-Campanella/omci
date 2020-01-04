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

// CesServiceProfileClassID is the 16-bit ID for the OMCI
// Managed entity CES service profile
const CesServiceProfileClassID ClassID = ClassID(21)

var cesserviceprofileBME *ManagedEntityDefinition

// CesServiceProfile (class ID #21)
//	NOTE - In [ITU-T G.984.4], this ME is called a CES service profile-G.
//
//	An instance of this ME organizes data that describe the CES service functions of the ONU.
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of a GEM IW TP.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Ces Buffered Cdv Tolerance
//			CES buffered CDV tolerance: This attribute represents the duration of user data that must be
//			buffered by the CES IW entity to offset packet delay variation. It is expressed in 10-us
//			increments. 75 (750-vs) is suggested as a default value. (R,-W, setbycreate) (mandatory)
//			(2-bytes)
//
//		Channel Associated Signalling Cas
//			(R,-W, setbycreate) (optional) (1-byte)
//
type CesServiceProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	cesserviceprofileBME = &ManagedEntityDefinition{
		Name:    "CesServiceProfile",
		ClassID: 21,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: Uint16Field("CesBufferedCdvTolerance", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2: ByteField("ChannelAssociatedSignallingCas", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 2),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewCesServiceProfile (class ID 21) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewCesServiceProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*cesserviceprofileBME, params...)
}
