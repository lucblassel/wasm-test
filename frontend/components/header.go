package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

func Header(text string) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("standard-header")),
		elem.Heading2(vecty.Text(text)),
	)
}
