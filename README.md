# getmyip

installation
`go get -u github.com/AndrusGerman/getmyip`


### Private
```go
package main

import (
	"fmt"

	"github.com/AndrusGerman/getmyip"
)

func main() {
	ip := getmyip.GetLocalIP()
	fmt.Println(ip)
}

```


### Public
```go
package main

import (
	"fmt"

	"github.com/AndrusGerman/getmyip"
)

func main() {
	ip := getmyip.GetPublicIP()
	fmt.Println(ip)
}
```
