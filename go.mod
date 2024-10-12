module github.com/pasataleo/go-collections

go 1.23

toolchain go1.23.2

replace github.com/pasataleo/go-objects => ../go-objects
replace github.com/pasataleo/go-errors => ../go-errors
replace github.com/pasataleo/go-testing => ../go-testing

require (
	github.com/pasataleo/go-errors v0.1.2
	github.com/pasataleo/go-objects v0.1.3
	github.com/pasataleo/go-testing v0.1.4
)

require github.com/google/go-cmp v0.5.9 // indirect
