// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package logger

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

var _ = (*callFrameMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c callFrame) MarshalJSON() ([]byte, error) {
	type callFrame struct {
		From  common.Address      `json:"from"`
		To    common.Address      `json:"to"`
		Input hexutil.Bytes       `json:"input,omitempty"`
		Gas   math.HexOrDecimal64 `json:"gas"`
		Value *hexutil.Big        `json:"value"`
		Type  string              `json:"type"`
	}
	var enc callFrame
	enc.From = c.From
	enc.To = c.To
	enc.Input = c.Input
	enc.Gas = math.HexOrDecimal64(c.Gas)
	enc.Value = (*hexutil.Big)(c.Value)
	enc.Type = c.Type()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *callFrame) UnmarshalJSON(input []byte) error {
	type callFrame struct {
		From  *common.Address      `json:"from"`
		To    *common.Address      `json:"to"`
		Input *hexutil.Bytes       `json:"input,omitempty"`
		Gas   *math.HexOrDecimal64 `json:"gas"`
		Value *hexutil.Big         `json:"value"`
	}
	var dec callFrame
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.To != nil {
		c.To = *dec.To
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.Gas != nil {
		c.Gas = uint64(*dec.Gas)
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	return nil
}
