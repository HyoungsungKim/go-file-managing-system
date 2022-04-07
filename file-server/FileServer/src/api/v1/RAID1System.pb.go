//
//RAID1System.proto includes rpc to run RAID1 system.
//When client send a file to filer server, file server send file to RAID1 server using gRPC.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: RAID1System.proto

package v1

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

//
//message FileInfo gives a file directory of file, which is stored in the File server.
//- message FileInfo includes two filelds: fileName, filePath.
//- Using these fields, RAID1 server has same URI of file, which is stored in file server.
//
//Assume file is stored in `./storage/:dirID/fileName`,
//fileName field denotes, `fileName` of fileName,
//filePath field denotes, `:dirID`.
//
//For example, if `helloworld.txt` is stored in './storage/ZGlyMS9kaXIyL2RpcjM=/helloworld.txt
//fileName field is `helloworld.txt`, filePath field is `ZGlyMS9kaXIyL2RpcjM=`
//message FileInfo does not includes root directory, such as `./storage`, beucase root directory of RAID1 server is not guranteeded.
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RAID1System_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RAID1System_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_RAID1System_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfo) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

//
//message FileStreamRequest has a chunkData field, which is used for streaming file.
//Files are streamed by message FileStreamRequest
type FileStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkData []byte `protobuf:"bytes,1,opt,name=chunkData,proto3" json:"chunkData,omitempty"`
}

func (x *FileStreamRequest) Reset() {
	*x = FileStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RAID1System_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStreamRequest) ProtoMessage() {}

func (x *FileStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_RAID1System_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStreamRequest.ProtoReflect.Descriptor instead.
func (*FileStreamRequest) Descriptor() ([]byte, []int) {
	return file_RAID1System_proto_rawDescGZIP(), []int{1}
}

func (x *FileStreamRequest) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type DeleteFileSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteFileSignal bool `protobuf:"varint,1,opt,name=deleteFileSignal,proto3" json:"deleteFileSignal,omitempty"`
}

func (x *DeleteFileSignal) Reset() {
	*x = DeleteFileSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RAID1System_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileSignal) ProtoMessage() {}

func (x *DeleteFileSignal) ProtoReflect() protoreflect.Message {
	mi := &file_RAID1System_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileSignal.ProtoReflect.Descriptor instead.
func (*DeleteFileSignal) Descriptor() ([]byte, []int) {
	return file_RAID1System_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFileSignal) GetDeleteFileSignal() bool {
	if x != nil {
		return x.DeleteFileSignal
	}
	return false
}

//
//message Ack is defined to give a acknowledge to server from client.
//When RAID1 server gets a Fileinfo or FileStreamRequest message,
//RAID1 server send back message ACK to server.
//
//field ackStatusCode has same status code of http status code
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckStatusCode    uint32 `protobuf:"varint,1,opt,name=ackStatusCode,proto3" json:"ackStatusCode,omitempty"`
	AckStatusMessage string `protobuf:"bytes,2,opt,name=ackStatusMessage,proto3" json:"ackStatusMessage,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RAID1System_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_RAID1System_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_RAID1System_proto_rawDescGZIP(), []int{3}
}

func (x *Ack) GetAckStatusCode() uint32 {
	if x != nil {
		return x.AckStatusCode
	}
	return 0
}

func (x *Ack) GetAckStatusMessage() string {
	if x != nil {
		return x.AckStatusMessage
	}
	return ""
}

var File_RAID1System_proto protoreflect.FileDescriptor

var file_RAID1System_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x41, 0x49, 0x44, 0x31, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x31, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x2a,
	0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22, 0x57, 0x0a, 0x03, 0x41, 0x63,
	0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x28, 0x01, 0x12, 0x28, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04,
	0x2e, 0x41, 0x63, 0x6b, 0x28, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x28, 0x01, 0x42,
	0x1b, 0x5a, 0x19, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RAID1System_proto_rawDescOnce sync.Once
	file_RAID1System_proto_rawDescData = file_RAID1System_proto_rawDesc
)

func file_RAID1System_proto_rawDescGZIP() []byte {
	file_RAID1System_proto_rawDescOnce.Do(func() {
		file_RAID1System_proto_rawDescData = protoimpl.X.CompressGZIP(file_RAID1System_proto_rawDescData)
	})
	return file_RAID1System_proto_rawDescData
}

var file_RAID1System_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_RAID1System_proto_goTypes = []interface{}{
	(*FileInfo)(nil),          // 0: FileInfo
	(*FileStreamRequest)(nil), // 1: FileStreamRequest
	(*DeleteFileSignal)(nil),  // 2: DeleteFileSignal
	(*Ack)(nil),               // 3: Ack
}
var file_RAID1System_proto_depIdxs = []int32{
	0, // 0: FileStreamService.SendFileInfo:input_type -> FileInfo
	1, // 1: FileStreamService.StreamFile:input_type -> FileStreamRequest
	2, // 2: FileStreamService.DeleteFile:input_type -> DeleteFileSignal
	3, // 3: FileStreamService.SendFileInfo:output_type -> Ack
	3, // 4: FileStreamService.StreamFile:output_type -> Ack
	3, // 5: FileStreamService.DeleteFile:output_type -> Ack
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RAID1System_proto_init() }
func file_RAID1System_proto_init() {
	if File_RAID1System_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RAID1System_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_RAID1System_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStreamRequest); i {
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
		file_RAID1System_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileSignal); i {
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
		file_RAID1System_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
			RawDescriptor: file_RAID1System_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RAID1System_proto_goTypes,
		DependencyIndexes: file_RAID1System_proto_depIdxs,
		MessageInfos:      file_RAID1System_proto_msgTypes,
	}.Build()
	File_RAID1System_proto = out.File
	file_RAID1System_proto_rawDesc = nil
	file_RAID1System_proto_goTypes = nil
	file_RAID1System_proto_depIdxs = nil
}
