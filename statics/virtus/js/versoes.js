function resetVersaoForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}

function openCreateVersao(e){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NomeVersaoForInsert').focus();
}

function deleteVersao(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var versaoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('VersaoIdForDelete').value = versaoId;
}

function editVersao(e){
	resetVersaoForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var versaoId = e.parentNode.parentNode.childNodes[3].innerText;
    var versaoNome = e.parentNode.parentNode.childNodes[5].innerText;
    var versaoObjetivo = e.parentNode.parentNode.childNodes[7].innerText;
    var versaoDefinicaoPronto = e.parentNode.parentNode.childNodes[9].innerText;
    var versaoIniciaEm = e.parentNode.parentNode.childNodes[11].innerText;
    var versaoTerminaEm = e.parentNode.parentNode.childNodes[13].innerText;
    var versaoAutor = e.parentNode.parentNode.childNodes[15].innerText;
    var versaoCriadoEm = e.parentNode.parentNode.childNodes[17].innerText;
    var versaoStatus = e.parentNode.parentNode.childNodes[19].innerText;
	document.getElementById('VersaoIdForUpdate').value = versaoId;
    document.getElementById('NomeVersaoForUpdate').value = versaoNome;
    document.getElementById('ObjetivoVersaoForUpdate').value = versaoObjetivo;
    document.getElementById('DefinicaoProntoVersaoForUpdate').value = versaoDefinicaoPronto;
    document.getElementById('IniciaEmVersaoForUpdate').value = formatarData(versaoIniciaEm);
    document.getElementById('TerminaEmVersaoForUpdate').value = formatarData(versaoTerminaEm);
	document.getElementById('AuthorNameVersaoForUpdate').value = versaoAutor;
    document.getElementById('CriadoEmVersaoForUpdate').value = versaoCriadoEm;
    document.getElementById('StatusVersaoForUpdate').value = versaoStatus;
    document.getElementById('NomeVersaoForUpdate').focus();
	loadPilaresByCicloId(versaoId);
}