package service

import (
	"bytes"
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"io"
	"net/http"
	"option_get/config"
	"testing"

	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
// 复制该 Demo 后, 需要将 "YOUR_APP_ID", "YOUR_APP_SECRET" 替换为自己应用的 APP_ID, APP_SECRET.
// 以下示例代码是根据 API 调试台参数自动生成，如果存在代码问题，请在 API 调试台填上相关必要参数后再使用
func TestSDK(t *testing.T) {
	// 创建 Client
	client := lark.NewClient("cli_a62b0a38a77ad00e", "G3TPCEOJxkqWPP5rgbcHWdGvX8SqyvcD")
	// 创建请求对象
	emptyBody, err := larkbitable.NewSearchAppTableRecordPathReqBodyBuilder().Build()
	if err != nil {
		t.Error(err)
	}
	req := larkbitable.NewSearchAppTableRecordReqBuilder().
		AppToken(`B49bbdDZLarCLmsDToecjt1gn4g`).
		TableId(`tbl9yZdFV7JUAO0W`).
		PageSize(20).Body(emptyBody).
		Build()

	// 发起请求
	resp, err := client.Bitable.AppTableRecord.Search(context.Background(), req, larkcore.WithUserAccessToken(config.Token))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}
	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}
	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}

func TestRequest(t *testing.T) {
	client := http.Client{}
	request, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/bitable/v1/apps/B49bbdDZLarCLmsDToecjt1gn4g/tables/tbl9yZdFV7JUAO0W/records/search", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Error(err)
	}
	request.Header.Set("Authorization", "Bearer "+config.Token)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.StatusCode)
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(">>>>>>", string(all))
}
