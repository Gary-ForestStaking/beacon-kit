// SPDX-License-Identifier: MIT
//
// # Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
//nolint:errcheck // do not check for err returns
package ssz_test

import (
	"os"
	"testing"

	sszv2 "github.com/berachain/beacon-kit/mod/primitives/pkg/ssz/v2/lib"
	ssz "github.com/ferranbt/fastssz"
	"github.com/stretchr/testify/require"
)

// https://goerli.beaconcha.in/slot/4744352
// Test fixture from fastssz.
const TestFileName = "fixtures/beacon_state_bellatrix.ssz"

//nolint:gochecknoglobals // test debug print toggle
var debug = false

type TestLogger interface {
	Logf(format string, args ...any)
}

func debugPrint(debug bool, t TestLogger, s1 string, s ...any) {
	if debug {
		t.Logf(s1, s...)
	}
}

func runBench(b *testing.B, cb func()) {
	b.ResetTimer()
	for range b.N {
		cb()
	}
}

func getSszState() (*sszv2.BeaconStateBellatrix, error) {
	// A checkpt is the simplest field.
	data, err := os.ReadFile(TestFileName)
	if err != nil {
		return nil, err
	}

	sszState := sszv2.BeaconStateBellatrix{}
	err2 := sszState.UnmarshalSSZ(data)
	if err2 != nil {
		return nil, err2
	}
	return &sszState, nil
}

func getVariableLengthItemSerialized() []byte {
	// magic nums from the generated ssz.go file for historicalRoots
	startPos := 2736633
	endPos := startPos + 18528

	return sliceBlockData(startPos, endPos)
}

func sliceBlockData(start int, end int) []byte {
	s, err := getSszState()
	if err != nil {
		return nil
	}
	data, _ := s.MarshalSSZ()
	return data[start:end]
}

func getStruct(bb *sszv2.BeaconStateBellatrix) *sszv2.Checkpoint {
	return bb.CurrentJustifiedCheckpoint
}

func getU64(bb *sszv2.BeaconStateBellatrix) uint64 {
	return getStruct(bb).Epoch
}

func getByteArray32(bb *sszv2.BeaconStateBellatrix) []byte {
	return getStruct(bb).Root
}

func getByteArray32Serialized(bb *sszv2.BeaconStateBellatrix) ([]byte, error) {
	res, err := getStruct(bb).MarshalSSZ()
	if err != nil {
		return nil, err
	}
	// type Checkpoint struct {
	// 	Epoch uint64 `json:"epoch"`
	// 	Root  []byte `json:"root" ssz-size:"32"`
	// }
	// We grab the buf section from the serialized by fastSSZ buffer.
	// See bellatrix.ssz.go for buffer read done by fastssz codegen.
	// Since uint64 serializes to 8 bits. we grab the remaining bits of len 32.

	return res[8:], nil
}

func getEth1DataVotesSerialized(bb *sszv2.BeaconStateBellatrix) []byte {
	dst := make([]byte, 0)
	for ii := range len(bb.Eth1DataVotes) {
		dst, _ = bb.Eth1DataVotes[ii].MarshalSSZTo(dst)
	}
	return dst
}

// Todo: full object serialization
// func TestParityBellatrix(t *testing.T) {
// 	sszState, err := getSszState()
// 	require.NoError(t, err)

// 	s := sszv2.NewSerializer()
// 	o2, err3 := s.MarshalSSZ(sszState)
// 	require.NoError(t, err3)

// 	res, err4 := sszState.MarshalSSZ()
// 	require.NoError(t, err4)

// 	require.Equal(t, o2, res, "local output and fastssz output doesn't match")
// }

func TestParitySliceOfStructs(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)

	s := sszv2.NewSerializer()
	o2, err3 := s.MarshalSSZ(sszState.Eth1DataVotes)
	require.NoError(t, err3)

	res := getEth1DataVotesSerialized(sszState)

	require.Equal(t, o2, res, "local & fastssz output doesn't match")
}

func TestParityVariableLengthItem2(t *testing.T) {
	state, err := getSszState()
	require.NoError(t, err)
	bals := state.Balances

	s := sszv2.NewSerializer()
	o2, err3 := s.MarshalSSZ(bals)
	require.NoError(t, err3)

	// get balances via direct slice
	res := sliceBlockData(51139110, 54333774)
	require.Equal(t, o2, res, "local output and fastssz output doesn't match")
}

func TestParityVariableLengthItem1(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)
	historicalRoots := sszState.HistoricalRoots

	s := sszv2.NewSerializer()
	o2, err3 := s.MarshalSSZ(historicalRoots)
	require.NoError(t, err3)

	res := getVariableLengthItemSerialized()
	require.Equal(t, o2, res, "local output and fastssz output doesn't match")
}

func TestParityStruct(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)

	testStruct := getStruct(sszState)

	s := sszv2.NewSerializer()
	o2, err3 := s.MarshalSSZ(testStruct)
	require.NoError(t, err3)
	debugPrint(debug, t, "Local Serializer output: ", o2, err)

	res, _ := testStruct.MarshalSSZ()
	debugPrint(debug, t, "FastSSZ Output: ", res)
	require.Equal(t, o2, res, "local output and fastssz output doesnt match")
}

func BenchmarkNativeStruct(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)

	testStruct := getStruct(sszState)

	s := sszv2.NewSerializer()
	runBench(b, func() {
		s.MarshalSSZ(testStruct)
	})
}

func BenchmarkFastSSZStruct(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)

	testStruct := getStruct(sszState)

	runBench(b, func() {
		testStruct.MarshalSSZ()
	})
}

func TestParityUint64(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)

	testU64 := getU64(sszState)

	s := sszv2.NewSerializer()
	o2, err3 := s.MarshalSSZ(testU64)
	require.NoError(t, err3)
	debugPrint(debug, t, "Local Serializer output:", o2, err)

	res := make([]byte, 0)
	res = ssz.MarshalUint64(res, testU64)
	debugPrint(debug, t, "FastSSZ Output:", res)
	require.Equal(t, o2, res, "local output and fastssz output doesnt match")
}

func BenchmarkNativeUint64(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)

	testU64 := getU64(sszState)

	s := sszv2.NewSerializer()
	runBench(b, func() {
		s.MarshalSSZ(testU64)
	})
}

func BenchmarkFastSSZUint64(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)

	testU64 := getU64(sszState)
	res := make([]byte, 0)

	runBench(b, func() {
		res = ssz.MarshalUint64(res, testU64)
	})
}

func TestParityByteArray(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)
	testByteArr := getByteArray32(sszState)
	s := sszv2.NewSerializer()

	exp, err3 := s.MarshalSSZ(testByteArr)
	require.NoError(t, err3)
	debugPrint(debug, t, "Local Serializer output:", exp, err)

	res, err4 := getByteArray32Serialized(sszState)
	require.NoError(t, err4)
	debugPrint(debug, t, "FastSSZ Output:", res)

	require.Equal(t, exp, res, "local output and fastssz output doesnt match")
}

func BenchmarkNativeByteArray(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)
	testByteArr := getByteArray32(sszState)
	s := sszv2.NewSerializer()

	runBench(b, func() {
		s.MarshalSSZ(testByteArr)
	})
}

func BenchmarkFastSSZByteArray(b *testing.B) {
	sszState, err := getSszState()
	require.NoError(b, err)

	runBench(b, func() {
		sszState.PreviousJustifiedCheckpoint.MarshalSSZ()
	})
}

func TestParityByteArrayLarge2D(t *testing.T) {
	data, err := os.ReadFile(TestFileName)
	require.NoError(t, err)

	sszState := sszv2.BeaconStateBellatrix{}

	err = sszState.UnmarshalSSZ(data)
	require.NoError(t, err)

	s := sszv2.NewSerializer()
	exp, err2 := s.MarshalSSZ(sszState.StateRoots)
	require.NoError(t, err2)
	// We test serialized output. This may be lacking checks for offsets.
	debugPrint(debug, t, "Local Serializer output:", exp, err2)
	// fast ssz: len 262144 []uint8  | cap: 58065320.
	// local: len 262144 []uint8  |  cap:278528.

	res, err3 := sszState.MarshalSSZ()
	require.NoError(t, err3)
	prInRes := res[262320:524464]
	debugPrint(debug, t, "Local Serializer output length:", len(exp))
	debugPrint(debug, t, "FastSSZ Serializer output length:", len(prInRes))
	require.Equal(
		t,
		exp,
		prInRes,
		"local output and fastssz output doesnt match",
	)
}

func BenchmarkNativeByteArrayLarge(b *testing.B) {
	data, err := os.ReadFile(TestFileName)
	require.NoError(b, err)

	sszState := sszv2.BeaconStateBellatrix{}
	err2 := sszState.UnmarshalSSZ(data)
	require.NoError(b, err2)

	s := sszv2.NewSerializer()

	runBench(b, func() {
		s.MarshalSSZ(sszState.StateRoots)
	})
}

func BenchmarkFastSSZByteArrayLarge(b *testing.B) {
	data, err := os.ReadFile(TestFileName)
	require.NoError(b, err)

	sszState := sszv2.BeaconStateBellatrix{}
	err2 := sszState.UnmarshalSSZ(data)
	require.NoError(b, err2)

	runBench(b, func() {
		sszState.MarshalSSZ()
	})
}

func TestParityU64Array(t *testing.T) {
	sszState, err := getSszState()
	require.NoError(t, err)

	u64Arr := sszState.Slashings

	s := sszv2.NewSerializer()

	debugPrint(debug, t, "Local Serializer input len:", len(u64Arr), err)
	exp, err3 := s.MarshalSSZ(u64Arr)
	require.NoError(t, err3)
	debugPrint(debug, t, "Local Serializer output len:", len(exp), err)
	// slashings := make([]byte, (8192 * 8))
	res, err3 := sszState.MarshalSSZ()
	// See bellatrix.ssz.go generated file in unmarshalSSZ
	slashings := make([]byte, 0)
	if len(res) >= 2687248 {
		slashings = res[2621712:2687248]
	}
	require.NoError(t, err3)
	debugPrint(debug, t, "FastSSZ Output len:", len(slashings))

	require.Equal(
		t,
		len(exp),
		len(slashings),
		"local output and fastssz output length doesnt match",
	)
	require.Equal(
		t,
		exp,
		slashings,
		"local output and fastssz output doesnt match",
	)
}
