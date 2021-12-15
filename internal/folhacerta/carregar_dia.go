package folhacerta

type CarregarDiaResponse struct {
	BaseResponse
	MostrarBarraCabecalho    bool   `json:"MostrarBarraCabecalho"`
	Dia                      Dia    `json:"Dia"`
	Mes                      Mes    `json:"Mes"`
	HoraServidor             string `json:"HoraServidor"`
	URLFotoPontoIndisponivel string `json:"UrlFotoPontoIndisponivel"`
}

type Dia struct {
	Data                        string            `json:"Data"`
	Ano                         int64             `json:"Ano"`
	Mes                         int64             `json:"Mes"`
	Dia                         int64             `json:"Dia"`
	Titulo                      string            `json:"Titulo"`
	Descricao                   string            `json:"Descricao"`
	HorariosMarcacoes           []HorarioMarcacao `json:"HorariosMarcacoes"`
	HorarioMarcado              Horario           `json:"HorarioMarcado"`
	HorarioJornada              Horario           `json:"HorarioJornada"`
	AusenciaProgramada          bool              `json:"AusenciaProgramada"`
	MotivoAusenciaProgramada    string            `json:"MotivoAusenciaProgramada"`
	DescricaoAusenciaProgramada interface{}       `json:"DescricaoAusenciaProgramada"`
	AusenciaProgramadaBloqueada bool              `json:"AusenciaProgramadaBloqueada"`
	AbonoHoras                  bool              `json:"AbonoHoras"`
	MotivoAbonoHoras            string            `json:"MotivoAbonoHoras"`
	MarcacaoPendente            bool              `json:"MarcacaoPendente"`
	Fechado                     bool              `json:"Fechado"`
	Resumo                      Resumo            `json:"Resumo"`
	RegistrosLog                []interface{}     `json:"RegistrosLog"`
	ObservacaoIntervalo         string            `json:"ObservacaoIntervalo"`
	ConfirmarMarcacao           bool              `json:"ConfirmarMarcacao"`
	MensagemConfirmarMarcacao   string            `json:"MensagemConfirmarMarcacao"`
	TipoJornada                 string            `json:"TipoJornada"`
	MarcacoesFaltando           int64             `json:"MarcacoesFaltando"`
}

type Horario struct {
	HoraEntrada               *string       `json:"HoraEntrada"`
	DiaEntrada                *string       `json:"DiaEntrada"`
	ObservacaoHoraEntrada     interface{}   `json:"ObservacaoHoraEntrada"`
	MarcacoesEntrada          []interface{} `json:"MarcacoesEntrada"`
	HoraSaidaAlmoco           *string       `json:"HoraSaidaAlmoco"`
	DiaSaidaAlmoco            *string       `json:"DiaSaidaAlmoco"`
	ObservacaoHoraSaidaAlmoco interface{}   `json:"ObservacaoHoraSaidaAlmoco"`
	MarcacoesSaidaAlmoco      []interface{} `json:"MarcacoesSaidaAlmoco"`
	HoraVoltaAlmoco           *string       `json:"HoraVoltaAlmoco"`
	DiaVoltaAlmoco            *string       `json:"DiaVoltaAlmoco"`
	ObservacaoHoraVoltaAlmoco interface{}   `json:"ObservacaoHoraVoltaAlmoco"`
	MarcacoesVoltaAlmoco      []interface{} `json:"MarcacoesVoltaAlmoco"`
	HoraSaida                 *string       `json:"HoraSaida"`
	DiaSaida                  *string       `json:"DiaSaida"`
	ObservacaoHoraSaida       interface{}   `json:"ObservacaoHoraSaida"`
	MarcacoesSaida            []interface{} `json:"MarcacoesSaida"`
}

type HorarioMarcacao struct {
	ID               int64       `json:"Id"`
	Tipo             int64       `json:"Tipo"`
	DescricaoTipo    string      `json:"DescricaoTipo"`
	Subtipo          int64       `json:"Subtipo"`
	DescricaoSubtipo interface{} `json:"DescricaoSubtipo"`
	Hora             string      `json:"Hora"`
	Dia              *string     `json:"Dia"`
	Observacao       interface{} `json:"Observacao"`
	Foto             interface{} `json:"Foto"`
	Latitude         interface{} `json:"Latitude"`
	Longitude        interface{} `json:"Longitude"`
	AppID            *string     `json:"AppId"`
	AppDiaID         *string     `json:"AppDiaId"`
	Marcacoes        []Marcacao  `json:"Marcacoes"`
}

type Marcacao struct {
	ID                   int64       `json:"Id"`
	Tipo                 int64       `json:"Tipo"`
	DescricaoTipo        string      `json:"DescricaoTipo"`
	DataHora             string      `json:"DataHora"`
	DescricaoDataHora    string      `json:"DescricaoDataHora"`
	Data                 string      `json:"Data"`
	DescricaoData        string      `json:"DescricaoData"`
	Hora                 string      `json:"Hora"`
	DataCadastro         string      `json:"DataCadastro"`
	DescricaoDataCadasto string      `json:"DescricaoDataCadasto"`
	Motivo               *int64      `json:"Motivo"`
	DescricaoMotivo      string      `json:"DescricaoMotivo"`
	NomeUsuarioGestor    string      `json:"NomeUsuarioGestor"`
	Status               *int64      `json:"Status"`
	DescricaoStatus      string      `json:"DescricaoStatus"`
	MarcacaoOffLine      interface{} `json:"MarcacaoOffLine"`
}

type Resumo struct {
	Entrada                                 bool        `json:"Entrada"`
	SaidaAlmoco                             bool        `json:"SaidaAlmoco"`
	VoltaAlmoco                             bool        `json:"VoltaAlmoco"`
	Saida                                   bool        `json:"Saida"`
	RetornoNoturno                          bool        `json:"RetornoNoturno"`
	FotoEntrada                             interface{} `json:"FotoEntrada"`
	FotoSaidaAlmoco                         interface{} `json:"FotoSaidaAlmoco"`
	FotoVoltaAlmoco                         interface{} `json:"FotoVoltaAlmoco"`
	FotoSaida                               interface{} `json:"FotoSaida"`
	LatitudeEntrada                         interface{} `json:"LatitudeEntrada"`
	LongitudeEntrada                        interface{} `json:"LongitudeEntrada"`
	LatitudeSaidaAlmoco                     interface{} `json:"LatitudeSaidaAlmoco"`
	LongitudeSaidaAlmoco                    interface{} `json:"LongitudeSaidaAlmoco"`
	LatitudeVoltaAlmoco                     interface{} `json:"LatitudeVoltaAlmoco"`
	LongitudeVoltaAlmoco                    interface{} `json:"LongitudeVoltaAlmoco"`
	LatitudeSaida                           interface{} `json:"LatitudeSaida"`
	LongitudeSaida                          interface{} `json:"LongitudeSaida"`
	HorasNormais                            string      `json:"HorasNormais"`
	HorasAtraso                             string      `json:"HorasAtraso"`
	HorasAtrasoEntrada                      string      `json:"HorasAtrasoEntrada"`
	HorasEntradaAntecipada                  string      `json:"HorasEntradaAntecipada"`
	HorasSaidaAntecipada                    string      `json:"HorasSaidaAntecipada"`
	HorasAtrasoSaida                        string      `json:"HorasAtrasoSaida"`
	HorasExtras                             string      `json:"HorasExtras"`
	HorasFaltasAtrasos                      string      `json:"HorasFaltasAtrasos"`
	HorasEsperadas                          string      `json:"HorasEsperadas"`
	HorasTrabalhadas                        string      `json:"HorasTrabalhadas"`
	HorasExtrasCompensaveis                 string      `json:"HorasExtrasCompensaveis"`
	HorasExtrasDireto                       string      `json:"HorasExtrasDireto"`
	HorasAcumuladasSaldo                    string      `json:"HorasAcumuladasSaldo"`
	HorasIntervaloEsperadas                 string      `json:"HorasIntervaloEsperadas"`
	HorasIntervalo                          string      `json:"HorasIntervalo"`
	HorasPrimeiroPeriodo                    string      `json:"HorasPrimeiroPeriodo"`
	HorasSegundoPeriodo                     string      `json:"HorasSegundoPeriodo"`
	HorasAbonadas                           string      `json:"HorasAbonadas"`
	HorasDescansoObrigatorio                string      `json:"HorasDescansoObrigatorio"`
	ParadasDescansoObrigatorioNaoRealizadas int64       `json:"ParadasDescansoObrigatorioNaoRealizadas"`
	HorasIntervaloReduzido                  string      `json:"HorasIntervaloReduzido"`
	HorasToleranciaPositiva                 string      `json:"HorasToleranciaPositiva"`
	HorasToleranciaNegativa                 string      `json:"HorasToleranciaNegativa"`
}

type Mes struct {
	ResumoMesID             int64  `json:"ResumoMes_Id"`
	DiasEsperados           int64  `json:"DiasEsperados"`
	DiasRealizados          int64  `json:"DiasRealizados"`
	HorasEsperadas          string `json:"HorasEsperadas"`
	HorasTrabalhadas        string `json:"HorasTrabalhadas"`
	HorasNormais            string `json:"HorasNormais"`
	HorasExtras             string `json:"HorasExtras"`
	HorasExtrasDireto       string `json:"HorasExtrasDireto"`
	HorasAtrasosFaltas      string `json:"HorasAtrasosFaltas"`
	HorasAcumuladasSaldo    string `json:"HorasAcumuladasSaldo"`
	HorasAbonadas           string `json:"HorasAbonadas"`
	SaldoBancoHoras         string `json:"SaldoBancoHoras"`
	SaldoBancoHorasAnterior string `json:"SaldoBancoHorasAnterior"`
	CicloBancoHoras         string `json:"CicloBancoHoras"`
	StatusPonto             int64  `json:"StatusPonto"`
	UsuarioStatus           string `json:"UsuarioStatus"`
	DescricaoStatusPonto    string `json:"DescricaoStatusPonto"`
	DataStatus              string `json:"DataStatus"`
	HoraStatus              string `json:"HoraStatus"`
	PontosPendentes         int64  `json:"PontosPendentes"`
}
