package keeper

import (
	"github.com/RaspiRepo/leaderboard/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
