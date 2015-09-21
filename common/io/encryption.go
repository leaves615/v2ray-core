package io

import (
	"crypto/cipher"
	"io"

	"github.com/v2ray/v2ray-core/common/log"
)

// CryptionReader is a general purpose reader that applies
// block cipher on top of a regular reader.
type CryptionReader struct {
	stream cipher.Stream
	reader io.Reader
}

func NewCryptionReader(stream cipher.Stream, reader io.Reader) *CryptionReader {
	return &CryptionReader{
		stream: stream,
		reader: reader,
	}
}

// Read reads blocks from underlying reader, the length of blocks must be
// a multiply of BlockSize()
func (reader CryptionReader) Read(blocks []byte) (int, error) {
	nBytes, err := reader.reader.Read(blocks)
	if nBytes > 0 {
		reader.stream.XORKeyStream(blocks[:nBytes], blocks[:nBytes])
	}
	if err != nil && err != io.EOF {
		log.Error("Error reading blocks: %v", err)
	}
	return nBytes, err
}

// Cryption writer is a general purpose of byte stream writer that applies
// block cipher on top of a regular writer.
type CryptionWriter struct {
	stream cipher.Stream
	writer io.Writer
}

func NewCryptionWriter(stream cipher.Stream, writer io.Writer) *CryptionWriter {
	return &CryptionWriter{
		stream: stream,
		writer: writer,
	}
}

func (writer CryptionWriter) Crypt(blocks []byte) {
	writer.stream.XORKeyStream(blocks, blocks)
}

// Write writes the give blocks to underlying writer. The length of the blocks
// must be a multiply of BlockSize()
func (writer CryptionWriter) Write(blocks []byte) (int, error) {
	writer.Crypt(blocks)
	return writer.writer.Write(blocks)
}
