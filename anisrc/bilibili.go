package anisrc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var client = &http.Client{}

func ListResp() {

	uri := "https://api.bilibili.com/pgc/season/index/result"

	method := "GET"
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

	req, err := http.NewRequest(method, uri, nil)
	req.URL.RawQuery = u.Encode()

	if err != nil {
		fmt.Println(err)
		return
	}
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
