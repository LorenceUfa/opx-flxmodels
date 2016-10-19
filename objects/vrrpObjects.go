//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

/*
 * This DS will be used while Created/Deleting Vrrp Intf Config
 */
type VrrpV4Intf struct {
	baseObj
	//	IfIndex               int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: ""Interface index for which VRRP Config needs to be done"`
	IntfRef               string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: ""Interface (name) for which VRRP Version 2 aka VRRP with ipv4 Config needs to be done"`
	VRID                  int32  `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	Priority              int32  `DESCRIPTION: "Sending VRRP router's priority for the virtual router", DEFAULT:"100", MIN:"1", MAX:"255"`
	VirtualIPv4Addr       string `DESCRIPTION: "Virtual Router Identifier", STRLEN:"17"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:"1", MIN:"1", MAX:"4095"`
	PreemptMode           bool   `DESCRIPTION: "Controls whether a (starting or restarting) higher-priority Backup router preempts a lower-priority Master router", DEFAULT: "true"`
	AcceptMode            bool   `DESCRIPTION: "Controls whether a virtual router in Master state will accept packets addressed to the address owner's IPvX address as its own if it is not the IPvX address owner.", DEFAULT:"false"`
}

type VrrpV6Intf struct {
	baseObj
	//	IfIndex               int32  `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: ""Interface index for which VRRP Config needs to be done"`
	IntfRef               string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: ""Interface (name) for which VRRP Version 3 aka VRRP with ipv6 Config needs to be done"`
	VRID                  int32  `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	Priority              int32  `DESCRIPTION: "Sending VRRP router's priority for the virtual router", DEFAULT:"100", MIN:"1", MAX:"255"`
	VirtualIPv6Addr       string `DESCRIPTION: "Virtual Router Identifier", STRLEN:"43"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:"1", MIN:"1", MAX:"4095"`
	PreemptMode           bool   `DESCRIPTION: "Controls whether a (starting or restarting) higher-priority Backup router preempts a lower-priority Master router", DEFAULT: "true"`
	AcceptMode            bool   `DESCRIPTION: "Controls whether a virtual router in Master state will accept packets addressed to the address owner's IPvX address as its own if it is not the IPvX address owner.", DEFAULT:"false"`
}

type VrrpV4IntfState struct {
	baseObj
	//IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Interface index for which VRRP state is requested"`
	IntfRef                 string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: ""Interface (name) for which VRRP Version 2 aka VRRP with ipv4 state information needs to be retreived"`
	VRID                    int32  `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	IntfIpAddr              string `DESCRIPTION: "Ip Address of Interface where VRRP is configured"`
	Priority                int32  `DESCRIPTION: "Virtual router's Priority"`
	VirtualIPv4Addr         string `DESCRIPTION: "Ip Address of Virtual Router"`
	AdvertisementInterval   int32  `DESCRIPTION: "Time interval between Advertisements"`
	PreemptMode             bool   `DESCRIPTION: "States Whether Preempt is Supported or not"`
	VirtualRouterMACAddress string `DESCRIPTION: "VRRP router's Mac Address"`
	SkewTime                int32  `DESCRIPTION: "Time to skew Master Down Interval"`
	MasterDownTimer         int32  `DESCRIPTION: "Time interval for Backup to declare Master down"`
	VrrpState               string `DESCRIPTION: "Current vrrp state i.e. backup or master"`
}

type VrrpV6IntfState struct {
	baseObj
	//IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Interface index for which VRRP state is requested"`
	IntfRef                 string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: ""Interface (name) for which VRRP Version 3 aka VRRP with ipv4 state information needs to be retreived"`
	VRID                    int32  `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	IntfIpAddr              string `DESCRIPTION: "Ipv6 Address of Interface where VRRP is configured"`
	Priority                int32  `DESCRIPTION: "Virtual router's Priority"`
	VirtualIPv6Addr         string `DESCRIPTION: "Ipv6 Address of Virtual Router"`
	AdvertisementInterval   int32  `DESCRIPTION: "Time interval between Advertisements"`
	PreemptMode             bool   `DESCRIPTION: "States Whether Preempt is Supported or not"`
	VirtualRouterMACAddress string `DESCRIPTION: "VRRP router's Mac Address"`
	SkewTime                int32  `DESCRIPTION: "Time to skew Master Down Interval"`
	MasterDownTimer         int32  `DESCRIPTION: "Time interval for Backup to declare Master down"`
	VrrpState               string `DESCRIPTION: "Current vrrp state i.e. backup or master"`
}

/*
type VrrpVridState struct {
	baseObj
	IfIndex          int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Interface index for which VRRP state is requested"`
	VRID             int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Virtual Router's Unique Identifier""`
	AdverRx          int32  `DESCRIPTION:"Total number of advertisement packets received"`
	AdverTx          int32  `DESCRIPTION:"Total number of advertisement packets send"`
	LastAdverRx      string `DESCRIPTION:"Time when last advertisement packet was received"`
	LastAdverTx      string `DESCRIPTION:"Time when last advertisement packet was send out"`
	MasterIp         string `DESCRIPTION:"Ip Address of the Master VRRP"`
	CurrentState     string `DESCRIPTION:"Current State of Local VRRP"`
	PreviousState    string `DESCRIPTION:"Previous State of Local VRRP"`
	TransitionReason string `DESCRIPTION:"Reason why we moved from previous state -> current state"`
}
*/
