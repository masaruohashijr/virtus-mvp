function editComponente(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get componente id to update
    var componenteId = e.parentNode.parentNode.childNodes[3].innerText;
    var componenteNome = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('ComponenteIdToUpdate').value = componenteId;
    document.getElementById('ComponenteNome').value = componenteNome;
}

function deleteComponente(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var componenteId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ComponenteIdToDelete').value = componenteId;
}

function loadFeaturesByRoleId(roleId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var componentesEdit = JSON.parse(xmlhttp.responseText);
				selectOptionsFeaturesForUpdate(componentesEdit);
			}
	}
	xmlhttp.open("GET","/loadFeaturesByRoleId?roleId="+roleId,true);
	xmlhttp.send();
}

function selectOptionsFeaturesForUpdate(componentesEdit){
	let s = document.getElementById("FeaturesForUpdate");
	for(n=0;n<componentesEdit.length;n++){
		for(m=0;m<s.options.length;m++){
			if(s.options[m].value == componentesEdit[n].id){
				s.options[m].selected = 'selected';
				break;
			}
		}
	}
}

