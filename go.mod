module github.com/chrisguida/go-cln-plugin

go 1.20

require github.com/elementsproject/glightning v0.0.0-20221013194807-73978c84cee8

require (
	github.com/chrisguida/go-cln-plugin/util v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.2 // indirect
)

replace github.com/elementsproject/glightning => github.com/chrisguida/glightning v0.0.0-20230418225814-59d9185171f7

replace github.com/chrisguida/go-cln-plugin/util => ./util
