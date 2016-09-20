//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
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
 * This DS will be used while Created/Deleting Platform Config
 */

type PlatformState struct {
	baseObj
	ObjName      string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "ObjName", DEFAULT: "Platform"`
	ProductName  string `DESCRIPTION: "Product Number"`
	SerialNum    string `DESCRIPTION: "Serial Number"`
	Manufacturer string `DESCRIPTION: "Manufacturer"`
	Vendor       string `DESCRIPTION: "Vendor"`
	Release      string `DESCRIPTION: "Relese version"`
	PlatformName string `DESCRIPTION: "Platform Number"`
	Version      string `DESCRIPTION: "Platform Driver version, in case of ONLP(ONIE Version) and OpenBMC(BMC Version)"`
}

type Sfp struct {
	baseObj
	SfpId      int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "SFP id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin PORT UP/DOWN(TX OFF)"`
}

type SfpState struct {
	baseObj
	SfpId      int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "SFP id", DEFAULT:0`
	SfpSpeed   string `DESCRIPTION: "SFP speed in MBPS"`
	SfpLOS     string `DESCRIPTION: "SFP status RX LOS"`
	SfpPresent string `DESCRIPTION: "SFP status PRESENT/MISSING"`
	SfpType    string `DESCRIPTION: "SFP type Copper/Optical"`
	SerialNum  string `DESCRIPTION: "SFP SerialNum"`
	EEPROM     string `DESCRIPTION: "SFP eeprom"`
}

type ThermalState struct {
	baseObj
	ThermalId                 int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Thermal sensor id", DEFAULT:0`
	Location                  string `DESCRIPTION: "Thermal sensor location CPU/PSU/Motherboard"`
	Temperature               string `DESCRIPTION: "Temperature current"`
	LowerWatermarkTemperature string `DESCRIPTION: "Temperature warning"`
	UpperWatermarkTemperature string `DESCRIPTION: "Temperature error"`
	ShutdownTemperature       string `DESCRIPTION: "Temperature panic"`
}

type Psu struct {
	baseObj
	PsuId      int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "PSU id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin UP/DOWN PSU"`
}

type PsuState struct {
	baseObj
	PsuId      int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PSU id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Admin UP/DOWN PSU"`
	ModelNum   string `DESCRIPTION: "Model Number"`
	SerialNum  string `DESCRIPTION: "Serial Number"`
	Vin        int32  `DESCRIPTION: "Voltage in"`
	Vout       int32  `DESCRIPTION: "Voltage out"`
	Iin        int32  `DESCRIPTION: "Current in"`
	Iout       int32  `DESCRIPTION: "Current out"`
	Pin        int32  `DESCRIPTION: "Power in"`
	Pout       int32  `DESCRIPTION: "power out"`
	Fan        string `DESCRIPTION: "Fan PRESENT/MISSING"`
	FanId      int32  `DESCRIPTION: "Fan Info"`
	LedId      int32  `DESCRIPTION: "LED Info"`
}

type Led struct {
	baseObj
	LedId       int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "LED id", DEFAULT:0`
	LedAdmin    string `DESCRIPTION: "LED ON/OFF"`
	LedSetColor string `DESCRIPTION: "LED set color"`
}

type LedState struct {
	baseObj
	LedId       int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "LED id", DEFAULT:0`
	LedIdentify string `DESCRIPTION: "LED represents FAN/PSU/RESET etc"`
	LedState    string `DESCRIPTION: "LED State ON/OFF"`
	LedColor    string `DESCRIPTION: "LED Color"`
}

type Fan struct {
	baseObj
	FanId      int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Fan unit id", DEFAULT:0`
	AdminState string `DESCRIPTION: "Fan admin ON/OFF"`
	AdminSpeed int32  `DESCRIPTION: "Fan set speed in rpm"`
}

type FanState struct {
	baseObj
	FanId         int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fan unit id", DEFAULT:0`
	OperMode      string `DESCRIPTION: "Operational state of Fan", SELECTION: ON/OFF`
	OperSpeed     int32  `DESCRIPTION: "Fan operational speed in rpm"`
	OperDirection string `DESCRIPTION: "Air flow caused because of fan rotation", SELECTION: B2F/F2B"`
	Status        string `DESCRIPTION: "Fan status PRESENT/MISSING/FAILED/NORMAL"`
	Model         string `DESCRIPTION: "Model of Fan"`
	SerialNum     string `DESCRIPTION: "Serial Number"`
	LedId         int32  `DESCRIPTION: "LED Info"`
}

type FanSensor struct {
	baseObj
	Name                   string `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Fan Sensor Name"`
	AdminState             string `DESCRIPTION: "Enable/Disable", DEFAULT: Enable, SELECTION: Enable/Disable`
	HigherAlarmThreshold   int32  `DESCRIPTION: "Higher Alarm Threshold for TCA"`
	HigherWarningThreshold int32  `DESCRIPTION: "Higher Warning Threshold for TCA"`
	LowerWarningThreshold  int32  `DESCRIPTION: "Lower Warning Threshold for TCA"`
	LowerAlarmThreshold    int32  `DESCRIPTION: "Lower Alarm Threshold for TCA"`
	PMClassAAdminState     string `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassBAdminState     string `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassCAdminState     string `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
}

type FanSensorState struct {
	baseObj
	Name         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fan Sensor Name"`
	CurrentSpeed int32  `DESCRIPTION: "Fan Current Speed"`
}

type FanSensorPMData struct {
	TimeStamp string `DESCRIPTION: "Timestamp at which data is collected"`
	Value     int32  `DESCRIPTION: "Fan PM Data Value in RPM"`
}

type FanSensorPMDataState struct {
	baseObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fan Sensor Name"`
	Class string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data  []FanSensorPMData
}

type TemperatureSensor struct {
	baseObj
	Name                   string  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Temperature Sensor Name"`
	AdminState             string  `DESCRIPTION: "Enable/Disable", DEFAULT: Enable, SELECTION: Enable/Disable`
	HigherAlarmThreshold   float64 `DESCRIPTION: "Higher Alarm Threshold for TCA"`
	HigherWarningThreshold float64 `DESCRIPTION: "Higher Warning Threshold for TCA"`
	LowerWarningThreshold  float64 `DESCRIPTION: "Lower Warning Threshold for TCA"`
	LowerAlarmThreshold    float64 `DESCRIPTION: "Lower Alarm Threshold for TCA"`
	PMClassAAdminState     string  `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassBAdminState     string  `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassCAdminState     string  `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
}

type TemperatureSensorState struct {
	baseObj
	Name               string  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Temperature Sensor Name"`
	CurrentTemperature float64 `DESCRIPTION: "Current Temperature Value"`
}

type TemperatureSensorPMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "Temperature Sensor PM Data Value in Degree Celsius"`
}

type TemperatureSensorPMDataState struct {
	baseObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Temperature Sensor Name"`
	Class string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data  []TemperatureSensorPMData
}

type VoltageSensor struct {
	baseObj
	Name                   string  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Voltage Sensor Name"`
	AdminState             string  `DESCRIPTION: "Enable/Disable", DEFAULT: Enable, SELECTION: Enable/Disable`
	HigherAlarmThreshold   float64 `DESCRIPTION: "Higher Alarm Threshold for TCA"`
	HigherWarningThreshold float64 `DESCRIPTION: "Higher Warning Threshold for TCA"`
	LowerWarningThreshold  float64 `DESCRIPTION: "Lower Warning Threshold for TCA"`
	LowerAlarmThreshold    float64 `DESCRIPTION: "Lower Alarm Threshold for TCA"`
	PMClassAAdminState     string  `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassBAdminState     string  `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassCAdminState     string  `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
}

type VoltageSensorState struct {
	baseObj
	Name           string  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Voltage Sensor Name"`
	CurrentVoltage float64 `DESCRIPTION: "Current Voltage Value"`
}

type VoltageSensorPMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "Voltage Sensor PM Data Value in Volts"`
}

type VoltageSensorPMDataState struct {
	baseObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Voltage Sensor Name"`
	Class string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data  []VoltageSensorPMData
}

type PowerConverterSensor struct {
	baseObj
	Name                   string  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Power Converter Sensor Name"`
	AdminState             string  `DESCRIPTION: "Enable/Disable", DEFAULT: Enable, SELECTION: Enable/Disable`
	HigherAlarmThreshold   float64 `DESCRIPTION: "Higher Alarm Threshold for TCA"`
	HigherWarningThreshold float64 `DESCRIPTION: "Higher Warning Threshold for TCA"`
	LowerWarningThreshold  float64 `DESCRIPTION: "Lower Warning Threshold for TCA"`
	LowerAlarmThreshold    float64 `DESCRIPTION: "Lower Alarm Threshold for TCA"`
	PMClassAAdminState     string  `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassBAdminState     string  `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
	PMClassCAdminState     string  `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Enable, SELECTION: Enable/Disable`
}

type PowerConverterSensorState struct {
	baseObj
	Name         string  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Power Converter Sensor Name"`
	CurrentPower float64 `DESCRIPTION: "Current Output Power Value"`
}

type PowerConverterSensorPMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "Power Converter Sensor PM Data Value in Watts"`
}

type PowerConverterSensorPMDataState struct {
	baseObj
	Name  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Power Converter Sensor Name"`
	Class string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data  []PowerConverterSensorPMData
}

type Qsfp struct {
	baseObj
	QsfpId                   int32   `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Qsfp Id"`
	AdminState               string  `DESCRIPTION: "Enable/Disable", DEFAULT: Disable, SELECTION: Enable/Disable`
	HigherAlarmTemperature   float64 `DESCRIPTION: "Higher Alarm temperature threshold for TCA"`
	HigherAlarmVoltage       float64 `DESCRIPTION: "Higher Alarm Voltage threshold for TCA"`
	HigherWarningTemperature float64 `DESCRIPTION: "Higher Warning temperature threshold for TCA"`
	HigherWarningVoltage     float64 `DESCRIPTION: "Higher Warning Voltage threshold for TCA"`
	LowerAlarmTemperature    float64 `DESCRIPTION: "Lower Alarm temperature threshold for TCA"`
	LowerAlarmVoltage        float64 `DESCRIPTION: "Lower Alarm Voltage threshold for TCA"`
	LowerWarningTemperature  float64 `DESCRIPTION: "Lower Warning temperature threshold for TCA"`
	LowerWarningVoltage      float64 `DESCRIPTION: "Lower Warning Voltage threshold for TCA"`
	PMClassAAdminState       string  `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
	PMClassBAdminState       string  `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
	PMClassCAdminState       string  `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
}

type QsfpState struct {
	baseObj
	QsfpId             int32   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "QSFP Id"`
	Present            bool    `DESCRIPTION: "Present or Not Value: true if present and false if not present"`
	VendorName         string  `DESCRIPTION: Vendor Name"`
	VendorOUI          string  `DESCRIPTION: Vendor OUI"`
	VendorPartNumber   string  `DESCRIPTION: Vendor Part Number"`
	VendorRevision     string  `DESCRIPTION: Vendor Revision"`
	VendorSerialNumber string  `DESCRIPTION: Vendor Serial Number"`
	DataCode           string  `DESCRIPTION: Data Code"`
	Temperature        float64 `DESCRIPTION: "Current temperature"`
	Voltage            float64 `DESCRIPTION: "Current Voltage"`
}

type QsfpPMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "Qsfp PM Data Value"`
}

type QsfpPMDataState struct {
	baseObj
	QsfpId   int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "QSFP Id"`
	Resource string `SNAPROUTE: "KEY", DESCRIPTION: "QSFP PM Resource Name"`
	Class    string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data     []QsfpPMData
}

type QsfpChannel struct {
	baseObj
	QsfpId               int32   `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: "Qsfp Id"`
	ChannelNum           int32   `SNAPROUTE: "KEY", DESCRIPTION: "Qsfp Channel Number"`
	AdminState           string  `DESCRIPTION: "Enable/Disable", DEFAULT: Disable, SELECTION: Enable/Disable`
	HigherAlarmRXPower   float64 `DESCRIPTION: "Higher Alarm Rx power Threshold for TCA"`
	HigherAlarmTXPower   float64 `DESCRIPTION: "Higher Alarm Rx power for TCA"`
	HigherAlarmTXBias    float64 `DESCRIPTION: "Higher Alarm Tx Current Bias for TCA"`
	HigherWarningRXPower float64 `DESCRIPTION: "Higher Warning Rx power Threshold for TCA"`
	HigherWarningTXPower float64 `DESCRIPTION: "Higher Warning Rx power for TCA"`
	HigherWarningTXBias  float64 `DESCRIPTION: "Higher Warning Tx Current Bias for TCA"`
	LowerAlarmRXPower    float64 `DESCRIPTION: "Lower Alarm Rx power Threshold for TCA"`
	LowerAlarmTXPower    float64 `DESCRIPTION: "Lower Alarm Rx power for TCA"`
	LowerAlarmTXBias     float64 `DESCRIPTION: "Lower Alarm Tx Current Bias for TCA"`
	LowerWarningRXPower  float64 `DESCRIPTION: "Lower Warning Rx power Threshold for TCA"`
	LowerWarningTXPower  float64 `DESCRIPTION: "Lower Warning Rx power for TCA"`
	LowerWarningTXBias   float64 `DESCRIPTION: "Lower Warning Tx Current Bias for TCA"`
	PMClassAAdminState   string  `DESCRIPTION: "PM Class-A Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
	PMClassBAdminState   string  `DESCRIPTION: "PM Class-B Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
	PMClassCAdminState   string  `DESCRIPTION: "PM Class-C Admin State", DEFAULT: Disable, SELECTION: Enable/Disable`
}

type QsfpChannelState struct {
	baseObj
	QsfpId     int32   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "QSFP Id"`
	ChannelNum int32   `SNAPROUTE: "KEY", DESCRIPTION: "Qsfp Channel Number"`
	Present    bool    `DESCRIPTION: "Present or Not Value: true if present and false if not present"`
	RXPower    float64 `DESCRIPTION: "Rx power on channel 1"`
	TXPower    float64 `DESCRIPTION: "Rx power on channel 1"`
	TXBias     float64 `DESCRIPTION: "Tx Current Bias on channel 1"`
}

type QsfpChannelPMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "Qsfp PM Data Value"`
}

type QsfpChannelPMDataState struct {
	baseObj
	QsfpId     int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "QSFP Id"`
	ChannelNum int32  `SNAPROUTE: "KEY", DESCRIPTION: "Qsfp Channel Number"`
	Resource   string `SNAPROUTE: "KEY", DESCRIPTION: "QSFP PM Resource Name"`
	Class      string `SNAPROUTE: "KEY", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data       []QsfpChannelPMData
}

type PlatformMgmtDeviceState struct {
	baseObj
	DeviceName  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Device Name", DEFAULT: "BMC"`
	Uptime      string `DESCRIPTION: "Uptime and load description"`
	Description string `DESCRIPTION: "Platform Description"`
	ResetReason string `DESCRIPTION: "Reset Reason"`
	MemoryUsage string `DESCRIPTION: "Memory Usage"`
	Version     string `DESCRIPTION: "Version"`
	CPUUsage    string `DESCRIPTION: "CPU Usage"`
}
