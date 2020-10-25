var elemento_componente_tobe_deleted;
	
class ElementoComponente {
	constructor(order, id, componenteId, elementoId, elementoNome, pesoPadrao, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.componenteId = componenteId;
		this.elementoId = elementoId;
		this.elementoNome = elementoNome;
		this.pesoPadrao = pesoPadrao;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.idVersaoOrigem = idVersaoOrigem;
		this.statusId = statusId;
		this.cStatus = cStatus;
	}
}

function criarElementoComponente(){
	console.log('criarElementoComponente');
	/* Campo Select para selecionar o Elemento vinculado ao Componente */
	let campoSelect = document.getElementById('ElementoComponenteForInsert');
	let elementoId = 0;
	let elementoNome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			elementoId = campoSelect.options[n].value;
			elementoNome = campoSelect.options[n].text;
			break;
		}
	}
	/* FIM */
	/* Validação do preenchimento do campo obrigatório SELECT de Elemento */
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o elemento.\n';
		alert(erros);
		return;
	}
	let pesoPadrao = document.getElementById('PesoPadraoForInsert').value;
	elementoComponenteId = getMaxId(elementosComponente);
	elementoComponente = new ElementoComponente(0, elementoComponenteId, 0, elementoId, elementoNome, pesoPadrao, '', '', '', '', '', '', '');
	elementosComponente.push(elementoComponente);
	addElementoComponenteRow("table-elementos-componente-"+contexto);
	limparCamposElementoComponenteForm();
	document.getElementById('create-elemento-componente-form').style.display='none';
}

function addElementoComponenteRow(tableID) {
	console.log('addElementoComponenteRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = elementosComponente.length-1;
	elementoComponente = elementosComponente[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(elementoComponente.elementoNome);
	let json = JSON.stringify(elementoComponente);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="elementoComponente'+elementoComponente.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="componenteId" value="'+elementoComponente.componenteId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="elementoId" value="'+elementoComponente.elementoId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+elementoComponente.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Peso Padrão
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(elementoComponente.pesoPadrao);
	newCell.appendChild(newText);
	// Criado Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(elementoComponente.criadoEm);
	newCell.appendChild(newText);
	// Autor
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(elementoComponente.autorNome);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+elementoComponente.autorId+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editElementoComponente(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteElementoComponenteForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposElementoComponenteForm(){
	console.log("limparCamposElementoComponenteForm");
	document.getElementById("formulario-elemento-componente-create").reset();
	document.getElementById("formulario-elemento-componente-edit").reset();
}

function editElementoComponente(e) {
	console.log('editElementoComponente');
	var editElementoComponenteForm = document.getElementById('edit-elemento-componente-form');
	editElementoComponenteForm.style.display = 'block';
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var elementoId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var componenteId = e.parentNode.parentNode.childNodes[0].childNodes[3].value;
//	var elementoNome = e.parentNode.parentNode.childNodes[0].innerText;
	var pesoPadrao = e.parentNode.parentNode.childNodes[1].innerText;
//	var criadoEm = e.parentNode.parentNode.childNodes[2].innerText;
//	var autorId = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
//	var autorNome = e.parentNode.parentNode.childNodes[3].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('ElementoComponenteForUpdate').value=elementoId;
	document.getElementById('PesoPadraoForUpdate').value=pesoPadrao;
	document.getElementById('Id-ECForUpdate').value=id;
	document.getElementById('Order-ECForUpdate').value=order;
	document.getElementById('ComponenteId-ECForUpdate').value=componenteId;
}

function updateElementoComponente() {
	console.log('updateElementoComponente');
	var id = document.getElementById('Id-ECForUpdate').value;
	var order = document.getElementById('Order-ECForUpdate').value;
	var componenteId = document.getElementById('ComponenteId-ECForUpdate').value;
	var pesoPadrao = document.getElementById('PesoPadraoForUpdate').value;
	let campoSelect = document.getElementById('ElementoComponenteForUpdate');
	let elementoId = 0;
	let elementoNome = '';
	console.log(campoSelect.options.length);
	console.log(campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			elementoId = campoSelect.options[n].value;
			elementoNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o elemento.\n';
		alert(erros);
		return;
	}
	elementoComponente = new ElementoComponente(order, id, componenteId, elementoId, elementoNome, pesoPadrao, '', '', '', '', '', '');
	elementosComponente[order] = elementoComponente;
	updateElementoComponenteRow("table-elementos-componente-"+contexto,order);
	limparCamposElementoComponenteForm();
	document.getElementById('edit-elemento-componente-form').style.display='none';
}

function updateElementoComponenteRow(tableID, order){
	console.log('updateElementoComponente');
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
	console.log(elementosComponente[order].elementoNome);
	celula.innerText = elementosComponente[order].elementoNome;
	let json = JSON.stringify(elementosComponente[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="elementoComponente'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('elementoComponente.elementoId: '+elementoComponente.elementoId);
	celula.innerHTML = '<input type="hidden" name="componenteId" value="'+elementoComponente.componenteId+'"/>'+celula.innerHTML;
	console.log('elementosComponente[order].id: '+elementosComponente[order].id);
	celula.innerHTML = '<input type="hidden" name="elementoId" value="'+elementoComponente.elementoId+'"/>'+celula.innerHTML;
	console.log('elementoComponente.componenteId: '+elementoComponente.elementoId);
	celula.innerHTML = '<input type="hidden" name="id" value="'+elementosComponente[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	celula.innerText = elementosComponente[order].pesoPadrao;
}

function showDeleteElementoComponenteForm(e){
	console.log('showDeleteElementoComponenteForm');
	var deleteElementoComponenteForm = document.getElementById('delete-elemento-componente-form');
	deleteElementoComponenteForm.style.display = 'block';
	elemento_componente_tobe_deleted = e;
}

function deleteElementoComponente() {
	console.log('deleteElementoComponente');
	var order = elemento_componente_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newElementosComponente = [];
	let tbl = elemento_componente_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<elementosComponente.length;i++){
		if(i != order){
			newElementosComponente.push(elementosComponente[i]);
		}
	}
	elementosComponente = newElementosComponente;
	var deleteElementoComponenteForm = document.getElementById('delete-elemento-componente-form');
	deleteElementoComponenteForm.style.display = 'none';
}