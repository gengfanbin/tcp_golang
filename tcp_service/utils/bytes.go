package utils

import "bytes"

func BytesJoin(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
