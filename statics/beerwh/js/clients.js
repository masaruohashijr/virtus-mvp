function updateClient(e) {
	var editForm = document.getElementById('edit-form');
	// display update form
	editForm.style.display = 'block';
	// get client id to update
	var clientId = e.parentNode.parentNode.childNodes[3].innerText;
	var clientName = e.parentNode.parentNode.childNodes[5].innerText;
	var clientUsername = e.parentNode.parentNode.childNodes[7].innerText;
	var clientEmail = e.parentNode.parentNode.childNodes[9].innerText;
	var clientMobile = e.parentNode.parentNode.childNodes[11].innerText;
	document.getElementById('clientIdToUpdate').value = clientId;
	document.getElementById('clientName').value = clientName;
	document.getElementById('clientUsername').value = clientUsername;
	document.getElementById('clientEmail').value = clientEmail;
	document.getElementById('clientMobile').value = clientMobile;
}

function deleteClient(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var clientId = e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById('clientIdToDelete').value = clientId;
}