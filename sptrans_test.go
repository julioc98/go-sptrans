package sptrans

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	token = "tokendetestesptransolhovivojc"
)

func mockingServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Path:", r.URL.Path)
		switch r.URL.Path {

		case "/Login/Autenticar":
			rToken := r.URL.Query().Get("token")
			if rToken == token {
				fmt.Fprintf(w, "true")
			}
		}
	}))
}

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

func TestSPTrans_Auth(t *testing.T) {
	type fields struct {
		BasePath string
		Token    string
	}

	mockBasePath := mockingServer().URL

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Autenticando corretamente na API da SPTrans",
			fields: fields{
				BasePath: mockBasePath,
				Token:    token,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := &SPTrans{
				BasePath: tt.fields.BasePath,
				Token:    tt.fields.Token,
			}
			if got := sp.Auth(); got != tt.want {
				t.Errorf("SPTrans.Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
