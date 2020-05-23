package nodes

import (
	"context"
	"fmt"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"
	"github.com/AnkushJadhav/kamaji-root/pkg/utils"
)

// RegisterNode updates a user with their node information post successful sign in at node
func RegisterNode(ctx context.Context, store store.Driver, uID, name, version, hostOS, hostDockerversion string) error {
	user, err := store.GetUserByID(ctx, uID)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user does not exist")
	}

	hostinfo := models.HostData{
		OS:            hostOS,
		DockerVersion: hostDockerversion,
	}
	node := models.Node{
		Name:     name,
		Version:  version,
		HostData: hostinfo,
	}
	node.ID = utils.GenerateUUID()
	node.TS = time.Now()

	_, err = store.AddNodeToUser(ctx, user.ID, node)
	if err != nil {
		return err
	}
	return nil
}
