package types

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"
	"github.com/qredo/fusionchain/testutil/sample"
)

func TestMsgUpdateKeyRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKeyRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKeyRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKeyRequest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
