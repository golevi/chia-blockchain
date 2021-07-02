// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/announcement.py

package types

// Announcement
//
// 		origin_info: bytes32
// 		message: bytes
//
type Announcement struct {
	OriginInfo [32]byte `json:"origin_info"`
	Message    []byte   `json:"message"`
}

// @TODO
// Name
func (a Announcement) Name() [32]byte {
	return [32]byte{}
}
