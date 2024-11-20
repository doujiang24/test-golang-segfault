package main

import (
	"fmt"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
	"google.golang.org/protobuf/types/known/anypb"
)

const Name = "random"

func init() {
	// debug.SetGCPercent(-1)
	// debug.SetMemoryLimit(math.MaxInt64)
	http.RegisterHttpFilterFactoryAndConfigParser(Name, filterFactory, &parser{})
}


type parser struct {
}


// Parse the filter configuration. We can call the ConfigCallbackHandler to control the filter's
// behavior
func (p *parser) Parse(any *anypb.Any, callbacks api.ConfigCallbackHandler) (interface{}, error) {
	fmt.Println("Start Plugin")
	return any, nil
}

// Merge configuration from the inherited parent configuration
func (p *parser) Merge(parent interface{}, child interface{}) interface{} {
	return parent
}

func filterFactory(c interface{}, callbacks api.FilterCallbackHandler) api.StreamFilter {
	return &filter{
		callbacks: callbacks,
	}
}
