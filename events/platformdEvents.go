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

package events

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FanSensorKey struct {
	Name string
}

type TemperatureSensorKey struct {
	Name string
}

type VoltageSensorKey struct {
	Name string
}

type PowerConverterSensorKey struct {
	Name string
}

const (
	FanHigherTCAAlarmClear            EventId = 1
	FanHigherTCAAlarm                 EventId = 2
	FanHigherTCAWarnClear             EventId = 3
	FanHigherTCAWarn                  EventId = 4
	FanLowerTCAAlarmClear             EventId = 5
	FanLowerTCAAlarm                  EventId = 6
	FanLowerTCAWarnClear              EventId = 7
	FanLowerTCAWarn                   EventId = 8
	TemperatureHigherTCAAlarmClear    EventId = 9
	TemperatureHigherTCAAlarm         EventId = 10
	TemperatureHigherTCAWarnClear     EventId = 11
	TemperatureHigherTCAWarn          EventId = 12
	TemperatureLowerTCAAlarmClear     EventId = 13
	TemperatureLowerTCAAlarm          EventId = 14
	TemperatureLowerTCAWarnClear      EventId = 15
	TemperatureLowerTCAWarn           EventId = 16
	VoltageHigherTCAAlarmClear        EventId = 17
	VoltageHigherTCAAlarm             EventId = 18
	VoltageHigherTCAWarnClear         EventId = 19
	VoltageHigherTCAWarn              EventId = 20
	VoltageLowerTCAAlarmClear         EventId = 21
	VoltageLowerTCAAlarm              EventId = 22
	VoltageLowerTCAWarnClear          EventId = 23
	VoltageLowerTCAWarn               EventId = 24
	PowerConverterHigherTCAAlarmClear EventId = 25
	PowerConverterHigherTCAAlarm      EventId = 26
	PowerConverterHigherTCAWarnClear  EventId = 27
	PowerConverterHigherTCAWarn       EventId = 28
	PowerConverterLowerTCAAlarmClear  EventId = 29
	PowerConverterLowerTCAAlarm       EventId = 30
	PowerConverterLowerTCAWarnClear   EventId = 31
	PowerConverterLowerTCAWarn        EventId = 32
)

var PlatformdEventKeyMap KeyMap = KeyMap{
	"FanSensor":            FanSensorKey{},
	"TemperatureSensor":    TemperatureSensorKey{},
	"VoltageSensor":        VoltageSensorKey{},
	"PowerConverterSensor": PowerConverterSensorKey{},
}

func (obj FanSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("FanSensor#%s", obj.Name), nil
}

func (obj TemperatureSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("TemperatureSensor#%s", obj.Name), nil
}

func (obj VoltageSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("VoltageSensor#%s", obj.Name), nil
}

func (obj PowerConverterSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("PowerConverterSensor#%s", obj.Name), nil
}
