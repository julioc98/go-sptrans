package sptrans

import (
	"reflect"
	"testing"
)

const (
	token = "tokendetestesptransolhovivojc"
)

func TestNew(t *testing.T) {

	type args struct {
		token string
	}

	tests := []struct {
		name   string
		args   args
		wantSp *SPTrans
	}{
		{
			name: "Criando nova estrutura preenchida da SPTrans",
			args: args{
				token: token,
			},
			wantSp: &SPTrans{
				BasePath: "http://api.olhovivo.sptrans.com.br/v2.1/",
				Token:    token,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSp := New(tt.args.token); !reflect.DeepEqual(gotSp, tt.wantSp) {
				t.Errorf("New() = %v, want %v", gotSp, tt.wantSp)
			}
		})
	}
}
