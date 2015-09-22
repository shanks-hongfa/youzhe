package android
import (
	"../common"
	"io/ioutil"
	"text/template"
	"bytes"
)
type Manager struct {
	common.FileInfo
	Model   string
	Request string
	Params  map[string]string
}
func (this *Manager)PrintCode() {
	out := renderTemplate("template/android/manager", this)
	out2File(this.Package,this.Name,out)
}

func renderTemplate(filename string, this interface{}) (out []byte) {
	buf, err := ioutil.ReadFile(filename)
	if err !=nil {
		panic(err)
	}
	templ := template.New("tem")
	templ.Funcs(template.FuncMap{"handlerParams":HandlerParams})
	templ.Parse(string(buf))
	doc := new(bytes.Buffer)
	templ.Execute(doc, this)
	return doc.Bytes()
}