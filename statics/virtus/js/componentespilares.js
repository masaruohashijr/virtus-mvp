var componente_pilpilaar_tobe_deleted;
	
class ComponentePilar {
	constructor(order, id, pilarId, componenteId, componenteNome, tipoMediaId, tipoMedia, sonda, pesoPadrao, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.pilarId = pilarId;
		this.componenteId = componenteId;
		this.componenteNome = componenteNome;
		this.tipoMediaId = tipoMediaId;
		this.tipoMedia = tipoMedia;
		this.sonda = sonda;
		this.pesoPadrao = pesoPadrao;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.idVersaoOrigem = idVersaoOrigem;
		this.statusId = statusId;
		this.cStatus = cStatus;
	}
}

function criarComponentePilar(){
	console.log('criarComponentePilar');
	let campoSelect = document.getElementById('ComponentePilarForInsert');
	let componenteId = 0;
	let componenteNome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			componenteId = campoSelect.options[n].value;
			componenteNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o componente.\n';
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
	let sonda = document.getElementById('SondaForInsert').value;
	let pesoPadrao = document.getElementById('PesoPadraoForInsert').value;
	componentePilarId = getMaxId(componentesPilar);
	componentePilar = new ComponentePilar(0, componentePilarId, 0, componenteId, componenteNome, tipoMediaId, tipoMedia, sonda, pesoPadrao, '', '', '', '', '', '');
	componentesPilar.push(componentePilar);
	addComponentePilarRow("table-componentes-pilar-"+contexto);
	limparCamposComponentePilarForm();
	document.getElementById('create-componente-pilar-form').style.display='none';
}

function addComponentePilarRow(tableID) {
	console.log('addComponentePilarRow');
	console.log(tableID);
	let tableRef = document.getElementById(tableID);
	console.log(tableRef);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	console.log(newRow);
	order = componentesPilar.length-1;
	componentePilar = componentesPilar[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(componentePilar.componenteNome);
	let json = JSON.stringify(componentePilar);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="componentePilar'+componentePilar.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="componenteId" value="'+componentePilar.componenteId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="pilarId" value="'+componentePilar.pilarId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+componentePilar.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	newCell.style="text-align:left"; 
	// Tipo de Média
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(componentePilar.tipoMedia);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+componentePilar.tipoMediaId+'"/>'+newCell.innerHTML;
	// Peso Padrão
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(componentePilar.pesoPadrao);
	newCell.appendChild(newText);
	// Sonda
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(componentePilar.sonda);
	newCell.appendChild(newText);
	// Autor
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(componentePilar.autorNome);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+componentePilar.autorId+'"/>'+newCell.innerHTML;
	// Criado Em
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(componentePilar.criadoEm);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(6);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editComponentePilar(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteComponentePilarForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposComponentePilarForm(){
	console.log('limparCamposComponentePilarForm');
	document.getElementById('formulario-componente-pilar-create').reset()
	document.getElementById('formulario-componente-pilar-edit').reset()
}

function editComponentePilar(e) {
	console.log('editComponentePilar');
	var editPlanoForm = document.getElementById('edit-componente-pilar-form');
	editPlanoForm.style.display = 'block';
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var pilarId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var componenteId = e.parentNode.parentNode.childNodes[0].childNodes[3].value;
	var tipoMediaId = e.parentNode.parentNode.childNodes[1].childNodes[0].value;
	var pesoPadrao = e.parentNode.parentNode.childNodes[2].innerText;
	var sonda = e.parentNode.parentNode.childNodes[3].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('pilarId-edit').value=pilarId;
	document.getElementById('ComponentePilarForUpdate').value=componenteId;
	document.getElementById('TipoMediaForUpdate').value=tipoMediaId;
	document.getElementById('PesoPadraoForUpdate').value=pesoPadrao;
	document.getElementById('SondaForUpdate').value=sonda;
}

function updateComponentePilar() {
	console.log('updateComponentePilar');
	var id = document.getElementById('id-edit').value;
	var order = document.getElementById('order-edit').value;
	var pilarId = document.getElementById('pilarId-edit').value;
	let campoSelect = document.getElementById('ComponentePilarForUpdate');
	let componenteId = 0;
	let componenteNome = '';
	console.log(campoSelect.options.length);
	console.log(campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			componenteId = campoSelect.options[n].value;
			componenteNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o componente.\n';
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
	let pesoPadrao = document.getElementById('PesoPadraoForUpdate').value;
	let sonda = document.getElementById('SondaForUpdate').value;
	console.log('*** Logando Componente Pilar ***');
	console.log(order + ', ' + id + ', ' + pilarId + ', ' + componenteId + ', ' + componenteNome + ', ' + tipoMediaId + ', ' + tipoMedia + ', ' + sonda + ', ' + pesoPadrao);
	componentePilar = new ComponentePilar(order, id, pilarId, componenteId, componenteNome, tipoMediaId, tipoMedia, sonda, pesoPadrao, '', '', '', '', '', '');
	componentesPilar[order] = componentePilar;
	updateComponentePilarRow("table-componentes-pilar-"+contexto,order);
	limparCamposComponentePilarForm();
	document.getElementById('edit-componente-pilar-form').style.display='none';
}

function updateComponentePilarRow(tableID, order){
	console.log('updateComponentePilar');
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
	console.log(componentesPilar[order].nome);
	celula.innerText = componentesPilar[order].componenteNome;
	let json = JSON.stringify(componentesPilar[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="componentePilar'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('componentePilar.componenteId: '+componentePilar.componenteId);
	celula.innerHTML = '<input type="hidden" name="componenteId" value="'+componentePilar.componenteId+'"/>'+celula.innerHTML;
	console.log('componentePilar.pilarId: '+componentePilar.pilarId);
	celula.innerHTML = '<input type="hidden" name="pilarId" value="'+componentePilar.pilarId+'"/>'+celula.innerHTML;
	console.log('componentesPilar[order].id: '+componentesPilar[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+componentesPilar[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	console.log('componentesPilar[order].tipoMedia: '+componentesPilar[order].tipoMedia);
	celula.innerText = componentesPilar[order].tipoMedia;
	celula.innerHTML = '<input type="hidden" value="'+componentesPilar[order].tipoMediaId+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('componentesPilar[order].pesoPadrao: '+componentesPilar[order].pesoPadrao);
	celula.innerText = componentesPilar[order].pesoPadrao;
	celula = row.childNodes[3];
	console.log('componentesPilar[order].sonda: '+componentesPilar[order].sonda);
	celula.innerText = componentesPilar[order].sonda;
}

function showDeleteComponentePilarForm(e){
	console.log('showDeleteComponentePilarForm');
	var deleteComponentePilarForm = document.getElementById('delete-componente-pilar-form');
	deleteComponentePilarForm.style.display = 'block';
	componente_pilar_tobe_deleted = e;
}

function deleteComponentePilar() {
	console.log('deleteComponentePilar');
	var order = componente_pilar_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newComponentesPilar = [];
	let tbl = componente_pilar_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<componentesPilar.length;i++){
		if(i != order){
			newComponentesPilar.push(componentesPilar[i]);
		}
	}
	componentesPilar = newComponentesPilar;
	var deleteComponentePilarForm = document.getElementById('delete-componente-pilar-form');
	deleteComponentePilarForm.style.display = 'none';
}