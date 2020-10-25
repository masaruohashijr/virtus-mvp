function editComponente(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get componente id to update
    var componenteId = e.parentNode.parentNode.childNodes[3].innerText;
    var componenteNome = e.parentNode.parentNode.childNodes[5].innerText;
    var componenteDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('ComponenteIdForUpdate').value = componenteId;
    document.getElementById('ComponenteNomeForUpdate').value = componenteNome;
    document.getElementById('ComponenteDescricaoForUpdate').value = componenteDescricao;
	loadElementosByComponenteId(componenteId);
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
				wipeRows("table-elemento-componente-edit", elementosComponente)
				elementosComponente = [];
				for(order = 0;elementosComponenteJson != null && order<elementosComponenteJson.length;order++){
					elementosComponente[order]=elementosComponenteJson[order];
					addElementoComponenteRow("table-elementos-componente-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadElementosByComponenteId?componenteId="+componenteId,true);
	xmlhttp.send(); 
}
