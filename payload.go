package getuigo

type Payload interface {
	GetPushTitle() string
	SetPushTitle(title string)
	GetPushBody() string
	SetPushBody(body string)
	GetIsShowNotify() int
	SetIsShowNotify(isshow int)
	String() string
}

type BasicPayload struct {
	PushTitle    string `json:"push_title"`
	PushBody     string `json:"push_body"`
	IsShowNotify int    `json:"is_show_notify"`
	Ext          string `json:"ext"`
}

func (p *BasicPayload) GetPushTitle() string {
	return p.PushTitle
}

func (p *BasicPayload) SetPushTitle(title string) {
	p.PushTitle = title
}

func (p *BasicPayload) GetPushBody() string {
	return p.PushBody
}

func (p *BasicPayload) SetPushBody(body string) {
	p.PushBody = body
}

func (p *BasicPayload) GetIsShowNotify() int {
	return p.IsShowNotify
}

func (p *BasicPayload) SetIsShowNotify(isshow int) {
	p.IsShowNotify = isshow
}

func (p *BasicPayload) GetExt() string {
	return p.Ext
}

func (p *BasicPayload) SetExt(ext string) {
	p.Ext = ext
}

func (p *BasicPayload) String() string {
	return p.PushTitle
}
