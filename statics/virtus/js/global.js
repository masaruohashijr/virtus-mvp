var isChefe = false;
var actions_array = [];
var activity;
var activities = [];
var atualizacao;
var atualizacoes = [];
var anotacao;
var anotacoes = [];
var anotacaoRadar;
var anotacoesRadar = [];
var anotacoesMap;
var anotacoesSiglasMap;
var comentario;
var comentarios = [];
var item;
var itens = [];
var contexto;
var opened = false;
var plano;
var planos = [];
var entidadesMap;
var ciclosMap;
var pilaresMap;
var planosMap;
var componentesMap;
var tiposNotasMap;
var elementosMap;
var itensMap;
var auditoresMap;
var siglasMap;
var cicloEntidade;
var ciclosEntidade = [];
var pilarCiclo;
var pilaresCiclo = [];
var elementoComponente;
var elementosComponente = [];
var componentePilar;
var componentesPilar = [];
var jurisdicao;
var jurisdicoes = [];
var historico;
var historicos = [];
var membro;
var membros = [];
var tipo;
var tipos = [];
var matriz;

function parseNome2Valor(nome){
	let letra = nome.substr(0,1);
	return letra;
}