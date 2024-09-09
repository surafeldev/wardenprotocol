package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/warden-protocol/wardenprotocol/warden/x/act/types/v1beta1"
)

func (k msgServer) VoteForAction(goCtx context.Context, msg *types.MsgVoteForAction) (*types.MsgVoteForActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	act, err := k.ActionKeeper.Get(ctx, msg.ActionId)

	if err != nil {
		return nil, err
	}

	if act.TimeoutHeight > 0 && act.TimeoutHeight < uint64(ctx.BlockHeight()) {
		if err := act.SetStatus(ctx, types.ActionStatus_ACTION_STATUS_TIMEOUT); err != nil {
			return nil, err
		}
		if err := k.ActionKeeper.Set(ctx, act); err != nil {
			return nil, err
		}

		return &types.MsgVoteForActionResponse{
			Status: act.Status.String(),
		}, nil
	}

	if err := act.AddOrUpdateVote(ctx, msg.Participant, msg.VoteType); err != nil {
		return nil, err
	}

	if err := k.ActionKeeper.Set(ctx, act); err != nil {
		return nil, err
	}

	if msg.VoteType == types.ActionVoteType_VOTE_TYPE_APPROVED {
		if err := k.TryExecuteVotedAction(ctx, &act); err != nil {
			return nil, err
		}
	} else if msg.VoteType == types.ActionVoteType_VOTE_TYPE_REJECTED {
		if err := k.TryRejectVotedAction(ctx, &act); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unhandled VoteType value: %v", msg.VoteType)
	}

	return &types.MsgVoteForActionResponse{Status: act.Status.String()}, nil
}