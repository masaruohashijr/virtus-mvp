var pilar_ciclo_tobe_deleted;
	
class PilarCiclo {
	constructor(order, id, cicloId, pilarId, pilarNome, tipoMediaId, tipoMedia, pesoPadrao, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.cicloId = cicloId;
		this.pilarId = pilarId;
		this.pilarNome = pilarNome;
		this.tipoMediaId = tipoMediaId;
		this.tipoMedia = tipoMedia;
		this.pesoPadrao = pesoPadrao;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.idVersaoOrigem = idVersaoOrigem;
		this.statusId = statusId;
		this.cStatus = cStatus;
	}
}

function criarPilarCiclo(){
	console.log('criarPilarCiclo');
	let campoSelect = document.getElementById('PilarCicloForInsert');
	let pilarId = 0;
	let pilarNome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			pilarId = campoSelect.options[n].value;
			pilarNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o pilar.\n';
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
	let pesoPadrao = document.getElementById('PesoPadraoForInsert').value;
	console.log("variável pilarNome: "+pilarNome);
	pilarCicloId = getMaxId(pilaresCiclo);
	pilarCiclo = new PilarCiclo(0, pilarCicloId, 0, pilarId, pilarNome, tipoMediaId, tipoMedia, pesoPadrao, '', '', '', '', '', '');
	pilaresCiclo.push(pilarCiclo);
	console.log(contexto);
	addPilarCicloRow("table-pilar-ciclo-"+contexto);
	limparCamposPilarCicloForm();
	document.getElementById('create-pilar-ciclo-form').style.display='none';
}

function addPilarCicloRow(tableID) {
	console.log('addPilarCicloRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = pilaresCiclo.length-1;
	pilarCiclo = pilaresCiclo[order];
	// Nome do Pilar
	let newCell = newRow.insertCell(0);
	console.log('pilarCiclo.pilarNome: '+pilarCiclo.pilarNome);
	let newText = document.createTextNode(pilarCiclo.pilarNome);
	let json = JSON.stringify(pilarCiclo);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="pilarCiclo'+pilarCiclo.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="pilarId" value="'+pilarCiclo.pilarId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="cicloId" value="'+pilarCiclo.cicloId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+pilarCiclo.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	newCell.style = "vertical-align: middle; text-align: left";
	// Tipo de Média
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(pilarCiclo.tipoMedia);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+pilarCiclo.tipoMediaId+'"/>'+newCell.innerHTML;
	// Peso Padrão
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(pilarCiclo.pesoPadrao);
	newCell.appendChild(newText);
	// Autor
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(pilarCiclo.autorNome);
	newCell.appendChild(newText);
	// Criado Em
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(pilarCiclo.criadoEm);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(5);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editPilarCiclo(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeletePilarCicloForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposPilarCicloForm(){
	console.log('limparCamposPilarCicloForm');
	document.getElementById('formulario-pilar-ciclo-create').reset()
	document.getElementById('formulario-pilar-ciclo-edit').reset()
}

function editPilarCiclo(e) {
	console.log('editPilarCiclo');
	var editPilarCicloForm = document.getElementById('edit-pilar-ciclo-form');
	editPilarCicloForm.style.display = 'block';
	var linha = e.parentNode.parentNode;
	var order = linha.childNodes[0].childNodes[0].value;
	var id = linha.childNodes[0].childNodes[1].value;
	var cicloId = linha.childNodes[0].childNodes[2].value;
	var linha = e.parentNode.parentNode;
	var pilarId = linha.childNodes[0].childNodes[3].value;
	var tipoMediaId = linha.childNodes[1].childNodes[0].value;
	var pesoPadrao = linha.childNodes[2].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('Id-PCForUpdate').value=id;
	document.getElementById('Order-PCForUpdate').value=order;
	document.getElementById('CicloId-PCForUpdate').value=cicloId;
	document.getElementById('PilarCicloForUpdate').value=pilarId;
	document.getElementById('TipoMediaForUpdate').value=tipoMediaId;
	document.getElementById('PesoPadraoForUpdate').value=pesoPadrao.split(" ")[0];
}

function updatePilarCiclo() {
	console.log('updatePilarCiclo');
	var id = document.getElementById('Id-PCForUpdate').value;
	var order = document.getElementById('Order-PCForUpdate').value;
	var cicloId = document.getElementById('CicloId-PCForUpdate').value;
	var pesoPadrao = document.getElementById('PesoPadraoForUpdate').value;
	let campoSelect = document.getElementById('PilarCicloForUpdate');
	let pilarId = 0;
	let pilarNome = '';
	console.log('campoSelect.options.length: '+campoSelect.options.length);
	console.log('campoSelect.options.selectedIndex: '+campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			pilarId = campoSelect.options[n].value;
			pilarNome = campoSelect.options[n].text;
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
	pilarCiclo = new PilarCiclo(order, id, cicloId, pilarId, pilarNome, tipoMediaId, tipoMedia, pesoPadrao, '', '', '', '', '', '');
	// console.log('MARCA order: '+order);
	pilaresCiclo[order] = pilarCiclo;
	updatePilarCicloRow("table-pilar-ciclo-"+contexto,order);
	limparCamposPilarCicloForm();
	document.getElementById('edit-pilar-ciclo-form').style.display='none';
}

function updatePilarCicloRow(tableID, order){
	console.log('updatePilarCicloRow');
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
					// console.log("y: "+y);
					row = linhas[y];
					break;
				}
			}
		}
	}
	let celula = row.childNodes[0];
	console.log(pilaresCiclo[order].pilarNome);
	celula.innerText = pilaresCiclo[order].pilarNome;
	let json = JSON.stringify(pilaresCiclo[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="pilarCiclo'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('pilarCiclo.pilarId: '+pilarCiclo.pilarId);
	celula.innerHTML = '<input type="hidden" name="pilarId" value="'+pilarCiclo.pilarId+'"/>'+celula.innerHTML;
	console.log('pilarCiclo.cicloId: '+pilarCiclo.cicloId);
	celula.innerHTML = '<input type="hidden" name="cicloId" value="'+pilarCiclo.cicloId+'"/>'+celula.innerHTML;
	console.log('pilaresCiclo[order].id: '+pilaresCiclo[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+pilaresCiclo[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	console.log('pilaresCiclo[order].tipoMedia: '+pilaresCiclo[order].tipoMedia);
	celula.innerText = pilaresCiclo[order].tipoMedia;
	celula.innerHTML = '<input type="hidden" value="'+pilaresCiclo[order].tipoMediaId+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('pilaresCiclo[order].iniciaEm: '+pilaresCiclo[order].iniciaEm);
	celula.innerText = pilaresCiclo[order].pesoPadrao+" %";
}

function showDeletePilarCicloForm(e){
	console.log('showDeletePilarCicloForm');
	var deletePilarCicloForm = document.getElementById('delete-pilar-ciclo-form');
	deletePilarCicloForm.style.display = 'block';
	pilar_ciclo_tobe_deleted = e;
}

function deletePilarCiclo() {
	console.log('deletePilarCiclo');
	var order = pilar_ciclo_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newPilaresCiclo = [];
	let tbl = pilar_ciclo_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<pilaresCiclo.length;i++){
		if(i != order){
			newPilaresCiclo.push(pilaresCiclo[i]);
		}
	}
	pilaresCiclo = newPilaresCiclo;
	var deletePilarCicloForm = document.getElementById('delete-pilar-ciclo-form');
	deletePilarCicloForm.style.display = 'none';
}