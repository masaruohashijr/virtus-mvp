function editPilar(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get pilar id to update
    var pilarId = e.parentNode.parentNode.childNodes[3].innerText;
    var pilarNome = e.parentNode.parentNode.childNodes[5].innerText;
    var pilarDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('PilarIdForUpdate').value = pilarId;
    document.getElementById('PilarNomeForUpdate').value = pilarNome;
    document.getElementById('PilarDescricaoForUpdate').value = pilarDescricao;
}

function deletePilar(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var pilarId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('PilarIdToDelete').value = pilarId;
}
