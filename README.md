# lox 🥯
Extends [samber/lo](https://github.com/samber/lo) with error handling and other useful utilities. 

## Install
```sh
go get github.com/otterize/lox
```

## 💡 Usage
```go
package main

import (
	"fmt"
	"github.com/otterize/lox"
	"strings"
)

func main() {
	names, err := lox.MapErr([]string{"Otter", "Other", "Utter"}, func(s string, i int) (string, error) {
		if s == "" {
			return "", fmt.Errorf("empty name")
		}
		return s, nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Names: %s", strings.Join(names, ","))
}
```


## 👤 Authors
- [Amit Lichtenberg](https://github.com/amitlicht)
