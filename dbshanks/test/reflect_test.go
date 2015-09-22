package test
import (
	"testing"
	"reflect"
	"runtime"
/*
	"youzhe/dbshanks"
*/
	"html/template"

	"bytes"
)

func Test_reflect(t *testing.T) {
	i := int32(3)
	ii := reflect.TypeOf(i)
	t.Log(ii.String()=="int32")

	t.Log("就是不通过")
	t.Log(runtime.NumCPU())

	//dbshanks.Init()



}

type Person struct {
	UserName string
}
/*public class BidItem  implements IMTOPDataObject {
public long currentPrice;
public long nextPrice;
public int status;
public String buyId;
public boolean isMeWin;
public String lotNumber;
public String depositUrl;

}*/

type AndroidModel struct   {
	ClassName string
}

func Test_templ(t *testing.T){
	t_ := template.New("fieldname example")
	t_, _ = t_.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
//	t_.Execute(os.Stdout, p)

	////
	var doc bytes.Buffer
	t1:=template.New("android")
	t1,_=t1.Parse("public class {{.UserName}}  implements IMTOPDataObject{\n}")
	p = Person{UserName: "Astaxie"}
	doc.WriteString("public {{.key}} {{.name}};\n")
	t1.Execute(&doc, p)
	t.Log(doc.String())
	t.Log("++++++")
}