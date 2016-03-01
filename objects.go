package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	CreateDBTable(dbHdl *sql.DB) error
	StoreObjectInDb(dbHdl *sql.DB) (int64, error)
	DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error
	GetKey() (string, error)
	GetSqlKeyStr(string) (string, error)
	GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl *sql.DB) error
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	BaseObj
	DestinationNw     string `SNAPROUTE: "KEY"`
	NetworkMask       string `SNAPROUTE: "KEY"`
	NextHopIp         string `SNAPROUTE: "KEY"`
	Cost              uint32
	OutgoingIntfType  string
	OutgoingInterface string
	Protocol          string
}

func (obj IPV4Route) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Route IPV4Route
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Route); err != nil {
			fmt.Println("### Trouble in unmarshaling IPV4Route from Json", body)
		}
	}
	return v4Route, err
}

type IPV4RouteState struct {
	BaseObj
	DestinationNw    string `SNAPROUTE: "KEY"`
	NetworkMask      string `SNAPROUTE: "KEY"`
	PolicyList       []string
	RouteCreatedTime string
	RouteUpdatedTime string
}

func (obj IPV4RouteState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4RouteState IPV4RouteState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4RouteState); err != nil {
			fmt.Println("### Trouble in unmarshaling IPV4RouteState from Json", body)
		}
	}
	return v4RouteState, err
}

type IPV4EventState struct {
	BaseObj
	TimeStamp string
	EventInfo string
}

func (obj IPV4EventState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4EventState IPV4EventState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4EventState); err != nil {
			fmt.Println("### Trouble in unmarshaling IPV4EventState from Json", body)
		}
	}
	return v4EventState, err
}

type BGPGlobalConfig struct {
	BaseObj
	ASNum               uint32
	RouterId            string `SNAPROUTE: "KEY"`
	UseMultiplePaths    bool
	EBGPMaxPaths        uint32
	EBGPAllowMultipleAS bool
	IBGPMaxPaths        uint32
}

func (obj BGPGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BGPGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

type BGPGlobalState struct {
	BaseObj
	AS                  uint32
	RouterId            string
	UseMultiplePaths    bool
	EBGPMaxPaths        uint32
	EBGPAllowMultipleAS bool
	IBGPMaxPaths        uint32
	TotalPaths          uint32
	TotalPrefixes       uint32
}

func (obj BGPGlobalState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gState BGPGlobalState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gState); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gState, err
}

type PeerType int

const (
	PeerTypeInternal PeerType = iota
	PeerTypeExternal
)

type BgpCounters struct {
	Update       uint64
	Notification uint64
}

type BGPMessages struct {
	Sent     BgpCounters
	Received BgpCounters
}

type BGPQueues struct {
	Input  uint32
	Output uint32
}

type BGPNeighborConfig struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	AuthPassword            string
	Description             string
	NeighborAddress         string `SNAPROUTE: "KEY"`
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
	MultiHopEnable          bool
	MultiHopTTL             uint8
	ConnectRetryTime        uint32
	HoldTime                uint32
	KeepaliveTime           uint32
	PeerGroup               string
	BfdEnable               bool
}

func (obj BGPNeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nConf BGPNeighborConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nConf); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborConfig from Json", body)
		}
	}
	return nConf, err
}

type BGPPeerGroup struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	AuthPassword            string
	Description             string
	Name                    string `SNAPROUTE: "KEY"`
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
	MultiHopEnable          bool
	MultiHopTTL             uint8
	ConnectRetryTime        uint32
	HoldTime                uint32
	KeepaliveTime           uint32
}

func (obj BGPPeerGroup) UnmarshalObject(body []byte) (ConfigObj, error) {
	var group BGPPeerGroup
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &group); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPPeerGroup from Json", body)
		}
	}
	return group, err
}

type BGPNeighborState struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	PeerType                PeerType
	AuthPassword            string
	Description             string
	NeighborAddress         string
	SessionState            uint32
	Messages                BGPMessages
	Queues                  BGPQueues
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
	MultiHopEnable          bool
	MultiHopTTL             uint8
	ConnectRetryTime        uint32
	HoldTime                uint32
	KeepaliveTime           uint32
	BfdNeighborState        string
}

func (obj BGPNeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nState BGPNeighborState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nState); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborState from Json", body)
		}
	}
	return nState, err
}

type BGPRoute struct {
	BaseObj
	Network   string
	CIDRLen   uint16
	NextHop   string
	Metric    uint32
	LocalPref uint32
	Path      []uint32
	Updated   string
}

func (obj BGPRoute) UnmarshalObject(body []byte) (ConfigObj, error) {
	var bgpRoute BGPRoute
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &bgpRoute); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPRoute from Json", body)
		}
	}
	return bgpRoute, err
}

type BGPAggregate struct {
	BaseObj
	IPPrefix        string `SNAPROUTE: "KEY"`
	GenerateASSet   bool
	SendSummaryOnly bool
}

func (obj BGPAggregate) UnmarshalObject(body []byte) (ConfigObj, error) {
	var bgpAgg BGPAggregate
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &bgpAgg); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPRoute from Json", body)
		}
	}
	return bgpAgg, err
}

/* Start asicd owned objects */
/*
 * Vlan object and Route objects are exception cases, i.e
 * they are not seperated into config and state, sub-objects.
 * This approach is followed for objects that can be both
 * statically and dynamically created.
 */
type Vlan struct {
	BaseObj
	VlanId           int32 `SNAPROUTE: "KEY"`
	VlanName         string
	OperState        string
	IfIndex          int32
	IfIndexList      string
	UntagIfIndexList string
}

type IPv4Intf struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY"`
	IfIndex int32
}

type PortConfig struct {
	BaseObj
	PortNum     int32 `SNAPROUTE: "KEY"`
	Description string
	PhyIntfType string
	AdminState  string
	MacAddr     string
	Speed       int32
	Duplex      string
	Autoneg     string
	MediaType   string
	Mtu         int32
}

type PortState struct {
	BaseObj
	PortNum           int32
	IfIndex           int32
	Name              string
	OperState         string
	IfInOctets        int64
	IfInUcastPkts     int64
	IfInDiscards      int64
	IfInErrors        int64
	IfInUnknownProtos int64
	IfOutOctets       int64
	IfOutUcastPkts    int64
	IfOutDiscards     int64
	IfOutErrors       int64
}

func (obj PortState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf PortState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling PortState from Json", body)
		}
	}
	return gConf, err
}

/* End asicd owned objects */

/* ARP */
type ArpConfig struct {
	BaseObj
	// placeholder to create a key
	ArpConfigKey string `SNAPROUTE: "KEY"`
	Timeout      int32
}

func (obj ArpConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpConfigObj ArpConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}

	return arpConfigObj, err
}

type ArpEntry struct {
	BaseObj
	IpAddr         string `SNAPROUTE: "KEY"`
	MacAddr        string
	Vlan           uint32
	Intf           string
	ExpiryTimeLeft string
}

func (obj ArpEntry) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpEntryObj ArpEntry
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpEntryObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}
	return arpEntryObj, err
}

type UserConfig struct {
	BaseObj
	UserName    string `SNAPROUTE: "KEY"`
	Password    string
	Description string
	Previledge  string
}

func (obj UserConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userConfigObj UserConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserConfig from Json", body)
		}
	}

	return userConfigObj, err
}

type UserState struct {
	BaseObj
	UserName      string
	LastLoginTime time.Time
	LastLoginIp   string
	NumAPICalled  uint32
}

func (obj UserState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userStateObj UserState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserState from Json", body)
		}
	}

	return userStateObj, err
}

type Login struct {
	BaseObj
	UserName string
	Password string
}

func (obj Login) UnmarshalObject(body []byte) (ConfigObj, error) {
	var loginObj Login
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &loginObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Login from Json", body)
		}
	}

	return loginObj, err
}

type Logout struct {
	BaseObj
	UserName  string
	SessionId uint32
}

func (obj Logout) UnmarshalObject(body []byte) (ConfigObj, error) {
	var logoutObj Logout
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &logoutObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Logout from Json", body)
		}
	}

	return logoutObj, err
}
