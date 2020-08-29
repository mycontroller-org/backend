package node

import (
	ml "github.com/mycontroller-org/backend/v2/pkg/model"
	nml "github.com/mycontroller-org/backend/v2/pkg/model/node"
	pml "github.com/mycontroller-org/backend/v2/pkg/model/pagination"
	svc "github.com/mycontroller-org/backend/v2/pkg/service"
	ut "github.com/mycontroller-org/backend/v2/pkg/util"
)

// ListNodes by filter and pagination
func ListNodes(f []pml.Filter, p pml.Pagination) ([]nml.Node, error) {
	out := make([]nml.Node, 0)
	svc.STG.Find(ml.EntityNode, f, p, &out)
	return out, nil
}

// GetNode returns a Node
func GetNode(f []pml.Filter) (nml.Node, error) {
	out := nml.Node{}
	err := svc.STG.FindOne(ml.EntityNode, f, &out)
	return out, err
}

// Save Node config into disk
func Save(g *nml.Node) error {
	if g.ID == "" {
		g.ID = ut.RandID()
	}
	return svc.STG.Upsert(ml.EntityNode, nil, g)
}