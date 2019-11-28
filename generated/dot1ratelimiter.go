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

// Dot1RateLimiterClassID is the 16-bit ID for the OMCI
// Managed entity Dot1 rate limiter
const Dot1RateLimiterClassID ClassID = ClassID(298)

var dot1ratelimiterBME *ManagedEntityDefinition

// Dot1RateLimiter (class ID #298)
//	This ME allows rate limits to be defined for various types of upstream traffic that are
//	processed by IEEE 802.1 bridges or related structures.
//
//	Relationships
//		An instance of this ME may be linked to an instance of a MAC bridge service profile or an IEEE
//		802.1p mapper.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Parent Me Pointer
//			Parent ME pointer: This attribute points to an instance of a ME. The type of ME is determined by
//			the TP type attribute. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Tp Type
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Upstream Unicast Flood Rate Pointer
//			Upstream unicast flood rate pointer: This attribute points to an instance of the traffic
//			descriptor that governs the rate of upstream unicast packets whose DA is unknown to the bridge.
//			A null pointer specifies that no administrative limit is to be imposed. (R,-W, setbycreate)
//			(optional) (2-bytes)
//
//		Upstream Broadcast Rate Pointer
//			Upstream broadcast rate pointer: This attribute points to an instance of the traffic descriptor
//			that governs the rate of upstream broadcast packets. A null pointer specifies that no
//			administrative limit is to be imposed. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Upstream Multicast Payload Rate Pointer
//			Upstream multicast payload rate pointer: This attribute points to an instance of the traffic
//			descriptor that governs the rate of upstream multicast payload packets. A null pointer specifies
//			that no administrative limit is to be imposed. (R,-W, setbycreate) (optional) (2-bytes)
//
type Dot1RateLimiter struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1ratelimiterBME = &ManagedEntityDefinition{
		Name:    "Dot1RateLimiter",
		ClassID: 298,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("ParentMePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: ByteField("TpType", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("UpstreamUnicastFloodRatePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 3),
			4: Uint16Field("UpstreamBroadcastRatePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 4),
			5: Uint16Field("UpstreamMulticastPayloadRatePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 5),
		},
		Access:  UnknownAccess,
		Support: UnknownSupport,
	}
}

// NewDot1RateLimiter (class ID 298 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1RateLimiter(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1ratelimiterBME, params...)
}
