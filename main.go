//
// @project Subscribe Project 2015
//

package main

import (
  "flag"
  "fmt"
  "log"
  "math/rand"
  "net"
  "net/http"
  "net/http/fcgi"
  "path/filepath"
  "time"

  "github.com/gopk/config"
  "github.com/gopk/templates"

  "./router"
)

var (
  flagDebug        = flag.Bool("debug", false, "Debug mode")
  flagHost         = flag.String("host", "", "Listen host address")
  flagBaseDir      = flag.String("basedir", "", "Base dir is prefix for all other paths")
  flagConfigPath   = flag.String("config", "conf/main.conf", "Config file path")
  flagTemplatePath = flag.String("templates", "templates/", "Template dirrectory")
  flagFastCGI      = flag.Bool("fastcgi", false, "Use FastCGI protocol")
)

// INIT Program
func init() {
  // Parse command line
  flag.Parse()

  // Init random
  rand.Seed(time.Now().Unix())

  // Load config
  conf, err := config.GlobalByFile(path(*flagConfigPath), "yaml")
  fatalError(err, "Service config not found!!!")

  conf["path"] = path("")

  // Update config from subsection
  if *flagDebug {
    conf.UpdateByPath("dev")
  } else {
    conf.UpdateByPath("prod")
  }
  delete(conf, "dev")
  delete(conf, "prod")
  conf.Prepare("{{", "}}")

  if *flagDebug {
    j, _ := conf.JSONPrettify()
    fmt.Println("conf", string(j))
  }

  if len(*flagHost) < 1 {
    *flagHost = conf.StringOrDefault("server.listen", ":9001")
  }

  // ** [Init templates]
  templates.InitGlobalRender(path(*flagTemplatePath), "", !*flagDebug)
  templates.RegisterHandler(500, router.Http500Handler)
}

///////////////////////////////////////////////////////////////////////////////
/// MAIN
///////////////////////////////////////////////////////////////////////////////

func main() {
  // Run HTTP server
  fmt.Println("Domain: " + config.String("server.domain"))
  fmt.Println("Start rotator server: " + *flagHost)

  if *flagFastCGI {
    fmt.Println("FastCGI Mode: ON")
  }
  if *flagDebug {
    fmt.Println("Debug Mode: ON")
  }

  // Run server
  if *flagFastCGI {
    if tcp, err := net.Listen("tcp", *flagHost); nil == err {
      fcgi.Serve(tcp, router.NewRouter(*flagDebug))
    } else if err != nil {
      log.Fatal(err)
    }
  } else {
    var err error
    http.Handle("/", router.NewRouter(*flagDebug))
    http.Handle("/public/", http.FileServer(http.Dir(*flagBaseDir)))
    if err = http.ListenAndServe(*flagHost, nil); nil != err {
      log.Fatal(err)
    }
  }
}

///////////////////////////////////////////////////////////////////////////////
/// Helpers
///////////////////////////////////////////////////////////////////////////////

func path(path string) string {
  if filepath.IsAbs(path) {
    return path
  }
  return filepath.Join(*flagBaseDir, path)
}

func fatalError(err error, msg string) {
  if nil != err {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}
