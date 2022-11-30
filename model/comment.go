package model

type Comment struct {
	MID     string //所属评论
	CID     string //主键
	Sender  string //
	Time    string //
	Text    string //
	Deleted string //是否删除状态
	RID     string //是哪条评论的回复，没有就为0
}
