function editRole(e) {
	// resetando
	resetRolesEditForm();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var roleId = e.parentNode.parentNode.childNodes[3].innerText;
    var roleName = e.parentNode.parentNode.childNodes[5].innerText;
    var roleDescription = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('RoleIdForUpdate').value = roleId;
    document.getElementById('RoleNameForUpdate').value = roleName;
    document.getElementById('RoleDescriptionForUpdate').value = roleDescription;
    document.getElementById('RoleNameForUpdate').focus();
	// carregar as features do papel
	loadFeaturesByRoleId(roleId);
}

function deleteRole(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var roleId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('perfilIdToDelete').value = roleId;
}

function resetRolesEditForm(){
	document.getElementById('formulario-perfil-edit').reset();
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

function openCreatePerfil(){
	document.getElementById('create-form').style.display='block';
	document.getElementById('RoleNameForInsert').focus();
}