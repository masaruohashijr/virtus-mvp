var elemento_componente_tobe_deleted;
	
class TipoNota {
	constructor(order, id, tipoNotaId, componenteId, nome, descricao, letra, corLetra, pesoPadrao, authorId, authorName, createdAt, c_createdAt, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.tipoNotaId = tipoNotaId;
		this.componenteId = componenteId;
		this.nome = nome;
		this.descricao = descricao;
		this.letra = letra;
		this.corLetra = corLetra;
		this.pesoPadrao = pesoPadrao;
		this.authorId = authorId;
		this.authorName = authorName;
		this.createdAt = createdAt;
		this.c_createdAt = c_createdAt;
		this.idVersaoOrigem = idVersaoOrigem;
		this.statusId = statusId;
		this.cStatus = cStatus;
	}
}
