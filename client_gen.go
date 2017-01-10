//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/johnprather/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type Client struct {
	rpc *xmlrpc.Client
	Session SessionClass
	Auth AuthClass
	Subject SubjectClass
	Role RoleClass
	Task TaskClass
	Event EventClass
	Pool PoolClass
	PoolPatch PoolPatchClass
	VM VMClass
	VMMetrics VMMetricsClass
	VMGuestMetrics VMGuestMetricsClass
	VMPP VMPPClass
	VMAppliance VMApplianceClass
	DRTask DRTaskClass
	Host HostClass
	HostCrashdump HostCrashdumpClass
	HostPatch HostPatchClass
	HostMetrics HostMetricsClass
	HostCPU HostCPUClass
	Network NetworkClass
	VIF VIFClass
	VIFMetrics VIFMetricsClass
	PIF PIFClass
	PIFMetrics PIFMetricsClass
	Bond BondClass
	VLAN VLANClass
	SM SMClass
	SR SRClass
	VDI VDIClass
	VBD VBDClass
	VBDMetrics VBDMetricsClass
	PBD PBDClass
	Crashdump CrashdumpClass
	VTPM VTPMClass
	Console ConsoleClass
	User UserClass
	DataSource DataSourceClass
	Blob BlobClass
	Message MessageClass
	Secret SecretClass
	Tunnel TunnelClass
	PCI PCIClass
	PGPU PGPUClass
	GPUGroup GPUGroupClass
	VGPU VGPUClass
	VGPUType VGPUTypeClass
	Async struct { 
		Session AsyncSessionClass
		Auth AsyncAuthClass
		Subject AsyncSubjectClass
		Role AsyncRoleClass
		Task AsyncTaskClass
		Event AsyncEventClass
		Pool AsyncPoolClass
		PoolPatch AsyncPoolPatchClass
		VM AsyncVMClass
		VMMetrics AsyncVMMetricsClass
		VMGuestMetrics AsyncVMGuestMetricsClass
		VMPP AsyncVMPPClass
		VMAppliance AsyncVMApplianceClass
		DRTask AsyncDRTaskClass
		Host AsyncHostClass
		HostCrashdump AsyncHostCrashdumpClass
		HostPatch AsyncHostPatchClass
		HostMetrics AsyncHostMetricsClass
		HostCPU AsyncHostCPUClass
		Network AsyncNetworkClass
		VIF AsyncVIFClass
		VIFMetrics AsyncVIFMetricsClass
		PIF AsyncPIFClass
		PIFMetrics AsyncPIFMetricsClass
		Bond AsyncBondClass
		VLAN AsyncVLANClass
		SM AsyncSMClass
		SR AsyncSRClass
		VDI AsyncVDIClass
		VBD AsyncVBDClass
		VBDMetrics AsyncVBDMetricsClass
		PBD AsyncPBDClass
		Crashdump AsyncCrashdumpClass
		VTPM AsyncVTPMClass
		Console AsyncConsoleClass
		User AsyncUserClass
		DataSource AsyncDataSourceClass
		Blob AsyncBlobClass
		Message AsyncMessageClass
		Secret AsyncSecretClass
		Tunnel AsyncTunnelClass
		PCI AsyncPCIClass
		PGPU AsyncPGPUClass
		GPUGroup AsyncGPUGroupClass
		VGPU AsyncVGPUClass
		VGPUType AsyncVGPUTypeClass
	}
}

func prepClient(rpc *xmlrpc.Client) *Client {
	var client Client
	client.rpc = rpc
	client.Session = SessionClass{&client}
	client.Async.Session = AsyncSessionClass{&client}
	client.Auth = AuthClass{&client}
	client.Async.Auth = AsyncAuthClass{&client}
	client.Subject = SubjectClass{&client}
	client.Async.Subject = AsyncSubjectClass{&client}
	client.Role = RoleClass{&client}
	client.Async.Role = AsyncRoleClass{&client}
	client.Task = TaskClass{&client}
	client.Async.Task = AsyncTaskClass{&client}
	client.Event = EventClass{&client}
	client.Async.Event = AsyncEventClass{&client}
	client.Pool = PoolClass{&client}
	client.Async.Pool = AsyncPoolClass{&client}
	client.PoolPatch = PoolPatchClass{&client}
	client.Async.PoolPatch = AsyncPoolPatchClass{&client}
	client.VM = VMClass{&client}
	client.Async.VM = AsyncVMClass{&client}
	client.VMMetrics = VMMetricsClass{&client}
	client.Async.VMMetrics = AsyncVMMetricsClass{&client}
	client.VMGuestMetrics = VMGuestMetricsClass{&client}
	client.Async.VMGuestMetrics = AsyncVMGuestMetricsClass{&client}
	client.VMPP = VMPPClass{&client}
	client.Async.VMPP = AsyncVMPPClass{&client}
	client.VMAppliance = VMApplianceClass{&client}
	client.Async.VMAppliance = AsyncVMApplianceClass{&client}
	client.DRTask = DRTaskClass{&client}
	client.Async.DRTask = AsyncDRTaskClass{&client}
	client.Host = HostClass{&client}
	client.Async.Host = AsyncHostClass{&client}
	client.HostCrashdump = HostCrashdumpClass{&client}
	client.Async.HostCrashdump = AsyncHostCrashdumpClass{&client}
	client.HostPatch = HostPatchClass{&client}
	client.Async.HostPatch = AsyncHostPatchClass{&client}
	client.HostMetrics = HostMetricsClass{&client}
	client.Async.HostMetrics = AsyncHostMetricsClass{&client}
	client.HostCPU = HostCPUClass{&client}
	client.Async.HostCPU = AsyncHostCPUClass{&client}
	client.Network = NetworkClass{&client}
	client.Async.Network = AsyncNetworkClass{&client}
	client.VIF = VIFClass{&client}
	client.Async.VIF = AsyncVIFClass{&client}
	client.VIFMetrics = VIFMetricsClass{&client}
	client.Async.VIFMetrics = AsyncVIFMetricsClass{&client}
	client.PIF = PIFClass{&client}
	client.Async.PIF = AsyncPIFClass{&client}
	client.PIFMetrics = PIFMetricsClass{&client}
	client.Async.PIFMetrics = AsyncPIFMetricsClass{&client}
	client.Bond = BondClass{&client}
	client.Async.Bond = AsyncBondClass{&client}
	client.VLAN = VLANClass{&client}
	client.Async.VLAN = AsyncVLANClass{&client}
	client.SM = SMClass{&client}
	client.Async.SM = AsyncSMClass{&client}
	client.SR = SRClass{&client}
	client.Async.SR = AsyncSRClass{&client}
	client.VDI = VDIClass{&client}
	client.Async.VDI = AsyncVDIClass{&client}
	client.VBD = VBDClass{&client}
	client.Async.VBD = AsyncVBDClass{&client}
	client.VBDMetrics = VBDMetricsClass{&client}
	client.Async.VBDMetrics = AsyncVBDMetricsClass{&client}
	client.PBD = PBDClass{&client}
	client.Async.PBD = AsyncPBDClass{&client}
	client.Crashdump = CrashdumpClass{&client}
	client.Async.Crashdump = AsyncCrashdumpClass{&client}
	client.VTPM = VTPMClass{&client}
	client.Async.VTPM = AsyncVTPMClass{&client}
	client.Console = ConsoleClass{&client}
	client.Async.Console = AsyncConsoleClass{&client}
	client.User = UserClass{&client}
	client.Async.User = AsyncUserClass{&client}
	client.DataSource = DataSourceClass{&client}
	client.Async.DataSource = AsyncDataSourceClass{&client}
	client.Blob = BlobClass{&client}
	client.Async.Blob = AsyncBlobClass{&client}
	client.Message = MessageClass{&client}
	client.Async.Message = AsyncMessageClass{&client}
	client.Secret = SecretClass{&client}
	client.Async.Secret = AsyncSecretClass{&client}
	client.Tunnel = TunnelClass{&client}
	client.Async.Tunnel = AsyncTunnelClass{&client}
	client.PCI = PCIClass{&client}
	client.Async.PCI = AsyncPCIClass{&client}
	client.PGPU = PGPUClass{&client}
	client.Async.PGPU = AsyncPGPUClass{&client}
	client.GPUGroup = GPUGroupClass{&client}
	client.Async.GPUGroup = AsyncGPUGroupClass{&client}
	client.VGPU = VGPUClass{&client}
	client.Async.VGPU = AsyncVGPUClass{&client}
	client.VGPUType = VGPUTypeClass{&client}
	client.Async.VGPUType = AsyncVGPUTypeClass{&client}
	return &client
}
