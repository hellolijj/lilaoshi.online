package crawler

import (
	"cookie-shop-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)
type CookiesRep struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		HasMore bool `json:"hasMore"`
		List    []struct {
			Title                    string `json:"title"`
			SubTitle                 string `json:"subTitle"`
			ImageURL                 string `json:"imageUrl"`
			Price                    string `json:"price"`
			URL                      string `json:"url"`
		} `json:"list"`
	} `json:"data"`
}

var urlMap = map[string]string {
	"经典法式": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=361kl9s9iyza2&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"精品慕斯": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=2fp47hr5iket6&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"奶油戚风": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=35wnk1ubypjmi&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"情人系列": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=2g1e3inbekasa&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"儿童系列": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=ntawhz1m&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"女神系列": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=8fq08g351&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"男士系列": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=92pyp4eg&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"长辈系列": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=auc8us951&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
	"高级定制": `https://shop360322.youzan.com/wscshop/goods-api/goodsByTagAlias.json?alias=2xa7rlct9kday&page=1&pageSize=20&offlineId=0&json=1&activityPriceIndependent=1&uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8`,
}

var kind = map[string]int {
"经典法式": 1,
"精品慕斯": 2,
"奶油戚风": 3,
"情人系列": 4,
"儿童系列": 5,
"女神系列": 6,
"男士系列": 7,
"长辈系列": 8,
"高级定制": 9,
}

func fetchCookies(url string) (*CookiesRep, error) {
	client := &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent","Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Add("Cookie","yz_log_uuid=e9a2a9d5-e75f-c7dd-9cbf-73505fae1dd8; yz_log_ftime=1631419664204; _kdt_id_=168154; KDTSESSIONID=YZ1038124355618041856YZvVFhvhE7; nobody_sign=YZ1038124355618041856YZvVFhvhE7; yz_log_seqb=1667549735297; trace_sdk_context_pv_id=/wscshop/showcase/feature~3b634850-6f6f-4a45-b5dc-02eebe3c7631; yz_log_seqn=13")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cookieResp CookiesRep
	err = json.Unmarshal(body, &cookieResp)
	if err != nil {
		return nil, err
	}

	return &cookieResp, nil
}

func FetchCookies() []models.Good {
	var res []models.Good
	for k, u := range urlMap {
		resp ,err := fetchCookies(u)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, item := range resp.Data.List {
			good := models.Good{
				Name:   item.Title,
				Cover:  item.ImageURL,
				Image1: "",
				Image2: "",
				Price:  item.Price,
				Intro:  item.SubTitle,
				Stock:  10,
				TypeId: kind[k],
			}
			res = append(res, good)
		}
	}
	return res
}


func main() {



}
