package dtos

// type BlockMeta struct {}

type Block struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Content string      `json:"content"`
	GroupID string      `json:"groupID"`
	Meta    interface{} `json:"meta"`
}
