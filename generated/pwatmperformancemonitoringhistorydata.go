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

// PwAtmPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity PW ATM performance monitoring history data
const PwAtmPerformanceMonitoringHistoryDataClassID ClassID = ClassID(338)

var pwatmperformancemonitoringhistorydataBME *ManagedEntityDefinition

// PwAtmPerformanceMonitoringHistoryData (class ID #338)
//	This ME collects PM data associated with an ATM pseudowire. Instances of this ME are created and
//	deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of the PW ATM configuration data ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to the instance of the PW ATM configuration data ME.
//			(R, setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Downstream Missing Packets Counter
//			Downstream missing packets counter: This attribute counts missing packets, as detected via
//			control word sequence number gaps. (R) (mandatory) (4-bytes)
//
//		Downstream Reordered Packets Counter
//			Downstream reordered packets counter: This attribute counts packets detected out of sequence via
//			the control word sequence number, but successfully reordered. Some implementations may not
//			support this feature. (R) (optional) (4-bytes)
//
//		Downstream Misordered Packets Counter
//			Downstream misordered packets counter: This attribute counts packets detected out of order via
//			the control word sequence numbers. (R) (mandatory) (4-bytes)
//
//		Upstream Timeout Packets Counter
//			Upstream timeout packets counter: This attribute counts packets transmitted due to timeout
//			expiration while attempting to collect cells. (R) (mandatory) (4-bytes)
//
//		Upstream Transmitted Cells Counter
//			Upstream transmitted cells counter: This attribute counts transmitted cells. (R) (mandatory)
//			(4-bytes)
//
//		Upstream Dropped Cells Counter
//			Upstream dropped cells counter: This attribute counts dropped cells. (R) (mandatory) (4-bytes)
//
//		Upstream Received Cells Counter
//			Upstream received cells counter: This attribute counts received cells. (R) (mandatory) (4-bytes)
//
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
		AllowedAttributeMask: 0xff80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("DownstreamMissingPacketsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 3),
			4: Uint32Field("DownstreamReorderedPacketsCounter", 0, mapset.NewSetWith(Read), false, true, true, false, 4),
			5: Uint32Field("DownstreamMisorderedPacketsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 5),
			6: Uint32Field("UpstreamTimeoutPacketsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 6),
			7: Uint32Field("UpstreamTransmittedCellsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 7),
			8: Uint32Field("UpstreamDroppedCellsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 8),
			9: Uint32Field("UpstreamReceivedCellsCounter", 0, mapset.NewSetWith(Read), false, true, false, false, 9),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewPwAtmPerformanceMonitoringHistoryData (class ID 338) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewPwAtmPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*pwatmperformancemonitoringhistorydataBME, params...)
}
