package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	router "marwan.io/vecty-router"
)

type pageView struct {
	vecty.Core
	Input string
}

func (p *pageView) Render() vecty.ComponentOrHTML {
	exactMatchOpts := router.NewRouteOpts{ExactMatch: true}
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			router.NewRoute("/", &home{}, exactMatchOpts),
		),
	)
}

type Markdown struct {
	vecty.Core
	Input string `vecty:"prop"`
}

func (m *Markdown) Render() vecty.ComponentOrHTML {

	unsafeHTML := blackfriday.Run([]byte(m.Input))

	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}

func main() {

	vecty.SetTitle("Hello from Vecty!")
	vecty.RenderBody(&pageView{})

}
