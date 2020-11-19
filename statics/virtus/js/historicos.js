class Historico {
	constructor(id, entidadeId, cicloId, pilarId, componenteId, elementoId, nota, metodo, peso, autorId, autorNome, alteradoEm, motivacao, tipoAlteracao) {
		this.id = id;
		this.entidadeId = entidadeId;
		this.cicloId = cicloId;
		this.pilarId = pilarId;
		this.componenteId = componenteId;
		this.elementoId = elementoId;
		this.nota = nota;
		this.metodo = metodo;
		this.peso = peso;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.alteradoEm = alteradoEm;
		this.motivacao = motivacao;
		this.tipoAlteracao = tipoAlteracao;
	}
}

function loadHistoricos(btn){
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
	let componenteId = valores[4];
	let tipoNotaId = valores[5];
	let elementoId = valores[6];
	xmlhttp.open("GET","/loadHistoricos?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&componenteId="+componenteId+"&tipoNotaId="+tipoNotaId+"&elementoId="+elementoId,true);
	xmlhttp.send();
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
	document.getElementById('mot_text').value=e.parentNode.parentNode.childNodes[5].childNodes[0].value;
}