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

const FastLineInventoryAndStatusDataPart2ClassId uint16 = 436

var fastlineinventoryandstatusdatapart2BME *BaseManagedEntityDefinition

// FastLineInventoryAndStatusDataPart2 (class ID #436) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type FastLineInventoryAndStatusDataPart2 struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	fastlineinventoryandstatusdatapart2BME = &BaseManagedEntityDefinition{
		Name:    "FastLineInventoryAndStatusDataPart2",
		ClassID: 436,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0XF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read, false, false, false),
			1: MultiByteField("DateTimeStampingOfLastSuccessfulDownstreamFraOperationStampFrads", 7, nil, Read, false, false, true),
			2: MultiByteField("DateTimeStampingOfLastSuccessfulUpstreamFraOperationStampFraus", 7, nil, Read, false, false, true),
			3: MultiByteField("DateTimeStampingOfLastSuccessfulDownstreamRpaOperationStampRpads", 7, nil, Read, false, false, true),
			4: MultiByteField("DateTimeStampingOfLastSuccessfulUpstreamRpaOperationStampRpaus", 7, nil, Read, false, false, true),
			5: MultiByteField("DateTimeStampingOfLastSuccessfulDownstreamTigaOperationStampTiga", 7, nil, Read, false, false, true),
		},
	}
}

// NewFastLineInventoryAndStatusDataPart2 (class ID 436 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastLineInventoryAndStatusDataPart2(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: fastlineinventoryandstatusdatapart2BME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
