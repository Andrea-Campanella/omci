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

// XdslXtuCPerformanceMonitoringHistoryDataPart2ClassID is the 16-bit ID for the OMCI
// Managed entity xDSL xTU-C performance monitoring history data part 2
const XdslXtuCPerformanceMonitoringHistoryDataPart2ClassID ClassID = ClassID(408)

var xdslxtucperformancemonitoringhistorydatapart2BME *ManagedEntityDefinition

// XdslXtuCPerformanceMonitoringHistoryDataPart2 (class ID #408)
//	This ME collects PM data on the xTUC to xTUR path as seen from the xTU-C. Instances of this ME
//	are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contain PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Leftr Defect Seconds
//			"leftr" defect seconds: If retransmission is used, this attribute is a count of the seconds with
//			a near-end ''leftr'' defect present - see clause 7.2.1.1.6 of [ITU-T G.997.1]. (R) (mandatory)
//			(2-bytes)
//
//		Error_Free Bits Counter
//			Error-free bits counter: If retransmission is used, this attribute is a count of the number of
//			error-free bits passed over the B1 reference point, divided by 216 - see clause-7.2.1.1.7 of
//			[ITU-T G.997.1]. (R) (mandatory) (4-bytes)
//
//		Minimum Error_Free Throughput Mineftr
//			Minimum error-free throughput (MINEFTR): If retransmission is used, this attribute is the
//			minimum error-free throughput in bits per second - see clause 7.2.1.1.8 of [ITUT-G.997.1]. (R)
//			(mandatory) (4-bytes)
//
type XdslXtuCPerformanceMonitoringHistoryDataPart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslxtucperformancemonitoringhistorydatapart2BME = &ManagedEntityDefinition{
		Name:    "XdslXtuCPerformanceMonitoringHistoryDataPart2",
		ClassID: 408,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("LeftrDefectSeconds", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("ErrorFreeBitsCounter", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("MinimumErrorFreeThroughputMineftr", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
		},
	}
}

// NewXdslXtuCPerformanceMonitoringHistoryDataPart2 (class ID 408 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslXtuCPerformanceMonitoringHistoryDataPart2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslxtucperformancemonitoringhistorydatapart2BME, params...)
}
