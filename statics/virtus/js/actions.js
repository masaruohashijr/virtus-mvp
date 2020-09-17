function updateaction(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get action id to update
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    var actionName = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('actionIdToUpdate').value = actionId;
    document.getElementById('actionName').value = actionName;
}

function deleteaction(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('actionIdToDelete').value = actionId;
}