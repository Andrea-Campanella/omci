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

const EthernetFramePerformanceMonitoringHistoryDataDownstreamClassId uint16 = 321

// EthernetFramePerformanceMonitoringHistoryDataDownstream (class ID #321) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EthernetFramePerformanceMonitoringHistoryDataDownstream struct {
	BaseManagedEntityDefinition
}

// NewEthernetFramePerformanceMonitoringHistoryDataDownstream (class ID 321 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFramePerformanceMonitoringHistoryDataDownstream(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EthernetFramePerformanceMonitoringHistoryDataDownstream",
		ClassID:  321,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			1:  ByteField("IntervalEndTime", 0, Read),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write),
			3:  Uint32Field("DropEvents", 0, Read),
			4:  Uint32Field("Octets", 0, Read),
			5:  Uint32Field("Packets", 0, Read),
			6:  Uint32Field("BroadcastPackets", 0, Read),
			7:  Uint32Field("MulticastPackets", 0, Read),
			8:  Uint32Field("CrcErroredPackets", 0, Read),
			9:  Uint32Field("UndersizePackets", 0, Read),
			10: Uint32Field("OversizePackets", 0, Read),
			11: Uint32Field("Packets64Octets", 0, Read),
			12: Uint32Field("Packets65To127Octets", 0, Read),
			13: Uint32Field("Packets128To255Octets", 0, Read),
			14: Uint32Field("Packets256To511Octets", 0, Read),
			15: Uint32Field("Packets512To1023Octets", 0, Read),
			16: Uint32Field("Packets1024To1518Octets", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &EthernetFramePerformanceMonitoringHistoryDataDownstream{entity}, nil
}
