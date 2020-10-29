function resetCicloForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}

function iniciarCiclo(e) {
	resetCicloForms();
    var editForm = document.getElementById('edit-iniciar-form');
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
	loadPilaresByCicloId(cicloId);
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

function loadPilaresByCicloId(cicloId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var pilaresJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-pilar-ciclo-edit", pilaresCiclo)
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