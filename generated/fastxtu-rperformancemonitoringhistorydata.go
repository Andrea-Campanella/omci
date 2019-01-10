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

const FastXtuRPerformanceMonitoringHistoryDataClassId uint16 = 438

// FastXtuRPerformanceMonitoringHistoryData (class ID #438) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type FastXtuRPerformanceMonitoringHistoryData struct {
	BaseManagedEntityDefinition
}

// NewFastXtuRPerformanceMonitoringHistoryData (class ID 438 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastXtuRPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "FastXtuRPerformanceMonitoringHistoryData",
		ClassID:  438,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			1: ByteField("IntervalEndTime", 0, Read),
			2: Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write),
			3: Uint32Field("SuccessfulFraCounter", 0, Read),
			4: Uint32Field("SuccessfulRpaCounter", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &FastXtuRPerformanceMonitoringHistoryData{entity}, nil
}
