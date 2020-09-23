function updateEntidade(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get entidade id to update
    var entidadeId = e.parentNode.parentNode.childNodes[3].innerText;
    var entidadeTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('EntidadeIdToUpdate').value = entidadeId;
    document.getElementById('EntidadeTitulo').value = entidadeTitulo;
}

function deleteEntidade(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var entidadeId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('EntidadeIdToDelete').value = entidadeId;
}
