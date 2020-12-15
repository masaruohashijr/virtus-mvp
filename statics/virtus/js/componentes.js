function editComponente(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
	document.getElementById('formulario-componente-create').reset();
	document.getElementById('formulario-componente-edit').reset();
    var componenteId = e.parentNode.parentNode.childNodes[3].innerText;
    var componenteNome = e.parentNode.parentNode.childNodes[5].innerText;
    var componenteDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var componenteReferencia = e.parentNode.parentNode.childNodes[9].innerText;
    var componentePGA = e.parentNode.parentNode.childNodes[11].innerText;
	document.getElementById('ComponenteIdForUpdate').value = componenteId;
    document.getElementById('ComponenteNomeForUpdate').value = componenteNome;
    document.getElementById('ComponenteDescricaoForUpdate').value = componenteDescricao;
    document.getElementById('ComponenteReferenciaForUpdate').value = componenteReferencia;
	if(componentePGA=="Sim"){
	    document.getElementById('ComponentePGAForUpdate').checked = true;
	}
	document.getElementById('ComponenteNomeForUpdate').focus();
	loadElementosByComponenteId(componenteId);
	loadTiposNotaByComponenteId(componenteId);
}

function deleteComponente(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var componenteId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ComponenteIdToDelete').value = componenteId;
}

function loadElementosByComponenteId(componenteId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var elementosComponenteJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-elementos-componente-edit")
				elementosComponente = [];
				for(order = 0;elementosComponenteJson != null && order<elementosComponenteJson.length;order++){
					elementosComponente[order]=elementosComponenteJson[order];
					console.log(elementosComponente[order].tipoNotaId);
					console.log(elementosComponente[order].tipoNotaNome);
					console.log(elementosComponente[order].pesoPadrao);
					addElementoComponenteRow("table-elementos-componente-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadElementosByComponenteId?componenteId="+componenteId,true);
	xmlhttp.send(); 
}

function loadTiposNotaByComponenteId(componenteId){
	console.log('loadTiposNotaByComponenteId');
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var tiposNotaJson = JSON.parse(xmlhttp.responseText);
				tipos = [];
				for(order = 0;tiposNotaJson != null && order<tiposNotaJson.length;order++){
					tipos[order]=tiposNotaJson[order];
					console.log("TipoNota_"+tiposNotaJson[order].Id);
					document.getElementById("TipoNota_"+tiposNotaJson[order].tipoNotaId).value=tiposNotaJson[order].pesoPadrao;
				}
			}
	}
	xmlhttp.open("GET","/loadTiposNotaByComponenteId?componenteId="+componenteId,true);
	xmlhttp.send(); 
}

function openCreateComponente(){
	document.getElementById('create-form').style.display='block'; 
	document.getElementById('NomeComponenteForInsert').focus();
}