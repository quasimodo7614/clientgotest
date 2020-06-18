package descriptors

import (
	"github.com/quasimodo7614/clientgotest/pkg/deployment"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/messages",
		Definitions: []def.Definition{listMessages},
	}, {
		Path:        "/messages/{deployment}",
		Definitions: []def.Definition{getMessage},
	},
	}...)
}

var listMessages = def.Definition{
	Method:      def.List,
	Summary:     "List Messages",
	Description: "Query a specified number of messages and returns an array",
	Function:    deployment.ListMessages,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "count",
			Default:     10,
			Description: "Number of messages",
		},
	},
	Results: def.DataErrorResults("A list of messages"),
}

var getMessage = def.Definition{
	Method:      def.Get,
	Summary:     "Get Message",
	Description: "Get a deployment by id",
	Function:    deployment.GetMessage,
	Parameters: []def.Parameter{
		def.PathParameterFor("deployment", "Message id"),
	},
	Results: def.DataErrorResults("A deployment"),
}
