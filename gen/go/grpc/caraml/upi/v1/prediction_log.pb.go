// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: caraml/upi/v1/prediction_log.proto

package upiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PredictionLog stores information of prediction request handled by a specific model version.
type PredictionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of prediction.
	PredictionId string `protobuf:"bytes,1,opt,name=prediction_id,json=predictionId,proto3" json:"prediction_id,omitempty"`
	// Target name / concept to be predicted by the prediction.
	TargetName string `protobuf:"bytes,2,opt,name=target_name,json=targetName,proto3" json:"target_name,omitempty"`
	// Project name that host the model performing prediction.
	ProjectName string `protobuf:"bytes,3,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// Model name perforing the prediction.
	ModelName string `protobuf:"bytes,4,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// Model version performing the prediction.
	ModelVersion string `protobuf:"bytes,5,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// Input of the prediction process
	Input *ModelInput `protobuf:"bytes,10,opt,name=input,proto3" json:"input,omitempty"`
	// Output of the prediction process
	Output *ModelOutput `protobuf:"bytes,11,opt,name=output,proto3" json:"output,omitempty"`
	// Timestamp of the corresponding prediction request
	RequestTimestamp *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=request_timestamp,json=requestTimestamp,proto3" json:"request_timestamp,omitempty"`
	// Schema version of raw_features, features, entities, and prediction results fields are formatted
	// I.e. for version 1, all of those fields will be formatted as a modified JSON SPLIT representation of a table
	TableSchemaVersion uint32 `protobuf:"varint,100,opt,name=table_schema_version,json=tableSchemaVersion,proto3" json:"table_schema_version,omitempty"`
}

func (x *PredictionLog) Reset() {
	*x = PredictionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionLog) ProtoMessage() {}

func (x *PredictionLog) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionLog.ProtoReflect.Descriptor instead.
func (*PredictionLog) Descriptor() ([]byte, []int) {
	return file_caraml_upi_v1_prediction_log_proto_rawDescGZIP(), []int{0}
}

func (x *PredictionLog) GetPredictionId() string {
	if x != nil {
		return x.PredictionId
	}
	return ""
}

func (x *PredictionLog) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *PredictionLog) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *PredictionLog) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *PredictionLog) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *PredictionLog) GetInput() *ModelInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *PredictionLog) GetOutput() *ModelOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *PredictionLog) GetRequestTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestTimestamp
	}
	return nil
}

func (x *PredictionLog) GetTableSchemaVersion() uint32 {
	if x != nil {
		return x.TableSchemaVersion
	}
	return 0
}

// Model input stores information of all input for prediction process.
// The information in model input are extracted from the prediction request received by model.
type ModelInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON-representation of features table. "table_schema_version" field describe the encoding of this field.
	FeaturesTable *structpb.Struct `protobuf:"bytes,1,opt,name=features_table,json=featuresTable,proto3" json:"features_table,omitempty"`
	// JSON-representation of entities table. "table_schema_version" field describe the encoding of this field.
	EntitiesTable *structpb.Struct `protobuf:"bytes,2,opt,name=entities_table,json=entitiesTable,proto3" json:"entities_table,omitempty"`
	// JSON-representation of raw_features table. "table_schema_version" field describe the encoding of this field.
	RawFeatures *structpb.Struct `protobuf:"bytes,3,opt,name=raw_features,json=rawFeatures,proto3" json:"raw_features,omitempty"`
	// Context of the prediction request.
	PredictionContext []*Variable `protobuf:"bytes,4,rep,name=prediction_context,json=predictionContext,proto3" json:"prediction_context,omitempty"`
	// map containing request headers/metadata
	Headers map[string]string `protobuf:"bytes,10,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ModelInput) Reset() {
	*x = ModelInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelInput) ProtoMessage() {}

func (x *ModelInput) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelInput.ProtoReflect.Descriptor instead.
func (*ModelInput) Descriptor() ([]byte, []int) {
	return file_caraml_upi_v1_prediction_log_proto_rawDescGZIP(), []int{1}
}

func (x *ModelInput) GetFeaturesTable() *structpb.Struct {
	if x != nil {
		return x.FeaturesTable
	}
	return nil
}

func (x *ModelInput) GetEntitiesTable() *structpb.Struct {
	if x != nil {
		return x.EntitiesTable
	}
	return nil
}

func (x *ModelInput) GetRawFeatures() *structpb.Struct {
	if x != nil {
		return x.RawFeatures
	}
	return nil
}

func (x *ModelInput) GetPredictionContext() []*Variable {
	if x != nil {
		return x.PredictionContext
	}
	return nil
}

func (x *ModelInput) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Model output stores information of all output produced from prediction process.
// Model output is extracted from the prediction response sent by model.
type ModelOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON-representation of prediction result table. "table_schema_version" field describe the encoding of this field.
	PredictionResultsTable *structpb.Struct `protobuf:"bytes,1,opt,name=prediction_results_table,json=predictionResultsTable,proto3" json:"prediction_results_table,omitempty"`
	// Context of the prediction response.
	PredictionContext []*Variable `protobuf:"bytes,2,rep,name=prediction_context,json=predictionContext,proto3" json:"prediction_context,omitempty"`
	// map containing response headers/metadata
	Headers map[string]string `protobuf:"bytes,10,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// grpc status of the response from model (see https://grpc.github.io/grpc/core/md_doc_statuscodes.html)
	Status uint32 `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	// grpc message
	Message string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ModelOutput) Reset() {
	*x = ModelOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelOutput) ProtoMessage() {}

func (x *ModelOutput) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_upi_v1_prediction_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelOutput.ProtoReflect.Descriptor instead.
func (*ModelOutput) Descriptor() ([]byte, []int) {
	return file_caraml_upi_v1_prediction_log_proto_rawDescGZIP(), []int{2}
}

func (x *ModelOutput) GetPredictionResultsTable() *structpb.Struct {
	if x != nil {
		return x.PredictionResultsTable
	}
	return nil
}

func (x *ModelOutput) GetPredictionContext() []*Variable {
	if x != nil {
		return x.PredictionContext
	}
	return nil
}

func (x *ModelOutput) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ModelOutput) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ModelOutput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_caraml_upi_v1_prediction_log_proto protoreflect.FileDescriptor

var file_caraml_upi_v1_prediction_log_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x75, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x75, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x03, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x32, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x30, 0x0a,
	0x14, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x8e, 0x03, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3e,
	0x0a, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3e,
	0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3a,
	0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x72,
	0x61, 0x77, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x12, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e,
	0x75, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x51, 0x0a, 0x18, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x16, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x11, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xce, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x75, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x75, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x75, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x55, 0x58, 0xaa,
	0x02, 0x0d, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x55, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0d, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x5c, 0x55, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x19, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x5c, 0x55, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x43, 0x61,
	0x72, 0x61, 0x6d, 0x6c, 0x3a, 0x3a, 0x55, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_caraml_upi_v1_prediction_log_proto_rawDescOnce sync.Once
	file_caraml_upi_v1_prediction_log_proto_rawDescData = file_caraml_upi_v1_prediction_log_proto_rawDesc
)

func file_caraml_upi_v1_prediction_log_proto_rawDescGZIP() []byte {
	file_caraml_upi_v1_prediction_log_proto_rawDescOnce.Do(func() {
		file_caraml_upi_v1_prediction_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_caraml_upi_v1_prediction_log_proto_rawDescData)
	})
	return file_caraml_upi_v1_prediction_log_proto_rawDescData
}

var file_caraml_upi_v1_prediction_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_caraml_upi_v1_prediction_log_proto_goTypes = []interface{}{
	(*PredictionLog)(nil),         // 0: caraml.upi.v1.PredictionLog
	(*ModelInput)(nil),            // 1: caraml.upi.v1.ModelInput
	(*ModelOutput)(nil),           // 2: caraml.upi.v1.ModelOutput
	nil,                           // 3: caraml.upi.v1.ModelInput.HeadersEntry
	nil,                           // 4: caraml.upi.v1.ModelOutput.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
	(*Variable)(nil),              // 7: caraml.upi.v1.Variable
}
var file_caraml_upi_v1_prediction_log_proto_depIdxs = []int32{
	1,  // 0: caraml.upi.v1.PredictionLog.input:type_name -> caraml.upi.v1.ModelInput
	2,  // 1: caraml.upi.v1.PredictionLog.output:type_name -> caraml.upi.v1.ModelOutput
	5,  // 2: caraml.upi.v1.PredictionLog.request_timestamp:type_name -> google.protobuf.Timestamp
	6,  // 3: caraml.upi.v1.ModelInput.features_table:type_name -> google.protobuf.Struct
	6,  // 4: caraml.upi.v1.ModelInput.entities_table:type_name -> google.protobuf.Struct
	6,  // 5: caraml.upi.v1.ModelInput.raw_features:type_name -> google.protobuf.Struct
	7,  // 6: caraml.upi.v1.ModelInput.prediction_context:type_name -> caraml.upi.v1.Variable
	3,  // 7: caraml.upi.v1.ModelInput.headers:type_name -> caraml.upi.v1.ModelInput.HeadersEntry
	6,  // 8: caraml.upi.v1.ModelOutput.prediction_results_table:type_name -> google.protobuf.Struct
	7,  // 9: caraml.upi.v1.ModelOutput.prediction_context:type_name -> caraml.upi.v1.Variable
	4,  // 10: caraml.upi.v1.ModelOutput.headers:type_name -> caraml.upi.v1.ModelOutput.HeadersEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_caraml_upi_v1_prediction_log_proto_init() }
func file_caraml_upi_v1_prediction_log_proto_init() {
	if File_caraml_upi_v1_prediction_log_proto != nil {
		return
	}
	file_caraml_upi_v1_variable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_caraml_upi_v1_prediction_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictionLog); i {
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
		file_caraml_upi_v1_prediction_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelInput); i {
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
		file_caraml_upi_v1_prediction_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelOutput); i {
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
			RawDescriptor: file_caraml_upi_v1_prediction_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_caraml_upi_v1_prediction_log_proto_goTypes,
		DependencyIndexes: file_caraml_upi_v1_prediction_log_proto_depIdxs,
		MessageInfos:      file_caraml_upi_v1_prediction_log_proto_msgTypes,
	}.Build()
	File_caraml_upi_v1_prediction_log_proto = out.File
	file_caraml_upi_v1_prediction_log_proto_rawDesc = nil
	file_caraml_upi_v1_prediction_log_proto_goTypes = nil
	file_caraml_upi_v1_prediction_log_proto_depIdxs = nil
}
