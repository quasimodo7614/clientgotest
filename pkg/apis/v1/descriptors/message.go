package descriptors

import (
	"github.com/quasimodo7614/clientgotest/pkg/deployment"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/deployments",
		Definitions: []def.Definition{listDeployments, createDeployment},
		Children: []def.Descriptor{
			{
				Path:        "{deployment}",
				Definitions: []def.Definition{getDeployment, updateDeployment, deleteDeployment},
			},
		},
	},
	}...)
}

var listDeployments = def.Definition{
	Method:      def.List,
	Summary:     "List Deployments",
	Description: "Query a specified number of deployments and returns an array",
	Function:    deployment.ListDeployments,
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

var getDeployment = def.Definition{
	Method:      def.Get,
	Summary:     "Get Message",
	Description: "Get a deployment by id",
	Function:    deployment.GetDeployment,
	Parameters: []def.Parameter{
		def.PathParameterFor("deployment", "Message id"),
	},
	Results: def.DataErrorResults("A deployment"),
}

var deleteDeployment = def.Definition{
	Method:      def.Delete,
	Summary:     "Delete Message",
	Description: "Delete a deployment by id",
	Function:    deployment.DeleteDeployment,
	Parameters: []def.Parameter{
		def.PathParameterFor("deployment", "Message id"),
	},
	Results: []def.Result{def.ErrorResult()},
}

var createDeployment = def.Definition{
	Method:      def.Create,
	Summary:     "Create Message",
	Description: "Create a deployment ",
	Function:    deployment.CreateDeployment,
	Parameters: []def.Parameter{
		def.BodyParameterFor("a deployment"),
	},
	Results: def.DataErrorResults("A deployment"),
}

var updateDeployment = def.Definition{
	Method:      def.Update,
	Summary:     "Update Message",
	Description: "Update a deployment ",
	Function:    deployment.UpdateDeployment,
	Parameters: []def.Parameter{
		def.PathParameterFor("deployment", "Message id"),
	},
	Results: def.DataErrorResults("A deployment"),
}
