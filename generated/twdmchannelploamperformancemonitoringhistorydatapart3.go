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

// TwdmChannelPloamPerformanceMonitoringHistoryDataPart3ClassID is the 16-bit ID for the OMCI
// Managed entity TWDM channel PLOAM performance monitoring history data part 3
const TwdmChannelPloamPerformanceMonitoringHistoryDataPart3ClassID ClassID = ClassID(448)

var twdmchannelploamperformancemonitoringhistorydatapart3BME *ManagedEntityDefinition

// TwdmChannelPloamPerformanceMonitoringHistoryDataPart3 (class ID #448)
//	This ME collects remaining PLOAM-related PM data associated with the slot/circuit pack, hosting
//	one or more ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted
//	by the OLT.
//
//	This ME contains the counters related to transmitted upstream PLOAM messages. All these counters
//	are characterized as optional in clause 14 of [ITU-T- G.989.3].
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
//		Upstream Ploam Message Count
//			Upstream PLOAM message count: The aggregate counter of PLOAM messages, other than AK PLOAM MT,
//			transmitted by the given ONU. (R) (mandatory) (4-byte)
//
//		Serial_Number_Onu In_Band Message Count
//			Serial_Number_ONU (in-band) message count: The counter of transmitted in-band Serial_Number_ONU
//			PLOAM messages. (R) (mandatory) (4-byte)
//
//		Serial_Number_Onu Amcc Message Count
//			Serial_Number_ONU (AMCC) message count: The counter of transmitted auxiliary management and
//			control channel (AMCC) Serial_Number_ONU PLOAM messages. (R) (mandatory) (4-byte)
//
//		Registration Message Count
//			Registration message count: The counter of transmitted Registration PLOAM messages. (R)
//			(mandatory) (4-byte)
//
//		Key_Report Message Count
//			Key_Report message count: The counter of transmitted Key_Report PLOAM messages. (R) (mandatory)
//			(4-byte)
//
//		Acknowledgement Message Count
//			Acknowledgement message count: The counter of transmitted Registration PLOAM messages. (R)
//			(mandatory) (4-byte)
//
//		Sleep_Request Message Count
//			Sleep_Request message count: The counter of transmitted Sleep_Request PLOAM messages. (R)
//			(mandatory) (4-byte)
//
//		Tuning_Response Ack_Nack Message Count
//			Tuning_Response (ACK/NACK) message count: The counter of transmitted Tuning_Response PLOAM
//			messages with ACK/NACK operation code. (R) (mandatory) (4-byte)
//
//		Tuning_Response Complete_U_Rollback Message Count
//			Tuning_Response (Complete_u/Rollback) message count: The counter of transmitted Tuning_Response
//			PLOAM messages with Complete_u/Rollback operation code. (R) (mandatory) (4-byte)
//
//		Power_Consumption_Report Message Count
//			Power_Consumption_Report message count: The counter of transmitted Power_Consumption_Report
//			PLOAM messages. (R) (mandatory) (4-byte)
//
//		Change_Power_Level Parameter Error Count
//			Change_Power_Level parameter error count: The counter of transmitted Acknowledgement PLOAM
//			messages with Parameter Error completion code in response to Change_Power_Level PLOAM message.
//			(R) (mandatory) (4-byte)
//
type TwdmChannelPloamPerformanceMonitoringHistoryDataPart3 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	twdmchannelploamperformancemonitoringhistorydatapart3BME = &ManagedEntityDefinition{
		Name:    "TwdmChannelPloamPerformanceMonitoringHistoryDataPart3",
		ClassID: 448,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetCurrentData,
			Set,
		),
		AllowedAttributeMask: 0XFFF8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint32Field("UpstreamPloamMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  Uint32Field("SerialNumberOnuInBandMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  Uint32Field("SerialNumberOnuAmccMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  Uint32Field("RegistrationMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  Uint32Field("KeyReportMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  Uint32Field("AcknowledgementMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
			9:  Uint32Field("SleepRequestMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: Uint32Field("TuningResponseAckNackMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: Uint32Field("TuningResponseCompleteURollbackMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 11),
			12: Uint32Field("PowerConsumptionReportMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 12),
			13: Uint32Field("ChangePowerLevelParameterErrorCount", 0, mapset.NewSetWith(Read), false, false, false, false, 13),
		},
	}
}

// NewTwdmChannelPloamPerformanceMonitoringHistoryDataPart3 (class ID 448 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTwdmChannelPloamPerformanceMonitoringHistoryDataPart3(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*twdmchannelploamperformancemonitoringhistorydatapart3BME, params...)
}
