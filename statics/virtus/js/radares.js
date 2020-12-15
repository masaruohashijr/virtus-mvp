function resetRadarForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}

function openCreateRadar(btn){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NomeRadarForInsert').focus();
}

function deleteRadar(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var radarId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('RadarIdForDelete').value = radarId;
}

function editRadar(e){
	resetRadarForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var radarId = e.parentNode.parentNode.childNodes[3].innerText;
    var radarNome = e.parentNode.parentNode.childNodes[5].innerText;
    var radarDataRadar = e.parentNode.parentNode.childNodes[7].innerText;
    var radarDescricao = e.parentNode.parentNode.childNodes[9].innerText;
    var radarReferencia = e.parentNode.parentNode.childNodes[11].innerText;
    var radarAutor = e.parentNode.parentNode.childNodes[13].innerText;
    var radarCriadoEm = e.parentNode.parentNode.childNodes[15].innerText;
    var radarStatus = e.parentNode.parentNode.childNodes[17].innerText;
	document.getElementById('RadarIdForUpdate').value = radarId;
    document.getElementById('NomeRadarForUpdate').value = radarNome;
    document.getElementById('DataRadarForUpdate').value = formatarData(radarDataRadar);
    document.getElementById('DescricaoRadarForUpdate').value = radarDescricao;
    document.getElementById('ReferenciaRadarForUpdate').value = radarReferencia;
	document.getElementById('AuthorNameRadarForUpdate').value = radarAutor;
    document.getElementById('CriadoEmRadarForUpdate').value = radarCriadoEm;
    document.getElementById('StatusRadarForUpdate').value = radarStatus;
    document.getElementById('NomeRadarForUpdate').focus();
	loadAnotacoesRadaresByRadarId(radarId);	
}

function loadAnotacoesRadaresByRadarId(radarId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var anotacoesRadarJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-anotacoes-radar-edit")
				anotacoesRadar = [];
				for(order = 0;anotacoesRadarJson != null && order<anotacoesRadarJson.length;order++){
					anotacoesRadar[order]=anotacoesRadarJson[order];
					addAnotacaoRadarRow("table-anotacoes-radar-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadAnotacoesRadaresByRadarId?radarId="+radarId,true);
	xmlhttp.send();
}
