var integrante_tobe_deleted;
	
class Integrante {
	constructor(order, id, entidadeId, cicloId, usuarioId, usuarioNome, usuarioPerfil, iniciaEm, terminaEm, autorId, autorNome, criadoEm, idVersaoOrigem, statusId, cStatus) {
		this.order = order;
		this.id = id;
		this.entidadeId = entidadeId;
		this.cicloId = cicloId;
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

function criarIntegrante(){
	console.log('criarIntegrante');
	let campoSelect = document.getElementById('UsuarioForInsert');
	let usuarioId = 0;
	let usuarioNome = '';
	let usuarioPerfil = '';
	for(n=0;n<campoSelect.options.length;n++){
		if(campoSelect.options[n].selected){
			usuarioId = campoSelect.options[n].value;
			text = campoSelect.options[n].text;
			usuarioPerfil = text.split("[")[1].trim(); 
			usuarioPerfil = usuarioPerfil.substr(0,usuarioPerfil.length-1);
			usuarioNome = text.split("[")[0].trim();
			break;
		}
	}
	let erros = '';
	if(campoSelect.selectedIndex==0){
		erros += 'Falta vincular o integrante.\n';
		alert(erros);
		return;
	}
	let iniciaEm = formatarData(document.getElementById('IniciaEmForInsert').value);
	let terminaEm = formatarData(document.getElementById('TerminaEmForInsert').value);
	integranteId = getMaxId(integrantes);
	integrante = new Integrante(0, integranteId, 0, 0, usuarioId, usuarioNome, usuarioPerfil, iniciaEm, terminaEm, '', '', '', '', '', '', '');
	integrantes.push(integrante);
	addIntegranteRow("table-integrantes-"+contexto);
	limparCamposIntegranteForm();
	document.getElementById('create-integrante-form').style.display='none';
}

function addIntegranteRow(tableID) {
	console.log('addIntegranteRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = integrantes.length-1;
	integrante = integrantes[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(integrante.usuarioNome);
	let json = JSON.stringify(integrante);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="integrante'+integrante.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="usuarioId" value="'+integrante.usuarioId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+integrante.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="cicloId" value="'+integrante.cicloId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+integrante.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Perfil
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(integrante.usuarioPerfil);
	newCell.appendChild(newText);
	// Inicio Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(integrante.iniciaEm);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(integrante.terminaEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+integrante.autorId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" value="'+integrante.criadoEm+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editIntegrante(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteIntegranteForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function limparCamposIntegranteForm(){
	console.log('limparCamposIntegranteForm');
	document.getElementById('formulario-integrante-create').reset()
	// document.getElementById('formulario-integrante-edit').reset()
}

function editEquipe(e) {
	console.log('editEquipe');
	limparCamposIntegranteForm();
	let editEquipeForm = document.getElementById('edit-designar-equipe-form');
	editEquipeForm.style.display = 'block';
	let entidadeId = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
	let entidadeNome = e.parentNode.parentNode.childNodes[5].innerText;
	let campoSelect = e.parentNode.parentNode.childNodes[9].childNodes[1];
	let cicloId = campoSelect.options[campoSelect.selectedIndex].value;
	let cicloNome = campoSelect.options[campoSelect.selectedIndex].text;
	document.getElementById("EntidadeIdForUpdate").value = entidadeId;
	document.getElementById("CicloIdForUpdate").value = cicloId;
	document.getElementById("EntidadeNomeForUpdate").value = entidadeNome;
	document.getElementById("CicloNomeForUpdate").value = cicloNome;
	loadSupervisorByEntidadeCicloId(entidadeId, cicloId);
	loadIntegrantesByEntidadeCicloId(entidadeId, cicloId);
}

function editIntegrante(e) {
	console.log('editIntegrante');
	limparCamposIntegranteForm();
	var editIntegranteForm = document.getElementById('edit-integrante-form');
	editIntegranteForm.style.display = 'block';
	var linha = e.parentNode.parentNode;
	var order = linha.childNodes[0].childNodes[0].value;
	var usuarioId = linha.childNodes[0].childNodes[4].value;
	var entidadeId = linha.childNodes[0].childNodes[3].value;
	var cicloId = linha.childNodes[0].childNodes[2].value;
	// var tipoMedia = linha.childNodes[1].innerText;
	var iniciaEm = linha.childNodes[2].innerText;
	console.log(iniciaEm);
	var terminaEm = linha.childNodes[3].innerText;
	console.log(terminaEm);
	// Atribuindo os valores de edit-item-form
	document.getElementById('UsuarioMEForUpdate').value=usuarioId;
	document.getElementById('Order-MEForUpdate').value=order;
	document.getElementById('EntidadeId-MEForUpdate').value=entidadeId;
	document.getElementById('CicloId-MEForUpdate').value=cicloId;
	document.getElementById('IniciaEmForUpdate').value=iniciaEm;
	document.getElementById('TerminaEmForUpdate').value=terminaEm;
}


function updateIntegrante() {
	console.log('updateIntegrante');
	let id = document.getElementById('Id-MEForUpdate').value;
	let order = document.getElementById('Order-MEForUpdate').value;
	let entidadeId = document.getElementById('EntidadeId-MEForUpdate').value;
	let cicloId = document.getElementById('CicloId-MEForUpdate').value;
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
	let iniciaEm = formatarData(document.getElementById('IniciaEmMEForUpdate').value);
	let terminaEm = formatarData(document.getElementById('TerminaEmMEForUpdate').value);
	integrante = new Integrante(order, id, entidadeId, cicloId, usuarioId, usuarioNome, '', iniciaEm, terminaEm, '', '', '', '', '', '');
	integrantes[order] = integrante;
	updateIntegranteRow("table-integrantes-"+contexto,order);
	limparCamposIntegranteForm();
	document.getElementById('edit-integrante-form').style.display='none';
}

function updateIntegranteRow(tableID, order){
	console.log('updateIntegrante');
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
	console.log(integrantes[order].usuarioNome);
	celula.innerText = integrantes[order].usuarioNome;
	let json = JSON.stringify(integrantes[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="integrante'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('integrante.cicloId: '+integrante.cicloId);
	celula.innerHTML = '<input type="hidden" name="usuarioId" value="'+integrante.usuarioId+'"/>'+celula.innerHTML;
	console.log('integrante.entidadeId: '+integrante.escritorioId);
	celula.innerHTML = '<input type="hidden" name="escritorioId" value="'+integrante.escritorioId+'"/>'+celula.innerHTML;
	console.log('integrantes[order].id: '+integrantes[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+integrantes[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('integrantes[order].iniciaEm: '+integrantes[order].iniciaEm);
	celula.innerText = integrantes[order].iniciaEm;
	celula = row.childNodes[3];
	console.log('integrantes[order].terminaEm: '+integrantes[order].terminaEm);
	celula.innerText = integrantes[order].terminaEm;
}

function showDeleteIntegranteForm(e){
	console.log('showDeleteIntegranteForm');
	var deleteIntegranteForm = document.getElementById('delete-integrante-form');
	deleteIntegranteForm.style.display = 'block';
	integrante_tobe_deleted = e;
}

function deleteIntegrante() {
	console.log('deleteIntegrante');
	var order = integrante_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newIntegrantes = [];
	let tbl = integrante_tobe_deleted.parentNode.parentNode.parentNode;
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
	for(i=0;i<integrantes.length;i++){
		if(i != order){
			newIntegrantes.push(integrantes[i]);
		}
	}
	integrantes = newIntegrantes;
	var deleteIntegranteForm = document.getElementById('delete-integrante-form');
	deleteIntegranteForm.style.display = 'none';
}

function loadIntegrantesByEntidadeCicloId(entidadeId, cicloId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var integrantesJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-integrantes-edit")
				integrantes = [];
				for(order = 0;integrantesJson != null && order<integrantesJson.length;order++){
					integrantes[order]=integrantesJson[order];
					removerOpcoesUsuarios("UsuarioForInsert",integrantes[order].usuarioId);
					addIntegranteRow("table-integrantes-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadIntegrantesByEntidadeIdCicloId?entidadeId="+entidadeId+"&cicloId="+cicloId,true);
	xmlhttp.send();
}

function loadSupervisorByEntidadeCicloId(entidadeId, cicloId){
	console.log('loadSupervisorByEntidadeCicloId');
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				let supervisor = JSON.parse(xmlhttp.responseText);
				document.getElementById('SupervisorEquipeForUpdate').value = supervisor.id; 
			}
	}
	xmlhttp.open("GET","/loadSupervisorByEntidadeIdCicloId?entidadeId="+entidadeId+"&cicloId="+cicloId,true);
	xmlhttp.send();
}

function removerOpcoesUsuarios(nomeCampo,usuarioId){
	console.log('removerOpcoesUsuarios');
	let campoSelect = document.getElementById(nomeCampo);
	for(i=0;i<campoSelect.length;i++){
		if(campoSelect.options[i].value == usuarioId){
			campoSelect.remove(i);
		}
	}
}