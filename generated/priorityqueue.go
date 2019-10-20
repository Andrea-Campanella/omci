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

// PriorityQueueClassID is the 16-bit ID for the OMCI
// Managed entity Priority queue
const PriorityQueueClassID ClassID = ClassID(277)

var priorityqueueBME *ManagedEntityDefinition

// PriorityQueue (class ID #277)
//	NOTE 1 - In [ITU-T G.984.4], this is called a priority queue-G.
//
//	This ME specifies the priority queue used by a GEM port network CTP in the upstream direction.
//	The upstream priority queue ME is also related to a T-CONT ME. By default, this relationship is
//	fixed by the ONU hardware architecture, but some ONUs may also permit the relationship to be
//	configured through the OMCI, as indicated by the QoS configuration flexibility attribute of the
//	ONU2G ME.
//
//	In the downstream direction, priority queues are associated with UNIs. Again, the association is
//	fixed by default, but some ONUs may permit the association to be configured through the OMCI.
//
//	If an ONU as a whole contains priority queues, it instantiates these queues autonomously.
//	Priority queues may also be localized to pluggable circuit packs, in which case the ONU creates
//	and deletes them in accordance with circuit pack pre-provisioning and the equipped
//	configuration.
//
//	The OLT can find all the queues by reading the priority queue ME instances. If the OLT tries to
//	retrieve a non-existent priority queue, the ONU denies the get action with an error indication.
//
//	See also Appendix II.
//
//	Priority queues can exist in the ONU core and circuit packs serving both UNI and ANI functions.
//	Therefore, they can be indirectly created and destroyed through cardholder provisioning actions.
//
//	In the upstream direction, the weight attribute permits the configuring of an optional traffic
//	scheduler. Several attributes support back pressure operation, whereby a back-pressure signal is
//	sent backwards and causes the attached terminal to temporarily suspend sending data.
//
//	In the downstream direction, strict priority discipline among the queues serving a given UNI is
//	the default, with priorities established through the related port attribute. If two or more non-
//	empty queues have the same priority, capacity is allocated among them in proportion to their
//	weights. Note that the details of the downstream model differ from those of the upstream model.
//
//	The yellow packet drop thresholds specify the drop probability for a packet that has been marked
//	yellow (drop eligible) by a traffic descriptor or by external equipment such as a residential
//	gateway (RG). If the current average queue occupancy is less than the minimum threshold, the
//	yellow packet drop probability is zero. If the current average queue occupancy is greater than
//	or equal to the maximum threshold, the yellow packet drop probability is one. The yellow drop
//	probability increases linearly between 0 and max_p as the current average queue occupancy
//	increases from the minimum to the maximum threshold.
//
//	The same model can be configured for green packets, those regarded as being within the traffic
//	contract.
//
//	Drop precedence colour marking indicates the method by which a packet is marked as drop eligible
//	(yellow). For discard eligibility indicator (DEI) and priority code point (PCP) marking, a drop
//	eligible indicator is equivalent to yellow colour; otherwise, the colour is green. For
//	differentiated services code point (DSCP) assured forwarding (AF) marking, the lowest drop
//	precedence is equivalent to green; otherwise, the colour is yellow.
//
//	Relationships
//		One or more instances of this ME are associated with the ONU-G ME to model upstream priority
//		queues if the traffic management option attribute in the ONU-G ME is 0 or 2.////		One or more instances of this ME are associated with a PPTP UNI ME as downstream priority
//		queues. Downstream priority queues may or may not be provided for a virtual Ethernet interface
//		point (VEIP).
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The MSB
//			represents the direction (1: upstream, 0:-downstream). The 15 LSBs represent a queue ID. The
//			queue ID is numbered in ascending order by the ONU itself. It is strongly encouraged that the
//			queue ID be formulated to simplify finding related queues. One way to do this is to number the
//			queues such that the related port attributes are in ascending order (for the downstream and
//			upstream queues separately). The range of downstream queue ids is 0 to 0x7FFF and the range of
//			upstream queue ids is 0x8000 to 0xFFFF. (R) (mandatory) (2-bytes)
//
//		Queue Configuration Option
//			Queue configuration option: This attribute identifies the buffer partitioning policy. The value
//			1 means that several queues share one buffer of maximum queue size, while the value 0 means that
//			each queue has an individual buffer of maximum queue size. (R) (mandatory) (1-byte)
//
//		Maximum Queue Size
//			NOTE 2 - In this and the other similar attributes of the priority queue ME, some legacy
//			implementations may take the queue scale factor from the GEM block length attribute of the ANI-G
//			ME. This option is discouraged in new implementations.
//
//		Allocated Queue Size
//			Allocated queue size: This attribute identifies the allocated size of this queue, in bytes,
//			scaled by the priority queue scale factor attribute of the ONU2G. (R, W) (mandatory) (2 bytes)
//
//		Discard_Block Counter Reset Interval
//			Discard-block counter reset interval: This attribute represents the interval in milliseconds at
//			which the counter resets itself. (R,-W) (optional) (2-bytes)
//
//		Threshold Value For Discarded Blocks Due To Buffer Overflow
//			Threshold value for discarded blocks due to buffer overflow: This attribute specifies the
//			threshold for the number of bytes (scaled by the priority queue scale factor attribute of the
//			ONU2G) discarded on this queue due to buffer overflow. Its value controls the declaration of the
//			block loss alarm. (R, W) (optional) (2-bytes)
//
//		Related Port
//			If flexible configuration is not supported, the ONU should reject an attempt to set the related
//			port with a parameter error result-reason code.
//
//		Traffic Scheduler Pointer
//			The ONU should reject an attempt to violate these conditions with a parameter error result-
//			reason code.
//
//		Weight
//			Weight:	This attribute represents weight for WRR scheduling. At a given priority level, capacity
//			is distributed to non-empty queues in proportion to their weights. In the upstream direction,
//			this weight is meaningful if several priority queues are associated with a traffic scheduler or
//			T-CONT whose policy is WRR. In the downstream direction, this weight is used by a UNI in a WRR
//			fashion. Upon ME instantiation, the ONU sets this attribute to 1. (R,-W) (mandatory) (1-byte)
//
//		Back Pressure Operation
//			Back pressure operation: This attribute enables (0) or disables (1) back pressure operation. Its
//			default value is 0. (R,-W) (mandatory) (2-bytes)
//
//		Back Pressure Time
//			Back pressure time: This attribute specifies the duration in microseconds of the backpressure
//			signal. It can be used as a pause time for an Ethernet UNI. Upon ME instantiation, the ONU sets
//			this attribute to 0. (R,-W) (mandatory) (4-bytes)
//
//		Back Pressure Occur Queue Threshold
//			Back pressure occur queue threshold: This attribute identifies the threshold queue occupancy, in
//			bytes, scaled by the priority queue scale factor attribute of the ONU2G, to start sending a
//			back-pressure signal. (R, W) (mandatory) (2-bytes)
//
//		Back Pressure Clear Queue Threshold
//			Back pressure clear queue threshold: This attribute identifies the threshold queue occupancy, in
//			bytes, scaled by the priority queue scale factor attribute of the ONU2G, to stop sending a back-
//			pressure signal. (R, W) (mandatory) (2-bytes)
//
//		Packet Drop Queue Thresholds
//			Packet drop queue thresholds: This attribute is a composite of four 2-byte values, a minimum and
//			a maximum threshold, measured in bytes, scaled by the priority queue scale factor attribute of
//			the ONU2-G, for green and yellow packets. The first value is the minimum green threshold, the
//			queue occupancy below which all green packets are admitted to the queue. The second value is the
//			maximum green threshold, the queue occupancy at or above which all green packets are discarded.
//			The third value is the minimum yellow threshold, the queue occupancy below which all yellow
//			packets are admitted to the queue. The fourth value is the maximum yellow threshold, the queue
//			occupancy at or above which all yellow packets are discarded. The default is that all thresholds
//			take the value of the maximum queue size. (R,-W) (optional) (8-bytes)
//
//		Packet Drop Max_P
//			Packet drop max_p: This attribute is a composite of two 1-byte values, the probability of
//			dropping a coloured packet when the queue occupancy lies just below the maximum threshold for
//			packets of that colour. The first value is the green packet max_p, and the second value is the
//			yellow packet max_p. The probability, max_p, is determined by adding one to the unsigned value
//			(0..255) of this attribute and dividing the result by 256. The default for each value is 255.
//			(R,-W) (optional) (2-bytes)
//
//		Queue Drop W_Q
//			Queue drop w_q: This attribute determines the averaging coefficient, w_q, as described in
//			[b-Floyd]. The averaging coefficient, w_q, is equal to 2Queue_drop_w_q. For example, when queue
//			drop_w_q has the value 9, the averaging coefficient, w_q, is 1/512-= 0.001-9. The default value
//			is 9. (R,-W) (optional) (1-byte)
//
//		Drop Precedence Colour Marking
//			(R,-W) (optional) (1-byte)
//
type PriorityQueue struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	priorityqueueBME = &ManagedEntityDefinition{
		Name:    "PriorityQueue",
		ClassID: 277,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("QueueConfigurationOption", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("MaximumQueueSize", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3:  Uint16Field("AllocatedQueueSize", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  Uint16Field("DiscardBlockCounterResetInterval", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 4),
			5:  Uint16Field("ThresholdValueForDiscardedBlocksDueToBufferOverflow", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 5),
			6:  Uint32Field("RelatedPort", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  Uint16Field("TrafficSchedulerPointer", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  ByteField("Weight", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  Uint16Field("BackPressureOperation", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 9),
			10: Uint32Field("BackPressureTime", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 10),
			11: Uint16Field("BackPressureOccurQueueThreshold", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 11),
			12: Uint16Field("BackPressureClearQueueThreshold", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 12),
			13: Uint64Field("PacketDropQueueThresholds", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 13),
			14: Uint16Field("PacketDropMaxP", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 14),
			15: ByteField("QueueDropWQ", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 15),
			16: ByteField("DropPrecedenceColourMarking", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 16),
		},
	}
}

// NewPriorityQueue (class ID 277 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPriorityQueue(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*priorityqueueBME, params...)
}
