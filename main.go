package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

)



func main() {
	//dbshanks.Helper()

	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())
	//////可以读取数据库   来分享数据
	m.Get("((/api)|(/bpi))/xx/.hml", func(r render.Render) {
		value := map[string]string{"name":"xx","age":"12"}
		r.JSON(200, map[string]interface{}{"hello": "world", "key":value})
		//
	})
	m.Get("/api/ll", func(r render.Render) {
		value := map[string]string{"name":"api.ll","age":"12"}
		r.JSON(200, map[string]interface{}{"hello": "world", "key":value})
		//
	})

	m.Get("/api/.*", func(r render.Render) {
		value := map[string]string{"name":"api/.*","age":"12"}
		r.JSON(200, map[string]interface{}{"hello": "world", "key":value})
		//
	})

	m.Get("/api.*", func(r render.Render) {
		value := map[string]string{"name":"api.*","age":"12"}
		r.JSON(200, map[string]interface{}{"hello": "world", "key":value})
		//
	})

	m.Get("", func(r render.Render) {
		value := map[string]string{"name":"xx","age":"12"}
		r.JSON(200, map[string]interface{}{"hello": "world", "key":value})
		//
	})
	m.Get("/hello/:name(/:id)?", func(params martini.Params) string {
		return "Hello ---yy" + params["name"]+params["id"]
	})

	m.Run()



}
