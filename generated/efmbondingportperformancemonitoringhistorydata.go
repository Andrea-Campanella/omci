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

// EfmBondingPortPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity EFM bonding port performance monitoring history data
const EfmBondingPortPerformanceMonitoringHistoryDataClassID ClassID = ClassID(424)

var efmbondingportperformancemonitoringhistorydataBME *ManagedEntityDefinition

// EfmBondingPortPerformanceMonitoringHistoryData (class ID #424)
//	This ME collects PM data as seen at the xTU-C. Instances of this ME are created and deleted by
//	the OLT.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The two MSBs of
//			the first byte are the bearer channel ID. Excluding the first 2-bits of the first byte, the
//			remaining part of the ME ID is identical to that of this ME's parent PPTP xDSL UNI part 1. (R,
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
//		Rx Frames
//			Rx frames: Number of Ethernet frames received over this port. (R) (mandatory) (4-bytes)
//
//		Tx Frames
//			Tx frames: Number of Ethernet frames transmitted over this port. (R) (mandatory) (4-bytes)
//
//		Rx Bytes
//			Rx bytes: Number of bytes contained in the Ethernet frames received over this port. (R)
//			(mandatory) (4-bytes)
//
//		Tx Bytes
//			Tx bytes: Number of bytes contained in the Ethernet frames transmitted over this port. (R)
//			(mandatory) (4-bytes)
//
//		Tx Discarded Frames
//			Tx discarded frames: Number of Ethernet frames discarded by the port transmit function. (R)
//			(mandatory) (4-bytes)
//
//		Tx Discarded Bytes
//			Tx discarded bytes: Number of bytes contained in the Ethernet frames discarded by the port
//			transmit function. (R) (mandatory) (4-bytes)
//
type EfmBondingPortPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	efmbondingportperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "EfmBondingPortPerformanceMonitoringHistoryData",
		ClassID: 424,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xFF00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("RxFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("TxFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("RxBytes", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6: Uint32Field("TxBytes", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7: Uint32Field("TxDiscardedFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8: Uint32Field("TxDiscardedBytes", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewEfmBondingPortPerformanceMonitoringHistoryData (class ID 424 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEfmBondingPortPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*efmbondingportperformancemonitoringhistorydataBME, params...)
}
