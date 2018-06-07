// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package aggregation

import (
	"fmt"

	"github.com/m3db/m3metrics/generated/proto/schema"
)

const (
	// IDLen is the length of the ID.
	// The IDLen will be 1 when maxTypeID <= 63.
	IDLen = (maxTypeID)/64 + 1

	// ID uses an array of int64 to represent aggregation types.
	idBitShift = 6
	idBitMask  = 63
)

var (
	// DefaultID is a default ID.
	DefaultID ID
)

// ID represents a compressed view of Types.
type ID [IDLen]uint64

// NewIDFromSchema creates an ID from schema.
func NewIDFromSchema(input []schema.AggregationType) (ID, error) {
	aggTypes, err := NewTypesFromSchema(input)
	if err != nil {
		return DefaultID, err
	}

	// TODO(cw): consider pooling these compressors,
	// this allocates one extra slice of length one per call.
	id, err := NewIDCompressor().Compress(aggTypes)
	if err != nil {
		return DefaultID, err
	}
	return id, nil
}

// MustCompressTypes compresses a list of aggregation types to
// an ID, it panics if an error was encountered.
func MustCompressTypes(aggTypes ...Type) ID {
	res, err := NewIDCompressor().Compress(aggTypes)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// IsDefault checks if the ID is the default aggregation type.
func (id ID) IsDefault() bool {
	return id == DefaultID
}

// Contains checks if the given aggregation type is contained in the aggregation id.
func (id ID) Contains(aggType Type) bool {
	if !aggType.IsValid() {
		return false
	}
	idx := int(aggType) >> idBitShift   // aggType / 64
	offset := uint(aggType) & idBitMask // aggType % 64
	return (id[idx] & (1 << offset)) > 0
}

// Types returns the aggregation types defined by the id.
func (id ID) Types() (Types, error) {
	return NewIDDecompressor().Decompress(id)
}

// String for debugging.
func (id ID) String() string {
	aggTypes, err := id.Types()
	if err != nil {
		return fmt.Sprintf("[invalid ID: %v]", err)
	}
	return aggTypes.String()
}
