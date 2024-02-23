// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package reporter

import (
	_ "cosmossdk.io/api/amino"
	v1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_OracleReporter              protoreflect.MessageDescriptor
	fd_OracleReporter_reporter     protoreflect.FieldDescriptor
	fd_OracleReporter_total_tokens protoreflect.FieldDescriptor
	fd_OracleReporter_commission   protoreflect.FieldDescriptor
	fd_OracleReporter_jailed       protoreflect.FieldDescriptor
	fd_OracleReporter_jailed_until protoreflect.FieldDescriptor
)

func init() {
	file_layer_reporter_oracle_reporter_proto_init()
	md_OracleReporter = File_layer_reporter_oracle_reporter_proto.Messages().ByName("OracleReporter")
	fd_OracleReporter_reporter = md_OracleReporter.Fields().ByName("reporter")
	fd_OracleReporter_total_tokens = md_OracleReporter.Fields().ByName("total_tokens")
	fd_OracleReporter_commission = md_OracleReporter.Fields().ByName("commission")
	fd_OracleReporter_jailed = md_OracleReporter.Fields().ByName("jailed")
	fd_OracleReporter_jailed_until = md_OracleReporter.Fields().ByName("jailed_until")
}

var _ protoreflect.Message = (*fastReflection_OracleReporter)(nil)

type fastReflection_OracleReporter OracleReporter

func (x *OracleReporter) ProtoReflect() protoreflect.Message {
	return (*fastReflection_OracleReporter)(x)
}

func (x *OracleReporter) slowProtoReflect() protoreflect.Message {
	mi := &file_layer_reporter_oracle_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_OracleReporter_messageType fastReflection_OracleReporter_messageType
var _ protoreflect.MessageType = fastReflection_OracleReporter_messageType{}

type fastReflection_OracleReporter_messageType struct{}

func (x fastReflection_OracleReporter_messageType) Zero() protoreflect.Message {
	return (*fastReflection_OracleReporter)(nil)
}
func (x fastReflection_OracleReporter_messageType) New() protoreflect.Message {
	return new(fastReflection_OracleReporter)
}
func (x fastReflection_OracleReporter_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_OracleReporter
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_OracleReporter) Descriptor() protoreflect.MessageDescriptor {
	return md_OracleReporter
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_OracleReporter) Type() protoreflect.MessageType {
	return _fastReflection_OracleReporter_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_OracleReporter) New() protoreflect.Message {
	return new(fastReflection_OracleReporter)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_OracleReporter) Interface() protoreflect.ProtoMessage {
	return (*OracleReporter)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_OracleReporter) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Reporter != "" {
		value := protoreflect.ValueOfString(x.Reporter)
		if !f(fd_OracleReporter_reporter, value) {
			return
		}
	}
	if x.TotalTokens != "" {
		value := protoreflect.ValueOfString(x.TotalTokens)
		if !f(fd_OracleReporter_total_tokens, value) {
			return
		}
	}
	if x.Commission != nil {
		value := protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
		if !f(fd_OracleReporter_commission, value) {
			return
		}
	}
	if x.Jailed != false {
		value := protoreflect.ValueOfBool(x.Jailed)
		if !f(fd_OracleReporter_jailed, value) {
			return
		}
	}
	if x.JailedUntil != nil {
		value := protoreflect.ValueOfMessage(x.JailedUntil.ProtoReflect())
		if !f(fd_OracleReporter_jailed_until, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_OracleReporter) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "layer.reporter.OracleReporter.reporter":
		return x.Reporter != ""
	case "layer.reporter.OracleReporter.total_tokens":
		return x.TotalTokens != ""
	case "layer.reporter.OracleReporter.commission":
		return x.Commission != nil
	case "layer.reporter.OracleReporter.jailed":
		return x.Jailed != false
	case "layer.reporter.OracleReporter.jailed_until":
		return x.JailedUntil != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OracleReporter) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "layer.reporter.OracleReporter.reporter":
		x.Reporter = ""
	case "layer.reporter.OracleReporter.total_tokens":
		x.TotalTokens = ""
	case "layer.reporter.OracleReporter.commission":
		x.Commission = nil
	case "layer.reporter.OracleReporter.jailed":
		x.Jailed = false
	case "layer.reporter.OracleReporter.jailed_until":
		x.JailedUntil = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_OracleReporter) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "layer.reporter.OracleReporter.reporter":
		value := x.Reporter
		return protoreflect.ValueOfString(value)
	case "layer.reporter.OracleReporter.total_tokens":
		value := x.TotalTokens
		return protoreflect.ValueOfString(value)
	case "layer.reporter.OracleReporter.commission":
		value := x.Commission
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "layer.reporter.OracleReporter.jailed":
		value := x.Jailed
		return protoreflect.ValueOfBool(value)
	case "layer.reporter.OracleReporter.jailed_until":
		value := x.JailedUntil
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OracleReporter) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "layer.reporter.OracleReporter.reporter":
		x.Reporter = value.Interface().(string)
	case "layer.reporter.OracleReporter.total_tokens":
		x.TotalTokens = value.Interface().(string)
	case "layer.reporter.OracleReporter.commission":
		x.Commission = value.Message().Interface().(*v1beta1.Commission)
	case "layer.reporter.OracleReporter.jailed":
		x.Jailed = value.Bool()
	case "layer.reporter.OracleReporter.jailed_until":
		x.JailedUntil = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OracleReporter) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "layer.reporter.OracleReporter.commission":
		if x.Commission == nil {
			x.Commission = new(v1beta1.Commission)
		}
		return protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
	case "layer.reporter.OracleReporter.jailed_until":
		if x.JailedUntil == nil {
			x.JailedUntil = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.JailedUntil.ProtoReflect())
	case "layer.reporter.OracleReporter.reporter":
		panic(fmt.Errorf("field reporter of message layer.reporter.OracleReporter is not mutable"))
	case "layer.reporter.OracleReporter.total_tokens":
		panic(fmt.Errorf("field total_tokens of message layer.reporter.OracleReporter is not mutable"))
	case "layer.reporter.OracleReporter.jailed":
		panic(fmt.Errorf("field jailed of message layer.reporter.OracleReporter is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_OracleReporter) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "layer.reporter.OracleReporter.reporter":
		return protoreflect.ValueOfString("")
	case "layer.reporter.OracleReporter.total_tokens":
		return protoreflect.ValueOfString("")
	case "layer.reporter.OracleReporter.commission":
		m := new(v1beta1.Commission)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "layer.reporter.OracleReporter.jailed":
		return protoreflect.ValueOfBool(false)
	case "layer.reporter.OracleReporter.jailed_until":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: layer.reporter.OracleReporter"))
		}
		panic(fmt.Errorf("message layer.reporter.OracleReporter does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_OracleReporter) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in layer.reporter.OracleReporter", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_OracleReporter) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_OracleReporter) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_OracleReporter) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_OracleReporter) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*OracleReporter)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Reporter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TotalTokens)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Commission != nil {
			l = options.Size(x.Commission)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Jailed {
			n += 2
		}
		if x.JailedUntil != nil {
			l = options.Size(x.JailedUntil)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*OracleReporter)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.JailedUntil != nil {
			encoded, err := options.Marshal(x.JailedUntil)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Jailed {
			i--
			if x.Jailed {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
		}
		if x.Commission != nil {
			encoded, err := options.Marshal(x.Commission)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TotalTokens) > 0 {
			i -= len(x.TotalTokens)
			copy(dAtA[i:], x.TotalTokens)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TotalTokens)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Reporter) > 0 {
			i -= len(x.Reporter)
			copy(dAtA[i:], x.Reporter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Reporter)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*OracleReporter)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: OracleReporter: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: OracleReporter: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Reporter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalTokens", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TotalTokens = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Commission == nil {
					x.Commission = &v1beta1.Commission{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Commission); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Jailed = bool(v != 0)
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.JailedUntil == nil {
					x.JailedUntil = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.JailedUntil); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: layer/reporter/oracle_reporter.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OracleReporter is the struct that holds the data for a reporter
type OracleReporter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reporter is the address of the reporter
	Reporter string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	// tokens is the amount of tokens the reporter has
	TotalTokens string `protobuf:"bytes,2,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
	// commission for the reporter
	Commission *v1beta1.Commission `protobuf:"bytes,3,opt,name=commission,proto3" json:"commission,omitempty"`
	// jailed is a bool weather the reporter is jailed or not
	Jailed bool `protobuf:"varint,4,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// jailed_until is the time the reporter is jailed until
	JailedUntil *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=jailed_until,json=jailedUntil,proto3" json:"jailed_until,omitempty"`
}

func (x *OracleReporter) Reset() {
	*x = OracleReporter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer_reporter_oracle_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleReporter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleReporter) ProtoMessage() {}

// Deprecated: Use OracleReporter.ProtoReflect.Descriptor instead.
func (*OracleReporter) Descriptor() ([]byte, []int) {
	return file_layer_reporter_oracle_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *OracleReporter) GetReporter() string {
	if x != nil {
		return x.Reporter
	}
	return ""
}

func (x *OracleReporter) GetTotalTokens() string {
	if x != nil {
		return x.TotalTokens
	}
	return ""
}

func (x *OracleReporter) GetCommission() *v1beta1.Commission {
	if x != nil {
		return x.Commission
	}
	return nil
}

func (x *OracleReporter) GetJailed() bool {
	if x != nil {
		return x.Jailed
	}
	return false
}

func (x *OracleReporter) GetJailedUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.JailedUntil
	}
	return nil
}

var File_layer_reporter_oracle_reporter_proto protoreflect.FileDescriptor

var file_layer_reporter_oracle_reporter_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d,
	0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x12, 0x4e, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x15,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74,
	0x68, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x42, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x0c,
	0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0d,
	0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x0b, 0x6a,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x42, 0xa3, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x42, 0x13, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4c, 0x52, 0x58, 0xaa,
	0x02, 0x0e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0xca, 0x02, 0x0e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0xe2, 0x02, 0x1a, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_layer_reporter_oracle_reporter_proto_rawDescOnce sync.Once
	file_layer_reporter_oracle_reporter_proto_rawDescData = file_layer_reporter_oracle_reporter_proto_rawDesc
)

func file_layer_reporter_oracle_reporter_proto_rawDescGZIP() []byte {
	file_layer_reporter_oracle_reporter_proto_rawDescOnce.Do(func() {
		file_layer_reporter_oracle_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_layer_reporter_oracle_reporter_proto_rawDescData)
	})
	return file_layer_reporter_oracle_reporter_proto_rawDescData
}

var file_layer_reporter_oracle_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_layer_reporter_oracle_reporter_proto_goTypes = []interface{}{
	(*OracleReporter)(nil),        // 0: layer.reporter.OracleReporter
	(*v1beta1.Commission)(nil),    // 1: cosmos.staking.v1beta1.Commission
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_layer_reporter_oracle_reporter_proto_depIdxs = []int32{
	1, // 0: layer.reporter.OracleReporter.commission:type_name -> cosmos.staking.v1beta1.Commission
	2, // 1: layer.reporter.OracleReporter.jailed_until:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_layer_reporter_oracle_reporter_proto_init() }
func file_layer_reporter_oracle_reporter_proto_init() {
	if File_layer_reporter_oracle_reporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layer_reporter_oracle_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleReporter); i {
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
			RawDescriptor: file_layer_reporter_oracle_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_layer_reporter_oracle_reporter_proto_goTypes,
		DependencyIndexes: file_layer_reporter_oracle_reporter_proto_depIdxs,
		MessageInfos:      file_layer_reporter_oracle_reporter_proto_msgTypes,
	}.Build()
	File_layer_reporter_oracle_reporter_proto = out.File
	file_layer_reporter_oracle_reporter_proto_rawDesc = nil
	file_layer_reporter_oracle_reporter_proto_goTypes = nil
	file_layer_reporter_oracle_reporter_proto_depIdxs = nil
}
