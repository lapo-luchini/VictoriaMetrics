// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package prompbmarshal

import (
	"encoding/binary"
	"math"
)

type Sample struct {
	Value     float64
	Timestamp int64
}

type Exemplar struct {
	// Optional, can be empty.
	Labels []Label
	Value  float64
	// timestamp is in ms format, see model/timestamp/timestamp.go for
	// conversion from time.Time to Prometheus timestamp.
	Timestamp int64
}

func (m *Exemplar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *Exemplar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	return n
}

// TimeSeries represents samples and labels for a single time series.
type TimeSeries struct {
	Labels    []Label
	Samples   []Sample
	Exemplars []Exemplar
}

type Label struct {
	Name  string
	Value string
}

func (m *Sample) MarshalToSizedBuffer(dst []byte) (int, error) {
	i := len(dst)
	if m.Timestamp != 0 {
		i = encodeVarint(dst, i, uint64(m.Timestamp))
		i--
		dst[i] = 0x10
	}
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dst[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dst[i] = 0x9
	}
	return len(dst) - i, nil
}

func (m *TimeSeries) MarshalToSizedBuffer(dst []byte) (int, error) {
	i := len(dst)
	for j := len(m.Exemplars) - 1; j >= 0; j-- {
		size, err := m.Exemplars[j].MarshalToSizedBuffer(dst[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dst, i, uint64(size))
		i--
		dst[i] = 0x1a
	}
	for j := len(m.Samples) - 1; j >= 0; j-- {
		size, err := m.Samples[j].MarshalToSizedBuffer(dst[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dst, i, uint64(size))
		i--
		dst[i] = 0x12
	}
	for j := len(m.Labels) - 1; j >= 0; j-- {
		size, err := m.Labels[j].MarshalToSizedBuffer(dst[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dst, i, uint64(size))
		i--
		dst[i] = 0xa
	}
	return len(dst) - i, nil
}

func (m *Label) MarshalToSizedBuffer(dst []byte) (int, error) {
	i := len(dst)
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dst[i:], m.Value)
		i = encodeVarint(dst, i, uint64(len(m.Value)))
		i--
		dst[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dst[i:], m.Name)
		i = encodeVarint(dst, i, uint64(len(m.Name)))
		i--
		dst[i] = 0xa
	}
	return len(dst) - i, nil
}

func (m *Sample) Size() (n int) {
	if m == nil {
		return 0
	}
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	return n
}

func (m *TimeSeries) Size() (n int) {
	if m == nil {
		return 0
	}
	for _, e := range m.Labels {
		l := e.Size()
		n += 1 + l + sov(uint64(l))
	}
	for _, e := range m.Samples {
		l := e.Size()
		n += 1 + l + sov(uint64(l))
	}
	for _, e := range m.Exemplars {
		l := e.Size()
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func (m *Label) Size() (n int) {
	if m == nil {
		return 0
	}
	if l := len(m.Name); l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if l := len(m.Value); l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	return n
}
