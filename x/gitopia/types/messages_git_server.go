package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAuthorizeGitServer = "authorize_git_server"
)

var _ sdk.Msg = &MsgAuthorizeGitServer{}

func NewMsgAuthorizeGitServer(creator string, provider string) *MsgAuthorizeGitServer {
	return &MsgAuthorizeGitServer{
		Creator:  creator,
		Provider: provider,
	}
}

func (msg *MsgAuthorizeGitServer) Route() string {
	return RouterKey
}

func (msg *MsgAuthorizeGitServer) Type() string {
	return TypeMsgAuthorizeGitServer
}

func (msg *MsgAuthorizeGitServer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAuthorizeGitServer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuthorizeGitServer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	return nil
}