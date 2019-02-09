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

const OnuGClassId uint16 = 256

var onugBME *BaseManagedEntityDefinition

// OnuG (class ID #256) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type OnuG struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	onugBME = &BaseManagedEntityDefinition{
		Name:    "OnuG",
		ClassID: 256,
		MessageTypes: mapset.NewSetWith(
			Get,
			Reboot,
			Set,
			SynchronizeTime,
			Test,
		),
		AllowedAttributeMask: 0XFFF8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false),
			1:  Uint32Field("VendorId", 0, Read, false, false, false),
			2:  MultiByteField("Version", 14, nil, Read, false, false, false),
			3:  Uint64Field("SerialNumber", 0, Read, false, false, false),
			4:  ByteField("TrafficManagementOption", 0, Read, false, false, false),
			5:  ByteField("Deprecated", 0, Read, false, false, true),
			6:  ByteField("BatteryBackup", 0, Read|Write, false, false, false),
			7:  ByteField("AdministrativeState", 0, Read|Write, false, false, false),
			8:  ByteField("OperationalState", 0, Read, false, false, true),
			9:  ByteField("OnuSurvivalTime", 0, Read, false, false, true),
			10: MultiByteField("LogicalOnuId", 24, nil, Read, false, false, true),
			11: MultiByteField("LogicalPassword", 12, nil, Read, false, false, true),
			12: ByteField("CredentialsStatus", 0, Read|Write, false, false, true),
			13: Uint16Field("ExtendedTcLayerOptions", 0, Read, false, false, true),
		},
	}
}

// NewOnuG (class ID 256 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewOnuG(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: onugBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
