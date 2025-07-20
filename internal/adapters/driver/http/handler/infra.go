package handler

import (
	"text/template"

	"1337b04rd/internal/ports/inbound"
)

type handler struct {
	templates  *template.Template
	middleware inbound.MiddlewareSessionContext
	use        inbound.Service
}

func InitHandler(middleware inbound.MiddlewareSessionContext, use inbound.Service) (inbound.HandlerInter, error) {
	templates, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		return nil, err
	}
	// соңғысының аты болады
	// fmt.Println(templates.Name())
	// for _, t := range templates.Templates() {
	// 	fmt.Println(t.Name())
	// }
	return &handler{templates, middleware, use}, nil
}
