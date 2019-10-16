package cache

import (
	// "fmt"
	"os"
	"io/ioutil"
	"strings"
	"encoding/base64"
	"backend/utils"
	// "strconv"
	// jsoniter "github.com/json-iterator/go"
)

func FastWrite(filepath string,content []byte){
	w1, _:=os.OpenFile("cache",os.O_CREATE|os.O_TRUNC,0644)
	_, _=w1.Write(content)
	_=w1.Close()
}

func FastAppend(filepath string,content []byte){
	w1, _:=os.OpenFile("cache",os.O_APPEND,0644)
	_, _=w1.Write(content)
	_=w1.Close()
}

func OfflineCacheInit(){
	_,err:=os.Stat(utils.CACHE_DIR+"cache_name")
	if os.IsNotExist(err){
		_=os.Mkdir(utils.CACHE_DIR+"cache_name",os.ModePerm)
	}
	_,err=os.Stat(utils.CACHE_DIR+"cache_title")
	if os.IsNotExist(err){
		_=os.Mkdir(utils.CACHE_DIR+"cache_title",os.ModePerm)
	}
}

func OfflineCacheAppend(keyword string,gif []utils.Gifs){
	w1, _:=os.OpenFile(utils.CACHE_DIR+"cache_name/"+base64.URLEncoding.EncodeToString([]byte(keyword)),os.O_CREATE|os.O_TRUNC,0644)
	// _, _=w1.Write([]byte(strconv.FormatInt(int64(len(gif)),10)+"#"))
	for i:=0;i<len(gif);i++{
		_, _=w1.Write([]byte(gif[i].Name+"#"))
	}
	_=w1.Close()
	w1, _=os.OpenFile(utils.CACHE_DIR+"cache_title/"+base64.URLEncoding.EncodeToString([]byte(keyword)),os.O_CREATE|os.O_TRUNC,0644)
	// _, _=w1.Write([]byte(strconv.FormatInt(int64(len(gif)),10)+"#"))
	for i:=0;i<len(gif);i++{
		_, _=w1.Write([]byte(gif[i].Title+"#"))
	}
	_=w1.Close()
}

func OfflineCacheQuery(keyword string) []string{
	var res []string
	fname:=utils.CACHE_DIR+"cache_name/"+base64.URLEncoding.EncodeToString([]byte(keyword))
	_,err:=os.Stat(fname)
	if os.IsNotExist(err){
		res=append(res,"Failed")
		return res
	}
	res=append(res,"Succeed")
	ind,_:=ioutil.ReadFile(fname)
	res=append(res,strings.Split(string(ind),"#")...)
	return res
}

func OfflineCacheReload() map[string][]utils.Gifs{
	// m_name:=make(map[string][]string)
	// m_title:=make(map[string][]string)
	m:=make(map[string][]utils.Gifs)
	gif:=new(utils.Gifs)
	var gifs []utils.Gifs
	// var res []string
	dir,_:=ioutil.ReadDir(utils.CACHE_DIR+"cache_name/")
	for _,fi:=range dir{
		if fi.IsDir(){

		}else{
			gifs=make([]utils.Gifs,0)
			b, _:=base64.URLEncoding.DecodeString(fi.Name())
			b0,_:=ioutil.ReadFile(utils.CACHE_DIR+"cache_name/"+fi.Name())
			lis_name:=strings.Split(string(b0),"#")
			b0,_=ioutil.ReadFile(utils.CACHE_DIR+"cache_title/"+fi.Name())
			lis_title:=strings.Split(string(b0),"#")
			for i:=0;i<len(lis_name);i++{
				gif.Name=lis_name[i]
				gif.Title=lis_title[i]
				gifs=append(gifs,*gif)
			}
			m[string(b)]=gifs[0:len(gifs)-1]
		}
	}
	return m;
}

// func main(){
// 	ind,_:=ioutil.ReadFile("cache")
// 	lis:=strings.Fields(string(ind))
// 	fmt.Println(lis)
// 	// fmt.Println(base64.URLEncoding.EncodeToString([]byte("哈哈")))
// 	var gif0 []utils.Gifs
// 	gif0=append(gif0,make([]utils.Gifs,3)...)
// 	gif0[0].Name="111"
// 	gif0[1].Name="222"
// 	gif0[2].Name="333"
// 	OfflineCacheInit()
// 	// OfflineCacheAppend("哈哈哈",gif0)
// 	fmt.Println(OfflineCacheQuery("哈哈哈"))
// }

