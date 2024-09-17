package service

type DoSomeStuff struct {
	aa aa
}

type aa interface{}

func NewDoSomeStuff(aa aa) *DoSomeStuff {
	return &DoSomeStuff{
		aa: aa,
	}
}
