package nsbcli

//func TestNSBClient_AddAction(t *testing.T) {
//	type fields struct {
//		handler    *request.RequestClient
//		bufferPool *bytespool.BytesPool
//	}
//	type args struct {
//		user       types.Signer
//		toAddress  []byte
//		iscAddress []byte
//		tid        uint64
//		aid        uint64
//		stype      uint32
//		content    []byte
//		signature  []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		{name:"test_easy", fields: getNormalField(), args: args{
//			signer,
//			nil,
//
//		}}
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			nc := &NSBClient{
//				handler:    tt.fields.handler,
//				bufferPool: tt.fields.bufferPool,
//			}
//			got, err := nc.AddAction(tt.args.user, tt.args.toAddress, tt.args.iscAddress, tt.args.tid, tt.args.aid, tt.args.stype, tt.args.content, tt.args.signature)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("AddAction() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("AddAction() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
