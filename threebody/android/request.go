package android
import (
	"../common"
)

type Request struct {
	common.FileInfo
	ApiName     string
	Version     string
	NeedEcode   bool
	NeedSession bool
	Params      map[string]string
}


func (this *Request)PrintCode() {
	out := renderTemplate("template/android/request", this)
	out2File(this.Package,this.Name,out)
}