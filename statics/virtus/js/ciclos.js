function updateCiclo(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get ciclo id to update
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    var cicloTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('CicloIdToUpdate').value = cicloId;
    document.getElementById('CicloTitulo').value = cicloTitulo;
}

function deleteCiclo(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('CicloIdToDelete').value = cicloId;
}