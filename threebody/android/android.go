package android
import (
	"strings"
	"os"
	"io/ioutil"
)

const (
	OutputPath = "output/android/"
	AndroidPackage = "com.taobao.auction.model"
	PathModel = ".model"
)


///处理参数变量
func HandlerParams(params map[string]string) (s string) {
	var str string = ""
	length := len(params)
	for key, value := range params {
		str=str+value+" "+key
		length--
		if (length!=0) {
			str=str+","
		}
	}
	return str
}

func out2File(packageName string, name string, out []byte) {
	dir := OutputPath+strings.Replace(packageName, ".", "/", -1)+"/"
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(dir+name+".java", out, os.ModePerm)
}



