# dgocty (Transform Dgo to Cty and back)

[![](https://goreportcard.com/badge/github.com/lyraproj/dgocty)](https://goreportcard.com/report/github.com/lyraproj/dgocty)
[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/lyraproj/dgocty)
[![](https://github.com/lyraproj/dgocty/workflows/Dgo%20Build/badge.svg)](https://github.com/lyraproj/dgocty/actions)

This module provides conversion routines to convert [Dgo](/lyraproj/dgo) values into [Cty](/zclconf/go-cty) values
and vice versa.

Among other things, the cty values are used by [Terraform](/hashicorp/terraform) to represent configurations and states.

### Using dgocty as a library
To use dgo, first install the latest version of the library:
```sh
go get github.com/lyraproj/dgocty
```
