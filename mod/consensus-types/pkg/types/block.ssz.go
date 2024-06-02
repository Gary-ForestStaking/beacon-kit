// Code generated by fastssz. DO NOT EDIT.
// Hash: 00868f49a172e8fa31a78febdb69b1ce6c90e4d303d182e90b1cab3bf8794b95
// Version: 0.1.3
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockDeneb object to a target array
func (b *BeaconBlockDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(84)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, b.ProposerIndex)

	// Field (2) 'ParentBlockRoot'
	dst = append(dst, b.ParentBlockRoot[:]...)

	// Field (3) 'StateRoot'
	dst = append(dst, b.StateRoot[:]...)

	// Offset (4) 'Body'
	dst = ssz.WriteOffset(dst, offset)

	// Field (4) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 84 {
		return ssz.ErrSize
	}

	tail := buf
	var o4 uint64

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'ParentBlockRoot'
	copy(b.ParentBlockRoot[:], buf[16:48])

	// Field (3) 'StateRoot'
	copy(b.StateRoot[:], buf[48:80])

	// Offset (4) 'Body'
	if o4 = ssz.ReadOffset(buf[80:84]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 84 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (4) 'Body'
	{
		buf = tail[o4:]
		if b.Body == nil {
			b.Body = new(BeaconBlockBodyDeneb)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) SizeSSZ() (size int) {
	size = 84

	// Field (4) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockDeneb object with a hasher
func (b *BeaconBlockDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (1) 'ProposerIndex'
	hh.PutUint64(b.ProposerIndex)

	// Field (2) 'ParentBlockRoot'
	hh.PutBytes(b.ParentBlockRoot[:])

	// Field (3) 'StateRoot'
	hh.PutBytes(b.StateRoot[:])

	// Field (4) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
