package android
import (
	"../common"
)
type Model struct {
	common.FileInfo
	Properties map[string]string ///key value
}

func (this *Model)PrintCode() {

	out := renderTemplate("template/android/model", this)
	out2File(this.Package, this.Name, out)

}

