//
// @project Subscriber 2015
// @author Dmitry Ponomarev <demdxx@gmail.com>
//

package router

import (
  "net/http"

  "github.com/go-martini/martini"
  "github.com/gopk/templates"

  "../views"
)

func NewRouter(debug bool) http.Handler {
  m := martini.Classic()
  m.Get("/", templates.HttpHandler(nil, views.Index))
  m.Post("/", templates.HttpHandler(nil, views.Subscribe))
  m.NotFound(NotFoundHandler)
  return m
}

func Http500Handler(resp *templates.HttpResponse) error {
  return templates.Render(resp.Writer, resp.Context, "50x.html")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
  global := map[string]interface{}{
    "query": r.URL.Query(),
  }
  templates.Render(w, global, "404.html")
}
