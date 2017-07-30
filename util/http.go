package util 

import(
    "errors"
    "strings"
    "net/http"
    "net/url"
    "io/ioutil"
    "fmt"

    "lmq/util/logger"
)

func HttpGet(uri string)(string,error){
    resp,err:=http.Get(uri);
    if err != nil {
        logger.Logger.Errorf("httpGet[%s] failed:%s", uri, err)
        return "",err
    }
    body,err:=ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logger.Logger.Errorf("httpGet[%s] get resp body failed:%s", uri, err)
        return "",err
    }
    return string(body),nil 
}

func HttpPostJson(uri, inputParam string) (string,error) {
    client := &http.Client{}
    req, err := http.NewRequest("POST", uri, strings.NewReader(inputParam))
    if err != nil {
        logger.Logger.Errorf("HttpPostJson failed url=%s, err=%s",uri, err.Error())
        return "", err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded;param=value")
    // 完成后断开连接
    req.Header.Set("Connection", "close")
    res, err := client.Do(req)
    if err != nil {
        logger.Logger.Errorf("HttpPostJson failed url=%s, err=%s",uri, err.Error())
        return "", err
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        logger.Logger.Errorf("HttpPostJson failed url=%s, err=%s",uri, err.Error())
        return "", err
    }
    bodystr := string(body)
    if res.StatusCode == 200 {
        logger.Logger.Tracef("HttpPostJson success url=%s, ret=%s",uri, bodystr)
        return bodystr, nil
    } else {
        logger.Logger.Errorf("HttpPostJson failed url=%s, ret=%s",uri, bodystr)
        return "", errors.New(fmt.Sprintf("statusCode=%d", res.StatusCode))
    }
}

func HttpPostMap(uri string, data map[string]string) (string, error) {
    values := url.Values{}
    for k, v := range data {
        values.Add(k, v)
    }
    return HttpPostJson(uri, values.Encode())
}






