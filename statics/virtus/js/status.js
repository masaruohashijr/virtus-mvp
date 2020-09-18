function updatestatus(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get status id to update
    var statusId = e.parentNode.parentNode.childNodes[3].innerText;
    var statusName = e.parentNode.parentNode.childNodes[5].innerText;
    var statusStereotype = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('statusIdForUpdate').value = statusId;
    document.getElementById('statusNameForUpdate').value = statusName;
    document.getElementById('statusStereotypeForUpdate').value = statusStereotype;
}

function deletestatus(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var statusId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('statusIdToDelete').value = statusId;
}