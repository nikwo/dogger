# :dog: Dogger - Go logger

### Quickstart

---
```shell
go get github.com/nikwo/dogger
```

### Usage

---

:dog2: Basic example:
```go
package main
// example basic usage

import (
	"context"
	"github.com/nikwo/dogger"
)

func main() {
	ctx := context.Background()
	dogger.WithContext(ctx).Trace("hello")
	dogger.Debug("world")
	dogger.WithContext(ctx).WithFields("chiba", "inu").Info(emoji.Sprintf("i :heart:  pugs!"))
	dogger.Info("dogs aren't cute")
	dogger.WithContext(ctx).Error("error: dogs are cute, you've lied!")
}
```
It produces output in stdout:
```text
🐶 [trace] 2023-03-08T18:08:42+03:00 /$HOME/dogger/example/main.go(main.main:12) message="hello"
🐶 [debug] 2023-03-08T18:08:42+03:00 /$HOME/dogger/example/main.go(main.main:13) message="world"
🐶 [info] 2023-03-08T18:08:42+03:00 /$HOME/dogger/example/main.go(main.main:14) entry="chiba" value="inu" message="i ❤️  pugs!"
🐶 [info] 2023-03-08T18:08:42+03:00 /$HOME/dogger/example/main.go(main.main:15) message="dogs aren't cute"
🐶 [error] 2023-03-08T18:08:42+03:00 /$HOME/dogger/example/main.go(main.main:16) message="error: dogs are cute, you've lied!"
```

Log level could be updated in any time with
```go
dogger.SetLevel(lvl level.Level)
```

You can set custom writer with 
```go
dogger.SetWriter(w io.Writer)
```

Also you can set your own custom formatter by implementing interface Format from github.com/nikwo/dogger/format and call 
```go
dogger.SetFormatter(f format.Format)
```