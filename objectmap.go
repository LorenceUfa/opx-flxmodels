package models

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
var ConfigObjectMap = map[string]ConfigObj{
	"IPV4Route":                          &IPV4Route{},         // created before auto YANG
	"Vlan":                               &Vlan{},              // created before auto YANG
	"IPv4Intf":                           &IPv4Intf{},          // created before auto YANG
	"IPv4Neighbor":                       &IPv4Neighbor{},      // created before auto YANG
        "ArpConfig":                          &ArpConfig{},         // created before auto YANG
	"BGPGlobalConfig":                    &BGPGlobalConfig{},   // created before auto YANG
	"BGPGlobalState":                     &BGPGlobalState{},    // created before auto YANG
	"BGPNeighborConfig":                  &BGPNeighborConfig{}, // created before auto YANG
	"BGPNeighborState":                   &BGPNeighborState{},  // created before auto YANG
	"AggregationLacpConfig":              &AggregationLacpConfig{},
	"EthernetConfig":                     &EthernetConfig{},
	"AggregationLacpMemberStateCounters": &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":               &AggregationLacpState{},
}
