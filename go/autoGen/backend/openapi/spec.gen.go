// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RSwW7UMBD9lerB0WqWFnHwDThVQsClp9UeTDKbDkps156AVpH/HY2zod3VCoE49RJb",
	"npm8N++9GW0YY/DkJcPOyO0Dja5eP9PP+0xJrzGFSEmYasH1pIccIsHCT+M3SigG3o3PC1kS+x6lGCR6",
	"nDhRB7tdukz9ya4YTEcINwxf9rDbGa8T7WHxqnli1hxpNSunYs5JcafffUijE1iwl3dvYVYu7IV6nTsj",
	"wx12ZVf0mf0+VPYsg44oztX7r3cw+EEpc/CweHO9ud7oriGSd5FhcVufDKKTh8qk0ZXqrSfRQ2k64eDv",
	"Olh84iz3tUNnkhtJavt2Rke5TRxlwRp4ZIESg8XjROmAVePftaddJE1kjvadK3F7c1GJc8ConlzGO5b+",
	"B26n4zkGnxfDbjYbPdrghXzVycU4cFuVar5npTQ/Q2ChsQ7+KR/TEo4V3aXkDqj2nq6qfVcrnZoKcb2a",
	"gMUaTWYM+YJ9HxM5odVAFYSyfAjd4Z+W+auMn4ZVFS8vS8NSyq8AAAD//4fv3JBgBAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}