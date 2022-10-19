package spec

//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
//go:generate oapi-codegen -old-config-style -package spec -generate types,skip-prune -o types.gen.go  ../../../../../api/internal_admin/openapi.yaml
//go:generate oapi-codegen -old-config-style -package spec -generate server           -o server.gen.go ../../../../../api/internal_admin/openapi.yaml
