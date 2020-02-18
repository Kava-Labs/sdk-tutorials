package types

import sdk "github.com/cosmos/cosmos-sdk/types"


// the whois struct contains all the metadata for a name
type Whois struct{
	Value string `json:"value"`
	Owner sdk.AccAddress `json:"ownder"`
	Price sdk.Coins `json:"price"`
}