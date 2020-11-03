var ciclo_entidade_tobe_deleted;
	
class CicloEntidade {
	constructor(order, id, entidadeId, cicloId, nome, tipoMediaId, tipoMedia, iniciaEm, terminaEm, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.entidadeId = entidadeId;
		this.cicloId = cicloId;
		this.nome = nome;
		this.tipoMediaId = tipoMediaId;
		this.tipoMedia = tipoMedia;
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

function criarCicloEntidade(){
	console.log('criarCicloEntidade');
	let campoSelect = document.getElementById('CicloEntidadeForInsert');
	let cicloId = 0;
	let nome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			cicloId = campoSelect.options[n].value;
			nome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o ciclo.\n';
		alert(erros);
		return;
	}
	let tipoMediaId = 0;
	let tipoMedia = '';
	campoSelect = document.getElementById('TipoMediaForInsert');
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			tipoMediaId = campoSelect.options[n].value;
			tipoMedia = campoSelect.options[n].text;
			break;
		}
	}
	let iniciaEm = document.getElementById('IniciaEmForInsert').value;
	let terminaEm = document.getElementById('TerminaEmForInsert').value;
	cicloEntidadeId = getMaxId(ciclosEntidade);
	cicloEntidade = new CicloEntidade(0, cicloEntidadeId, 0, cicloId, nome, tipoMediaId, tipoMedia, iniciaEm, terminaEm, '', '', '', '', '', '', '');
	ciclosEntidade.push(cicloEntidade);
	addCicloEntidadeRow("table-ciclos-entidade-"+contexto);
	limparCamposCicloEntidadeForm();
	document.getElementById('create-ciclo-entidade-form').style.display='none';
}

function addCicloEntidadeRow(tableID) {
	console.log('addCicloEntidadeRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = ciclosEntidade.length-1;
	cicloEntidade = ciclosEntidade[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(cicloEntidade.nome);
	let json = JSON.stringify(cicloEntidade);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="cicloEntidade'+cicloEntidade.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="cicloId" value="'+cicloEntidade.cicloId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+cicloEntidade.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+cicloEntidade.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Tipo de Média
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(cicloEntidade.tipoMedia);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+cicloEntidade.tipoMediaId+'"/>'+newCell.innerHTML;
	// Inicio Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(cicloEntidade.iniciaEm);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(cicloEntidade.terminaEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+cicloEntidade.autorId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" value="'+cicloEntidade.criadoEm+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editCicloEntidade(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteCicloEntidadeForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposCicloEntidadeForm(){
	console.log('limparCamposCicloEntidadeForm');
	document.getElementById('formulario-ciclo-entidade-create').reset()
	document.getElementById('formulario-ciclo-entidade-edit').reset()
}

function editCicloEntidade(e) {
	console.log('editCicloEntidade');
	let editCicloEntidadeForm = document.getElementById('edit-ciclo-entidade-form');
	editCicloEntidadeForm.style.display = 'block';
	let linha = e.parentNode.parentNode;
	let order = linha.childNodes[0].childNodes[0].value;
	let id = linha.childNodes[0].childNodes[1].value;
	let entidadeId = linha.childNodes[0].childNodes[2].value;
	let cicloId = linha.childNodes[0].childNodes[3].value;
	let tipoMediaId = linha.childNodes[1].childNodes[0].value;
	// let tipoMedia = linha.childNodes[1].innerText;
	let iniciaEm = formatarData(linha.childNodes[2].innerText);
	console.log(iniciaEm);
	let terminaEm = formatarData(linha.childNodes[3].innerText);
	console.log(terminaEm);
	// Atribuindo os valores de edit-item-form
	document.getElementById('Id-CEForUpdate').value=id;
	document.getElementById('Order-CEForUpdate').value=order;
	document.getElementById('EntidadeId-CEForUpdate').value=entidadeId;
	document.getElementById('CicloEntidadeForUpdate').value=cicloId;
	document.getElementById('TipoMediaForUpdate').value=tipoMediaId;
	document.getElementById('IniciaEmForUpdate').value=iniciaEm;
	document.getElementById('TerminaEmForUpdate').value=terminaEm;
}

function updateCicloEntidade() {
	console.log('updateCicloEntidade');
	let id = document.getElementById('Id-CEForUpdate').value;
	let order = document.getElementById('Order-CEForUpdate').value;
	let entidadeId = document.getElementById('EntidadeId-CEForUpdate').value;
	let campoSelect = document.getElementById('CicloEntidadeForUpdate');
	let cicloId = 0;
	let nome = '';
	console.log(campoSelect.options.length);
	console.log(campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			cicloId = campoSelect.options[n].value;
			nome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o ciclo.\n';
		alert(erros);
		return;
	}
	let tipoMediaId = 0;
	let tipoMedia = '';
	campoSelect = document.getElementById('TipoMediaForUpdate');
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			tipoMediaId = campoSelect.options[n].value;
			tipoMedia = campoSelect.options[n].text;
			break;
		}
	}
	let iniciaEm = document.getElementById('IniciaEmForUpdate').value;
	let terminaEm = document.getElementById('TerminaEmForUpdate').value;
	cicloEntidade = new CicloEntidade(order, id, entidadeId, cicloId, nome, tipoMediaId, tipoMedia, iniciaEm, terminaEm, '', '', '', '', '', '');
	ciclosEntidade[order] = cicloEntidade;
	updateCicloEntidadeRow("table-ciclos-entidade-"+contexto,order);
	limparCamposCicloEntidadeForm();
	document.getElementById('edit-ciclo-entidade-form').style.display='none';
}

function updateCicloEntidadeRow(tableID, order){
	console.log('updateCicloEntidade');
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
	console.log(ciclosEntidade[order].nome);
	celula.innerText = ciclosEntidade[order].nome;
	let json = JSON.stringify(ciclosEntidade[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="cicloEntidade'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('cicloEntidade.cicloId: '+cicloEntidade.cicloId);
	celula.innerHTML = '<input type="hidden" name="cicloId" value="'+cicloEntidade.cicloId+'"/>'+celula.innerHTML;
	console.log('cicloEntidade.entidadeId: '+cicloEntidade.entidadeId);
	celula.innerHTML = '<input type="hidden" name="entidadeId" value="'+cicloEntidade.entidadeId+'"/>'+celula.innerHTML;
	console.log('ciclosEntidade[order].id: '+ciclosEntidade[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+ciclosEntidade[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	console.log('ciclosEntidade[order].tipoMedia: '+ciclosEntidade[order].tipoMedia);
	celula.innerText = ciclosEntidade[order].tipoMedia;
	celula.innerHTML = '<input type="hidden" value="'+ciclosEntidade[order].tipoMediaId+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('ciclosEntidade[order].iniciaEm: '+ciclosEntidade[order].iniciaEm);
	celula.innerText = ciclosEntidade[order].iniciaEm;
	celula = row.childNodes[3];
	console.log('ciclosEntidade[order].terminaEm: '+ciclosEntidade[order].terminaEm);
	celula.innerText = ciclosEntidade[order].terminaEm;
}

function showDeleteCicloEntidadeForm(e){
	console.log('showDeleteCicloEntidadeForm');
	let deleteCicloEntidadeForm = document.getElementById('delete-ciclo-entidade-form');
	deleteCicloEntidadeForm.style.display = 'block';
	ciclo_entidade_tobe_deleted = e;
}

function deleteCicloEntidade() {
	console.log('deleteCicloEntidade');
	let order = ciclo_entidade_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	let newCicloEntidades = [];
	let tbl = ciclo_entidade_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<ciclosEntidade.length;i++){
		if(i != order){
			newCicloEntidades.push(ciclosEntidade[i]);
		}
	}
	ciclosEntidade = newCicloEntidades;
	let deleteCicloEntidadeForm = document.getElementById('delete-ciclo-entidade-form');
	deleteCicloEntidadeForm.style.display = 'none';
}

