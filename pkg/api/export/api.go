package export

import (
	ml "github.com/mycontroller-org/backend/v2/pkg/model"
	nml "github.com/mycontroller-org/backend/v2/pkg/model/node"
	pml "github.com/mycontroller-org/backend/v2/pkg/model/pagination"
	stgml "github.com/mycontroller-org/backend/v2/pkg/model/storage"
	svc "github.com/mycontroller-org/backend/v2/pkg/service"
	ut "github.com/mycontroller-org/backend/v2/pkg/util"
)

// List by filter and pagination
func List(f []pml.Filter, p *pml.Pagination) (*pml.Result, error) {
	out := make([]nml.Node, 0)
	return svc.STG.Find(ml.EntityNode, &out, f, p)
}

// Get returns a Node
func Get(f []pml.Filter) (nml.Node, error) {
	out := nml.Node{}
	err := svc.STG.FindOne(ml.EntityNode, &out, f)
	return out, err
}

// Save Node config into disk
func Save(node *nml.Node) error {
	if node.ID == "" {
		node.ID = ut.RandUUID()
	}
	f := []pml.Filter{
		{Key: ml.KeyID, Value: node.ID},
	}
	return svc.STG.Upsert(ml.EntityNode, node, f)
}

// GetByIDs returns a node details by gatewayID and nodeId of a message
func GetByIDs(gatewayID, nodeID string) (*nml.Node, error) {
	f := []pml.Filter{
		{Key: ml.KeyGatewayID, Value: gatewayID},
		{Key: ml.KeyNodeID, Value: nodeID},
	}
	out := &nml.Node{}
	err := svc.STG.FindOne(ml.EntityNode, out, f)
	return out, err
}

// Delete node
func Delete(IDs []string) (int64, error) {
	f := []pml.Filter{{Key: ml.KeyID, Operator: stgml.OperatorIn, Value: IDs}}
	return svc.STG.Delete(ml.EntityNode, f)
}