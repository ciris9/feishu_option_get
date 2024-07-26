package model

type ApproveRequest struct {
	UserID        string         `json:"user_id"`
	EmployeeID    string         `json:"employee_id"`
	Query         string         `json:"query"`
	LinkageParams *LinkageParams `json:"linkage_params"`
	PageToken     string         `json:"page_token"`
	Locale        string         `json:"locale"`
	Token         string         `json:"token"`
}

type LinkageParams struct {
	KV map[string]string
}

func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{
		UserID:     "",
		EmployeeID: "",
		Query:      "",
		LinkageParams: &LinkageParams{
			KV: make(map[string]string),
		},
		PageToken: "",
		Locale:    "",
		Token:     "",
	}
}

type ApproveResponse struct {
	Msg  string       `json:"msg"`
	Code int64        `json:"code"`
	Data *ApproveData `json:"data"`
}

type ApproveData struct {
	Result *Result `json:"result"`
}

type Result struct {
	I18NResources []*I18NResource  `json:"i18nResources"`
	NextPageToken string           `json:"nextPageToken"`
	Options       []*ApproveOption `json:"options"`
	HasMore       bool             `json:"hasMore"`
}

type I18NResource struct {
	IsDefault bool   `json:"isDefault"`
	Texts     *Texts `json:"texts"`
	Locale    string `json:"locale"`
}

type Texts struct {
	I18NOptions1Name3 string `json:"@i18n@options_1_name_3"`
	I18NOptions1Name2 string `json:"@i18n@options_1_name_2"`
	I18NOptions1Name1 string `json:"@i18n@options_1_name_1"`
}

type ApproveOption struct {
	IsDefault bool   `json:"isDefault,omitempty"`
	ID        string `json:"id"`
	Value     string `json:"value"`
}

func NewApproveResponse() *ApproveResponse {
	return &ApproveResponse{
		Msg:  "",
		Code: 0,
		Data: &ApproveData{
			Result: &Result{
				I18NResources: []*I18NResource{},
				NextPageToken: "",
				Options:       []*ApproveOption{},
				HasMore:       false,
			},
		},
	}
}
