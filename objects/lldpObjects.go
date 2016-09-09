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

type LLDPGlobal struct {
	baseObj
	Vrf    string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "LLDP Global Config For Default VRF", DEFAULT:"default", AUTOCREATE:"true"`
	Enable bool   `DESCRIPTION: "Enable/Disable LLDP Globally", DEFAULT:false`
}

type LLDPIntf struct {
	baseObj
	//IfIndex int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*",DESCRIPTION: "IfIndex where lldp needs is enabled/disabled", AUTODISCOVER:"true"`
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", DESCRIPTION: "IfIndex where lldp needs is enabled/disabled", DEFAULT: "None", AUTODISCOVER:"true"`
	Enable  bool   `DESCRIPTION: "Enable/Disable lldp config Per Port", DEFAULT:true`
}

type LLDPIntfState struct {
	baseObj
	//IfIndex             int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*",DESCRIPTION: "IfIndex where lldp needs to be configured"`
	IntfRef             string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*",DESCRIPTION: "IntfRef where lldp is configured"`
	IfIndex             int32  `DESCRIPTION: "IfIndex where lldp needs to be configured"`
	Enable              bool   `DESCRIPTION: "Enable/Disable lldp config"`
	LocalPort           string `DESCRIPTION: "Local interface"`
	PeerMac             string `DESCRIPTION: "Mac address of direct connection"`
	PeerPort            string `DESCRIPTION: "Name of directtly connected pors"`
	PeerHostName        string `DESCRIPTION: "Name of the peer host`
	HoldTime            string `DESCRIPTION: "Validity of the peer information"`
	SystemDescription   string `DESCRIPTION: "System Description of the peer port"`
	SystemCapabilities  string `DESCRIPTION: "System Capabilities of the peer port"`
	EnabledCapabilities string `DESCRIPTION: "Enabled Capabilities of the peer port"`
}
