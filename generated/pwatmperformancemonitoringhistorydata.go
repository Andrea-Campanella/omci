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

const PwAtmPerformanceMonitoringHistoryDataClassId ClassID = ClassID(338)

var pwatmperformancemonitoringhistorydataBME *ManagedEntityDefinition

// PwAtmPerformanceMonitoringHistoryData (class ID #338) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PwAtmPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	pwatmperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "PwAtmPerformanceMonitoringHistoryData",
		ClassID: 338,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, Read, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false, 2),
			3: Uint32Field("DownstreamMissingPacketsCounter", 0, Read, false, false, false, 3),
			4: Uint32Field("DownstreamReorderedPacketsCounter", 0, Read, false, false, true, 4),
			5: Uint32Field("DownstreamMisorderedPacketsCounter", 0, Read, false, false, false, 5),
			6: Uint32Field("UpstreamTimeoutPacketsCounter", 0, Read, false, false, false, 6),
			7: Uint32Field("UpstreamTransmittedCellsCounter", 0, Read, false, false, false, 7),
			8: Uint32Field("UpstreamDroppedCellsCounter", 0, Read, false, false, false, 8),
			9: Uint32Field("UpstreamReceivedCellsCounter", 0, Read, false, false, false, 9),
		},
	}
}

// NewPwAtmPerformanceMonitoringHistoryData (class ID 338 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPwAtmPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(pwatmperformancemonitoringhistorydataBME, params...)
}
