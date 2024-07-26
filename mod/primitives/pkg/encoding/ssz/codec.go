package ssz

import (
	"io"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/types"
)

type Codec struct {
	// one of:
	encoder *Encoder
	schema  *Schema
	hasher  *Hasher
	treeer  *Treeer
}

type SSZObject interface {
	DefineSSZ(*Codec)
}

func SchemaCodec() *Codec {
	return &Codec{schema: &Schema{}}
}

type Schema struct {
}

type Encoder struct {
	outWriter io.Writer
	outBuffer []byte
	err       error

	buf [32]byte
}

type Hasher struct {
}

type Treeer struct {
}

func encode[VectorT types.MinimalSSZType](c *Codec, t types.MinimalSSZType) {
	switch f := t.(type) {
	case Vector[VectorT]:
		f.
	}

}

func DefineField(c *Codec, t types.MinimalSSZType) {
	switch {
	case c.encoder != nil:
		c.encode(t)
	case c.schema != nil:
		RegisterBasic(c.schema, t)
	case c.hasher != nil:
	case c.treeer != nil:
		panic("not implemented")
	}
}

// --------------------------------------------------------------------------
// 							  Encoding
// --------------------------------------------------------------------------

type BufferWriter struct {
	out []byte
}

func (bw *BufferWriter) Write(buf []byte) (int, error) {
	n := copy(bw.out, buf)
	bw.out = bw.out[len(buf):]
	return n, nil
}

// --------------------------------------------------------------------------
// 							  Schema Registration
// --------------------------------------------------------------------------

func RegisterBasic(s *Schema, t types.MinimalSSZType) {
}
