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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TCPSetting defines TCP forward settings
type TCPSetting struct {
	// +kubebuilder:default="tcp"
	NetworkType *string `json:"NetworkType,omitempty"`
	// +kubebuilder:default="8081"
	ListenPort *string `json:"ListenPort,omitempty"`
}

// MQTTSetting defines MQTT specific settings when connecting to an EdgeDevice
type MQTTSetting struct {
	MQTTTopic          *string `json:"MQTTTopic,omitempty"`
	MQTTServerAddress  *string `json:"MQTTServerAddress,omitempty"`
	MQTTServerSecret   *string `json:"MQTTServerSecret,omitempty"`
	MQTTServerUserName string  `json:"MQTTServerUserName,omitempty"`
	MQTTServerPassword string  `json:"MQTTServerPassword,omitempty"`
}

// OPCUASetting defines OPC UA specific settings when connecting to an OPC UA endpoint
type OPCUASetting struct {
	OPCUAEndpoint                   *string `json:"OPCUAEndpoint,omitempty"`
	SecurityMode                    *string `json:"SecurityMode,omitempty"`
	AuthenticationMode              *string `json:"AuthenticationMode,omitempty"`
	CertificateFileName             *string `json:"CertificateFileName,omitempty"`
	PrivateKeyFileName              *string `json:"PrivateKeyFileName,omitempty"`
	ConfigmapName                   *string `json:"ConfigmapName,omitempty"`
	IssuedToken                     *string `json:"IssuedToken,omitempty"`
	SecurityPolicy                  *string `json:"SecurityPolicy,omitempty"`
	Username                        *string `json:"Username,omitempty"`
	Password                        *string `json:"Password,omitempty"`
	ConnectionTimeoutInMilliseconds *int64  `json:"ConnectionTimeoutInMilliseconds,omitempty"`
}

type PLC4XSetting struct {
	Protocol *Plc4xProtocol `json:"protocol,omitempty"`
}

type Plc4xProtocol string

const (
	Plc4xProtocolS7          Plc4xProtocol = "s7"
	Plc4xProtocolADS         Plc4xProtocol = "ads"
	Plc4xProtocolBACnet      Plc4xProtocol = "bacnet"
	Plc4xProtocolCBus        Plc4xProtocol = "cbus"
	Plc4xProtocolEip         Plc4xProtocol = "eip"
	Plc4xProtocolKnx         Plc4xProtocol = "knx"
	Plc4xProtocolModbusAscii Plc4xProtocol = "modbus-ascii"
	Plc4xProtocolModbusRTU   Plc4xProtocol = "modbus-rtu"
	Plc4xProtocolModbusTcp   Plc4xProtocol = "modbus-tcp"
)

// SocketSetting defines Socket specific settings when connecting to an EdgeDevice
type SocketSetting struct {
	// +kubebuilder:default="utf-8"
	Encoding *Encoding `json:"encoding,omitempty"`
	// +kubebuilder:default=1024
	BufferLength *int64  `json:"bufferLength,omitempty"`
	NetworkType  *string `json:"networkType,omitempty"`
}

type SecurityMode string

const (
	// SecurityModeNoSec No security
	SecurityModeNone SecurityMode = "None"
	// SecurityModePSK Pre-Shared Key
	SecurityModeDTLS SecurityMode = "DTLS"
)

type DTLSMode string

const (
	// DTLSModePSK Pre-Shared Key
	DTLSModePSK DTLSMode = "PSK"
	// DTLSModeRPK Raw Public Key
	DTLSModeRPK DTLSMode = "RPK"
	// DTLSModeX509 X.509
	DTLSModeX509 DTLSMode = "X.509"
)

type LwM2MSetting struct {
	// +kubebuilder:validation:Required
	EndpointName string `json:"endpointName,omitempty"`
	// +kubebuilder:default="None"
	SecurityMode *SecurityMode `json:"securityMode,omitempty"`
	DTLSMode     *DTLSMode     `json:"dtlsMode,omitempty"`

	CipherSuites []CipherSuite `json:"cipherSuites,omitempty"`
	PSKIdentity  *string       `json:"pskIdentity,omitempty"`
	PSKKey       *string       `json:"pskKey,omitempty"`

	// +kubebuilder:default=-1
	PingIntervalSec int64 `json:"pingIntervalSec,omitempty"`
	// reference https://datatracker.ietf.org/doc/html/rfc7252#section-4.8.2
	// +kubebuilder:default=247
	LifeTimeSec int64 `json:"lifeTimeSec,omitempty"`
	// +kubebuilder:default=60
	// +kubebuilder:validation:Minimum=1
	UpdateIntervalSec int64 `json:"updateIntervalSec,omitempty"`
	// +kubebuilder:default=30
	ObserveIntervalSec int64 `json:"observeIntervalSec,omitempty"`
}

type CipherSuite string

// Reference:
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-4
// https://github.com/pion/dtls/blob/98a05d681d3affae2d055a70d3273cbb35425b5a/cipher_suite.go#L25-L45
const (
	// AES-128-CCM
	CipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CCM   CipherSuite = "TLS_ECDHE_ECDSA_WITH_AES_128_CCM"
	CipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8 CipherSuite = "TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8"

	// AES-128-GCM-SHA256
	CipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 CipherSuite = "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	CipherSuite_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256   CipherSuite = "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"

	CipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 CipherSuite = "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	CipherSuite_TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384   CipherSuite = "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"

	// AES-256-CBC-SHA
	CipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA CipherSuite = "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"
	CipherSuite_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA   CipherSuite = "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"

	CipherSuite_TLS_PSK_WITH_AES_128_CCM        CipherSuite = "TLS_PSK_WITH_AES_128_CCM"
	CipherSuite_TLS_PSK_WITH_AES_128_CCM_8      CipherSuite = "TLS_PSK_WITH_AES_128_CCM_8"
	CipherSuite_TLS_PSK_WITH_AES_256_CCM_8      CipherSuite = "TLS_PSK_WITH_AES_256_CCM_8"
	CipherSuite_TLS_PSK_WITH_AES_128_GCM_SHA256 CipherSuite = "TLS_PSK_WITH_AES_128_GCM_SHA256"
	CipherSuite_TLS_PSK_WITH_AES_128_CBC_SHA256 CipherSuite = "TLS_PSK_WITH_AES_128_CBC_SHA256"

	CipherSuite_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256 CipherSuite = "TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256"
)

// ProtocolSettings defines protocol settings when connecting to an EdgeDevice
type ProtocolSettings struct {
	MQTTSetting   *MQTTSetting   `json:"MQTTSetting,omitempty"`
	OPCUASetting  *OPCUASetting  `json:"OPCUASetting,omitempty"`
	SocketSetting *SocketSetting `json:"SocketSetting,omitempty"`
	PLC4XSetting  *PLC4XSetting  `json:"PLC4XSetting,omitempty"`
	TCPSetting    *TCPSetting    `json:"TCPSetting,omitempty"`
	LwM2MSetting  *LwM2MSetting  `json:"LwM2MSetting,omitempty"`
}

// GatewaySettings defines gateway settings when connecting to an EdgeDevice
type GatewaySettings struct {
	Protocol     *string       `json:"protocol,omitempty"`
	Address      *string       `json:"address,omitempty"`
	LwM2MSetting *LwM2MSetting `json:"LwM2MSetting,omitempty"`
}

// EdgeDeviceSpec defines the desired state of EdgeDevice
type EdgeDeviceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of EdgeDevice
	// Important: Run "make" to regenerate code after modifying this file

	// Sku specifies the EdgeDevice's SKU.
	Sku              *string            `json:"sku,omitempty"`
	Connection       *Connection        `json:"connection,omitempty"`
	Address          *string            `json:"address,omitempty"`
	Protocol         *Protocol          `json:"protocol,omitempty"`
	ProtocolSettings *ProtocolSettings  `json:"protocolSettings,omitempty"`
	GatewaySettings  *GatewaySettings   `json:"gatewaySettings,omitempty"`
	CustomMetadata   *map[string]string `json:"customMetadata,omitempty"`

	// TODO: add other fields like disconnectTimemoutInSeconds
}

// EdgeDeviceStatus defines the observed state of EdgeDevice
type EdgeDeviceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of EdgeDevice
	// Important: Run "make" to regenerate code after modifying this file

	// TODO: EdgeDeiveIP
	// EdgeDeviceIP is the IP address of the EdgeDevice, if it has native IP support.
	// For non-IP connections, EdgeDeviceIP is the connected EdgeNode's IP address.
	// EdgeDeviceIP *string `json:"edgedeviceip"`

	EdgeDevicePhase *EdgeDevicePhase `json:"edgedevicephase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Connection specifies the EdgeDevice-EdgeNode connection type.
type Connection string

const (
	// ConnectionEthernet String
	ConnectionEthernet Connection = "Ethernet"
)

// Protocol specifies the EdgeDevice's communication protocol.
type Protocol string

// Protocol String
const (
	ProtocolHTTP            Protocol = "HTTP"
	ProtocolHTTPCommandline Protocol = "HTTPCommandline"
	ProtocolLwM2M           Protocol = "LwM2M"
	ProtocolMQTT            Protocol = "MQTT"
	ProtocolOPCUA           Protocol = "OPCUA"
	ProtocolPLC4X           Protocol = "PLC4X"
	ProtocolSocket          Protocol = "Socket"
	ProtocolTCP             Protocol = "TCP"
	ProtocolUSB             Protocol = "USB"
)

type Encoding string

// SocketEncodingStr
const (
	UTF8 Encoding = "utf-8"
	HEX  Encoding = "hex"
)

// EdgeDevicePhase is a simple, high-level summary of where the EdgeDevice is in its lifecycle.
type EdgeDevicePhase string

const (
	// EdgeDevicePending means the EdgeDevice has been accepted by the system but not ready yet.
	EdgeDevicePending EdgeDevicePhase = "Pending"
	// EdgeDeviceRunning means the EdgeDevice is able to interact with the system and user applications.
	EdgeDeviceRunning EdgeDevicePhase = "Running"
	// EdgeDeviceFailed means the EdgeDevice is failed state.
	EdgeDeviceFailed EdgeDevicePhase = "Failed"
	// EdgeDeviceUnknown means the EdgeDevice's info could not be obtained.
	// This is typically due to communication failures.
	EdgeDeviceUnknown EdgeDevicePhase = "Unknown"
)

//+kubebuilder:object:root=true

// EdgeDevice is the Schema for the edgedevices API
type EdgeDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeDeviceSpec   `json:"spec,omitempty"`
	Status EdgeDeviceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EdgeDeviceList contains a list of EdgeDevice
type EdgeDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeDevice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeDevice{}, &EdgeDeviceList{})
}
