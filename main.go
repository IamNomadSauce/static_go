package main

import (
  "html/template"
  "log"
  "net/http"
  "path/filepath"
)

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", serveTemplate)

  log.Print("Listening on :3000")

  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}
// package main
//
// import (
//   "fmt"
//   "html/template"
//   "net/http"
//   "path/filepath"
//   "github.com/oxtoacart/bpool"
// )
//
// var templates map[string]*template.Template
// var bufpool *bpool.BufferPool
//
// func init() {
//   if templates == nil {
//     templates = make(map[string]*template.Template)
//   }
//
//
//   templatesDir := "/"
//
//   layouts, err := filepath.Glob(templatesDir + "templates/*.tmpl")
//   if err != nil {
//     fmt.Printf("Error creating layouts: %v", err)
//   }
//
//   includes, err:= filepath.Glob(templatesDir + "includes/*.tmpl")
//   if err != nil {
//     fmt.Printf("Error with includes: %v",err)
//   }
//
//   for _, layout := range layouts {
//     files := append(includes, layout)
//     templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
//   }
// }
//
// func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
//   tmpl, ok := templates[name]
//   if !ok {
//     return fmt.Errorf("The template %s does not exist.", name)
//   }
//
//   buf := bufpool.Get()
//   defer bufpool.Put(buf)
//
//   err := tmpl.ExecuteTemplate(buf, "base.tmpl", data)
//   if err != nil {
//     return err
//   }
//   w.Header().Set("Content-Type", "text/html;")
//   buf.WriteTo(w)
//   return nil
// }
//
// func main() {
//   init()
//
// }
