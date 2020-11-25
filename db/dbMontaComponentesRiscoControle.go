package db

import ()

func initRiscoDeCredito() ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementos []ElementoESI
	var elemento ElementoESI

	var itens = []string{"Ativos com elevado risco e retorno.",
		"Qualidade da carteira de crédito",
		"Incapacidade/dificuldade do Patrocinador em honrar seus compromissos."}
	elemento.Nome = "Risco de Inadimplência"
	elemento.Itens = itens
	elementos = append(elementos, elemento)

	itens = []string{"Tolerância ao risco", "Concentração da Carteira"}
	elemento.Nome = "Risco de Concentração"
	elemento.Itens = itens
	elementos = append(elementos, elemento)

	itens = []string{"Critérios para diminuição de risco",
		"Eficácia dos mitigadores de risco", "Agência de Classificação de Risco de Crédito"}
	elemento.Nome = "Risco de Ineficácia das Garantias"
	elemento.Itens = itens
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Risco"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	// CONTROLE
	elementos = make([]ElementoESI, 0)
	itens = []string{"Estrutura da área de gestão do risco de crédito",
		"Políticas e estratégias",
		"Estrutura de responsabilidades e alçadas",
		"Comunicação"}
	elemento.Nome = "Ambiente de Controle"
	elemento.Itens = itens
	elementos = append(elementos, elemento)

	itens = []string{"Modelos de mensuração e classificação de risco",
		"Gestão do risco de inadimplência",
		"Gestão do risco de concentração",
		"Gestão de mitigadores de risco"}

	elemento.Itens = itens
	elemento.Nome = "Identificação, Avaliação e Mensuração"
	elementos = append(elementos, elemento)

	itens = []string{"Monitoramento do risco de crédito", "Revisão das estratégicas, políticas e limites"}
	elemento.Itens = itens
	elemento.Nome = "Monitoramento"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Controle"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = "Risco de Crédito"
	componente.TiposNotas = tiposNotas

	return componente
}

func initRiscoDeMercado() ComponenteESI {

	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementos []ElementoESI

	var itens = []string{"Perfil do Plano de Benefícios",
		"Taxas de juros domésticas", "Limites"}
	var elemento ElementoESI
	elemento.Itens = itens
	elemento.Nome = "Apetite para risco de mercado"
	elementos = append(elementos, elemento)

	itens = []string{"Risco de mercado na carteira", "Utilização de derivativos e de hedge"}
	elemento.Itens = itens
	elemento.Nome = "Nível de risco de mercado na carteira"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Risco"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	// CONTROLE
	elementos = make([]ElementoESI, 0)
	itens = []string{"Estratégias e políticas", "Órgãos Estatutários e demais órgãos de controle", "Manuais"}
	elemento.Itens = itens
	elemento.Nome = "Ambiente de Controle"
	elementos = append(elementos, elemento)

	itens = []string{"Sistemas de mensuração internos ou contrato de prestadores de serviços",
		"Marcação a mercado das posições",
		"Exposição dos investimentos - análise de volatilidade versus retorno",
		"Mensuração do risco de mercado e da exposição líquida",
		"Relatórios"}
	elemento.Itens = itens
	elemento.Nome = "Identificação, avaliação, mensuração e comunicação do risco de mercado da carteira"
	elementos = append(elementos, elemento)

	itens = []string{"Estrutura de limites",
		"Testes de estresse",
		"Função do controle de risco de mercado"}
	elemento.Itens = itens
	elemento.Nome = "Controles específicos"
	elementos = append(elementos, elemento)

	itens = []string{"Avaliação do modelo",
		"Revisão da estrutura de gestão de risco",
		"Auditoria interna"}
	elemento.Itens = itens
	elemento.Nome = "Monitoramento"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Controle"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = "Risco de Mercado"
	componente.TiposNotas = tiposNotas
	return componente
}

func initRiscoDeLiquidez() ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementos []ElementoESI

	var itens = []string{"Composição do colchão de liquidez",
		"Descasamento no fluxo de caixa",
		"Índice de Liquidez Ampla",
		"Índice de Liquidez Restrita"}
	var elemento ElementoESI
	elemento.Itens = itens
	elemento.Nome = "Nível de risco de liquidez"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Risco"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	// CONTROLE
	elementos = make([]ElementoESI, 0)
	itens = []string{"Órgãos estatutários e infra estatutários", "Estratégia e políticas"}
	elemento.Itens = itens
	elemento.Nome = "Ambiente de Controle"
	elementos = append(elementos, elemento)

	itens = []string{"Sistemas de mensuração", "Estudo ALM"}
	elemento.Itens = itens
	elemento.Nome = "Identificação, avaliação, mensuração"
	elementos = append(elementos, elemento)

	itens = []string{"Estrutura de limites",
		"Plano de contingência",
		"Função de controle do risco de liquidez"}
	elemento.Itens = itens
	elemento.Nome = "Controles específicos"
	elementos = append(elementos, elemento)

	itens = []string{"Relatórios e formalização do processo"}
	elemento.Itens = itens
	elemento.Nome = "Comunicação"
	elementos = append(elementos, elemento)

	itens = []string{"Gerência de controles internos",
		"Auditoria interna",
		"Conselho Fiscal"}
	elemento.Itens = itens
	elemento.Nome = "Monitoramento"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Controle"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = "Risco de Liquidez"
	componente.TiposNotas = tiposNotas
	return componente
}

func initRiscoAtuarial() ComponenteESI {
	var componente ComponenteESI
	var tiposNotas []TipoNotaESI
	var tipoNota TipoNotaESI
	var elementos []ElementoESI

	var itens = []string{"Erro operacional na inserção ou atualização de dados pela EFPC",
		"Erro operacional na inserção ou atualização de dados pela patrocinadora",
		"Fornecimento de informação falsa ou incorreta pelo participante ou assistido",
		"Fraude na concessão de benefícios",
		"Falta de atualização espontânea de informações pelo participante ou assistido"}
	var elemento ElementoESI
	elemento.Itens = itens
	elemento.Nome = "Risco de Cadastro"
	elementos = append(elementos, elemento)

	itens = []string{"Inconsistência nas bases de dados usadas para testes estatísticos",
		"Processos inadequados tecnicamente e/ou insuficientemente manualizados",
		"Falta de atenção e/ou falta de zelo e/ou falta de competência e/ou má fé dos colaboradores e terceirizados envolvidos; 	Falta de controle sobre os trabalhos dos colaboradores e/ou terceirizados envolvidos; 	Sistemas parametrizados de forma inadequada e/ou sistemas inadequados, produzindo erros sistemáticos e/ou aumentando a probabilidade de erro humano;	Compreensão insuficiente da alta governança da EFPC sobre o funcionamento das hipóteses atuariais",
		"Intenção de ocultar resultados negativos ou postergar medidas de equacionamento"}
	elemento.Itens = itens
	elemento.Nome = "Risco Operacional Atuarial"
	elementos = append(elementos, elemento)

	itens = []string{"Natureza estocástica dos eventos", "Probabilidade e Impacto"}
	elemento.Itens = itens
	elemento.Nome = "Risco dos Eventos"
	elementos = append(elementos, elemento)

	itens = []string{"Inconsistência entre a modelagem da avaliação atuarial e as disposições do regulamento",
		"Adoção de regimes financeiros e métodos de financiamento que produzem custos insuportáveis ao longo do tempo para participantes e assistidos, especialmente em planos fechados a novas adesões",
		"Adoção de hipóteses atuariais conservadoras por longo período de tempo",
		"Ausência de hipóteses para representação de determinados eventos."}
	elemento.Itens = itens
	elemento.Nome = "Risco de Modelagem"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Risco"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	// CONTROLE
	elementos = make([]ElementoESI, 0)
	itens = []string{"Estrutura da área de gestão do risco atuarial",
		"Estrutura de responsabilidades e alçadas"}
	elemento.Itens = itens
	elemento.Nome = "Ambiente de Controle"
	elementos = append(elementos, elemento)

	itens = []string{"Gestão do risco de cadastro",
		"Gestão do risco operacional atuarial",
		"Gestão do risco dos eventos",
		"Gestão do risco de modelagem",
		"Gestão de mitigadores de risco"}
	elemento.Itens = itens
	elemento.Nome = "Identificação, avaliação, mensuração"
	elementos = append(elementos, elemento)

	itens = []string{"Processo atuarial"}
	elemento.Itens = itens
	elemento.Nome = "Controles específicos"
	elementos = append(elementos, elemento)

	itens = []string{"Revisão das estratégias, políticas e limites",
		"Monitoramento do risco atuarial"}
	elemento.Itens = itens
	elemento.Nome = "Monitoramento"
	elementos = append(elementos, elemento)

	tipoNota.Nome = "Controle"
	tipoNota.Elementos = elementos
	tiposNotas = append(tiposNotas, tipoNota)

	componente.Nome = "Risco Atuarial"
	componente.TiposNotas = tiposNotas
	return componente
}
