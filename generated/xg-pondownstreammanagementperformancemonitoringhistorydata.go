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

// XgPonDownstreamManagementPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity XG-PON downstream management performance monitoring history data
const XgPonDownstreamManagementPerformanceMonitoringHistoryDataClassID ClassID = ClassID(345)

var xgpondownstreammanagementperformancemonitoringhistorydataBME *ManagedEntityDefinition

// XgPonDownstreamManagementPerformanceMonitoringHistoryData (class ID #345)
//	This ME collects PM data associated with the XG-PON TC layer. It collects counters associated
//	with downstream PLOAM and OMCI messages.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an ANI-G.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the ANI-G. (R, set-by-create)
//			(mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, set-by-create) (mandatory) (2-bytes)
//
//		Ploam Message Integrity Check Mic Error Count
//			PLOAM message integrity check (MIC) error count: This attribute counts MIC errors detected in
//			downstream PLOAM messages, either directed to this ONU or broadcast to all ONUs. (R) (optional)
//			(4-bytes)
//
//		Downstream Ploam Messages Count
//			Downstream PLOAM messages count: This attribute counts PLOAM messages received, either directed
//			to this ONU or broadcast to all ONUs. (R) (optional) (4-bytes)
//
//		Profile Messages Received
//			Profile messages received: This attribute counts the number of profile messages received, either
//			directed to this ONU or broadcast to all ONUs. In [ITU-T G.9807.1], this attribute is used for
//			received burst_profile message count.  (R) (optional) (4-bytes)
//
//		Ranging_Time Messages Received
//			Ranging_time messages received: This attribute counts the number of ranging_time messages
//			received, either directed to this ONU or broadcast to all ONUs. (R) (mandatory) (4-bytes)
//
//		Deactivate_Onu_Id Messages Received
//			Deactivate_ONU-ID messages received: This attribute counts the number of deactivate_ONU-ID
//			messages received, either directed to this ONU or broadcast to all ONUs. Deactivate_ONU-ID
//			messages do not reset this counter. (R) (optional) (4-bytes)
//
//		Disable_Serial_Number Messages Received
//			Disable_serial_number messages received: This attribute counts the number of
//			disable_serial_number messages received, whose serial number specified this ONU. (R) (optional)
//			(4-bytes)
//
//		Request_Registration Messages Received
//			Request_registration messages received: This attribute counts the number of request_registration
//			messages received. (R) (optional) (4-bytes)
//
//		Assign_Alloc_Id Messages Received
//			Assign_alloc-ID messages received: This attribute counts the number of assign_alloc-ID messages
//			received. (R) (optional) (4-bytes)
//
//		Key_Control Messages Received
//			Key_control messages received: This attribute counts the number of key_control messages
//			received, either directed to this ONU or broadcast to all ONUs. (R) (optional) (4-bytes)
//
//		Sleep_Allow Messages Received
//			Sleep_allow messages received: This attribute counts the number of sleep_allow messages
//			received, either directed to this ONU or broadcast to all ONUs. (R) (optional) (4-bytes)
//
//		Baseline Omci Messages Received Count
//			Baseline OMCI messages received count: This attribute counts the number of OMCI messages
//			received in the baseline message format. (R) (optional) (4-bytes)
//
//		Extended Omci Messages Received Count
//			Extended OMCI messages received count: This attribute counts the number of OMCI messages
//			received in the extended message format. (R) (optional) (4-bytes)
//
//		Assign_Onu_Id Messages Received
//			Assign_ONU-ID messages received: This attribute counts the number of assign_ONU-ID messages
//			received since the last re-boot. (R) (optional) (4-bytes)
//
//		Omci Mic Error Count
//			OMCI MIC error count: This attribute counts MIC errors detected in OMCI messages directed to
//			this ONU. (R) (optional) (4-bytes)
//
type XgPonDownstreamManagementPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xgpondownstreammanagementperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "XgPonDownstreamManagementPerformanceMonitoringHistoryData",
		ClassID: 345,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint32Field("PloamMessageIntegrityCheckMicErrorCount", 0, mapset.NewSetWith(Read), false, false, true, false, 3),
			4:  Uint32Field("DownstreamPloamMessagesCount", 0, mapset.NewSetWith(Read), false, false, true, false, 4),
			5:  Uint32Field("ProfileMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 5),
			6:  Uint32Field("RangingTimeMessagesReceived", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  Uint32Field("DeactivateOnuIdMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 7),
			8:  Uint32Field("DisableSerialNumberMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 8),
			9:  Uint32Field("RequestRegistrationMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 9),
			10: Uint32Field("AssignAllocIdMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 10),
			11: Uint32Field("KeyControlMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 11),
			12: Uint32Field("SleepAllowMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 12),
			13: Uint32Field("BaselineOmciMessagesReceivedCount", 0, mapset.NewSetWith(Read), false, false, true, false, 13),
			14: Uint32Field("ExtendedOmciMessagesReceivedCount", 0, mapset.NewSetWith(Read), false, false, true, false, 14),
			15: Uint32Field("AssignOnuIdMessagesReceived", 0, mapset.NewSetWith(Read), false, false, true, false, 15),
			16: Uint32Field("OmciMicErrorCount", 0, mapset.NewSetWith(Read), false, false, true, false, 16),
		},
	}
}

// NewXgPonDownstreamManagementPerformanceMonitoringHistoryData (class ID 345 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXgPonDownstreamManagementPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xgpondownstreammanagementperformancemonitoringhistorydataBME, params...)
}
