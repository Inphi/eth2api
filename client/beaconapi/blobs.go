package beaconapi

import (
	"context"

	"github.com/protolambda/eth2api"
)

func Blobs(ctx context.Context, cli eth2api.Client, blockId eth2api.BlockId, dest *eth2api.BlobsSidecar) (exists bool, err error) {
	return eth2api.SimpleRequest(ctx, cli, eth2api.FmtGET("/eth/v1/beacon/blobs_sidecar/%s", blockId.BlockId()), dest)
}
