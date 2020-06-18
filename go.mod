// go.mod example
//
// Refer to https://github.com/golang/go/wiki/Modules#gomod
// for detailed go.mod and go mod command documentation.
//
// module github.com/my/module/v3
//
// require (
//     github.com/some/dependency v1.2.3
//     github.com/another/dependency v0.1.0
//     github.com/additional/dependency/v4 v4.0.0
// )

module github.com/quasimodo7614/clientgotest

require (
	github.com/caicloud/nirvana v0.2.8
	github.com/imdario/mergo v0.3.9 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect

)

go 1.13
