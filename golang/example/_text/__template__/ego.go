package main
import (
"fmt"
"io"
)
//line sample.ego:1
 func MyTmpl(w io.Writer, e TemplateExec) error  {
//line sample.ego:2
_, _ = fmt.Fprintf(w, "\n\n## 表示\n")
//line sample.ego:4
_, _ = fmt.Fprintf(w, "%v",  e.Title )
//line sample.ego:5
_, _ = fmt.Fprintf(w, "\n\n## 変数の代入\n")
//line sample.ego:7
 title := e.Title 
//line sample.ego:8
_, _ = fmt.Fprintf(w, "\n$title = ")
//line sample.ego:8
_, _ = fmt.Fprintf(w, "%v",  title )
//line sample.ego:9
_, _ = fmt.Fprintf(w, "\n\n## for文\n")
//line sample.ego:11
 for _, foo := range e.FooList { 
//line sample.ego:11
_, _ = fmt.Fprintf(w, "%v",  foo )
//line sample.ego:12
_, _ = fmt.Fprintf(w, "\n")
//line sample.ego:12
 } 
//line sample.ego:13
_, _ = fmt.Fprintf(w, "\n\n## mapのfor文\n")
//line sample.ego:15
 for key, val := range e.FooMap { 
//line sample.ego:15
_, _ = fmt.Fprintf(w, "キー = ")
//line sample.ego:15
_, _ = fmt.Fprintf(w, "%v",  key )
//line sample.ego:15
_, _ = fmt.Fprintf(w, " 値 = ")
//line sample.ego:15
_, _ = fmt.Fprintf(w, "%v",  val )
//line sample.ego:15
_, _ = fmt.Fprintf(w, " 変数のスコープ = ")
//line sample.ego:15
_, _ = fmt.Fprintf(w, "%v",  title )
//line sample.ego:16
_, _ = fmt.Fprintf(w, "\n")
//line sample.ego:16
 } 
//line sample.ego:17
_, _ = fmt.Fprintf(w, "\n\n## templateの入れ子\n")
//line sample.ego:19
_, _ = fmt.Fprintf(w, "%v",  e.Content )
//line sample.ego:20
_, _ = fmt.Fprintf(w, "\n")
return nil
}
//line subsample.ego:1
 func MySubTmpl(w io.Writer, title string) error  {
//line subsample.ego:2
_, _ = fmt.Fprintf(w, "\n> 子要素への変数渡し ")
//line subsample.ego:2
_, _ = fmt.Fprintf(w, "%v",  title )
return nil
}
