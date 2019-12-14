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

// XdslImpulseNoiseMonitorPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity xDSL impulse noise monitor performance monitoring history data
const XdslImpulseNoiseMonitorPerformanceMonitoringHistoryDataClassID ClassID = ClassID(324)

var xdslimpulsenoisemonitorperformancemonitoringhistorydataBME *ManagedEntityDefinition

// XdslImpulseNoiseMonitorPerformanceMonitoringHistoryData (class ID #324)
//	This ME collects PM data from the impulse noise monitor function at both near and far ends.
//	Instances of this ME are created and deleted by the OLT. Note that, unlike most xDSL PM, [ITU-T
//	G.997.1] only requires current and previous 15-min interval storage; a longer view of this PM is
//	not expected at 15-min granularity.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME may be associated with an xDSL UNI. This ME is meaningful only for ITUT
//		G.993.2 VDSL2, [ITUT G.992.3] and [ITUT-G.992.5].
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The ME ID is
//			identical to that of this ME's parent PPTP xDSL UNI part 1. (R, setbycreate) (mandatory)
//			(2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: No thresholds are defined for this ME. For uniformity with other PM, the
//			attribute is retained and shown as mandatory, but it should be set to a null pointer. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Inm Inpeq Histogram Table
//			INM INPEQ histogram table: INMINPEQ1..17-L is a count of the near-end INMAINPEQi anomalies
//			occurring on the line during the accumulation period. This parameter is subject to inhibiting -
//			see clause 7.2.7.13 of [ITUT-G.997.1]. (R) (optional) (2-bytes * 17 entries-= 34-bytes)
//
//		Inm Total Measurement
//			INM total measurement: INMME-L is a count of the near-end INMAME anomalies occurring on the line
//			during the accumulation period. This parameter is subject to inhibiting - see clause 7.2.7.13 of
//			[ITUT G.997.1]. (R) (optional) (2-bytes)
//
//		Inm Iat Histogram
//			INM IAT histogram: INMIAT0..7-L is a count of the near-end INMAIATi anomalies occurring on the
//			line during the accumulation period. This parameter is subject to inhibiting - see clause
//			7.2.7.13 of [ITUT G.997.1]. (R) (optional) (2-bytes-* 8 entries-= 16-bytes)
//
//		Inm Inpeq Histogram Lfe Table
//			INM INPEQ histogram LFE table: INMINPEQ1..17-LFE is a count of the far-end INMAINPEQi anomalies
//			occurring on the line during the accumulation period. This parameter is subject to inhibiting -
//			see clause 7.2.7.13 of [ITUT-G.997.1]. (R) (optional) (2-bytes * 17 entries-= 34-bytes)
//
//		Inm Total Measurement Lfe
//			INM total measurement LFE: INMME-LFE is a count of the far-end INMAME anomalies occurring on the
//			line during the accumulation period. This parameter is subject to inhibiting - see clause
//			7.2.7.13 of [ITUT G.997.1]. (R) (optional) (2-bytes)
//
//		Inm Iat Histogram Lfe
//			INM IAT histogram LFE: INMIAT0..7-LFE is a count of the far-end INMAIATi anomalies occurring on
//			the line during the accumulation period. This parameter is subject to inhibiting - see clause
//			7.2.7.13 of [ITUT G.997.1]. (R) (optional) (2-bytes-* 8 entries-= 16-bytes)
//
type XdslImpulseNoiseMonitorPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslimpulsenoisemonitorperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "XdslImpulseNoiseMonitorPerformanceMonitoringHistoryData",
		ClassID: 324,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: TableField("InmInpeqHistogramTable", TableInfo{0, 2}, mapset.NewSetWith(Read), false, true, false, 3),
			4: Uint16Field("InmTotalMeasurement", 0, mapset.NewSetWith(Read), false, false, true, false, 4),
			5: Uint16Field("InmIatHistogram", 0, mapset.NewSetWith(Read), false, false, true, false, 5),
			6: TableField("InmInpeqHistogramLfeTable", TableInfo{0, 2}, mapset.NewSetWith(Read), false, true, false, 6),
			7: Uint16Field("InmTotalMeasurementLfe", 0, mapset.NewSetWith(Read), false, false, true, false, 7),
			8: Uint16Field("InmIatHistogramLfe", 0, mapset.NewSetWith(Read), false, false, true, false, 8),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewXdslImpulseNoiseMonitorPerformanceMonitoringHistoryData (class ID 324) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslImpulseNoiseMonitorPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslimpulsenoisemonitorperformancemonitoringhistorydataBME, params...)
}
