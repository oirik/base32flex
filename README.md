# base32flex

[![Build Status](https://travis-ci.org/oirik/base32flex.svg?branch=master)](https://travis-ci.org/oirik/base32flex)
[![GoDoc](https://godoc.org/github.com/oirik/base32flex?status.svg)](https://godoc.org/github.com/oirik/base32flex)
[![apache license](https://img.shields.io/badge/license-Apache-blue.svg)](LICENSE)

base32flex is a Go package which implements base32 Encoding.

Standard base32 could become more readable if it didn't contain 'I', 'l' (these are often confused with '1') and 'O', 'o' (these are often confused with '0').

So this libray encodes excluding these letters.

`LowerEncoding` excludes 'l' and 'o' for use of lower case.

`UpperEncoding` excludes 'I' and 'O' for use of upper case.

Both encodings include '8' and '9' that are excluded by Standard base32 encoding (but these could not be confused by any letters).

## GoDoc

https://godoc.org/github.com/oirik/base32flex
