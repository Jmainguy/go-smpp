// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

// Binary text codec, no encoding.
type Binary []byte

// Type implements the Codec interface.
func (s Binary) Type() DataCoding {
	return 0x02
}

// Encode raw text.
func (s Binary) Encode() []byte {
	return s
}

// Decode raw text.
func (s Binary) Decode() []byte {
	return s
}
