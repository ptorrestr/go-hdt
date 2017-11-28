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

