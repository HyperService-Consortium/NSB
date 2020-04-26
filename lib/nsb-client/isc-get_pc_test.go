package nsbcli

import (
	"encoding/base64"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

func TestNSBClient_ISCGetPC(t *testing.T) {
	type args struct {
		user            uip.Signer
		contractAddress []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{"easy", getNormalField(),
			args{
				user: signer,
				contractAddress: sugar.HandlerError(
					base64.StdEncoding.DecodeString("dSkSBKJzv87q+UnuOh6Po3ol46FDTSKIioM01uzI0cY=")).([]byte),
			}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.ISCGetPC(tt.args.user, tt.args.contractAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("ISCGetPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ISCGetPC() got = %v, want %v", got, tt.want)
			}
		})
	}
}
