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

import "fmt"

// ManagedEntityInfo provides ManagedEntity information
type ManagedEntityInfo struct {
	New func(params ...ParamData) (*ManagedEntity, error)
}

// ParamData can be passed to the 'New' function to dictate how the returned
// Managed Entity is created. You should supply either zero or one ParamData
// structure to 'New'.
//
// If No ParamData is passed, the returned Managed Entity can only be used for
// providing validation of other structures. This is commonly done in a packet
// encoder/decoder to assist in that process.
//
// If One ParamData is passed, the returned Managed Entity will be initialized
// with the given values/attributes and then validated. This is commonly done
// when you wish to create an ME for transmission, storage or removal from a
// persistent database, or some other similar purpose.
//
type ParamData struct {
	EntityID   uint16
	Attributes AttributeValueMap
}

// CreateME wraps a function that makes it a creator of a Managed Entity
type CreateME func(params ...ParamData) (*ManagedEntity, OmciErrors)

var classToManagedEntityMap map[ClassID]CreateME

func init() {
	// Create mapping of 16-bit managed entity class IDs to ME-type
	classToManagedEntityMap = make(map[ClassID]CreateME, 160)

	classToManagedEntityMap[2] = NewOnuData
	classToManagedEntityMap[5] = NewCardholder
	classToManagedEntityMap[6] = NewCircuitPack
	classToManagedEntityMap[7] = NewSoftwareImage
	classToManagedEntityMap[11] = NewPhysicalPathTerminationPointEthernetUni
	classToManagedEntityMap[12] = NewPhysicalPathTerminationPointCesUni
	classToManagedEntityMap[14] = NewInterworkingVccTerminationPoint
	classToManagedEntityMap[16] = NewAal5Profile
	classToManagedEntityMap[18] = NewAal5PerformanceMonitoringHistoryData
	classToManagedEntityMap[21] = NewCesServiceProfile
	classToManagedEntityMap[24] = NewEthernetPerformanceMonitoringHistoryData
	classToManagedEntityMap[45] = NewMacBridgeServiceProfile
	classToManagedEntityMap[46] = NewMacBridgeConfigurationData
	classToManagedEntityMap[47] = NewMacBridgePortConfigurationData
	classToManagedEntityMap[48] = NewMacBridgePortDesignationData
	classToManagedEntityMap[49] = NewMacBridgePortFilterTableData
	classToManagedEntityMap[50] = NewMacBridgePortBridgeTableData
	classToManagedEntityMap[51] = NewMacBridgePerformanceMonitoringHistoryData
	classToManagedEntityMap[52] = NewMacBridgePortPerformanceMonitoringHistoryData
	classToManagedEntityMap[53] = NewPhysicalPathTerminationPointPotsUni
	classToManagedEntityMap[58] = NewVoiceServiceProfile
	classToManagedEntityMap[62] = NewVpPerformanceMonitoringHistoryData
	classToManagedEntityMap[78] = NewVlanTaggingOperationConfigurationData
	classToManagedEntityMap[79] = NewMacBridgePortFilterPreAssignTable
	classToManagedEntityMap[82] = NewPhysicalPathTerminationPointVideoUni
	classToManagedEntityMap[83] = NewPhysicalPathTerminationPointLctUni
	classToManagedEntityMap[84] = NewVlanTaggingFilterData
	classToManagedEntityMap[89] = NewEthernetPerformanceMonitoringHistoryData2
	classToManagedEntityMap[90] = NewPhysicalPathTerminationPointVideoAni
	classToManagedEntityMap[98] = NewPhysicalPathTerminationPointXdslUniPart1
	classToManagedEntityMap[99] = NewPhysicalPathTerminationPointXdslUniPart2
	classToManagedEntityMap[100] = NewXdslLineInventoryAndStatusDataPart1
	classToManagedEntityMap[101] = NewXdslLineInventoryAndStatusDataPart2
	classToManagedEntityMap[102] = NewXdslChannelDownstreamStatusData
	classToManagedEntityMap[103] = NewXdslChannelUpstreamStatusData
	classToManagedEntityMap[105] = NewXdslLineConfigurationProfilePart2
	classToManagedEntityMap[106] = NewXdslLineConfigurationProfilePart3
	classToManagedEntityMap[107] = NewXdslChannelConfigurationProfile
	classToManagedEntityMap[108] = NewXdslSubcarrierMaskingDownstreamProfile
	classToManagedEntityMap[109] = NewXdslSubcarrierMaskingUpstreamProfile
	classToManagedEntityMap[112] = NewXdslXtuCPerformanceMonitoringHistoryData
	classToManagedEntityMap[113] = NewXdslXtuRPerformanceMonitoringHistoryData
	classToManagedEntityMap[114] = NewXdslXtuCChannelPerformanceMonitoringHistoryData
	classToManagedEntityMap[115] = NewXdslXtuRChannelPerformanceMonitoringHistoryData
	classToManagedEntityMap[116] = NewTcAdaptorPerformanceMonitoringHistoryDataXdsl
	classToManagedEntityMap[130] = NewIeee8021PMapperServiceProfile
	classToManagedEntityMap[131] = NewOltG
	classToManagedEntityMap[133] = NewOnuPowerShedding
	classToManagedEntityMap[134] = NewIpHostConfigData
	classToManagedEntityMap[135] = NewIpHostPerformanceMonitoringHistoryData
	classToManagedEntityMap[136] = NewTcpUdpConfigData
	classToManagedEntityMap[137] = NewNetworkAddress
	classToManagedEntityMap[138] = NewVoipConfigData
	classToManagedEntityMap[139] = NewVoipVoiceCtp
	classToManagedEntityMap[140] = NewCallControlPerformanceMonitoringHistoryData
	classToManagedEntityMap[141] = NewVoipLineStatus
	classToManagedEntityMap[142] = NewVoipMediaProfile
	classToManagedEntityMap[143] = NewRtpProfileData
	classToManagedEntityMap[144] = NewRtpPerformanceMonitoringHistoryData
	classToManagedEntityMap[145] = NewNetworkDialPlanTable
	classToManagedEntityMap[146] = NewVoipApplicationServiceProfile
	classToManagedEntityMap[147] = NewVoipFeatureAccessCodes
	classToManagedEntityMap[148] = NewAuthenticationSecurityMethod
	classToManagedEntityMap[150] = NewSipAgentConfigData
	classToManagedEntityMap[151] = NewSipAgentPerformanceMonitoringHistoryData
	classToManagedEntityMap[152] = NewSipCallInitiationPerformanceMonitoringHistoryData
	classToManagedEntityMap[153] = NewSipUserData
	classToManagedEntityMap[155] = NewMgcConfigData
	classToManagedEntityMap[156] = NewMgcPerformanceMonitoringHistoryData
	classToManagedEntityMap[160] = NewEquipmentExtensionPackage
	classToManagedEntityMap[162] = NewPhysicalPathTerminationPointMocaUni
	classToManagedEntityMap[163] = NewMocaEthernetPerformanceMonitoringHistoryData
	classToManagedEntityMap[171] = NewExtendedVlanTaggingOperationConfigurationData
	classToManagedEntityMap[256] = NewOnuG
	classToManagedEntityMap[257] = NewOnu2G
	classToManagedEntityMap[262] = NewTCont
	classToManagedEntityMap[263] = NewAniG
	classToManagedEntityMap[264] = NewUniG
	classToManagedEntityMap[266] = NewGemInterworkingTerminationPoint
	classToManagedEntityMap[268] = NewGemPortNetworkCtp
	classToManagedEntityMap[269] = NewVpNetworkCtp
	classToManagedEntityMap[272] = NewGalEthernetProfile
	classToManagedEntityMap[273] = NewThresholdData1
	classToManagedEntityMap[274] = NewThresholdData2
	classToManagedEntityMap[276] = NewGalEthernetPerformanceMonitoringHistoryData
	classToManagedEntityMap[277] = NewPriorityQueue
	classToManagedEntityMap[278] = NewTrafficScheduler
	classToManagedEntityMap[280] = NewTrafficDescriptor
	classToManagedEntityMap[281] = NewMulticastGemInterworkingTerminationPoint
	classToManagedEntityMap[282] = NewPseudowireTerminationPoint
	classToManagedEntityMap[283] = NewRtpPseudowireParameters
	classToManagedEntityMap[284] = NewPseudowireMaintenanceProfile
	classToManagedEntityMap[285] = NewPseudowirePerformanceMonitoringHistoryData
	classToManagedEntityMap[286] = NewEthernetFlowTerminationPoint
	classToManagedEntityMap[287] = NewOmci
	classToManagedEntityMap[288] = NewManagedEntityMe
	classToManagedEntityMap[289] = NewAttributeMe
	classToManagedEntityMap[290] = NewDot1XPortExtensionPackage
	classToManagedEntityMap[291] = NewDot1XConfigurationProfile
	classToManagedEntityMap[292] = NewDot1XPerformanceMonitoringHistoryData
	classToManagedEntityMap[293] = NewRadiusPerformanceMonitoringHistoryData
	classToManagedEntityMap[296] = NewEthernetPerformanceMonitoringHistoryData3
	classToManagedEntityMap[298] = NewDot1RateLimiter
	classToManagedEntityMap[299] = NewDot1AgMaintenanceDomain
	classToManagedEntityMap[300] = NewDot1AgMaintenanceAssociation
	classToManagedEntityMap[301] = NewDot1AgDefaultMdLevel
	classToManagedEntityMap[302] = NewDot1AgMep
	classToManagedEntityMap[305] = NewDot1AgCfmStack
	classToManagedEntityMap[310] = NewMulticastSubscriberConfigInfo
	classToManagedEntityMap[311] = NewMulticastSubscriberMonitor
	classToManagedEntityMap[313] = NewReAniG
	classToManagedEntityMap[316] = NewReDownstreamAmplifier
	classToManagedEntityMap[321] = NewEthernetFramePerformanceMonitoringHistoryDataDownstream
	classToManagedEntityMap[322] = NewEthernetFramePerformanceMonitoringHistoryDataUpstream
	classToManagedEntityMap[325] = NewXdslLineInventoryAndStatusDataPart5
	classToManagedEntityMap[328] = NewReCommonAmplifierParameters
	classToManagedEntityMap[329] = NewVirtualEthernetInterfacePoint
	classToManagedEntityMap[332] = NewEnhancedSecurityControl
	classToManagedEntityMap[333] = NewMplsPseudowireTerminationPoint
	classToManagedEntityMap[334] = NewEthernetFrameExtendedPm
	classToManagedEntityMap[335] = NewSnmpConfigurationData
	classToManagedEntityMap[336] = NewOnuDynamicPowerManagementControl
	classToManagedEntityMap[338] = NewPwAtmPerformanceMonitoringHistoryData
	classToManagedEntityMap[339] = NewPwEthernetConfigurationData
	classToManagedEntityMap[340] = NewBbfTr069ManagementServer
	classToManagedEntityMap[341] = NewGemPortNetworkCtpPerformanceMonitoringHistoryData
	classToManagedEntityMap[342] = NewTcpUdpPerformanceMonitoringHistoryData
	classToManagedEntityMap[343] = NewEnergyConsumptionPerformanceMonitoringHistoryData
	classToManagedEntityMap[344] = NewXgPonTcPerformanceMonitoringHistoryData
	classToManagedEntityMap[345] = NewXgPonDownstreamManagementPerformanceMonitoringHistoryData
	classToManagedEntityMap[346] = NewXgPonUpstreamManagementPerformanceMonitoringHistoryData
	classToManagedEntityMap[348] = NewMacBridgePortIcmpv6ProcessPreAssignTable
	classToManagedEntityMap[400] = NewEthernetPseudowireParameters
	classToManagedEntityMap[408] = NewXdslXtuCPerformanceMonitoringHistoryDataPart2
	classToManagedEntityMap[410] = NewVdsl2LineConfigurationExtensions3
	classToManagedEntityMap[412] = NewXdslChannelConfigurationProfilePart2
	classToManagedEntityMap[414] = NewXdslLineInventoryAndStatusDataPart8
	classToManagedEntityMap[420] = NewEfmBondingLink
	classToManagedEntityMap[421] = NewEfmBondingGroupPerformanceMonitoringHistoryData
	classToManagedEntityMap[422] = NewEfmBondingGroupPerformanceMonitoringHistoryDataPart2
	classToManagedEntityMap[423] = NewEfmBondingLinkPerformanceMonitoringHistoryData
	classToManagedEntityMap[424] = NewEfmBondingPortPerformanceMonitoringHistoryData
	classToManagedEntityMap[425] = NewEfmBondingPortPerformanceMonitoringHistoryDataPart2
	classToManagedEntityMap[426] = NewEthernetFrameExtendedPm64Bit
	classToManagedEntityMap[432] = NewFastChannelConfigurationProfile
	classToManagedEntityMap[433] = NewFastDataPathConfigurationProfile
	classToManagedEntityMap[434] = NewFastVectoringLineConfigurationExtensions
	classToManagedEntityMap[436] = NewFastLineInventoryAndStatusDataPart2
	classToManagedEntityMap[437] = NewFastXtuCPerformanceMonitoringHistoryData
	classToManagedEntityMap[438] = NewFastXtuRPerformanceMonitoringHistoryData
	classToManagedEntityMap[443] = NewTwdmChannelManagedEntity
	classToManagedEntityMap[444] = NewTwdmChannelPhyLodsPerformanceMonitoringHistoryData
	classToManagedEntityMap[445] = NewTwdmChannelXgemPerformanceMonitoringHistoryData
	classToManagedEntityMap[446] = NewTwdmChannelPloamPerformanceMonitoringHistoryDataPart1
	classToManagedEntityMap[447] = NewTwdmChannelPloamPerformanceMonitoringHistoryDataPart2
	classToManagedEntityMap[448] = NewTwdmChannelPloamPerformanceMonitoringHistoryDataPart3
	classToManagedEntityMap[449] = NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart1
	classToManagedEntityMap[450] = NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart2
	classToManagedEntityMap[451] = NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart3
	classToManagedEntityMap[452] = NewTwdmChannelOmciPerformanceMonitoringHistoryData
}

// LoadManagedEntityDefinition returns a function to create a Managed Entity for a specific
// Managed Entity class ID
func LoadManagedEntityDefinition(classID ClassID, params ...ParamData) (*ManagedEntity, OmciErrors) {
	newFunc, ok := classToManagedEntityMap[classID]
	if ok {
		return newFunc(params...)
	}
	return nil, NewUnknownEntityError(fmt.Sprintf("managed entity %d (%#x) definition not found",
		uint16(classID), uint16(classID)))
}

// GetSupportedClassIDs returns an array of Managed Entity Class IDs supported
func GetSupportedClassIDs() []ClassID {
	supported := make([]ClassID, 0, len(classToManagedEntityMap))
	for k := range classToManagedEntityMap {
		supported = append(supported, k)
	}
	return supported
}
