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

const EthernetPerformanceMonitoringHistoryData3ClassId uint16 = 296

// EthernetPerformanceMonitoringHistoryData3 (class ID #296) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EthernetPerformanceMonitoringHistoryData3 struct {
	BaseManagedEntityDefinition
}

// NewEthernetPerformanceMonitoringHistoryData3 (class ID 296 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetPerformanceMonitoringHistoryData3(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EthernetPerformanceMonitoringHistoryData3",
		ClassID:  296,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, false),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false, false),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false, false),
			3:  Uint32Field("DropEvents", 0, Read, false, false, false, false),
			4:  Uint32Field("Octets", 0, Read, false, false, false, false),
			5:  Uint32Field("Packets", 0, Read, false, false, false, false),
			6:  Uint32Field("BroadcastPackets", 0, Read, false, false, false, false),
			7:  Uint32Field("MulticastPackets", 0, Read, false, false, false, false),
			8:  Uint32Field("UndersizePackets", 0, Read, false, false, false, false),
			9:  Uint32Field("Fragments", 0, Read, false, false, false, false),
			10: Uint32Field("Jabbers", 0, Read, false, false, false, false),
			11: Uint32Field("Packets64Octets", 0, Read, false, false, false, false),
			12: Uint32Field("Packets65To127Octets", 0, Read, false, false, false, false),
			13: Uint32Field("Packets128To255Octets", 0, Read, false, false, false, false),
			14: Uint32Field("Packets256To511Octets", 0, Read, false, false, false, false),
			15: Uint32Field("Packets512To1023Octets", 0, Read, false, false, false, false),
			16: Uint32Field("Packets1024To1518Octets", 0, Read, false, false, false, false),
		},
	}
	entity.computeAttributeMask()
	return &EthernetPerformanceMonitoringHistoryData3{entity}, nil
}
