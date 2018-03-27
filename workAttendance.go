package main
import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    // "math/rand"
    "encoding/json"
    // "github.com/hokaccha/go-prettyjson"
)
func main() {
    // 构造HTTP Request
    req0, _ := http.NewRequest("POST", "https://e.xinrenxinshi.com/attendance/ajax-sign", strings.NewReader("{\"longitude\":\"116.27678\",\"latitude\":\"40.0479\",\"accuracy\":\"30.0\",\"timestamp\":1521024177811,\"signature\":\"AdhzKmKCTCn+5HI1KIqrtKrfZ8c=\"}"))
    req0.Header.Add("Connection", "keep-alive")
    req0.Header.Add("Origin", "https://e.xinrenxinshi.com")
    req0.Header.Add("X-CSRF-TOKEN", "NWRlNDczMzUEBgMMAlIHBQBTBAQDAQEFDVdTAVZVCw1RUgEDA1YBBA==")
    req0.Header.Add("isAjax", "true")
    req0.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/6.2 TBS/043909 Mobile Safari/537.36 MicroMessenger/6.6.5.1280(0x26060533) NetType/WIFI Language/zh_CN")
    req0.Header.Add("X-Tingyun-Id", "yhbknIt6i2Y;r=24177918")
    req0.Header.Add("Content-Type", "application/json;charset=UTF-8")
    req0.Header.Add("Accept", "application/json, text/plain, */*")
    req0.Header.Add("Referer", "https://e.xinrenxinshi.com/")
    req0.Header.Add("Accept-Language", "zh-CN,en-US;q=0.8")
    req0.Header.Add("Cookie", "TY_SESSION_ID=866f2eec-b93e-4e05-ac03-1f5f02c5e49f; JSESSIONID=ft2yoay5osfdhl898l5v011q; QJYDSID=b1710d47e9f14f8caaabff84e16fb8aa_cd8e6e0e1c3b49c98f85b9191452c824; ONEAPM_BI_sessionid=6285.629|1521020525746; sensorsdata2015jssdkcross={\"distinct_id\":\"cd8e6e0e1c3b49c98f85b9191452c824\",\"$device_id\":\"16208b839c376-0997b6a3ce4de7-6145272-230400-16208b839c89\",\"props\":{\"$latest_traffic_source_type\":\"直接流量\",\"$latest_referrer\":\"\",\"$latest_referrer_host\":\"\",\"$latest_search_keyword\":\"未取到值_直接打开\"},\"first_id\":\"16208b839c376-0997b6a3ce4de7-6145272-230400-16208b839c89\"}")
    req0.Header.Add("X-Requested-With", "com.tencent.mm")

    // 发射HTTP Request
    client0  := &http.Client{}
    resp0, _ := client0.Do(req0)
    body0, _ := ioutil.ReadAll(resp0.Body)
    resp0.Body.Close()

    // 打印结果
    // fmt.Printf("%s", body0)

    // body0Json := make(map[string]interface{})
    // body0DataJson := body0Json["data"].(map[string]interface{})
    var body0Json struct{
        Code int
        Status bool
        Message string
        Data struct{
            ParentDeps, Msg string
            Date int
        }
    }

    json.Unmarshal(body0, &body0Json)
    // body0JsonPretty, _ := prettyjson.Marshal(body0Json)
    body0JsonPretty, _ := json.MarshalIndent(body0Json, "", "    ")
    msg := []rune(body0Json.Data.Msg);
    hour := string(msg[5:7]);
    minute := string(msg[8:10]);
    fmt.Printf("%s\n", body0JsonPretty)


    // Server酱推送结果到手机微信端
    req1, _  := http.NewRequest("POST", "https://sc.ftqq.com/SCU2341T4e27c3e1eb7d16d48e89ae2217f2591257e240ab9a72d.send", strings.NewReader(fmt.Sprintf("text=工作考勤丨%s点%s分&desp=```%s```", hour, minute, body0)))
    // req1, _  := http.NewRequest("GET", fmt.Sprintf("https://sc.ftqq.com/SCU2341T4e27c3e1eb7d16d48e89ae2217f2591257e240ab9a72d.send?text=%d&desp=```%s```", rand.Intn(100), body0), nil)
    req1.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")
    req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    client1  := &http.Client{}
    resp1, _ := client1.Do(req1)
    // body1, _ := ioutil.ReadAll(resp1.Body)
    resp1.Body.Close()

    // fmt.Printf("%s\n", body1)
}
