function updaterole(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get role id to update
    var roleId = e.parentNode.parentNode.childNodes[3].innerText;
    var roleName = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('roleIdToUpdate').value = roleId;
    document.getElementById('roleName').value = roleName;
	// resetando
	resetRolesEditForm();
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