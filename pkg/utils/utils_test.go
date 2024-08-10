package utils

// import (
// 	"context"
// 	"testing"
// )

// func TestGetUserIDFromContext(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    int
// 		wantErr bool
// 	}{
// 		{name: "Successfully got userID", args:  , want: 7, wantErr: false}
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetUserIDFromContext(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUserIDFromContext() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("GetUserIDFromContext() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
