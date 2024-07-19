package model

type BridgeOrder struct {
	Data       []byte // msg binary data
	Hash       string
	VoteStatus bool //  vote 失败，成功
	Status     bool //  execute 失败，成功
}
