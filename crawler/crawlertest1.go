package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)


func httpResults(url string) (results string,err error) {
	//解析网页的内容
	client := http.Client{}
	provisionalResults,err := http.NewRequest("GET",url,nil)
	provisionalResults.Header.Add("User-Agent","client-go")
	if err !=nil {
		return "",errors.New("failed to load url")
	}
	response,err :=client.Do(provisionalResults)
	if err !=nil {
		return "",errors.New("an exception has occurred in accessing the page")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "",errors.New("parsing the web page failed")
	}
	results = string(body)
	defer response.Body.Close()
	return results,nil
}

func dataScreening(data string) {
	//将换行替换成空字符
	results := strings.Replace(data,"\n","",-1)
	//搜索出标有title的字符串，并找出电影名字
	movieName :=regexp.MustCompile(`<span class="title">(.*?)</span>`)
	//搜索出电影名称的序号
	serialNumber := regexp.MustCompile(`<a href="https://(.*?)/"`)
	//过滤出html页面中关于电影字段
	screening := regexp.MustCompile(`<div class="hd">(.*?)</div>`).FindAllStringSubmatch(results,-1)
	for _,temporary := range screening {
		fmt.Println(movieName.FindStringSubmatch(temporary[1])[1],"\t",serialNumber.FindStringSubmatch(temporary[1])[1])
	}

}



func main(){
	url := "https://movie.douban.com/top250?start="
	for i:=0;i<=9;i++ {
		ii := 25 * i
		urll := url + strconv.Itoa(ii)+ "&filter="
		data,err := httpResults(urll)
		if err !=nil {
			fmt.Println(err)
		}else {
			dataScreening(data)
		}

	}




}