package resourceviewer

import (
	"context"

	"github.com/vmware-tanzu/octant/pkg/view/component"
)

//go:generate mockgen -destination=./fake/mock_details.go -package=fake github.com/vmware-tanzu/octant/internal/resourceviewer Details

// Details generates details for resource viewer.
type Details interface {
	AdjacencyList() (*component.AdjList, error)
	Nodes(ctx context.Context) (component.Nodes, error)
}

// GenerateComponent generates a resource viewer component given details.
func GenerateComponent(ctx context.Context, details Details, selected string) (*component.ResourceViewer, error) {
	rv := component.NewResourceViewer("Resource Viewer")
	rv.SetAccessor("resourceViewer")

	nodes, err := details.Nodes(ctx)
	if err != nil {
		return nil, err
	}

	for id, node := range nodes {
		rv.AddNode(id, node)
	}

	edges, err := details.AdjacencyList()
	if err != nil {
		return nil, err
	}

	for k, list := range *edges {
		for i := range list {
			item := list[i]
			if err := rv.AddEdge(k, item.Node, item.Type); err != nil {
				return nil, err
			}
		}
	}

	rv.Select(selected)

	return rv, nil
}
