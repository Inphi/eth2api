package beaconapi

import (
	"context"

	"github.com/protolambda/eth2api"
	"github.com/protolambda/zrnt/eth2/beacon/altair"
	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/zrnt/eth2/beacon/phase0"
)

// Retrieves attestations known by the node but not necessarily incorporated into any block
func PoolAttestations(ctx context.Context, cli eth2api.Client, slot *common.Slot, committeeIndex *common.CommitteeIndex, dest *[]phase0.Attestation) error {
	var q eth2api.Query
	if slot != nil {
		if committeeIndex != nil {
			q = eth2api.Query{"slot": *slot, "committee_index": *committeeIndex}
		} else {
			q = eth2api.Query{"slot": *slot}
		}
	} else if committeeIndex != nil {
		q = eth2api.Query{"committee_index": *committeeIndex}
	}
	return eth2api.MinimalRequest(ctx, cli, eth2api.QueryGET(q, "/eth/v1/beacon/pool/attestations"), eth2api.Wrap(dest))
}

// Submits Attestation objects to the node.  Each attestation in the request body is processed individually.
//
// If an attestation is validated successfully the node MUST publish that attestation on the appropriate subnet.
//
// If one or more attestations fail validation the node MUST return a 400 error with details of which attestations have failed, and why.
// In that case, a non-nil list of errors will be returned, with entries pointing to original array indices of input attestations
func SubmitAttestations(ctx context.Context, cli eth2api.Client, attestations []phase0.Attestation) (failures []eth2api.IndexedErrorMessageItem, err error) {
	resp := cli.Request(ctx, eth2api.BodyPOST("/eth/v1/beacon/pool/attestations", attestations))
	_, err = resp.Decode(nil)
	if err != nil {
		if ierr, ok := err.(eth2api.IndexedError); ok {
			return ierr.IndexedErrors(), err
		}
		return nil, err
	}
	return nil, nil
}

// Retrieves attester slashings known by the node but not necessarily incorporated into any block.
func PoolAttesterSlashings(ctx context.Context, cli eth2api.Client, dest *[]phase0.AttesterSlashing) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.PlainGET("/eth/v1/beacon/pool/attester_slashings"), eth2api.Wrap(dest))
}

// Submits AttesterSlashing object to node's pool and if passes validation node MUST broadcast it to network.
func SubmitAttesterSlashing(ctx context.Context, cli eth2api.Client, attSlashing *phase0.AttesterSlashing) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.BodyPOST("/eth/v1/beacon/pool/attester_slashings", attSlashing), nil)
}

// Retrieves proposer slashings known by the node but not necessarily incorporated into any block.
func PoolProposerSlashings(ctx context.Context, cli eth2api.Client, dest *[]phase0.ProposerSlashing) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.PlainGET("/eth/v1/beacon/pool/proposer_slashings"), eth2api.Wrap(dest))
}

// Submits ProposerSlashing object to node's pool and if passes validation node MUST broadcast it to network.
func SubmitProposerSlashing(ctx context.Context, cli eth2api.Client, propSlashing *phase0.ProposerSlashing) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.BodyPOST("/eth/v1/beacon/pool/proposer_slashings", propSlashing), nil)
}

// Retrieves voluntary exits known by the node but not necessarily incorporated into any block.
func PoolVoluntaryExits(ctx context.Context, cli eth2api.Client, dest *[]phase0.SignedVoluntaryExit) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.PlainGET("/eth/v1/beacon/pool/voluntary_exits"), eth2api.Wrap(dest))
}

// Submits SignedVoluntaryExit object to node's pool and if passes validation node MUST broadcast it to network.
func SubmitVoluntaryExit(ctx context.Context, cli eth2api.Client, exit *phase0.SignedVoluntaryExit) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.BodyPOST("/eth/v1/beacon/pool/voluntary_exits", exit), nil)
}

// TODO: the API does not have a method to view the pool of sync-committee messages

// Submits sync committee signature objects to the node.
//
// If a sync committee signature is validated successfully the node MUST publish that sync committee signature on all applicable subnets.
//
// If one or more sync committee signatures fail validation the node MUST return a 400 error with details of which sync committee signatures have failed, and why.
func SubmitSyncCommitteeMessages(ctx context.Context, cli eth2api.Client, messages []altair.SyncCommitteeMessage) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.BodyPOST("/eth/v1/beacon/pool/sync_committees", messages), nil)
}
