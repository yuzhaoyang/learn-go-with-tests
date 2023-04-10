package fucntiontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFuction8(t *testing.T) {

	a := MyStruct{i: 30}
	b := &MyStruct{i: 40}
	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	myFunction(a, b)
	fmt.Printf("after calling  - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

type MyStruct struct {
	i int
}

func myFunction(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

func TestFuction9(t *testing.T) {

	var s string = "{\"i\":1,\"actual_amount\":\"1300\",\"author_id\":\"\",\"code_valid_end_time\":\"0\",\"code_valid_start_time\":\"0\",\"commodity_amount\":\"1300\",\"commodity_ec_discount_amount\":\"0\",\"create_time\":\"\",\"detail_enter_method\":\"\",\"detail_enter_page\":\"\",\"enter_source\":\"\",\"fee_type\":\"0\",\"first_category_name\":\"美食\",\"first_cid\":\"1000000\",\"item_id\":\"\",\"merchant_address\":\"中国（四川）自由贸易试验区成都高新区吉泰路20号1栋1层101号\",\"merchant_area_id\":\"\",\"merchant_city_id\":\"\",\"merchant_name\":\"四川书亦餐饮管理有限公司\",\"merchant_nickname\":\"四川书亦餐饮管理有限公司\",\"merchant_prov_id\":\"\",\"merchant_userid\":\"200_2749225662429262\",\"miniapp_id\":\"\",\"miniapp_name\":\"\",\"order_id\":\"\",\"order_item_id\":[],\"order_limit\":\"10\",\"order_scene\":\"\",\"order_source\":\"livebroadcasting\",\"poi_enter_page\":\"\",\"poi_num\":\"4315\",\"price\":\"1300\",\"product_id\":\"1761228659172392\",\"product_name\":\"【达人特惠】一桶水果茶\",\"product_type\":\"1\",\"quantity\":\"1\",\"room_anchor_id\":\"1002375244155992\",\"room_id\":\"7215087590919834429\",\"sale_end_time\":\"1680451199\",\"sale_start_time\":\"1679630100\",\"second_category_name\":\"饮品\",\"second_cid\":\"1015000\",\"service_provider_id\":\"\",\"service_provider_name\":\"\",\"shop_id\":\"6962037026589771776\",\"sku_id\":\"1761228659172392\",\"third_category_name\":\"茶饮果汁\",\"third_cid\":\"1015003\",\"trade_limit\":\"\",\"user_id\":\"92946585983\",\"user_limit\":\"10\"}"

	stct := &MyStruct{}
	if err := json.Unmarshal([]byte(s), stct); err != nil {

		fmt.Printf("err xxxxxxxxxxxx")

	} else {
		fmt.Printf("pass")
	}

	fmt.Printf("%%%%%%%%%%%%%%")

}

// 订单商品明细信息
type SubOrderInfo struct {
	ShopId                  string                 `json:"shop_id"`
	ProductId               string                 `json:"product_id"`
	Sku                     map[string]interface{} `json:"sku"`
	AuthorId                string                 `json:"author_id"`
	RoomId                  string                 `json:"room_id"`
	FirstCid                string                 `json:"first_cid"`
	SecondCid               string                 `json:"second_cid"`
	ThirdCid                string                 `json:"third_cid"`
	FourthCid               string                 `json:"fourth_cid"`
	TradeType               int                    `json:"trade_type"`
	Biz                     int                    `json:"biz"`
	CommodityAmount         string                 `json:"commodity_amount"`
	CommodityDisCountAmount string                 `json:"commodity_ec_discount_amount"`
	OrderId                 string                 `json:"order_id"`
	Stype                   string                 `json:"stype"`
	SubType                 string                 `json:"sub_type"`
	ProductBiz              int                    `json:"product_biz"`
	RoomAnchorId            string                 `json:"room_anchor_id"` //本地生活直播间id
	SpuId                   string                 `json:"spu_id"`         //本地生活商品id
	ProductType             string                 `json:"product_type"`   //本地生活商品类型，"15"代表次卡商品
	EcomSceneId             string                 `json:"ecom_scene_id"`  //电商购物来源信息，多个，英文逗号分隔
	FeeType                 string                 `json:"fee_type"`       //子单类型
}

type SubOrderInfo1 struct {
	ShopId                     string                 `json:"shop_id"`
	ProductId                  string                 `json:"product_id"`
	Sku                        map[string]interface{} `json:"sku"`
	AuthorId                   string                 `json:"author_id"`
	RoomId                     string                 `json:"room_id"`
	FirstCid                   string                 `json:"first_cid"`
	SecondCid                  string                 `json:"second_cid"`
	ThirdCid                   string                 `json:"third_cid"`
	FourthCid                  string                 `json:"fourth_cid"`
	TradeType                  int                    `json:"trade_type"`
	Biz                        int                    `json:"biz"`
	CommodityAmount            string                 `json:"commodity_amount"`
	CommodityDisCountAmount    string                 `json:"commodity_ec_discount_amount"`
	CommodityAmountInt         int64                  `json:"commodity_amount_int"`
	CommodityDisCountAmountInt int64                  `json:"commodity_ec_discount_amount_int"`
	OrderId                    string                 `json:"order_id"`
	Stype                      string                 `json:"stype"`
	SubType                    string                 `json:"sub_type"`
	ProductBiz                 int                    `json:"product_biz"`
	FxhPeriods                 []string               `json:"fxh_periods"`
	EcomSceneId                string                 `json:"ecom_scene_id"` //电商购物来源信息，多个，英文逗号分隔
	GroupList                  string                 `json:"group_list"`    //名单id列表
}

func TestFuction10(t *testing.T) {

	//stct := &SubOrderInfo{}
	//
	//marshal, _ := json.Marshal(stct)
	//
	//fmt.Printf("%s", marshal)

	stct1 := &SubOrderInfo1{}

	marshal1, _ := json.Marshal(stct1)

	fmt.Printf("%s", marshal1)

}
