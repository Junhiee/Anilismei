package anisrc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var proxyURL, _ = url.Parse("http://192.168.31.42:7890")

var tr = &http.Transport{
	Proxy: http.ProxyURL(proxyURL),
}

var client = &http.Client{
	Transport: tr,
}

type Bili struct {
	Code int `json:"code"`
	Data struct {
		HasNext int `json:"has_next"`
		List    []struct {
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType int    `json:"badge_type"`
			Cover     string `json:"cover"`
			FirstEp   struct {
				Cover string `json:"cover"`
				EpID  int    `json:"ep_id"`
			} `json:"first_ep"`
			IndexShow    string `json:"index_show"`
			IsFinish     int    `json:"is_finish"`
			Link         string `json:"link"`
			MediaID      string `json:"media_id"`
			Order        string `json:"order"`
			OrderType    string `json:"order_type"`
			Score        string `json:"score"`
			SeasonID     int    `json:"season_id"`
			SeasonStatus int    `json:"season_status"`
			SeasonType   int    `json:"season_type"`
			SubTitle     string `json:"subTitle"`
			Title        string `json:"title"`
			TitleIcon    string `json:"title_icon"`
		} `json:"list"`
		Num   int `json:"num"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

func ListResp() (mid string) {

	uri := "https://api.bilibili.com/pgc/season/index/result"

	u := url.Values{
		"st":             []string{"4"},
		"order":          []string{"3"},
		"season_version": []string{"-1"},
		"is_finish":      []string{"-1"},
		"copyright":      []string{"-1"},
		"season_status":  []string{"-1"},
		"year":           []string{"-1"},
		"style_id":       []string{"-1"},
		"sort":           []string{"0"},
		"page":           []string{"1"},
		"season_type":    []string{"4"},
		"pagesize":       []string{"20"},
		"type":           []string{"1"},
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		fmt.Println("构造 req 错误", err)
		return
	}

	req.URL.RawQuery = u.Encode()

	req.Header.Add("authority", "api.bilibili.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en")
	req.Header.Add("cookie", "b_lsid=CE5D91BA_18DBBAE42FB; _uuid=EBD510748-4368-4CA10-33106-1535B65F99BD76325infoc; buvid_fp=cf3761d6fa55fab80e5bc8dd44bc73c5; buvid3=A63B588A-ECC7-A9ED-2649-322E8F3DF36E75931infoc; b_nut=1708250776; buvid4=8CBA4D0B-F576-8DDE-9A75-BB0B19A1F5E475931-024021810-71Lmdf6eyFJyG5DAz4rYYA%3D%3D")
	req.Header.Add("origin", "https://www.bilibili.com")
	req.Header.Add("referer", "https://www.bilibili.com/guochuang/index/?from_spmid=666.5.index.0")
	req.Header.Add("sec-ch-ua", "\"Not A(Brand\";v=\"99\", \"Google Chrome\";v=\"121\", \"Chromium\";v=\"121\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("请求报错:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	bili := &Bili{}
	err = json.Unmarshal(body, bili)
	if err != nil {
		fmt.Println(err)
	}
	mid = bili.Data.List[0].MediaID
	fmt.Println(bili.Data.List[0].MediaID)
	return
}

func DetailResp(mid string) {
	rawURL := "https://www.bilibili.com/bangumi/media/"
	uri, _ := url.JoinPath(rawURL, "md"+mid)
	// fmt.Println(uri)

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.bilibili.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "b_lsid=B9F8C4FF_18DBBD21BD8; _uuid=E1010252B2-829C-8CB10-F76B-553F26D586E225604infoc; buvid3=195D9A4D-2C90-80B2-9DE7-BCB3E9F915FF25499infoc; b_nut=1708253126; buvid4=BB0D503C-3F42-0C6E-1DB2-4B77337CEC4B25499-024021810-71Lmdf6eyFJyG5DAz4rYYA%3D%3D; buvid_fp=cf3761d6fa55fab80e5bc8dd44bc73c5")
	req.Header.Add("sec-ch-ua", "\"Not A(Brand\";v=\"99\", \"Google Chrome\";v=\"121\", \"Chromium\";v=\"121\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// TODO 列表页、详情页，返回值结构化