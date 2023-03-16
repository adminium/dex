module github.com/adminium/dex/examples

go 1.17

require (
	github.com/adminium/dex/api/v2 v2.0.0
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/spf13/cobra v1.3.0
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783
	google.golang.org/grpc v1.52.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace github.com/adminium/dex/api/v2 => ../api/v2
