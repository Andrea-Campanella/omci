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

const ReDownstreamAmplifierClassId ClassID = ClassID(316)

var redownstreamamplifierBME *ManagedEntityDefinition

// ReDownstreamAmplifier (class ID #316) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type ReDownstreamAmplifier struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	redownstreamamplifierBME = &ManagedEntityDefinition{
		Name:    "ReDownstreamAmplifier",
		ClassID: 316,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
			Test,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, 0),
			1:  ByteField("AdministrativeState", 0, Read|Write, false, false, false, 1),
			2:  ByteField("OperationalState", 0, Read, true, false, true, 2),
			3:  ByteField("Arc", 0, Read|Write, true, false, true, 3),
			4:  ByteField("ArcInterval", 0, Read|Write, false, false, true, 4),
			5:  ByteField("OperationalMode", 0, Read|Write, false, false, false, 5),
			6:  Uint16Field("InputOpticalSignalLevel", 0, Read, false, false, true, 6),
			7:  ByteField("LowerInputOpticalThreshold", 0, Read|Write, false, false, true, 7),
			8:  ByteField("UpperInputOpticalThreshold", 0, Read|Write, false, false, true, 8),
			9:  Uint16Field("OutputOpticalSignalLevel", 0, Read, false, false, true, 9),
			10: ByteField("LowerOutputOpticalThreshold", 0, Read|Write, false, false, true, 10),
			11: ByteField("UpperOutputOpticalThreshold", 0, Read|Write, false, false, true, 11),
			12: ByteField("R'S'SplitterCouplingRatio", 0, Read, false, false, true, 12),
		},
	}
}

// NewReDownstreamAmplifier (class ID 316 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewReDownstreamAmplifier(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(redownstreamamplifierBME, params...)
}
