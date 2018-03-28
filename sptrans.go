package sptrans

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
