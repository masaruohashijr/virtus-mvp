function editTipoNota(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var tipoNotaId = e.parentNode.parentNode.childNodes[3].innerText;
    var tipoNotaNome = e.parentNode.parentNode.childNodes[5].innerText;
    var tipoNotaDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var tipoNotaReferencia = e.parentNode.parentNode.childNodes[9].innerText;
    var tipoNotaLetra = e.parentNode.parentNode.childNodes[11].innerText;
    var tipoNotaCorLetra = e.parentNode.parentNode.childNodes[13].innerText;
    var tipoNotaDominioComponente = e.parentNode.parentNode.childNodes[15].innerText;
	document.getElementById('TipoNotaIdForUpdate').value = tipoNotaId;
    document.getElementById('NomeForUpdate').value = tipoNotaNome;
    document.getElementById('DescricaoForUpdate').value = tipoNotaDescricao;
    document.getElementById('ReferenciaForUpdate').value = tipoNotaReferencia;
    document.getElementById('LetraForUpdate').value = tipoNotaLetra;
    document.getElementById('CorLetraForUpdate').value = tipoNotaCorLetra;
    document.getElementById('DominioComponenteForUpdate').value = tipoNotaDominioComponente;
    document.getElementById('NomeForUpdate').focus();
}

function deleteTipoNota(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var tipoNotaId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('TipoNotaIdForDelete').value = tipoNotaId;
}

function openCreateTipoNota(){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NomeForInsert').focus();
}