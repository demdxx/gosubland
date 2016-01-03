//
// @project Subscriber 2015
// @author Dmitry Ponomarev <demdxx@gmail.com>
//

package router

import (
  "net/http"

  "github.com/go-martini/martini"
  "github.com/gopk/config"
  "github.com/gopk/templates"

  "../views"
)

func NewRouter(debug bool) http.Handler {
  static := http.StripPrefix("/public/", http.FileServer(http.Dir(config.String("path")+"/public/")))

  r := martini.Classic()
  r.Get("/", templates.HttpHandler(nil, views.Index))
  r.Post("/", templates.HttpHandler(nil, views.Subscribe))
  r.Any("/public/**", func(w http.ResponseWriter, r *http.Request) { static.ServeHTTP(w, r) })
  r.NotFound(NotFoundHandler)
  return r
}

func Http500Handler(resp *templates.HttpResponse) error {
  return templates.Render(resp.Writer, resp.Context, "50x.html")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
  global := map[string]interface{}{
    "query": r.URL.Query(),
  }
  w.WriteHeader(http.StatusNotFound)
  templates.Render(w, global, "404.html")
}
