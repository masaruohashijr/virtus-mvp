function updatemeasure(e) {
	var editForm = document.getElementById('edit-form');
	// display update form
	editForm.style.display = 'block';
	// get measure id to update
	var measureId = e.parentNode.parentNode.childNodes[3].innerText;
	var measureName = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('measureIdToUpdate').value = measureId;
	document.getElementById('measureName').value = measureName;
}

function deletemeasure(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var measureId = e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById('measureIdToDelete').value = measureId;
}