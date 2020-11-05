var jurisdicao_tobe_deleted;
	
class Jurisdicao {
	constructor(order, id, escritorioId, entidadeId, entidadeNome, iniciaEm, terminaEm, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.escritorioId = escritorioId;
		this.entidadeId = entidadeId;
		this.entidadeNome = entidadeNome;
		this.iniciaEm = iniciaEm;
		this.terminaEm = terminaEm;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.idVersaoOrigem = idVersaoOrigem;
		this.statusId = statusId;
		this.cStatus = cStatus;
	}
}

function criarJurisdicao(){
	console.log('criarJurisdicao');
	let campoSelect = document.getElementById('EntidadeForInsert');
	let entidadeId = 0;
	let entidadeNome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			entidadeId = campoSelect.options[n].value;
			entidadeNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular a entidade.\n';
		alert(erros);
		return;
	}
	let iniciaEm = document.getElementById('IniciaEmForInsert').value;
	let terminaEm = document.getElementById('TerminaEmForInsert').value;
	jurisdicaoId = getMaxId(jurisdicoes);
	jurisdicao = new Jurisdicao(0, jurisdicaoId, 0, entidadeId, entidadeNome, iniciaEm, terminaEm, '', '', '', '', '', '', '');
	jurisdicoes.push(jurisdicao);
	addJurisdicaoRow("table-jurisdicoes-"+contexto);
	limparCamposJurisdicaoForm();
	document.getElementById('create-jurisdicao-form').style.display='none';
}

function addJurisdicaoRow(tableID) {
	console.log('addJurisdicaoRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = jurisdicoes.length-1;
	jurisdicao = jurisdicoes[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(jurisdicao.entidadeNome);
	let json = JSON.stringify(jurisdicao);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="jurisdicao'+jurisdicao.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="escritorioId" value="'+jurisdicao.escritorioId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+jurisdicao.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+jurisdicao.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Inicia Em
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(jurisdicao.iniciaEm);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(jurisdicao.terminaEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+jurisdicao.autorId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" value="'+jurisdicao.criadoEm+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(3);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editJurisdicao(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteJurisdicaoForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposJurisdicaoForm(){
	console.log('limparCamposJurisdicaoForm');
	document.getElementById('formulario-jurisdicao-create').reset()
	document.getElementById('formulario-jurisdicao-edit').reset()
}

function updateJurisdicao() {
	console.log('updateJurisdicao');
	var id = document.getElementById('Id-JUForUpdate').value;
	var order = document.getElementById('Order-JUForUpdate').value;
	var escritorioId = document.getElementById('EscritorioId-JUForUpdate').value;
	let campoSelect = document.getElementById('EntidadeForUpdate');
	let entidadeId = 0;
	let entidadeNome = '';
	console.log(campoSelect.options.length);
	console.log(campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			entidadeId = campoSelect.options[n].value;
			entidadeNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular a entidade.\n';
		alert(erros);
		return;
	}
	let iniciaEm = document.getElementById('IniciaEmForUpdate').value;
	let terminaEm = document.getElementById('TerminaEmForUpdate').value;
	jurisdicao = new Jurisdicao(order, id, escritorioId, entidadeId, entidadeNome, iniciaEm, terminaEm, '', '', '', '', '', '');
	jurisdicoes[order] = jurisdicao;
	updateJurisdicaoRow("table-jurisdicoes-"+contexto,order);
	limparCamposJurisdicaoForm();
	document.getElementById('edit-jurisdicao-form').style.display='none';
}

function updateJurisdicaoRow(tableID, order){
	console.log('updateJurisdicao');
	console.log('contexto: '+contexto);
	let tbl = document.getElementById(tableID);
	console.log('tableID: '+tableID);
	let linhas = tbl.childNodes[1].childNodes;
	console.log('linhas: '+linhas);
	let row = tbl.childNodes[0];
	console.log('row: '+row);
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			console.log(inputOrder);
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				console.log('tagName: '+inputOrder.tagName);
				console.log('value: '+inputOrder.value);
				console.log('order: '+order);
				if(inputOrder.value==order){
					console.log("y: "+y);
					row = linhas[y];
					break;
				}
			}
		}
	}
	let celula = row.childNodes[0];
	console.log(jurisdicoes[order].entidadeNome);
	celula.innerText = jurisdicoes[order].entidadeNome;
	let json = JSON.stringify(jurisdicoes[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="jurisdicao'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('jurisdicao.entidadeId: '+jurisdicao.entidadeId);
	celula.innerHTML = '<input type="hidden" name="entidadeId" value="'+jurisdicao.entidadeId+'"/>'+celula.innerHTML;
	console.log('jurisdicao.escritorioId: '+jurisdicao.escritorioId);
	celula.innerHTML = '<input type="hidden" name="escritorioId" value="'+jurisdicao.escritorioId+'"/>'+celula.innerHTML;
	console.log('jurisdicoes[order].id: '+jurisdicoes[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+jurisdicoes[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	let strIniciaEm = formatarData(jurisdicoes[order].iniciaEm);
	let strTerminaEm = formatarData(jurisdicoes[order].terminaEm);
	console.log('jurisdicoes[order].iniciaEm: '+strIniciaEm);
	celula.innerText = strIniciaEm;
	celula = row.childNodes[2];
	console.log('jurisdicoes[order].terminaEm: '+strTerminaEm);
	celula.innerText = strTerminaEm;
}

function showDeleteJurisdicaoForm(e){
	console.log('showDeleteJurisdicaoForm');
	var deleteJurisdicaoForm = document.getElementById('delete-jurisdicao-form');
	deleteJurisdicaoForm.style.display = 'block';
	jurisdicao_tobe_deleted = e;
}

function deleteJurisdicao() {
	console.log('deleteJurisdicao');
	var order = jurisdicao_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newJurisdicoes = [];
	let tbl = jurisdicao_tobe_deleted.parentNode.parentNode.parentNode;
	let linhas = tbl.childNodes;
	contadorLinha = 1;
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				if(inputOrder.value==order){
					if(inputOrder.value == order){
						tbl.deleteRow(contadorLinha);
						break;
					}
				}
				contadorLinha ++;
			}
		}
	}
	for(i=0;i<jurisdicoes.length;i++){
		if(i != order){
			newJurisdicoes.push(jurisdicoes[i]);
		}
	}
	jurisdicoes = newJurisdicoes;
	var deleteJurisdicaoForm = document.getElementById('delete-jurisdicao-form');
	deleteJurisdicaoForm.style.display = 'none';
}

function editJurisdicao(e) {
	console.log('editJurisdicao');
	var editJurisdicaoForm = document.getElementById('edit-jurisdicao-form');
	editJurisdicaoForm.style.display = 'block';
	var linha = e.parentNode.parentNode;
	var order = linha.childNodes[0].childNodes[0].value;
	var id = linha.childNodes[0].childNodes[1].value;
	var entidadeId = linha.childNodes[0].childNodes[2].value;
	var escritorioId = linha.childNodes[0].childNodes[3].value;
	var iniciaEm = formatarData(linha.childNodes[1].innerText);
	console.log(iniciaEm);
	var terminaEm = formatarData(linha.childNodes[2].innerText);
	console.log(terminaEm);
	// Atribuindo os valores de edit-item-form
	document.getElementById('Id-JUForUpdate').value=id;
	document.getElementById('Order-JUForUpdate').value=order;
	document.getElementById('EscritorioId-JUForUpdate').value=escritorioId;
	document.getElementById('EntidadeForUpdate').value=entidadeId;
	document.getElementById('IniciaEmForUpdate').value=iniciaEm;
	document.getElementById('TerminaEmForUpdate').value=terminaEm;
}

