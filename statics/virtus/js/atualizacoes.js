var atualizacao_tobe_deleted;
	
class Atualizacao {
	constructor(order, id, anotacaoId, texto, autor, c_criadoEm) {
		this.order = order;
		this.id = id;
		this.anotacaoId = anotacaoId;
		this.texto = texto;
		this.autor = autor;
		this.c_criadoEm = c_criadoEm;
	}
}

function openCreateAtualizacao(){
	document.getElementById('create-atualizacao-anotacao-form').style.display='block';
	document.getElementById('AtualizacaoForInsert').focus();
}

function criarAtualizacao(){
	console.log('criarAtualizacao');
	var atualizacao = document.getElementById('AtualizacaoForInsert').value;
	atualizacaoId = getMaxId(atualizacoes);
	atualizacao = new Atualizacao(0, atualizacaoId, atualizacao);
	atualizacoes.push(atualizacao);
	addAtualizacaoRow("table-atualizacoes-anotacao-"+contexto);	
	document.getElementById('AtualizacaoForInsert').value = '';
	document.getElementById('create-atualizacao-anotacao-form').style.display='none';
}

function addAtualizacaoRow(tableID) {
	console.log(tableID);
	let tableRef = document.getElementById(tableID);
	console.log('tableRef: '+tableRef);
	let newRow = tableRef.insertRow(-1);
	console.log('newRow: '+newRow);
	order = atualizacoes.length-1;
	atualizacao = atualizacoes[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(atualizacao.texto);
	newCell.style = "text-align: left";
	var jsonAtualizacao = JSON.stringify(atualizacao);
	jsonAtualizacao = jsonAtualizacao.split(',').join('#');
	jsonAtualizacao = jsonAtualizacao.split('"').join('');
	jsonAtualizacao = jsonAtualizacao.split('{').join('');
	jsonAtualizacao = jsonAtualizacao.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="atualizacao'+atualizacao.id+'" value="'+jsonItem+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="anotacaoId" value="'+atualizacao.anotacaoId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+atualizacao.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// atualizacao
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(atualizacao.texto);
	newCell.appendChild(newText);
	// autor
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(atualizacao.autor);
	newCell.appendChild(newText);
	// criado em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(atualizacao.c_criadoEm);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editAtualizacao(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteAtualizacaoForm(btnApagar)};
	newCell.appendChild(btnApagar);
}
