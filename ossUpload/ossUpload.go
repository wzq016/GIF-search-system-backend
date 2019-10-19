package ossUpload

import (
	"backend/utils"
	"fmt"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//上传对应的Gif文件。一次性执行即可。localpath为gif存储文件夹相对路径
func OssUpload(gif utils.Gifs, localpath string) {
	log.SetPrefix("Error:")

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4FduW6Yf6AZY8ysPGmB9", "2eayaXUYwzCzK8HuOv8yrqRvtmsxd9")
	if err != nil {
		log.Panicln(err)
	}

	bucket, err := client.Bucket("gif-dio")
	if err != nil {
		log.Panicln(err)
	}

	err = bucket.PutObjectFromFile(gif.Name+".gif", localpath+"\\"+gif.Name+".gif")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(gif.Title + "Uploaded")
}

//按文件名对特定的gif图分配签名链接，参数为Gifs类及签名链接有效时长（单位：秒）,链接支持http Get请求
func OssSignLink(gif utils.Gifs, timeSpan int64) string {
	log.SetPrefix("Error:")

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4FduW6Yf6AZY8ysPGmB9", "2eayaXUYwzCzK8HuOv8yrqRvtmsxd9")
	if err != nil {
		log.Panicln(err)
	}

	bucket, err := client.Bucket("gif-dio")
	if err != nil {
		log.Panicln(err)
	}

	signedURL, err := bucket.SignURL(gif.Name+".gif", oss.HTTPGet, timeSpan)
	if err != nil {
		log.Panicln(err)
	}

	return signedURL
}
