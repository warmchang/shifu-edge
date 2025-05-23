//go:build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDevice) DeepCopyInto(out *EdgeDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDevice.
func (in *EdgeDevice) DeepCopy() *EdgeDevice {
	if in == nil {
		return nil
	}
	out := new(EdgeDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceList) DeepCopyInto(out *EdgeDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceList.
func (in *EdgeDeviceList) DeepCopy() *EdgeDeviceList {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceSpec) DeepCopyInto(out *EdgeDeviceSpec) {
	*out = *in
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(Connection)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(Protocol)
		**out = **in
	}
	if in.ProtocolSettings != nil {
		in, out := &in.ProtocolSettings, &out.ProtocolSettings
		*out = new(ProtocolSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewaySettings != nil {
		in, out := &in.GatewaySettings, &out.GatewaySettings
		*out = new(GatewaySettings)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomMetadata != nil {
		in, out := &in.CustomMetadata, &out.CustomMetadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceSpec.
func (in *EdgeDeviceSpec) DeepCopy() *EdgeDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeDeviceStatus) DeepCopyInto(out *EdgeDeviceStatus) {
	*out = *in
	if in.EdgeDevicePhase != nil {
		in, out := &in.EdgeDevicePhase, &out.EdgeDevicePhase
		*out = new(EdgeDevicePhase)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeDeviceStatus.
func (in *EdgeDeviceStatus) DeepCopy() *EdgeDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySettings) DeepCopyInto(out *GatewaySettings) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.LwM2MSetting != nil {
		in, out := &in.LwM2MSetting, &out.LwM2MSetting
		*out = new(LwM2MSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySettings.
func (in *GatewaySettings) DeepCopy() *GatewaySettings {
	if in == nil {
		return nil
	}
	out := new(GatewaySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSetting) DeepCopyInto(out *HTTPSetting) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSetting.
func (in *HTTPSetting) DeepCopy() *HTTPSetting {
	if in == nil {
		return nil
	}
	out := new(HTTPSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LwM2MSetting) DeepCopyInto(out *LwM2MSetting) {
	*out = *in
	if in.SecurityMode != nil {
		in, out := &in.SecurityMode, &out.SecurityMode
		*out = new(SecurityMode)
		**out = **in
	}
	if in.DTLSMode != nil {
		in, out := &in.DTLSMode, &out.DTLSMode
		*out = new(DTLSMode)
		**out = **in
	}
	if in.CipherSuites != nil {
		in, out := &in.CipherSuites, &out.CipherSuites
		*out = make([]CipherSuite, len(*in))
		copy(*out, *in)
	}
	if in.PSKIdentity != nil {
		in, out := &in.PSKIdentity, &out.PSKIdentity
		*out = new(string)
		**out = **in
	}
	if in.PSKKey != nil {
		in, out := &in.PSKKey, &out.PSKKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LwM2MSetting.
func (in *LwM2MSetting) DeepCopy() *LwM2MSetting {
	if in == nil {
		return nil
	}
	out := new(LwM2MSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTSetting) DeepCopyInto(out *MQTTSetting) {
	*out = *in
	if in.MQTTTopic != nil {
		in, out := &in.MQTTTopic, &out.MQTTTopic
		*out = new(string)
		**out = **in
	}
	if in.MQTTServerAddress != nil {
		in, out := &in.MQTTServerAddress, &out.MQTTServerAddress
		*out = new(string)
		**out = **in
	}
	if in.MQTTServerSecret != nil {
		in, out := &in.MQTTServerSecret, &out.MQTTServerSecret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTSetting.
func (in *MQTTSetting) DeepCopy() *MQTTSetting {
	if in == nil {
		return nil
	}
	out := new(MQTTSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinIOSetting) DeepCopyInto(out *MinIOSetting) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(string)
		**out = **in
	}
	if in.RequestTimeoutMS != nil {
		in, out := &in.RequestTimeoutMS, &out.RequestTimeoutMS
		*out = new(int64)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.FileExtension != nil {
		in, out := &in.FileExtension, &out.FileExtension
		*out = new(string)
		**out = **in
	}
	if in.ServerAddress != nil {
		in, out := &in.ServerAddress, &out.ServerAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinIOSetting.
func (in *MinIOSetting) DeepCopy() *MinIOSetting {
	if in == nil {
		return nil
	}
	out := new(MinIOSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OPCUASetting) DeepCopyInto(out *OPCUASetting) {
	*out = *in
	if in.OPCUAEndpoint != nil {
		in, out := &in.OPCUAEndpoint, &out.OPCUAEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SecurityMode != nil {
		in, out := &in.SecurityMode, &out.SecurityMode
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationMode != nil {
		in, out := &in.AuthenticationMode, &out.AuthenticationMode
		*out = new(string)
		**out = **in
	}
	if in.CertificateFileName != nil {
		in, out := &in.CertificateFileName, &out.CertificateFileName
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeyFileName != nil {
		in, out := &in.PrivateKeyFileName, &out.PrivateKeyFileName
		*out = new(string)
		**out = **in
	}
	if in.ConfigmapName != nil {
		in, out := &in.ConfigmapName, &out.ConfigmapName
		*out = new(string)
		**out = **in
	}
	if in.IssuedToken != nil {
		in, out := &in.IssuedToken, &out.IssuedToken
		*out = new(string)
		**out = **in
	}
	if in.SecurityPolicy != nil {
		in, out := &in.SecurityPolicy, &out.SecurityPolicy
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.ConnectionTimeoutInMilliseconds != nil {
		in, out := &in.ConnectionTimeoutInMilliseconds, &out.ConnectionTimeoutInMilliseconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OPCUASetting.
func (in *OPCUASetting) DeepCopy() *OPCUASetting {
	if in == nil {
		return nil
	}
	out := new(OPCUASetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLC4XSetting) DeepCopyInto(out *PLC4XSetting) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(Plc4xProtocol)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLC4XSetting.
func (in *PLC4XSetting) DeepCopy() *PLC4XSetting {
	if in == nil {
		return nil
	}
	out := new(PLC4XSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolSettings) DeepCopyInto(out *ProtocolSettings) {
	*out = *in
	if in.MQTTSetting != nil {
		in, out := &in.MQTTSetting, &out.MQTTSetting
		*out = new(MQTTSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.OPCUASetting != nil {
		in, out := &in.OPCUASetting, &out.OPCUASetting
		*out = new(OPCUASetting)
		(*in).DeepCopyInto(*out)
	}
	if in.SocketSetting != nil {
		in, out := &in.SocketSetting, &out.SocketSetting
		*out = new(SocketSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.PLC4XSetting != nil {
		in, out := &in.PLC4XSetting, &out.PLC4XSetting
		*out = new(PLC4XSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPSetting != nil {
		in, out := &in.TCPSetting, &out.TCPSetting
		*out = new(TCPSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.LwM2MSetting != nil {
		in, out := &in.LwM2MSetting, &out.LwM2MSetting
		*out = new(LwM2MSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolSettings.
func (in *ProtocolSettings) DeepCopy() *ProtocolSettings {
	if in == nil {
		return nil
	}
	out := new(ProtocolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLConnectionSetting) DeepCopyInto(out *SQLConnectionSetting) {
	*out = *in
	if in.ServerAddress != nil {
		in, out := &in.ServerAddress, &out.ServerAddress
		*out = new(string)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.DBName != nil {
		in, out := &in.DBName, &out.DBName
		*out = new(string)
		**out = **in
	}
	if in.DBTable != nil {
		in, out := &in.DBTable, &out.DBTable
		*out = new(string)
		**out = **in
	}
	if in.DBType != nil {
		in, out := &in.DBType, &out.DBType
		*out = new(DBType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLConnectionSetting.
func (in *SQLConnectionSetting) DeepCopy() *SQLConnectionSetting {
	if in == nil {
		return nil
	}
	out := new(SQLConnectionSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSettings) DeepCopyInto(out *ServiceSettings) {
	*out = *in
	if in.HTTPSetting != nil {
		in, out := &in.HTTPSetting, &out.HTTPSetting
		*out = new(HTTPSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.MQTTSetting != nil {
		in, out := &in.MQTTSetting, &out.MQTTSetting
		*out = new(MQTTSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.SQLSetting != nil {
		in, out := &in.SQLSetting, &out.SQLSetting
		*out = new(SQLConnectionSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.MinIOSetting != nil {
		in, out := &in.MinIOSetting, &out.MinIOSetting
		*out = new(MinIOSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSettings.
func (in *ServiceSettings) DeepCopy() *ServiceSettings {
	if in == nil {
		return nil
	}
	out := new(ServiceSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SocketSetting) DeepCopyInto(out *SocketSetting) {
	*out = *in
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(Encoding)
		**out = **in
	}
	if in.BufferLength != nil {
		in, out := &in.BufferLength, &out.BufferLength
		*out = new(int64)
		**out = **in
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SocketSetting.
func (in *SocketSetting) DeepCopy() *SocketSetting {
	if in == nil {
		return nil
	}
	out := new(SocketSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSetting) DeepCopyInto(out *TCPSetting) {
	*out = *in
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.ListenPort != nil {
		in, out := &in.ListenPort, &out.ListenPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSetting.
func (in *TCPSetting) DeepCopy() *TCPSetting {
	if in == nil {
		return nil
	}
	out := new(TCPSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryRequest) DeepCopyInto(out *TelemetryRequest) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(EdgeDeviceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RawData != nil {
		in, out := &in.RawData, &out.RawData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.MQTTSetting != nil {
		in, out := &in.MQTTSetting, &out.MQTTSetting
		*out = new(MQTTSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.SQLConnectionSetting != nil {
		in, out := &in.SQLConnectionSetting, &out.SQLConnectionSetting
		*out = new(SQLConnectionSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.MinIOSetting != nil {
		in, out := &in.MinIOSetting, &out.MinIOSetting
		*out = new(MinIOSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryRequest.
func (in *TelemetryRequest) DeepCopy() *TelemetryRequest {
	if in == nil {
		return nil
	}
	out := new(TelemetryRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryService) DeepCopyInto(out *TelemetryService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryService.
func (in *TelemetryService) DeepCopy() *TelemetryService {
	if in == nil {
		return nil
	}
	out := new(TelemetryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TelemetryService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryServiceList) DeepCopyInto(out *TelemetryServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TelemetryService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryServiceList.
func (in *TelemetryServiceList) DeepCopy() *TelemetryServiceList {
	if in == nil {
		return nil
	}
	out := new(TelemetryServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TelemetryServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryServiceSpec) DeepCopyInto(out *TelemetryServiceSpec) {
	*out = *in
	if in.TelemetryServiceEndpoint != nil {
		in, out := &in.TelemetryServiceEndpoint, &out.TelemetryServiceEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ServiceSettings != nil {
		in, out := &in.ServiceSettings, &out.ServiceSettings
		*out = new(ServiceSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomMetadata != nil {
		in, out := &in.CustomMetadata, &out.CustomMetadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryServiceSpec.
func (in *TelemetryServiceSpec) DeepCopy() *TelemetryServiceSpec {
	if in == nil {
		return nil
	}
	out := new(TelemetryServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryServiceStatus) DeepCopyInto(out *TelemetryServiceStatus) {
	*out = *in
	if in.TelemetryServicePhase != nil {
		in, out := &in.TelemetryServicePhase, &out.TelemetryServicePhase
		*out = new(EdgeDevicePhase)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryServiceStatus.
func (in *TelemetryServiceStatus) DeepCopy() *TelemetryServiceStatus {
	if in == nil {
		return nil
	}
	out := new(TelemetryServiceStatus)
	in.DeepCopyInto(out)
	return out
}
