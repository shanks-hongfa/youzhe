package ios
import (
	"io/ioutil"
	"text/template"
	"bytes"
	"strings"
	"os"
)


var ios2JavaMap map[string]string
const prefix  = "@property (nonatomic, assign) "
const prefix4String  = "@property (nonatomic, retain) "
const prefix4Array  ="@property (nonatomic, strong) NSArray"
func init() {
	ios2JavaMap=make(map[string]string)
	ios2JavaMap["String"]="NSString"
	ios2JavaMap["long"]="long int"
	ios2JavaMap["int"]="int"
	ios2JavaMap["float"]="float"
	ios2JavaMap["double"]="double"
	ios2JavaMap["boolean"]="bool"
}
//@property (nonatomic, strong) NSArray<PMRecommendSellerModel> *
func FiledRender(properties map[string]string) (result string) {
	str:=""
//	interfaceStr:=""
	for key,value:=range properties{
		println("===========")
		if strings.EqualFold(value,"String") {
			str=str +prefix4String+ios2JavaMap[value]+" *"+key+";\n"
		} else if vv1,ok:=ios2JavaMap[value];ok {
			str=str+ prefix+vv1+" "+key+";\n"
		}else if strings.HasSuffix(value,"[]"){
			println("++++------+++++")
			subValue := string([]byte(value)[:len(value)-2])
			println("***"+value)
			if strings.EqualFold(subValue, "String") {
				str=str+"///这是String类型的数组\n"
				str=str+prefix4Array+" *"+key+";\n"
				println(str)
			}else if  vv1,ok:=ios2JavaMap[subValue];ok {
				str=str+"///这是"+vv1+"类型的数组\n"
				str=str+prefix4Array+" *"+key+";\n"
				println(str)
			}else {
				////需要定义协议
				//out:=renderTemplate("template/ios/protocol",&struct{Name string}{subValue})
				//interfaceStr=interfaceStr+string(out)+"\n"
				str=str+prefix4Array+"<"+subValue+">"+" *"+key+";\n"
				println(str)
			}
		}else{
			/////
			str=str+ prefix4Array+" *"+key+";\n"
		}
	}
	println(str)
	//str=str+"@end"+"\n"+interfaceStr
	return str
}

func ImportRender(value string) string {
	str := ""
	importMap := make(map[string]int8)
	 if (strings.HasSuffix(value, "[]")) {
		subValue := string([]byte(value)[:len(value)-2])
		if _, ok := ios2JavaMap[subValue]; !ok {
			///////
			if _ ,ok:=importMap[subValue];!ok{
				str=str+"#import \""+subValue+".h\"\n"
				importMap[value]=0
			}
		}
	} else if _, ok := ios2JavaMap[value]; !ok {
		/////////
		if _ ,ok:=importMap[value];!ok{
			str=str+"#import \""+value+".h\""
			importMap[value]=0
		}

	}
	return str;
}



func out2File( name string, out []byte) {
	dir := "output/ios/model/"
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(dir+name+".h", out, os.ModePerm)
}

func renderTemplate(filename string, this interface{}) (out []byte) {
	buf, err := ioutil.ReadFile(filename)
	if err !=nil {
		panic(err)
	}
	templ := template.New("tem")
	templ.Funcs(template.FuncMap{"filedRender":FiledRender})
	templ.Funcs(template.FuncMap{"importRender":ImportRender})
	templ.Parse(string(buf))
	doc := new(bytes.Buffer)
	templ.Execute(doc, this)
	return doc.Bytes()
}
