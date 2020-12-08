function resetCicloForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
	document.getElementById('formulario-iniciar').reset();
}

function iniciarCiclo(e) {
	resetCicloForms();
    var editForm = document.getElementById('edit-iniciar-ciclo-form');
    editForm.style.display = 'block';
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    var cicloNome = e.parentNode.parentNode.childNodes[5].innerText;
    var cicloDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var cicloAutor = e.parentNode.parentNode.childNodes[9].innerText;
    var cicloCriadoEm = e.parentNode.parentNode.childNodes[11].innerText;
	document.getElementById('CicloIdIniciarForUpdate').value = cicloId;
    document.getElementById('NomeCicloIniciarForUpdate').value = cicloNome;
    document.getElementById('DescricaoIniciarCicloForUpdate').value = cicloDescricao;
	document.getElementById('AuthorNameIniciarForUpdate').value = cicloAutor;
    document.getElementById('CriadoEmIniciarForUpdate').value = cicloCriadoEm;
}

function editCiclo(e) {
	resetCicloForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    var cicloNome = e.parentNode.parentNode.childNodes[5].innerText;
    var cicloDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var cicloAutor = e.parentNode.parentNode.childNodes[9].innerText;
    var cicloCriadoEm = e.parentNode.parentNode.childNodes[11].innerText;
	document.getElementById('CicloIdForUpdate').value = cicloId;
    document.getElementById('NomeCicloForUpdate').value = cicloNome;
    document.getElementById('DescricaoCicloForUpdate').value = cicloDescricao;
	document.getElementById('AuthorNameForUpdate').value = cicloAutor;
    document.getElementById('CriadoEmForUpdate').value = cicloCriadoEm;
	loadPilaresByCicloId(cicloId);
}

function deleteCiclo(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('CicloIdForDelete').value = cicloId;
}

function salvarIniciarCiclo(){
	let preenchidoIniciaEm = document.getElementById('IniciaEmForInsert').value != '';
	let preenchidoTerminaEm = document.getElementById('TerminaEmForInsert').value != '';
	if(!preenchidoTerminaEm || !preenchidoIniciaEm){
		document.getElementById("Errors").innerText = 'Falta preencher o período.';
		document.getElementById("error-message").style.display = 'block';
		if(!preenchidoIniciaEm){
			document.getElementById('IniciaEmForInsert').focus();
		} else {
			document.getElementById('TerminaEmForInsert').focus();
		}
	} else {
		document.getElementById('formulario-iniciar').submit();
	}
}

function loadPilaresByCicloId(cicloId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var pilaresJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-pilar-ciclo-edit")
				pilaresCiclo = [];
				for(order = 0;pilaresJson != null && order<pilaresJson.length;order++){
					pilaresCiclo[order]=pilaresJson[order];
					addPilarCicloRow("table-pilar-ciclo-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadPilaresByCicloId?cicloId="+cicloId,true);
	xmlhttp.send();
}

function validarPercentuais(){
	let tbl = document.getElementById("table-pilar-ciclo-"+contexto);
	let linhas = tbl.childNodes[1].childNodes;
	let row = tbl.childNodes[0];
	let total = 0;
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			console.log(linhas[y].childNodes[0]);
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				row = linhas[y];
				total = parseInt(row.childNodes[2].innerText.split(" ")[0]) + parseInt(total);
				//console.log(total);
			}
		}
	}
	
	let nomeVazio = true;
	if(contexto == 'create'){
		if (document.getElementById('NomeCicloForInsert').value != ''){
			nomeVazio = false;
		}
	} else {
		if (document.getElementById('NomeCicloForUpdate').value != ''){
			nomeVazio = false;
		}
	}
	if( nomeVazio || total != 100){
		let msg = '';
		if(nomeVazio){
			msg = 'Você deve preencher um Nome.\n\n';
		} 
		if(total != 100) {
			msg = msg+'A soma dos pesos deve ser 100%.';
		}
		document.getElementById("Errors").innerText = msg;
		document.getElementById("error-message").style.display = 'block';
		return;
	}
	document.getElementById('formulario-'+contexto).submit();
	return;
}