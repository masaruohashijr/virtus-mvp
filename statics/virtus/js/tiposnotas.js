function editTipoNota(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var tipoNotaId = e.parentNode.parentNode.childNodes[3].innerText;
    var tipoNotaNome = e.parentNode.parentNode.childNodes[5].innerText;
    var tipoNotaDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var tipoNotaLetra = e.parentNode.parentNode.childNodes[9].innerText;
    var tipoNotaCorLetra = e.parentNode.parentNode.childNodes[11].innerText;
    var tipoNotaDominioComponente = e.parentNode.parentNode.childNodes[13].innerText;
	document.getElementById('TipoNotaIdForUpdate').value = tipoNotaId;
    document.getElementById('NomeForUpdate').value = tipoNotaNome;
    document.getElementById('DescricaoForUpdate').value = tipoNotaDescricao;
    document.getElementById('LetraForUpdate').value = tipoNotaLetra;
    document.getElementById('CorLetraForUpdate').value = tipoNotaCorLetra;
    document.getElementById('DominioComponenteForUpdate').value = tipoNotaDominioComponente;
}

function deleteTipoNota(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var tipoNotaId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('TipoNotaIdForDelete').value = tipoNotaId;
}