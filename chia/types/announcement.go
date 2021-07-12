// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/announcement.py

package types

// Announcement
//
// 		origin_info: bytes32
// 		message: bytes
//
type Announcement struct {
	OriginInfo []byte `json:"origin_info"`
	Message    []byte `json:"message"`
}

// @TODO
// Name
func (a Announcement) Name() []byte {
	return []byte{}
}
