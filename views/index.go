//
// @project Subscriber 2015
// @author Dmitry Ponomarev <demdxx@gmail.com>
//

package views

import (
  "net/http"
  "os"
  "strings"
  "sync"

  "github.com/gopk/config"
  "github.com/gopk/templates"
)

var (
  writeLock sync.Mutex
)

func Index(w http.ResponseWriter, r *http.Request) *templates.HttpResponse {
  return templates.Response(200, "index.html", nil)
}

func Subscribe(w http.ResponseWriter, r *http.Request) *templates.HttpResponse {
  writeLock.Lock()
  defer writeLock.Unlock()

  params := map[string]interface{}{}
  good := false

  if f, err := os.OpenFile(config.String("storage.file"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666); nil == err {
    r.ParseForm()
    if email := strings.TrimSpace(r.Form.Get("email")); len(email) > 0 {
      _, err := f.WriteString(email + "\n")
      good = nil == err
    }
    f.Close()
  }

  if good {
    params["good"] = true
  } else {
    params["error"] = true
  }

  return templates.Response(200, "index.html", params)
}
