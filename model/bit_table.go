package model

type BiTableResponse struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
	Data *Data  `json:"data"`
}

type Data struct {
	Total   int64   `json:"total"`
	HasMore bool    `json:"has_more"`
	Items   []*Item `json:"items"`
}

type Item struct {
	Fields   map[string]any `json:"fields"`
	RecordID string         `json:"record_id"`
}

//type Member struct {
//	AvatarURL string `json:"avatar_url"`
//	Name      string `json:"name"`
//	EnName    string `json:"en_name"`
//	ID        string `json:"id"`
//	Email     string `json:"email"`
//}
//
//type Field struct {
//	Text             *Text `json:"text"`
//	Attachment       []*Attachment
//	MultipleOptions  []string
//	Location         *Location
//	Link             *Link
//	Checkbox         bool
//	Number           float64
//	AutoNumbering    string
//	Currency         int64
//	Barcode          []*Barcode
//	TelephoneNumber  string
//	Group            []*Group
//	LastModifiedDate int64
//	Score            int64
//	Date             int64
//	Personnel        []*Member
//	EditPersonnel    []*Member
//	CreatedDate      int64
//	OneWayLink       *WayLink
//	MultipleText     []*MultipleText
//	Progress         float64
//	Creator          []*Member
//	TwoWayLink       *WayLink
//	SingleOption     string
//}
//
//type WayLink struct {
//	LinkRecordIDS []string `json:"link_record_ids"`
//}
//
//type Location struct {
//	Address     string `json:"address"`
//	AdName      string `json:"adname"`
//	PName       string `json:"pname"`
//	CityName    string `json:"cityname"`
//	Name        string `json:"name"`
//	Location    string `json:"location"`
//	FullAddress string `json:"full_address"`
//}
//
//type MultipleText struct {
//	Text          string  `json:"text"`
//	Type          string  `json:"type"`
//	Name          *string `json:"name,omitempty"`
//	MentionNotify *bool   `json:"mentionNotify,omitempty"`
//	MentionType   *string `json:"mentionType,omitempty"`
//	Token         *string `json:"token,omitempty"`
//}
//
//type Barcode struct {
//	Text string `json:"text"`
//	Type string `json:"type"`
//}
//
//type Text struct {
//	Text string `json:"text"`
//	Type string `json:"type"`
//}
//
//type Group struct {
//	AvatarURL string `json:"avatar_url"`
//	Name      string `json:"name"`
//	ID        string `json:"id"`
//}
//
//type Link struct {
//	Link string `json:"link"`
//	Text string `json:"text"`
//}
//
//type Attachment struct {
//	Size      int64  `json:"size"`
//	Name      string `json:"name"`
//	FileToken string `json:"file_token"`
//	Type      string `json:"type"`
//	TmpURL    string `json:"tmp_url"`
//	URL       string `json:"url"`
//}

func NewBiTableResp() *BiTableResponse {
	return &BiTableResponse{
		Msg:  "",
		Code: 0,
		Data: &Data{
			Total:   0,
			HasMore: false,
			Items:   []*Item{},
		},
	}
}

type BiTableReq struct {
	AppToken string
	TableID  string
}

func NewBiTableReq() *BiTableReq {
	return &BiTableReq{
		AppToken: "",
		TableID:  "",
	}
}
