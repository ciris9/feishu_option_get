package model

// YApi QuickType插件生成，具体参考文档:https://plugins.jetbrains.com/plugin/18847-yapi-quicktype/documentation

type BiTableFieldResponse struct {
	Msg  string            `json:"msg"`
	Code int64             `json:"code"`
	Data *BiTableFieldData `json:"data"`
}

type BiTableFieldData struct {
	Total     int64               `json:"total"`
	PageToken string              `json:"page_token"`
	HasMore   bool                `json:"has_more"`
	Items     []*BiTableFieldItem `json:"items"`
}

type BiTableFieldItem struct {
	FieldID   string    `json:"field_id"`
	UIType    string    `json:"ui_type"`
	IsPrimary bool      `json:"is_primary"`
	Type      int64     `json:"type"`
	FieldName string    `json:"field_name"`
	Property  *Property `json:"property,omitempty"`
}

type Property struct {
	Multiple         *bool             `json:"multiple,omitempty"`
	Options          []*Option         `json:"options,omitempty"`
	DateFormatter    *string           `json:"date_formatter,omitempty"`
	AutoFill         *bool             `json:"auto_fill,omitempty"`
	TableID          *string           `json:"table_id,omitempty"`
	TableName        *string           `json:"table_name,omitempty"`
	BackFieldID      *string           `json:"back_field_id,omitempty"`
	BackFieldName    *string           `json:"back_field_name,omitempty"`
	Location         *Location         `json:"location,omitempty"`
	Formatter        *string           `json:"formatter,omitempty"`
	AllowedEditModes *AllowedEditModes `json:"allowed_edit_modes,omitempty"`
	AutoSerial       *AutoSerial       `json:"auto_serial,omitempty"`
	CurrencyCode     *string           `json:"currency_code,omitempty"`
	Min              *int64            `json:"min,omitempty"`
	Max              *int64            `json:"max,omitempty"`
	RangeCustomize   *bool             `json:"range_customize,omitempty"`
	Rating           *Rating           `json:"rating,omitempty"`
}

type AllowedEditModes struct {
	Scan   bool `json:"scan"`
	Manual bool `json:"manual"`
}

type AutoSerial struct {
	Type string `json:"type"`
}

type Location struct {
	InputType string `json:"input_type"`
}

type Option struct {
	Color int64  `json:"color"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type Rating struct {
	Symbol string `json:"symbol"`
}

func NewBITableFieldResponse() *BiTableFieldResponse {
	return &BiTableFieldResponse{
		Msg:  "",
		Code: 0,
		Data: &BiTableFieldData{
			Total:     0,
			PageToken: "",
			HasMore:   false,
			Items:     make([]*BiTableFieldItem, 0),
		},
	}
}
