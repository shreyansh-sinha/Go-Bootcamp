module example.com/Story1

go 1.22.3

require (
	example.com/common/utils v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace example.com/common/utils => ../utils
