package threebody
import (
	"testing"
	"time"
	"./android"
	"./ios"
)

func init() {
	println("========init======")
}
func Test_templ(t *testing.T) {

	model := new(android.Model)
	model.Name="Android"
	model.Package=android.AndroidPackage////独立出来
	model.Properties=make(map[string]string)
	model.Properties["year"]="long"
	model.Properties["page"]="int"
	model.Properties["album"]="Album"
	model.Author="shanks"
	model.CreateTime=time.Now().Format("2006-01-02 15:04:05")
	model.PrintCode()

	////子类
	childMode := new(android.Model)
	childMode.Name=model.Properties["album"]
	childMode.Package=model.Package
	childMode.Author=model.Author
	childMode.Properties=make(map[string]string)
	childMode.Properties["year"]="long"
	childMode.CreateTime=model.CreateTime
	childMode.PrintCode()
	/////model里面一定会有子类的引用

	//////
	////request
	request := new(android.Request)
	request.Package=android.AndroidPackage
	request.Name="AndroidRequest"
	request.ApiName="mtop.taobao.paimai.tbp2.getAllForLive"
	request.Version="1.0"
	request.NeedEcode=true
	request.NeedSession=true
	request.Params=make(map[string]string)
	request.Params["albumId"]="long"
	request.Params["msgId"]="long"
	request.Author="shanks"
	request.CreateTime=time.Now().Format("2006-01-02 15:04:05")
	request.PrintCode()
	//////manager
	manager := new(android.Manager)
	manager.Author=request.Author
	manager.CreateTime=request.CreateTime
	manager.Params=request.Params
	manager.Package=android.AndroidPackage
	manager.Name="AndroidManager"
	manager.Model=model.Name
	manager.Request=request.Name
	manager.PrintCode()

	//////ios
	modelios := new(ios.Model)
	modelios.Author="shanks"
	modelios.CreateTime=time.Now().Format("2006-01-02 15:04:05")
	modelios.Name="Test"
	modelios.Properties=make(map[string]string)
	modelios.Properties["id"]="int"
	modelios.Properties["name"]="String"
	modelios.Properties["nick"]="String[]"
	modelios.Properties["home"]="Home[]"
	modelios.PrintCode()

	modelioschild := new(ios.Model)
	modelioschild.Author="shanks"
	modelioschild.CreateTime=time.Now().Format("2006-01-02 15:04:05")
	modelioschild.Name="Home"
	modelioschild.Properties=make(map[string]string)
	modelioschild.Properties["id"]="int"
	modelioschild.Properties["name"]="String"
	modelioschild.Properties["nick"]="String[]"
	modelioschild.PrintCode()
}

/*
func Test_impl(t *testing.T) {
	maps:=make(map[string]string)
	maps["nick"]="foutai"
	xx:=maps["jj"]
	if xx==""{
		println("=====")
	}

}

*/
