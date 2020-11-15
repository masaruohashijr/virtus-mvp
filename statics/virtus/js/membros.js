var membro_tobe_deleted;
	
class Membro {
	constructor(order, id, escritorioId, usuarioId, usuarioNome, usuarioPerfil, iniciaEm, terminaEm, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.escritorioId = escritorioId;
		this.usuarioId = usuarioId;
		this.usuarioNome = usuarioNome;
		this.usuarioPerfil = usuarioPerfil;
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

function criarMembro(){
	console.log('criarMembro');
	let campoSelect = document.getElementById('UsuarioMEForInsert');
	let usuarioId = 0;
	let usuarioNome = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			usuarioId = campoSelect.options[n].value;
			usuarioNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o membro.\n';
		alert(erros);
		return;
	}
	let iniciaEm = formatarData(document.getElementById('IniciaEmMEForInsert').value);
	let terminaEm = formatarData(document.getElementById('TerminaEmMEForInsert').value);
	membroId = getMaxId(membros);
	membro = new Membro(0, membroId, 0, usuarioId, usuarioNome, '', iniciaEm, terminaEm, '', '', '', '', '', '');
	membros.push(membro);
	addMembroRow("table-membros-"+contexto);
	limparCamposMembroForm();
	document.getElementById('create-membro-form').style.display='none';
}

function addMembroRow(tableID) {
	console.log('addMembroRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = membros.length-1;
	membro = membros[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(membro.usuarioNome);
	let json = JSON.stringify(membro);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="membro'+membro.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="usuarioId" value="'+membro.usuarioId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="escritorioId" value="'+membro.escritorioId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+membro.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Perfil
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(membro.usuarioPerfil);
	newCell.appendChild(newText);
	// Inicio Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(membro.iniciaEm);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(membro.terminaEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+membro.autorId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" value="'+membro.criadoEm+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editMembro(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteMembroForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposMembroForm(){
	console.log('limparCamposMembroForm');
	document.getElementById('formulario-membro-create').reset()
	document.getElementById('formulario-membro-edit').reset()
}

function editMembro(e) {
	console.log('editMembro');
	var editMembroForm = document.getElementById('edit-membro-form');
	editMembroForm.style.display = 'block';
	var linha = e.parentNode.parentNode;
	var order = linha.childNodes[0].childNodes[0].value;
	var membroId = linha.childNodes[0].childNodes[1].value;
	var escritorioId = linha.childNodes[0].childNodes[2].value;
	var usuarioId = linha.childNodes[0].childNodes[3].value;
	// var tipoMedia = linha.childNodes[1].innerText;
	var iniciaEm = linha.childNodes[2].innerText;
	console.log(iniciaEm);
	var terminaEm = linha.childNodes[3].innerText;
	console.log(terminaEm);
	// Atribuindo os valores de edit-item-form
	document.getElementById('Id-MEForUpdate').value=membroId;
	document.getElementById('UsuarioMEForUpdate').value=usuarioId;
	document.getElementById('Order-MEForUpdate').value=order;
	document.getElementById('EscritorioId-MEForUpdate').value=escritorioId;
	document.getElementById('IniciaEmMEForUpdate').value=formatarData(iniciaEm);
	document.getElementById('TerminaEmMEForUpdate').value=formatarData(terminaEm);
}


function updateMembro() {
	console.log('updateMembro');
	var membroId = document.getElementById('Id-MEForUpdate').value;
	console.log('membroId: '+membroId);
	var order = document.getElementById('Order-MEForUpdate').value;
	var escritorioId = document.getElementById('EscritorioId-MEForUpdate').value;
	let campoSelect = document.getElementById('UsuarioMEForUpdate');
	let usuarioId = 0;
	let usuarioNome = '';
	console.log(campoSelect.options.length);
	console.log(campoSelect.options.selectedIndex);
	for(n=0;n<campoSelect.options.length;n++){
		console.log("n: "+n);
		console.log(campoSelect.options[n].selected);
		console.log(campoSelect.selectedIndex);
		if(campoSelect.options[n].selected){
			usuarioId = campoSelect.options[n].value;
			usuarioNome = campoSelect.options[n].text;
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o usuário.\n';
		alert(erros);
		return;
	}
	let iniciaEm = document.getElementById('IniciaEmMEForUpdate').value;
	let terminaEm = document.getElementById('TerminaEmMEForUpdate').value;
	membro = new Membro(order, membroId, escritorioId, usuarioId, usuarioNome, '', formatarData(iniciaEm), formatarData(terminaEm), '', '', '', '', '', '');
	membros[order] = membro;
	updateMembroRow("table-membros-"+contexto,order);
	limparCamposMembroForm();
	document.getElementById('edit-membro-form').style.display='none';
}

function updateMembroRow(tableID, order){
	console.log('updateMembro');
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
	console.log(membros[order].usuarioNome);
	celula.innerText = membros[order].usuarioNome;
	let json = JSON.stringify(membros[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="membro'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('membro.cicloId: '+membro.cicloId);
	celula.innerHTML = '<input type="hidden" name="usuarioId" value="'+membro.usuarioId+'"/>'+celula.innerHTML;
	console.log('membro.entidadeId: '+membro.escritorioId);
	celula.innerHTML = '<input type="hidden" name="escritorioId" value="'+membro.escritorioId+'"/>'+celula.innerHTML;
	console.log('membros[order].id: '+membros[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+membros[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('membros[order].iniciaEm: '+membros[order].iniciaEm);
	celula.innerText = membros[order].iniciaEm;
	celula = row.childNodes[3];
	console.log('membros[order].terminaEm: '+membros[order].terminaEm);
	celula.innerText = membros[order].terminaEm;
}

function showDeleteMembroForm(e){
	console.log('showDeleteMembroForm');
	var deleteMembroForm = document.getElementById('delete-membro-form');
	deleteMembroForm.style.display = 'block';
	membro_tobe_deleted = e;
}

function deleteMembro() {
	console.log('deleteMembro');
	var order = membro_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newMembros = [];
	let tbl = membro_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<membros.length;i++){
		if(i != order){
			newMembros.push(membros[i]);
		}
	}
	membros = newMembros;
	var deleteMembroForm = document.getElementById('delete-membro-form');
	deleteMembroForm.style.display = 'none';
}

