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

// TwdmChannelOmciPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity TWDM channel OMCI performance monitoring history data
const TwdmChannelOmciPerformanceMonitoringHistoryDataClassID ClassID = ClassID(452)

var twdmchannelomciperformancemonitoringhistorydataBME *ManagedEntityDefinition

// TwdmChannelOmciPerformanceMonitoringHistoryData (class ID #452)
//	This ME collects OMCI-related PM data associated with the slot/circuit pack, hosting one or more
//	ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted by the OLT.
//
//	The counters maintained by this ME are characterized as optional in Clause 14 of [ITU-
//	T-G.989.3].
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of TWDM channel ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the TWDM channel ME. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contains PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Omci Baseline Message Count
//			OMCI baseline message count: The counter of baseline format OMCI messages directed to the given
//			ONU. (R) (mandatory) (4-byte)
//
//		Omci Extended Message Count
//			OMCI extended message count: The counter of extended format OMCI messages directed to the given
//			ONU. (R) (mandatory) (4-byte)
//
//		Omci Mic Error Count
//			OMCI MIC error count: The counter of OMCI messages received with MIC errors. (R) (mandatory)
//			(4-byte)
//
type TwdmChannelOmciPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	twdmchannelomciperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "TwdmChannelOmciPerformanceMonitoringHistoryData",
		ClassID: 452,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetCurrentData,
			Set,
		),
		AllowedAttributeMask: 0XF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("OmciBaselineMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("OmciExtendedMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("OmciMicErrorCount", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
		},
	}
}

// NewTwdmChannelOmciPerformanceMonitoringHistoryData (class ID 452 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTwdmChannelOmciPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*twdmchannelomciperformancemonitoringhistorydataBME, params...)
}
