class AnotacaoRadar{
	constructor(order, id, radarId, entidadeId, anotacaoId, observacoes, registroEmAta, autor, criadoEm, status){
		this.order = order;		
		this.id = id;		
		this.radarId = radarId;		
		this.entidadeId = entidadeId;		
		this.anotacaoId = anotacaoId;		
		this.observacoes = observacoes;		
		this.registroEmAta = registroEmAta;		
		this.autor = autor;		
		this.criadoEm = criadoEm;		
		this.status = status;		
	}
}

function criarAnotacaoRadar(){
	console.log('criarAnotacaoRadar');
	let selectEFPC = document.getElementById('EFPCForInsert');
	let entidadeId = selectEFPC.value;
	let selectAnotacoes = document.getElementById('AnotacoesForInsert');
	let anotacaoId = 0;
	for(n=0;n<selectAnotacoes.options.length;n++){
		if(selectAnotacoes.options[n].selected){
			anotacaoId = selectAnotacoes.options[n].value;
			anotacaoAssunto = selectAnotacoes.options[n].text;
			break;
		}
	}
	let erros = '';
	if(selectAnotacoes.selectedIndex==0){
		erros += 'Falta informar a anotacao.\n';
		alert(erros);
		return;
	}
	let observacoes = document.getElementById('ObservacoesForInsert').value;
	let registroEmAta = document.getElementById('RegistroEmAtaForInsert').value;
	console.log("variável observações: "+observacoes);
	anotacaoRadarId = getMaxId(anotacoesRadar);
	anotacaoRadar = new AnotacaoRadar(0, anotacaoRadarId, 0, entidadeId, anotacaoId, observacoes, registroEmAta,'','','');
	anotacoesRadar.push(anotacaoRadar);
	console.log(contexto);
	addAnotacaoRadarRow("table-anotacoes-radar-"+contexto);
	//limparCamposAnotacaoRadarForm();
	document.getElementById(contexto+'-anotacao-radar-form').style.display='none';
}

function addAnotacaoRadarRow(tableID) {
	console.log('addAnotacaoRadarRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = anotacoesRadar.length-1;
	anotacaoRadar = anotacoesRadar[order];
	// Sigla
	let newCell = newRow.insertCell(0);
	let anotacao = anotacoesMap.get(anotacaoRadar.anotacaoId);
	let newText = document.createTextNode(anotacao.entidadeSigla);
	newCell.appendChild(newText);
	let json = JSON.stringify(anotacaoRadar);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+anotacaoRadar.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="anotacaoRadar'+anotacaoRadar.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="anotacaoId" value="'+anotacaoRadar.anotacaoId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="radarId" value="'+anotacaoRadar.radarId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+anotacaoRadar.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	newCell.style = "vertical-align: middle; text-align: left";
	// Assunto
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(anotacao.assunto);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="registroEmAta" value="'+anotacaoRadar.registroEmAta+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="observacoes" value="'+anotacaoRadar.observacoes+'"/>'+newCell.innerHTML;
	// Risco
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(anotacao.risco);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+anotacaoRadar.risco+'"/>'+newCell.innerHTML;
	// Tendência
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(anotacao.tendencia);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	// Autor
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(anotacaoRadar.autor);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	// Criado Em
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(anotacaoRadar.criadoEm);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	// Status
	newCell = newRow.insertCell(6);
	newText = document.createTextNode(anotacaoRadar.status);
	newCell.style = "vertical-align: middle; text-align: left";
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(7);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editAnotacaoRadar(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteAnotacaoRadarForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function editAnotacaoRadar(e) {
	console.log('editAnotacaoRadar');
	let editAnotacaoRadarForm = document.getElementById('edit-anotacao-radar-form');
	editAnotacaoRadarForm.style.display = 'block';
	let linha = e.parentNode.parentNode;
	let order = linha.childNodes[0].childNodes[0].value;
	let id = linha.childNodes[0].childNodes[1].value;
	let radarId = linha.childNodes[0].childNodes[2].value;
	let anotacaoId = linha.childNodes[0].childNodes[3].value;
	let entidadeId = linha.childNodes[0].childNodes[5].value;
	let observacoes = linha.childNodes[1].childNodes[0].value;
	let registroAta = linha.childNodes[1].childNodes[1].value;
	// Atribuindo os valores de edit-item-form
	document.getElementById('IdForUpdate').value=id;
	document.getElementById('OrderForUpdate').value=order;
	let efpc = document.getElementById('EFPCForUpdate');
	efpc.value=entidadeId;
	preencherAnotacoesSelect(efpc,'edit');
	document.getElementById('RadarIdForUpdate').value=radarId;
	document.getElementById('AnotacoesForUpdate').value=anotacaoId;
	document.getElementById('ObservacoesForUpdate').value=observacoes;
	document.getElementById('RegistroEmAtaForUpdate').value=registroAta;
}

function updateAnotacaoRadar() {
	console.log('updateAnotacaoRadar');
	let id = document.getElementById('IdForUpdate').value;
	let order = document.getElementById('OrderForUpdate').value;
	let selectEFPC = document.getElementById('EFPCForUpdate');
	let radarId = document.getElementById('RadarIdForUpdate').value;
	let entidadeId = selectEFPC.value;
	let selectAnotacoes = document.getElementById('AnotacoesForUpdate');
	let anotacaoId = 0;
	for(n=0;n<selectAnotacoes.options.length;n++){
		if(selectAnotacoes.options[n].selected){
			anotacaoId = selectAnotacoes.options[n].value;
			anotacaoAssunto = selectAnotacoes.options[n].text;
			break;
		}
	}
	let erros = '';
	if(selectAnotacoes.selectedIndex==0){
		erros += 'Falta informar a anotacao.\n';
		alert(erros);
		return;
	}
	let observacoes = document.getElementById('ObservacoesForUpdate');
	let registroEmAta = document.getElementById('RegistroEmAtaForUpdate');
	anotacaoRadar = new AnotacaoRadar(order, id, radarId, entidadeId, anotacaoId, observacoes, registroEmAta, '', '', '');
	anotacoesRadar[order] = anotacaoRadar;
	updateAnotacaoRadarRow("table-anotacoes-radar-"+contexto,order);
	//limparCamposAnotacaoRadarForm();
	alert(document.getElementById(contexto+'-anotacao-radar-form').id);
	document.getElementById('edit-anotacao-radar-form').style.display='none';
}

function updateAnotacaoRadarRow(tableID, order){
	console.log('updateAnotacaoRadarRow');
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
	console.log(anotacoesRadar[order].assunto);
	let anotacao = anotacoesMap.get(anotacaoRadar.anotacaoId);
	celula.innerText = anotacao.entidadeSigla;
	let json = JSON.stringify(anotacoesRadar[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	console.log(json);
	celula.innerHTML = '<input type="hidden" name="anotacaoRadar'+order+'" value="'+json+'"/>'+celula.innerHTML;
	console.log('anotacaoRadar.anotacaoId: '+anotacaoRadar.anotacaoId);
	celula.innerHTML = '<input type="hidden" name="anotacaoId" value="'+anotacaoRadar.anotacaoId+'"/>'+celula.innerHTML;
	console.log('anotacaoRadar.radarId: '+anotacaoRadar.radarId);
	celula.innerHTML = '<input type="hidden" name="radarId" value="'+anotacaoRadar.radarId+'"/>'+celula.innerHTML;
	console.log('anotacoesRadar[order].id: '+anotacoesRadar[order].id);
	celula.innerHTML = '<input type="hidden" name="id" value="'+anotacoesRadar[order].id+'"/>'+celula.innerHTML;
	console.log('order: '+order);
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	celula = row.childNodes[1];
	console.log('anotacoesRadar[order].assunto: '+anotacao.assunto);
	celula.innerText = anotacao.assunto;
	celula.innerHTML = '<input type="hidden" name="registroEmAta" value="'+anotacoesRadar[order].registroEmAta+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="observacoes" value="'+anotacoesRadar[order].observacoes+'"/>'+celula.innerHTML;
	celula = row.childNodes[2];
	console.log('anotacoesRadar[order].risco: '+anotacao.risco);
	celula.innerText = anotacao.risco;
	celula = row.childNodes[3];
	console.log('anotacoesRadar[order].tendencia: '+anotacao.tendencia);
	celula.innerText = anotacao.tendencia;
}

function showDeleteAnotacaoRadarForm(e){
	console.log('showDeleteAnotacaoRadarForm');
	let deleteAnotacaoRadarForm = document.getElementById('delete-anotacao-radar-form');
	deleteAnotacaoRadarForm.style.display = 'block';
	anotacao_radar_tobe_deleted = e;
}