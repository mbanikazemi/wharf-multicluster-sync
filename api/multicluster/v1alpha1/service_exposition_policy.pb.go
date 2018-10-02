// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multicluster/v1alpha1/service_exposition_policy.proto

package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `ServiceExpositionPolicy` describes an exposition policy for services
// available on the cluster where the policy is deployed. The cluster or
// mesh operator creates this policy object to selectively choose the specific
// service to be available to remote cluster. Each entry for exposed service is
// also accompanied with a list of cluster IDs that can access it. This ensures
// that only identified clusters can access the exposed services and only the
// services selected to be exposed.
//
// The following example exposes v1 of ServiceA from the cluster where it is
// deployed as service FooA to two remote clusters with IDs `clusterA` and
// `clusterB`.
//
// ```yaml
// apiVersion: multicluster.istio.io/v1alpha1
// kind: ServiceExpositionPolicy
// metadata:
//   name: sample1
//   namespace: mynamespace
// spec:
//   exposed:
//   - name: ServiceA
//     alias: FooA
//     subset: v1
//     port: 9080
//     clusters:
//     - clusterA
//     - clusterB
// ```
type ServiceExpositionPolicy struct {
	// REQUIRED: One or more exposed services. It is a list of services that
	// will be exposed by the cluster where this policy is deployed along with
	// the details for each service (e.g. alias, subset, etc).
	Exposed []*ServiceExpositionPolicy_ExposedService `protobuf:"bytes,1,rep,name=exposed" json:"exposed,omitempty"`
}

func (m *ServiceExpositionPolicy) Reset()         { *m = ServiceExpositionPolicy{} }
func (m *ServiceExpositionPolicy) String() string { return proto.CompactTextString(m) }
func (*ServiceExpositionPolicy) ProtoMessage()    {}
func (*ServiceExpositionPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptorServiceExpositionPolicy, []int{0}
}

func (m *ServiceExpositionPolicy) GetExposed() []*ServiceExpositionPolicy_ExposedService {
	if m != nil {
		return m.Exposed
	}
	return nil
}

// A single exposed service policy holds any information necessary for the
// configuration of both acceptor and donator clusters.
type ServiceExpositionPolicy_ExposedService struct {
	// REQUIRED: The name of the service to be exposed.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This is an alias that can be used for the exposed name of the service.
	// It allows the operator to hide names of in-cluster services and choose
	// descriptive names that acceptor clusters operators may find them more
	// informative.
	// This is an optional field. If not specified, the service name will be
	// used as the exposed service name.
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	// `subset` allows the operator to choose a specific subset (service
	// version) in cases when there are multiple subsets available for the
	// exposed service. Applicable only to services within the mesh. The subset
	//  must be defined in a corresponding DestinationRule.
	Subset string `protobuf:"bytes,3,opt,name=subset,proto3" json:"subset,omitempty"`
	// The port of the exposed service.
	// TODO: consider adding support for multiple ports, their types and names.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// A list of cluster IDs that are allowed to call the service exposed by
	// this cluster.
	Clusters []string `protobuf:"bytes,5,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ServiceExpositionPolicy_ExposedService) Reset() {
	*m = ServiceExpositionPolicy_ExposedService{}
}
func (m *ServiceExpositionPolicy_ExposedService) String() string { return proto.CompactTextString(m) }
func (*ServiceExpositionPolicy_ExposedService) ProtoMessage()    {}
func (*ServiceExpositionPolicy_ExposedService) Descriptor() ([]byte, []int) {
	return fileDescriptorServiceExpositionPolicy, []int{0, 0}
}

func (m *ServiceExpositionPolicy_ExposedService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceExpositionPolicy_ExposedService) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *ServiceExpositionPolicy_ExposedService) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *ServiceExpositionPolicy_ExposedService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServiceExpositionPolicy_ExposedService) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceExpositionPolicy)(nil), "istio.multicluster.v1alpha1.ServiceExpositionPolicy")
	proto.RegisterType((*ServiceExpositionPolicy_ExposedService)(nil), "istio.multicluster.v1alpha1.ServiceExpositionPolicy.ExposedService")
}
func (m *ServiceExpositionPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceExpositionPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Exposed) > 0 {
		for _, msg := range m.Exposed {
			dAtA[i] = 0xa
			i++
			i = encodeVarintServiceExpositionPolicy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ServiceExpositionPolicy_ExposedService) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceExpositionPolicy_ExposedService) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServiceExpositionPolicy(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Alias) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintServiceExpositionPolicy(dAtA, i, uint64(len(m.Alias)))
		i += copy(dAtA[i:], m.Alias)
	}
	if len(m.Subset) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintServiceExpositionPolicy(dAtA, i, uint64(len(m.Subset)))
		i += copy(dAtA[i:], m.Subset)
	}
	if m.Port != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintServiceExpositionPolicy(dAtA, i, uint64(m.Port))
	}
	if len(m.Clusters) > 0 {
		for _, s := range m.Clusters {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintServiceExpositionPolicy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ServiceExpositionPolicy) Size() (n int) {
	var l int
	_ = l
	if len(m.Exposed) > 0 {
		for _, e := range m.Exposed {
			l = e.Size()
			n += 1 + l + sovServiceExpositionPolicy(uint64(l))
		}
	}
	return n
}

func (m *ServiceExpositionPolicy_ExposedService) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovServiceExpositionPolicy(uint64(l))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovServiceExpositionPolicy(uint64(l))
	}
	l = len(m.Subset)
	if l > 0 {
		n += 1 + l + sovServiceExpositionPolicy(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovServiceExpositionPolicy(uint64(m.Port))
	}
	if len(m.Clusters) > 0 {
		for _, s := range m.Clusters {
			l = len(s)
			n += 1 + l + sovServiceExpositionPolicy(uint64(l))
		}
	}
	return n
}

func sovServiceExpositionPolicy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServiceExpositionPolicy(x uint64) (n int) {
	return sovServiceExpositionPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceExpositionPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceExpositionPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceExpositionPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceExpositionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exposed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Exposed = append(m.Exposed, &ServiceExpositionPolicy_ExposedService{})
			if err := m.Exposed[len(m.Exposed)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceExpositionPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceExpositionPolicy_ExposedService) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceExpositionPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExposedService: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExposedService: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceExpositionPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceExpositionPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServiceExpositionPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceExpositionPolicy
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceExpositionPolicy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthServiceExpositionPolicy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServiceExpositionPolicy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServiceExpositionPolicy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServiceExpositionPolicy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceExpositionPolicy   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("multicluster/v1alpha1/service_exposition_policy.proto", fileDescriptorServiceExpositionPolicy)
}

var fileDescriptorServiceExpositionPolicy = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x89, 0xfd, 0xa3, 0x8d, 0xe8, 0x21, 0x88, 0x86, 0x0a, 0xcb, 0xe2, 0x69, 0x2f, 0x4d,
	0xa8, 0xe2, 0x0b, 0x28, 0xbd, 0x79, 0x90, 0xf5, 0x56, 0x90, 0x92, 0xdd, 0x06, 0x77, 0x20, 0x69,
	0x42, 0x92, 0x2d, 0x7a, 0xf5, 0xe9, 0x3c, 0xfa, 0x08, 0xb2, 0xef, 0x21, 0x88, 0xd9, 0x6d, 0xa9,
	0xa0, 0xde, 0xe6, 0x37, 0xc3, 0xf7, 0x91, 0xc9, 0xe0, 0x6b, 0x5d, 0xab, 0x00, 0xa5, 0xaa, 0x7d,
	0x90, 0x8e, 0xaf, 0xa7, 0x42, 0xd9, 0x4a, 0x4c, 0xb9, 0x97, 0x6e, 0x0d, 0xa5, 0x5c, 0xc8, 0x67,
	0x6b, 0x3c, 0x04, 0x30, 0xab, 0x85, 0x35, 0x0a, 0xca, 0x17, 0x66, 0x9d, 0x09, 0x86, 0x9c, 0x83,
	0x0f, 0x60, 0xd8, 0x2e, 0xcc, 0x36, 0xf0, 0xc5, 0x27, 0xc2, 0x67, 0x0f, 0xad, 0x60, 0xb6, 0xe5,
	0xef, 0x23, 0x4e, 0x1e, 0xf1, 0x7e, 0x74, 0xca, 0x25, 0x45, 0x69, 0x2f, 0x3b, 0xbc, 0xbc, 0x65,
	0xff, 0xa8, 0xd8, 0x1f, 0x1a, 0x36, 0x6b, 0x1d, 0xdd, 0x38, 0xdf, 0x38, 0xc7, 0xaf, 0x08, 0x1f,
	0xff, 0x9c, 0x11, 0x82, 0xfb, 0x2b, 0xa1, 0x25, 0x45, 0x29, 0xca, 0x46, 0x79, 0xac, 0xc9, 0x09,
	0x1e, 0x08, 0x05, 0xc2, 0xd3, 0xbd, 0xd8, 0x6c, 0x03, 0x39, 0xc5, 0x43, 0x5f, 0x17, 0x5e, 0x06,
	0xda, 0x8b, 0xed, 0x2e, 0x7d, 0x1b, 0xac, 0x71, 0x81, 0xf6, 0x53, 0x94, 0x1d, 0xe5, 0xb1, 0x26,
	0x63, 0x7c, 0xd0, 0x3d, 0xd6, 0xd3, 0x41, 0xda, 0xcb, 0x46, 0xf9, 0x36, 0xdf, 0xcc, 0xdf, 0x9a,
	0x04, 0xbd, 0x37, 0x09, 0xfa, 0x68, 0x12, 0x34, 0xbf, 0x7b, 0x82, 0x50, 0xd5, 0x05, 0x83, 0x42,
	0xb3, 0xd2, 0x68, 0x1e, 0xd7, 0x9d, 0x38, 0xe9, 0xa5, 0x70, 0x65, 0xc5, 0x77, 0xf7, 0x9e, 0x38,
	0x23, 0x96, 0x5a, 0x58, 0x2e, 0x2c, 0xf0, 0x5f, 0x0f, 0x53, 0x0c, 0xe3, 0xff, 0x5f, 0x7d, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x20, 0xa5, 0xa5, 0x6a, 0xb8, 0x01, 0x00, 0x00,
}