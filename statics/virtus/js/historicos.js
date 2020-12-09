class Historico {
	constructor(id, entidadeId, cicloId, pilarId, componenteId, elementoId, nota, metodo, peso, auditorAnteriorId, auditorNovoId, autorId, autorNome, alteradoEm, motivacao, tipoAlteracao) {
		this.id = id;
		this.entidadeId = entidadeId;
		this.cicloId = cicloId;
		this.pilarId = pilarId;
		this.componenteId = componenteId;
		this.elementoId = elementoId;
		this.nota = nota;
		this.metodo = metodo;
		this.peso = peso;
		this.auditorAnteriorId = auditorAnteriorId;
		this.auditorNovoId = auditorNovoId;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.alteradoEm = alteradoEm;
		this.motivacao = motivacao;
		this.tipoAlteracao = tipoAlteracao;
	}
}

function addHistoricoRemocaoRow(tableID) {
	console.log('addHistoricoRemocaoRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = historicos.length-1;
	historico = historicos[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(auditoresMap.get(historico.auditorAnteriorId));
	newCell.appendChild(newText);
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(auditoresMap.get(historico.auditorNovoId));
	newCell.appendChild(newText);
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(historico.autorNome);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(historico.alteradoEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+historico.motivacao+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Motivo
	let btnMotivo = document.createElement('input');
	btnMotivo.type = "button";
	btnMotivo.className = "w3-btn w3-teal";
	btnMotivo.style = "margin-right: 10px";
	btnMotivo.value = "Motivo";
	btnMotivo.onclick = function() {openMotivoRemocao(btnMotivo)};
	newCell.appendChild(btnMotivo);
}

function addHistoricoRow(tableID) {
	console.log('addHistoricoRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = historicos.length-1;
	historico = historicos[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(historico.tipoAlteracao);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(historico.nota);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(historico.metodo);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(historico.peso);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(historico.autorNome);
	newCell.appendChild(newText);
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(historico.alteradoEm);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" value="'+historico.motivacao+'"/>'+newCell.innerHTML;
	// Botões
	newCell = newRow.insertCell(6);
	// Botão Motivo
	let btnMotivo = document.createElement('input');
	btnMotivo.type = "button";
	btnMotivo.className = "w3-btn w3-teal";
	btnMotivo.style = "margin-right: 10px";
	btnMotivo.value = "Motivo";
	btnMotivo.onclick = function() {openMotivo(btnMotivo)};
	newCell.appendChild(btnMotivo);
}

function openMotivo(e){
	document.getElementById('motivo-form').style.display='block';
	let tipoAlteracao = e.parentNode.parentNode.childNodes[0].innerText;
	let nota = e.parentNode.parentNode.childNodes[1].innerText;
	let metodo = e.parentNode.parentNode.childNodes[2].innerText;
	let peso = e.parentNode.parentNode.childNodes[3].innerText;
	let autor = e.parentNode.parentNode.childNodes[4].innerText;
	let alteradoEm = e.parentNode.parentNode.childNodes[5].innerText;
	let motivacao = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
	document.getElementById('mot_hist_alteracao').value=tipoAlteracao;
	document.getElementById('mot_hist_nota').value=nota;
	document.getElementById('mot_hist_metodo').value=metodo;
	document.getElementById('mot_hist_peso').value=peso;
	document.getElementById('mot_hist_autor').value=autor;
	document.getElementById('mot_hist_alterado_em').value=alteradoEm;
	document.getElementById('mot_text').value=motivacao;
}

function openMotivoRemocao(e){
	document.getElementById('motivo-remocao-form').style.display='block';
	let tipoAlteracao = 'Remoção';
	let de = e.parentNode.parentNode.childNodes[0].innerText;
	let para = e.parentNode.parentNode.childNodes[1].innerText;
	let autor = e.parentNode.parentNode.childNodes[2].innerText;
	let alteradoEm = e.parentNode.parentNode.childNodes[3].innerText;
	let motivacao = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
	document.getElementById('mot_hist_rem_alteracao').value=tipoAlteracao;
	document.getElementById('mot_hist_rem_de').value=de;
	document.getElementById('mot_hist_rem_para').value=para;
	document.getElementById('mot_hist_rem_autor').value=autor;
	document.getElementById('mot_hist_rem_alterado_em').value=alteradoEm;
	document.getElementById('mot_rem_text').value=motivacao;
}

function openHistElemento(btn){
	btn.disabled = true;
	let entidadeId = btn.name.split("_")[1];
	let cicloId = btn.name.split("_")[2];
	let pilarId = btn.name.split("_")[3];
	let planoId = btn.name.split("_")[4];
	let componenteId = btn.name.split("_")[5];
	let tipoNotaId = btn.name.split("_")[6];
	let elementoId = btn.name.split("_")[7];
	document.getElementById('hist-elemento-form').style.display='block';
	document.getElementById("histEFPC").value = entidadesMap.get(entidadeId);
	document.getElementById("histCiclo").value = ciclosMap.get(cicloId);
	document.getElementById("histPilar").value = pilaresMap.get(pilarId);
	document.getElementById("histPlano").value = planosMap.get(planoId);
	document.getElementById("histComponente").value = componentesMap.get(componenteId);
	document.getElementById("histTipoNota").value = tiposNotasMap.get(tipoNotaId);
	document.getElementById("histElemento").value = elementosMap.get(elementoId);
	loadHistoricosElemento(btn);
	return false;
}

function openHistPilar(btn){
	//btn.disabled = true;
	let entidadeId = btn.name.split("_")[1];
	let cicloId = btn.name.split("_")[2];
	let pilarId = btn.name.split("_")[3];
	document.getElementById('hist-pilar-form').style.display='block';
	document.getElementById("histPilarEFPC").value = entidadesMap.get(entidadeId);
	document.getElementById("histPilarCiclo").value = ciclosMap.get(cicloId);
	document.getElementById("histPilarPilar").value = pilaresMap.get(pilarId);
	loadHistoricosPilar(btn);
	return false;
}

function openHistComponente(btn){
	//btn.disabled = true;
	let entidadeId = btn.name.split("_")[1];
	let cicloId = btn.name.split("_")[2];
	let pilarId = btn.name.split("_")[3];
	let componenteId = btn.name.split("_")[4];
	document.getElementById('hist-componente-form').style.display='block';
	document.getElementById("histComponenteEFPC").value = entidadesMap.get(entidadeId);
	document.getElementById("histComponenteCiclo").value = ciclosMap.get(cicloId);
	document.getElementById("histComponentePilar").value = pilaresMap.get(pilarId);
	document.getElementById("histComponenteComponente").value = componentesMap.get(componenteId);
	loadHistoricosComponente(btn);
	return false;
}

function loadHistoricosElemento(btn){
	var xmlhttp;
	let valores = btn.name.split("_");
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var historicosJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-historicos-edit")
				historicos = [];
				for(i = 0;historicosJson != null && i<historicosJson.length;i++){
					historicos[i]=historicosJson[i];
					addHistoricoRow("table-historicos-edit");
				}
			}
	}
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	let planoId = valores[4];
	let componenteId = valores[5];
	let tipoNotaId = valores[6];
	let elementoId = valores[7];
	xmlhttp.open("GET","/loadHistoricosElemento?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&planoId="+planoId+"&componenteId="+componenteId+"&tipoNotaId="+tipoNotaId+"&elementoId="+elementoId,true);
	xmlhttp.send();
}

function loadHistoricosComponente(btn){
	var xmlhttp;
	let valores = btn.name.split("_");
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var historicosJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-historicos-componente-edit")
				historicos = [];
				for(i = 0;historicosJson != null && i<historicosJson.length;i++){
					historicos[i]=historicosJson[i];
					addHistoricoRemocaoRow("table-historicos-componente-edit");
				}
			}
	}
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	let componenteId = valores[4];
	xmlhttp.open("GET","/loadHistoricosComponente?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&componenteId="+componenteId,true);
	xmlhttp.send();
}

function loadHistoricosPilar(btn){
	var xmlhttp;
	let valores = btn.name.split("_");
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var historicosJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-historicos-pilar-edit")
				historicos = [];
				for(i = 0;historicosJson != null && i<historicosJson.length;i++){
					historicos[i]=historicosJson[i];
					addHistoricoRow("table-historicos-pilar-edit");
				}
			}
	}
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	xmlhttp.open("GET","/loadHistoricosPilar?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId,true);
	xmlhttp.send();
}