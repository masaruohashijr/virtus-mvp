function resetCicloForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}

function editCiclo(e) {
	resetCicloForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    // display update form
    // get ciclo id to update
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    var cicloNome = e.parentNode.parentNode.childNodes[5].innerText;
    var cicloDescricao = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var cicloAutor = e.parentNode.parentNode.childNodes[7].innerText;
    var cicloCriadoEm = e.parentNode.parentNode.childNodes[9].innerText;
	document.getElementById('CicloIdForUpdate').value = cicloId;
    document.getElementById('NomeCicloForUpdate').value = cicloNome;
    document.getElementById('DescricaoCicloForUpdate').value = cicloDescricao;
	document.getElementById('AuthorNameForUpdate').value = cicloAutor;
    document.getElementById('CriadoEmForUpdate').value = cicloCriadoEm;
}

function deleteCiclo(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var cicloId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('CicloIdForDelete').value = cicloId;
}