package ios
import "../common"
type Model  struct {
	common.FileInfo
	Properties map[string]string
}

func (this *Model)PrintCode() {

	out := renderTemplate("template/ios/model", this)
	out2File(this.Name, out)

}
