//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type VMApplianceOperation string

const (
	VMApplianceOperationStart VMApplianceOperation = "start"
	VMApplianceOperationCleanShutdown VMApplianceOperation = "clean_shutdown"
	VMApplianceOperationHardShutdown VMApplianceOperation = "hard_shutdown"
	VMApplianceOperationShutdown VMApplianceOperation = "shutdown"
)

type VMApplianceRecord struct {
	UUID string
	NameLabel string
	NameDescription string
	AllowedOperations []VMApplianceOperation
	CurrentOperations map[string]VMApplianceOperation
	VMs []VMRef
}

type VMApplianceRef string

type VMApplianceClass struct {
	client *Client
}

func (client *Client) VMAppliance() VMApplianceClass {
	return VMApplianceClass{client}
}

func (_class VMApplianceClass) GetAllRecords(sessionID SessionRef) (_retval map[VMApplianceRef]VMApplianceRecord, _err error) {
	_method := "VM_appliance.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToVMApplianceRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetAll(sessionID SessionRef) (_retval []VMApplianceRef, _err error) {
	_method := "VM_appliance.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) Recover(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef, force bool) (_err error) {
	_method := "VM_appliance.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	return
}

func (_class VMApplianceClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	_method := "VM_appliance.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) AssertCanBeRecovered(sessionID SessionRef, self VMApplianceRef, sessionTo SessionRef) (_err error) {
	_method := "VM_appliance.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	return
}

func (_class VMApplianceClass) Shutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	_method := "VM_appliance.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMApplianceClass) HardShutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	_method := "VM_appliance.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMApplianceClass) CleanShutdown(sessionID SessionRef, self VMApplianceRef) (_err error) {
	_method := "VM_appliance.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMApplianceClass) Start(sessionID SessionRef, self VMApplianceRef, paused bool) (_err error) {
	_method := "VM_appliance.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_pausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "paused"), paused)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _pausedArg)
	return
}

func (_class VMApplianceClass) SetNameDescription(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	_method := "VM_appliance.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMApplianceClass) SetNameLabel(sessionID SessionRef, self VMApplianceRef, value string) (_err error) {
	_method := "VM_appliance.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMApplianceClass) GetVMs(sessionID SessionRef, self VMApplianceRef) (_retval []VMRef, _err error) {
	_method := "VM_appliance.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetCurrentOperations(sessionID SessionRef, self VMApplianceRef) (_retval map[string]VMApplianceOperation, _err error) {
	_method := "VM_appliance.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVMApplianceOperationMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetAllowedOperations(sessionID SessionRef, self VMApplianceRef) (_retval []VMApplianceOperation, _err error) {
	_method := "VM_appliance.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMApplianceOperationSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetNameDescription(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	_method := "VM_appliance.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetNameLabel(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	_method := "VM_appliance.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetUUID(sessionID SessionRef, self VMApplianceRef) (_retval string, _err error) {
	_method := "VM_appliance.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMApplianceRef, _err error) {
	_method := "VM_appliance.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) Destroy(sessionID SessionRef, self VMApplianceRef) (_err error) {
	_method := "VM_appliance.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMApplianceClass) Create(sessionID SessionRef, args VMApplianceRecord) (_retval VMApplianceRef, _err error) {
	_method := "VM_appliance.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMApplianceRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMApplianceRef, _err error) {
	_method := "VM_appliance.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMApplianceClass) GetRecord(sessionID SessionRef, self VMApplianceRef) (_retval VMApplianceRecord, _err error) {
	_method := "VM_appliance.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRecordToGo(_method + " -> ", _result.Value)
	return
}