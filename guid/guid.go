package guid

import (
	"bytes"
	"errors"
	"log"
)

type group []byte

// GUID stores Microsoft binary guids. Since SQL Server and Active Directory
// uniqueidentifiers are stored in Big (network) Endian, it appears that the
// first half of a Guid is backwards if you take the raw byte slice and
// treat it like a string or a continuous 16 byte slice.
type GUID struct {
	group1 group // 4 bytes (MS stores as Big Endian uin32)
	group2 group // 2 bytes (MS stores as Big Endian uin16)
	group3 group // 2 bytes (MS stores as Big Endian uin16)
	group4 group // 8 bytes (MS stores as Big Endian an array of bytes)
}

// ParseRawBytes takes a the binary format of a Guid as a byte slice and
// returns a struct that can output the Guid as its string representation
// or Little Endian byte slice.
func ParseRawBytes(guid []byte) (GUID, error) {
	var result GUID
	// Remove dashes before attempting to parse.
	guid = bytes.Replace(guid, []byte{0x2D}, []byte{}, -1)
	if len(guid) != 16 {
		return result, errors.New("could not parse guid " + string(guid))
	}
	result.group1 = guid[:4]
	result.group2 = guid[4:6]
	result.group3 = guid[6:8]
	result.group4 = guid[8:]
	return result, nil
}

// MustParseRawBytes takes a the binary format of a Guid as a byte slice and
// returns a struct that can output the Guid as its string representation
// or Little Endian byte slice.
func MustParseRawBytes(guid []byte) GUID {
	result, err := ParseRawBytes(guid)
	if err != nil {
		log.Panicln("could not parse guid", guid)
	}
	return result
}

// UUIDBytes outputs a single 16 byte slice in Little Endian reflecting how
// non-Microsoft packages usually store UUIDs.
func (g GUID) UUIDBytes() []byte {
	part1 := append(g.group1.reverse(), g.group2.reverse()...)
	part1 = append(part1, g.group3.reverse()...)
	return append(part1, g.group4...)
}

func (g group) reverse() []byte {
	var reversed []byte
	for i := range g {
		reversed = append(reversed, g[len(g)-1-i])
	}
	return reversed
}

func (g group) isEqualTo(g2 group) bool {
	if len(g) != len(g2) {
		return false
	}
	for i, b1 := range g {
		if g2[i] != b1 {
			return false
		}
	}
	return true
}

// IsEqualTo checks that two GUIDs are equal.
func (g GUID) IsEqualTo(g2 GUID) bool {
	return g.group1.isEqualTo(g2.group1) &&
		g.group2.isEqualTo(g2.group2) &&
		g.group3.isEqualTo(g2.group3) &&
		g.group4.isEqualTo(g2.group4)
}
