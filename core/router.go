package core

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// handler 在 app 内注册的，注释写文档没有办法推理路由前缀
func (r *RouterGroup) WithDoc(doc *DocItem, handler func(ctx *gin.Context)) *RouterGroup {
	// register handler
	r.Handle(doc.Method, doc.Path, func(ctx *gin.Context) {
		// 似乎还可以在这里添加验证和直接填参数 handler(ctx,page,limit,params...)
		handler(ctx)
	})

	// write doc
	basePath := reflect.ValueOf(r).Elem().FieldByName("basePath")
	doc.Path = basePath.String() + doc.Path

	if doc.Body != nil {
		doc.Body = reflectSwagSpec(doc.Body)
	}

	if doc.Params != nil {
		doc.Params = reflectSwagSpec(doc.Params)
	}

	DocInstance.Add(doc)
	return r
}

// reflectSwagSpec reflects the struct and get the json tag and swag tag, and make it as a map.
// it will return a map like this:
//
//	{
//		"key1": {
//			"type": "string",
//			"label": "label1",
//			"required": true
//		},
//		"key2": {
//			"type": "integer",
//			"label": "label2",
//			"required": false
//		},
//		...
//	}
func reflectSwagSpec[T any](v T) map[string]map[string]any {
	var _map = make(map[string]map[string]any)
	var fieldLenth = reflect.TypeOf(v).NumField()
	var fields []reflect.StructField

	for i := 0; i < fieldLenth; i++ {
		fields = append(fields, reflect.TypeOf(v).Field(i))
	}
	for _, field := range fields {
		specstr := field.Tag.Get("swag")
		specs := strings.Split(specstr, ",")

		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}

		if _map[key] == nil {
			_map[key] = make(map[string]any)
		}

		for _, spec := range specs {
			switch spec {
			case "required":
				_map[key]["required"] = true
			case "string":
				_map[key]["type"] = "string"
			case "int":
				_map[key]["type"] = "integer"

			case "float64":
				_map[key]["type"] = "number"

			case "interface{}":
				_map[key]["type"] = "object"

			case "array":
				_map[key]["type"] = "array"

			default:
				_map[key]["label"] = spec
			}
		}

	}
	return _map

}

// GROUP
func (r *RouterGroup) Group(path string) *RouterGroup {
	return &RouterGroup{RouterGroup: r.RouterGroup.Group(path)}
}

// USE
func (r *RouterGroup) Use(middleware ...gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.Use(middleware...)
	return r
}

// GET
func (r *RouterGroup) GET(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.GET(path, handler)
	return r
}

// POST
func (r *RouterGroup) POST(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.POST(path, handler)
	return r
}

// PUT
func (r *RouterGroup) PUT(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.PUT(path, handler)
	return r
}

// DELETE
func (r *RouterGroup) DELETE(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.DELETE(path, handler)
	return r
}

// PATCH
func (r *RouterGroup) PATCH(path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.PATCH(path, handler)
	return r
}

// Handle
func (r *RouterGroup) Handle(method, path string, handler gin.HandlerFunc) *RouterGroup {
	r.RouterGroup.Handle(method, path, handler)
	return r
}
