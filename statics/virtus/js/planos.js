function updatePlano(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get plano id to update
    var planoId = e.parentNode.parentNode.childNodes[3].innerText;
    var planoTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('PlanoIdToUpdate').value = planoId;
    document.getElementById('PlanoTitulo').value = planoTitulo;
}

function deletePlano(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var planoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('PlanoIdToDelete').value = planoId;
}
