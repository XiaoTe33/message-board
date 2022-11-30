package model

type Message struct {
	MID      string //主键
	Sender   string //
	Receiver string //
	Time     string //
	Text     string //
	Deleted  string //是否删除状态
}
