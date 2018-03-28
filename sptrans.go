package sptrans

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// SPTrans é a pricipal estrutura onde se encontram as pricipar informaçoes para usar este pacote
type SPTrans struct {
	BasePath string
	Token    string
}

// New inicializa o componente principal SPTrans
func New(token string) (sp *SPTrans) {
	sp = &SPTrans{
		BasePath: "http://api.olhovivo.sptrans.com.br/v2.1/",
		Token:    token,
	}
	return
}

// Auth autentica na API da SPTrans Olho Vivo
func (sp *SPTrans) Auth() (bool, error) {

	url := sp.BasePath + "/Login/Autenticar?token=" + sp.Token

	resp, err := http.Post(url, "text/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if strings.Trim(string(body), "\r\n") == "true" {
		return true, nil
	}
	return false, nil
}
