function updateElemento(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get elemento id to update
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
    var elementoTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('ElementoIdToUpdate').value = elementoId;
    document.getElementById('ElementoTitulo').value = elementoTitulo;
}

function deleteElemento(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ElementoIdToDelete').value = elementoId;
}

function loadFeaturesByRoleId(roleId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var elementosEdit = JSON.parse(xmlhttp.responseText);
				selectOptionsFeaturesForUpdate(elementosEdit);
			}
	}
	xmlhttp.open("GET","/loadFeaturesByRoleId?roleId="+roleId,true);
	xmlhttp.send();
}

function selectOptionsFeaturesForUpdate(elementosEdit){
	let s = document.getElementById("FeaturesForUpdate");
	for(n=0;n<elementosEdit.length;n++){
		for(m=0;m<s.options.length;m++){
			if(s.options[m].value == elementosEdit[n].id){
				s.options[m].selected = 'selected';
				break;
			}
		}
	}
}

