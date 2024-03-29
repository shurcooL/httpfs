httpfs
======

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/httpfs.svg)](https://pkg.go.dev/github.com/shurcooL/httpfs)

Collection of Go packages for working with the [`http.FileSystem`](https://go.dev/pkg/net/http#FileSystem) interface.

Installation
------------

```sh
go get github.com/shurcooL/httpfs
```

Directories
-----------

| Path                                                                               | Synopsis                                                                                                   |
|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|
| [filter](https://pkg.go.dev/github.com/shurcooL/httpfs/filter)                     | Package filter offers an http.FileSystem wrapper with the ability to keep or skip files.                   |
| [html/vfstemplate](https://pkg.go.dev/github.com/shurcooL/httpfs/html/vfstemplate) | Package vfstemplate offers html/template helpers that use http.FileSystem.                                 |
| [httputil](https://pkg.go.dev/github.com/shurcooL/httpfs/httputil)                 | Package httputil implements HTTP utility functions for http.FileSystem.                                    |
| [path/vfspath](https://pkg.go.dev/github.com/shurcooL/httpfs/path/vfspath)         | Package vfspath implements utility routines for manipulating virtual file system paths.                    |
| [text/vfstemplate](https://pkg.go.dev/github.com/shurcooL/httpfs/text/vfstemplate) | Package vfstemplate offers text/template helpers that use http.FileSystem.                                 |
| [union](https://pkg.go.dev/github.com/shurcooL/httpfs/union)                       | Package union offers a simple http.FileSystem that can unify multiple filesystems at various mount points. |
| [vfsutil](https://pkg.go.dev/github.com/shurcooL/httpfs/vfsutil)                   | Package vfsutil implements some I/O utility functions for http.FileSystem.                                 |

License
-------

-	[MIT License](LICENSE)
