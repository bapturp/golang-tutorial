# Create a go package / module

To create a new package:

```sh
go mod init myapp
```

It creates a file that contains:

```txt
module myapp

go 1.20
```

Then we can create a new folder and file:

```sh
mkdir utils
touch utils/utils.go
```

In `utils/utils.go`, we define the package name the same as the folder and file name.

```go
package utils

var Name string = "Baptiste"
```

Then in `./main.go`, we can import that package and access to exported variables or functions. Exported variables and functions start with an uppercase letter, they are accessible outside the package they were declared in

```go
package main

import (utils "myapp/utils")

func main() {
    fmt.Println(utils.Name)
}
```
