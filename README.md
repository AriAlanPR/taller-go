# taller-go


## **Packages**

- Every Go program is made up of packages.
- Programs start running in package main.
- This program is using the packages with import paths "fmt" and "math/rand". By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```


**Exported names**

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

```
