function editEscritorio(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get escritorio id to edit
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    var escritorioNome = e.parentNode.parentNode.childNodes[5].innerText;
    var escritorioDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var escritorioChefeId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	document.getElementById('EscritorioIdForUpdate').value = escritorioId;
    document.getElementById('EscritorioNomeForUpdate').value = escritorioNome;
    document.getElementById('EscritorioDescricaoForUpdate').value = escritorioDescricao;
    document.getElementById('EscritorioChefeForUpdate').value = escritorioChefeId;
}

function deleteEscritorio(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('EscritorioIdToDelete').value = escritorioId;
}
