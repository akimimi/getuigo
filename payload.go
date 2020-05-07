package getuigo

const NotifyTypePayload = "0"
const NotifyTypeIntent = "1"
const NotifyTypeUrl = "2"

type Payload interface {
	GetNotifyTitle() string
	SetNotifyTitle(title string)
	GetNotifyBody() string
	SetNotifyBody(body string)
	GetIsShowNotify() int
	SetIsShowNotify(isshow int)
	GetIntent() string
	SetIntent(intent string)
	String() string
}

type BasicPayload struct {
	NotifyTitle  string `json:"push_title"`
	NotifyBody   string `json:"push_body"`
	IsShowNotify int    `json:"is_show_notify"`
	Ext          string `json:"ext"`
}

func (p *BasicPayload) GetNotifyTitle() string {
	return p.NotifyTitle
}

func (p *BasicPayload) SetNotifyTitle(title string) {
	p.NotifyTitle = title
}

func (p *BasicPayload) GetNotifyBody() string {
	return p.NotifyBody
}

func (p *BasicPayload) SetNotifyBody(body string) {
	p.NotifyBody = body
}

func (p *BasicPayload) GetIsShowNotify() int {
	return p.IsShowNotify
}

func (p *BasicPayload) SetIsShowNotify(isshow int) {
	p.IsShowNotify = isshow
}

func (p *BasicPayload) GetIntent() string {
	return ""
}

func (p *BasicPayload) SetIntent(_ string) {
	// just do nothing
}

func (p *BasicPayload) GetExt() string {
	return p.Ext
}

func (p *BasicPayload) SetExt(ext string) {
	p.Ext = ext
}

func (p *BasicPayload) String() string {
	return p.NotifyTitle
}
