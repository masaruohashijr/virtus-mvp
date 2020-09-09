function updateUser(e) {
	var editForm = document.getElementById('edit-form');
	// display update form
	editForm.style.display = 'block';
	// get user id to update
	var userId = e.parentNode.parentNode.childNodes[3].innerText;
	var userName = e.parentNode.parentNode.childNodes[5].innerText;
	var userUsername = e.parentNode.parentNode.childNodes[7].innerText;
	var userEmail = e.parentNode.parentNode.childNodes[9].innerText;
	var userMobile = e.parentNode.parentNode.childNodes[11].innerText;
	var userRole = e.parentNode.parentNode.childNodes[13].childNodes[1].value;
	document.getElementById('userIdToUpdate').value = userId;
	document.getElementById('userName').value = userName;
	document.getElementById('userUsername').value = userUsername;
	document.getElementById('userEmail').value = userEmail;
	document.getElementById('userMobile').value = userMobile;
	document.getElementById('RoleForUpdate').value = userRole;
}

function deleteUser(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var userId = e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById('userIdToDelete').value = userId;
}