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

const XgPonTcPerformanceMonitoringHistoryDataClassId ClassID = ClassID(344)

var xgpontcperformancemonitoringhistorydataBME *ManagedEntityDefinition

// XgPonTcPerformanceMonitoringHistoryData (class ID #344) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XgPonTcPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xgpontcperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "XgPonTcPerformanceMonitoringHistoryData",
		ClassID: 344,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false, 2),
			3:  Uint32Field("PsbdHecErrorCount", 0, Read, false, false, true, 3),
			4:  Uint32Field("XgtcHecErrorCount", 0, Read, false, false, true, 4),
			5:  Uint32Field("UnknownProfileCount", 0, Read, false, false, true, 5),
			6:  Uint32Field("TransmittedXgPonEncapsulationMethodXgemFrames", 0, Read, false, false, false, 6),
			7:  Uint32Field("FragmentXgemFrames", 0, Read, false, false, true, 7),
			8:  Uint32Field("XgemHecLostWordsCount", 0, Read, false, false, true, 8),
			9:  Uint32Field("XgemKeyErrors", 0, Read, false, false, false, 9),
			10: Uint32Field("XgemHecErrorCount", 0, Read, false, false, false, 10),
			11: Uint64Field("TransmittedBytesInNonIdleXgemFrames", 0, Read, false, false, false, 11),
			12: Uint64Field("ReceivedBytesInNonIdleXgemFrames", 0, Read, false, false, true, 12),
			13: Uint32Field("LossOfDownstreamSynchronizationLodsEventCount", 0, Read, false, false, true, 13),
			14: Uint32Field("LodsEventRestoredCount", 0, Read, false, false, true, 14),
			15: Uint32Field("OnuReactivationByLodsEvents", 0, Read, false, false, true, 15),
		},
	}
}

// NewXgPonTcPerformanceMonitoringHistoryData (class ID 344 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXgPonTcPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(xgpontcperformancemonitoringhistorydataBME, params...)
}
