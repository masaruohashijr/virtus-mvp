var ciclo_entidade_tobe_deleted;
	
class CicloEntidade {
	constructor(order, id, entidadeId, cicloId, nome, tipoMediaId, tipoMedia, nota, iniciaEm, terminaEm, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.entidadeId = entidadeId;
		this.cicloId = cicloId;
		this.nome = nome;
		this.tipoMediaId = tipoMediaId;
		this.tipoMedia = tipoMedia;
		this.nota = nota;
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
	let nota = document.getElementById('NotaForInsert').value;
	let iniciaEm = document.getElementById('IniciaEmForInsert').value;
	let terminaEm = document.getElementById('TerminaEmForInsert').value;
	cicloEntidade = new CicloEntidade(0, ciclos.length, 0, cicloId, nome, tipoMediaId, tipoMedia, nota, iniciaEm, terminaEm, '', '', '', '', '', '', '');
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
	// Nota
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(cicloEntidade.nota);
	newCell.appendChild(newText);
	// Inicio Em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(cicloEntidade.iniciaEm);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(cicloEntidade.terminaEm);
	newCell.appendChild(newText);
	// Autor
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(cicloEntidade.autorNome);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+cicloEntidade.autorId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" value="'+cicloEntidade.criadoEm+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(6);
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
	var editPlanoForm = document.getElementById('edit-ciclo-entidade-form');
	editPlanoForm.style.display = 'block';
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var entidadeId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var cicloId = e.parentNode.parentNode.childNodes[0].childNodes[3].value;
	var nome = e.parentNode.parentNode.childNodes[0].innerText;
	var tipoMediaId = e.parentNode.parentNode.childNodes[1].childNodes[0].value;
	var tipoMedia = e.parentNode.parentNode.childNodes[1].innerText;
	var nota = e.parentNode.parentNode.childNodes[2].innerText;
	var iniciaEm = e.parentNode.parentNode.childNodes[3].innerText;
	var terminaEm = e.parentNode.parentNode.childNodes[4].innerText;
	var criadoEm = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
	var autorId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
	var autorName = e.parentNode.parentNode.childNodes[5].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('Id-CEForUpdate').value=id;
	document.getElementById('Order-CEForUpdate').value=order;
	document.getElementById('EntidadeId-CEForUpdate').value=entidadeId;
	document.getElementById('CicloEntidadeForUpdate').value=cicloId;
	document.getElementById('TipoMediaForUpdate').value=tipoMediaId;
	document.getElementById('NotaForUpdate').value=nota;
	document.getElementById('IniciaEmForUpdate').value=iniciaEm;
	document.getElementById('TerminaEmForUpdate').value=terminaEm;
	document.getElementById('AutorForUpdate').value=autorName;
	document.getElementById('CriadoEmForUpdate').value=criadoEm;
	
}

function updateCicloEntidade() {
	console.log('updateCicloEntidade');
	var id = document.getElementById('Id-CEForUpdate').value;
	var order = document.getElementById('Order-CEForUpdate').value;
	var entidadeId = document.getElementById('EntidadeId-CEForUpdate').value;
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
	let nota = document.getElementById('NotaForUpdate').value;
	let iniciaEm = document.getElementById('IniciaEmForUpdate').value;
	let terminaEm = document.getElementById('TerminaEmForUpdate').value;
	cicloEntidade = new CicloEntidade(order, id, entidadeId, cicloId, nome, tipoMediaId, tipoMedia, nota, iniciaEm, terminaEm, '', '', '', '', '', '');
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
	console.log('ciclosEntidade[order].nota: '+ciclosEntidade[order].nota);
	celula.innerText = ciclosEntidade[order].nota;
	celula = row.childNodes[3];
	console.log('ciclosEntidade[order].iniciaEm: '+ciclosEntidade[order].iniciaEm);
	celula.innerText = ciclosEntidade[order].iniciaEm;
	celula = row.childNodes[4];
	console.log('ciclosEntidade[order].terminaEm: '+ciclosEntidade[order].terminaEm);
	celula.innerText = ciclosEntidade[order].terminaEm;
}
