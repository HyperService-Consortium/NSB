package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
	// "encoding/hex"
	"encoding/json"
)


func (nsb *NSBApplication) QueryIndex(req *types.RequestQuery) string {
	switch req.Path {
	case "acc_getAccInfo":
		return nsb.getAccInfo(req.Data, req.Height)
	default:
		return "unknown query type"
	}
}

func (nsb *NSBApplication) getAccInfo(paras []byte, height int64) string {
	// assuming height == 0 // latest
	bytesInfo, err := nsb.accMap.TryGet(paras)
	if err != nil || bytesInfo == nil {
		return "the account is not on this AccTrie"
	}
	var accInfo AccountInfo
	err = json.Unmarshal(bytesInfo, accInfo)
	if err != nil {
		return err.Error()
	}
	return accInfo.String()
}