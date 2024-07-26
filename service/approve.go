package service

import (
	"option_get/constants"
	"option_get/model"
)

/*
	return json format
{
  "code": 0,
  "data": {
    "has_more": false,
    "items": [
      {
        "created_by": {
          "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
          "email": "",
          "en_name": "测试1",
          "id": "ou_92945f86a98bba075174776959c90eda",
          "name": "测试1"
        },
        "created_time": 1691049973000,
        "fields": {
          "人员": [
            {
              "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/b2-7619-4b8a-b27b-c72d90b06a2j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
              "email": "zhangsan.leben@bytedance.com",
              "en_name": "ZhangSan",
              "id": "ou_2910013f1e6456f16a0ce75ede950a0a",
              "name": "张三"
            },
            {
              "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
              "email": "lisi.00@bytedance.com",
              "en_name": "LiSi",
              "id": "ou_e04138c9633dd0d2ea166d79f548ab5d",
              "name": "李四"
            }
          ],
          "修改人": [
            {
              "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
              "email": "",
              "en_name": "测试1",
              "id": "ou_92945f86a98bba075174776959c90eda",
              "name": "测试1"
            }
          ],
          "创建人": [
            {
              "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
              "email": "",
              "en_name": "测试1",
              "id": "ou_92945f86a98bba075174776959c90eda",
              "name": "测试1"
            }
          ],
          "创建时间": 1691049973000,
          "单向关联": {
            "link_record_ids": [
              "recnVYsuqV"
            ]
          },
          "单选": "选项1",
          "双向关联": {
            "link_record_ids": [
              "recqLvMaXT",
              "recrdld32q"
            ]
          },
          "地理位置": {
            "address": "东长安街",
            "adname": "东城区",
            "cityname": "北京市",
            "full_address": "天安门广场，北京市东城区东长安街",
            "location": "116.397755,39.903179",
            "name": "天安门广场",
            "pname": "北京市"
          },
          "复选框": true,
          "多行文本": [
            {
              "text": "多行文本内容1",
              "type": "text"
            },
            {
              "mentionNotify": false,
              "mentionType": "User",
              "name": "张三",
              "text": "@张三",
              "token": "ou_2910013f1e6456f16a0ce75ede950a0a",
              "type": "mention"
            }
          ],
          "多选": [
            "选项1",
            "选项2"
          ],
          "数字": 2323.2323,
          "日期": 1690992000000,
          "最后更新时间": 1702455191000,
          "条码": [
            {
              "text": "123",
              "type": "text"
            }
          ],
          "电话号码": "131xxxx6666",
          "自动编号": "17",
          "群组": [
            {
              "avatar_url": "https://internal-api-lark-file.feishu-boe.cn/static-resource/v1/v2_c8d2cd50-ba29-476f-b7f1-5b5917cb18ej~?image_size=72x72&amp;cut_type=&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
              "id": "oc_cd07f55f14d6f4a4f1b51504e7e97f48",
              "name": "武侠聊天组"
            }
          ],
          "评分": 3,
          "货币": 1,
          "超链接": {
            "link": "https://bitable.feishu.cn",
            "text": "飞书多维表格官网"
          },
          "进度": 0.66,
          "附件": [
            {
              "file_token": "Vl3FbVkvnowlgpxpqsAbBrtFcrd",
              "name": "飞书.jpeg",
              "size": 32975,
              "tmp_url": "https://open.feishu.cn/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=Vl3FbVk11owlgpxpqsAbBrtFcrd&amp;extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D",
              "type": "image/jpeg",
              "url": "https://open.feishu.cn/open-apis/drive/v1/medias/Vl3FbVk11owlgpxpqsAbBrtFcrd/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D"
            }
          ]
        },
        "last_modified_by": {
          "avatar_url": "https://internal-api-lark-file.feishu.cn/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&amp;cut_type=default-face&amp;quality=&amp;format=jpeg&amp;sticker_format=.webp",
          "email": "",
          "en_name": "测试1",
          "id": "ou_92945f86a98bba075174776959c90eda",
          "name": "测试1"
        },
        "last_modified_time": 1702455191000,
        "record_id": "recyOaMB2F"
      }
    ],
    "total":
  },
  "msg": "success"
}
*/

func universalApproveResponseInit(response *model.ApproveResponse) {
	response.Data.Result.HasMore = false
	response.Data.Result.NextPageToken = ""
	response.Data.Result.I18NResources = []*model.I18NResource{
		{
			IsDefault: true,
			Texts: &model.Texts{
				I18NOptions1Name3: "值3",
				I18NOptions1Name2: "值2",
				I18NOptions1Name1: "值1",
			},
			Locale: constants.ZhCNLocaleKey,
		},
		{
			IsDefault: false,
			Texts: &model.Texts{
				I18NOptions1Name3: "value3",
				I18NOptions1Name2: "value2",
				I18NOptions1Name1: "value1",
			},
			Locale: constants.EnUSLocalKey,
		},
	}
}
