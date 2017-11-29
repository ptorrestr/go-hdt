# Golang - HDT

Welcome to Golang-HDT. This is a driver for HDT in Go.

## Testing
Make sure you have a fresh Go environment. Also, you will need a C++ compiler.

This project requires a fresh version of HDT. You need to use the develop branch. Installed in your local machine.
If it is not installe in the standard folders, you will need to provide the folder to Go using `PKG_CONFIG_PATH`

The dependencies of this project can be installed:

	$

Then you can test

	$ PKG_CONFIG_PATH=/hdt/ go test

## Usage
This project can be then imported in another Go project to get access into a HDT file

```go
package main

import "github.com/ptorrestr/go-hdt"
import "fmt"

func main() {
	fmt.Printf("Example go-hdt\n")
	hdtMap := hdt.OpenHDT("/my/file.hdt")
	it := hdtMap.Search("/uri1/", "", "")
	for _, triple := range it.GetAll() {
		fmt.Printf("%s\n", triple)
	}
	hdtMap.Free()
}
```
When compiling your project, you need to provide access to your HDT library. 
If it was not installed in the standard folders, then you need to use `PKG_CONFIG_PATH` to indicate the folder.

	$ PKG_CONFIG_PATH=/my/hdt/lib/pkgconfig go run

Be aware that the C objects created during the executing have to be manually deleted.
This can be archived by executing the function `Free()` for each type associated to this package.
