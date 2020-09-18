function updaterole(e) {
	// resetando
	resetRolesEditForm();
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get role id to update
    var roleId = e.parentNode.parentNode.childNodes[3].innerText;
    var roleName = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('roleIdToUpdate').value = roleId;
    document.getElementById('roleName').value = roleName;
	// carregar as features do papel
	loadFeaturesByRoleId(roleId);
}

function deleterole(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var roleId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('roleIdToDelete').value = roleId;
}

function resetRolesEditForm(){
	document.getElementById('roleName').value = '';
	document.getElementById('FeaturesForUpdate').value='';
}

function loadRolesByActivityId(activityId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var rolesEdit = JSON.parse(xmlhttp.responseText);
				selectOptionsRolesForUpdate(rolesEdit);
			}
	}
	xmlhttp.open("GET","/loadRolesByActivityId?activityId="+activityId,true);
	xmlhttp.send();
}

function selectOptionsRolesForUpdate(rolesEdit){
	let s = document.getElementById("RolesForUpdate");
	for(n=0;n<rolesEdit.length;n++){
		for(m=0;m<s.options.length;m++){
			if(s.options[m].value == rolesEdit[n].id){
				s.options[m].selected = 'selected';
				break;
			}
		}
	}
}