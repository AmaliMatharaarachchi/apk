//  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 LLC. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/apkmgt/application.proto

package apkmgt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner         string            `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Attributes    map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subscriber    string            `protobuf:"bytes,5,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	Organization  string            `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
	Subscriptions []*Subscription   `protobuf:"bytes,7,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	ConsumerKeys  []*ConsumerKey    `protobuf:"bytes,8,rep,name=consumerKeys,proto3" json:"consumerKeys,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_apkmgt_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Application) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Application) GetSubscriber() string {
	if x != nil {
		return x.Subscriber
	}
	return ""
}

func (x *Application) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Application) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *Application) GetConsumerKeys() []*ConsumerKey {
	if x != nil {
		return x.ConsumerKeys
	}
	return nil
}

type ConsumerKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KeyManager string `protobuf:"bytes,2,opt,name=keyManager,proto3" json:"keyManager,omitempty"`
}

func (x *ConsumerKey) Reset() {
	*x = ConsumerKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerKey) ProtoMessage() {}

func (x *ConsumerKey) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerKey.ProtoReflect.Descriptor instead.
func (*ConsumerKey) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_apkmgt_application_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ConsumerKey) GetKeyManager() string {
	if x != nil {
		return x.KeyManager
	}
	return ""
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid               string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ApiUuid            string `protobuf:"bytes,2,opt,name=apiUuid,proto3" json:"apiUuid,omitempty"`
	PolicyId           string `protobuf:"bytes,3,opt,name=policyId,proto3" json:"policyId,omitempty"`
	SubscriptionStatus string `protobuf:"bytes,4,opt,name=subscriptionStatus,proto3" json:"subscriptionStatus,omitempty"`
	Organization       string `protobuf:"bytes,5,opt,name=organization,proto3" json:"organization,omitempty"`
	CreatedBy          string `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_apkmgt_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_apkmgt_application_proto_rawDescGZIP(), []int{2}
}

func (x *Subscription) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Subscription) GetApiUuid() string {
	if x != nil {
		return x.ApiUuid
	}
	return ""
}

func (x *Subscription) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *Subscription) GetSubscriptionStatus() string {
	if x != nil {
		return x.SubscriptionStatus
	}
	return ""
}

func (x *Subscription) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Subscription) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

var File_wso2_discovery_apkmgt_application_proto protoreflect.FileDescriptor

var file_wso2_discovery_apkmgt_application_proto_rawDesc = []byte{
	0x0a, 0x27, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x6b, 0x6d, 0x67, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x6b, 0x6d, 0x67, 0x74,
	0x22, 0xb5, 0x03, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x6b, 0x6d, 0x67, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x6b, 0x6d, 0x67, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x46, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x6b, 0x6d, 0x67, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x65, 0x79,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b,
	0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x70, 0x69, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x70, 0x69, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x42, 0x89, 0x01, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x6b,
	0x6d, 0x67, 0x74, 0x42, 0x16, 0x41, 0x70, 0x6b, 0x4d, 0x67, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x6b, 0x6d, 0x67, 0x74, 0x3b, 0x61, 0x70, 0x6b, 0x6d,
	0x67, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_apkmgt_application_proto_rawDescOnce sync.Once
	file_wso2_discovery_apkmgt_application_proto_rawDescData = file_wso2_discovery_apkmgt_application_proto_rawDesc
)

func file_wso2_discovery_apkmgt_application_proto_rawDescGZIP() []byte {
	file_wso2_discovery_apkmgt_application_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_apkmgt_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_apkmgt_application_proto_rawDescData)
	})
	return file_wso2_discovery_apkmgt_application_proto_rawDescData
}

var file_wso2_discovery_apkmgt_application_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wso2_discovery_apkmgt_application_proto_goTypes = []interface{}{
	(*Application)(nil),  // 0: wso2.discovery.apkmgt.Application
	(*ConsumerKey)(nil),  // 1: wso2.discovery.apkmgt.ConsumerKey
	(*Subscription)(nil), // 2: wso2.discovery.apkmgt.Subscription
	nil,                  // 3: wso2.discovery.apkmgt.Application.AttributesEntry
}
var file_wso2_discovery_apkmgt_application_proto_depIdxs = []int32{
	3, // 0: wso2.discovery.apkmgt.Application.attributes:type_name -> wso2.discovery.apkmgt.Application.AttributesEntry
	2, // 1: wso2.discovery.apkmgt.Application.subscriptions:type_name -> wso2.discovery.apkmgt.Subscription
	1, // 2: wso2.discovery.apkmgt.Application.consumerKeys:type_name -> wso2.discovery.apkmgt.ConsumerKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wso2_discovery_apkmgt_application_proto_init() }
func file_wso2_discovery_apkmgt_application_proto_init() {
	if File_wso2_discovery_apkmgt_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_apkmgt_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wso2_discovery_apkmgt_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wso2_discovery_apkmgt_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_apkmgt_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_apkmgt_application_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_apkmgt_application_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_apkmgt_application_proto_msgTypes,
	}.Build()
	File_wso2_discovery_apkmgt_application_proto = out.File
	file_wso2_discovery_apkmgt_application_proto_rawDesc = nil
	file_wso2_discovery_apkmgt_application_proto_goTypes = nil
	file_wso2_discovery_apkmgt_application_proto_depIdxs = nil
}
