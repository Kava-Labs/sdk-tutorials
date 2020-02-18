package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// custom error handler for names that do not exist
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)