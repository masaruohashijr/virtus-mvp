function editEscritorio(e) {
	resetDetalhesEscritorio();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    var escritorioNome = e.parentNode.parentNode.childNodes[5].innerText;
    var escritorioDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	if(e.parentNode.parentNode.childNodes[9].childNodes.length > 1){
	    var escritorioChefeId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	    document.getElementById('EscritorioChefeForUpdate').value = escritorioChefeId;
	}
	document.getElementById('EscritorioIdForUpdate').value = escritorioId;
    document.getElementById('EscritorioNomeForUpdate').value = escritorioNome;
    document.getElementById('EscritorioDescricaoForUpdate').value = escritorioDescricao;
}

function editJurisdicao(e) {
	resetDetalhesEscritorio();
    var editForm = document.getElementById('edit-jurisdicao-form');
    // display update form
    editForm.style.display = 'block';
    // get escritorio id to edit
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    var escritorioNome = e.parentNode.parentNode.childNodes[5].innerText;
    var escritorioDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	if(e.parentNode.parentNode.childNodes[9].childNodes.length > 1){
	    var escritorioChefeId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	    document.getElementById('EscritorioChefeJUForUpdate').value = escritorioChefeId;
	}
	document.getElementById('EscritorioIdJUForUpdate').value = escritorioId;
    document.getElementById('EscritorioNomeJUForUpdate').value = escritorioNome;
    document.getElementById('EscritorioDescricaoJUForUpdate').value = escritorioDescricao;
}

function editEquipe(e) {
	resetDetalhesEscritorio();
    var editForm = document.getElementById('edit-equipe-form');
    // display update form
    editForm.style.display = 'block';
    // get escritorio id to edit
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    var escritorioNome = e.parentNode.parentNode.childNodes[5].innerText;
    var escritorioDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	if(e.parentNode.parentNode.childNodes[9].childNodes.length > 1){
	    var escritorioChefeId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	    document.getElementById('EscritorioChefeESForUpdate').value = escritorioChefeId;
	}
	document.getElementById('EscritorioIdESForUpdate').value = escritorioId;
    document.getElementById('EscritorioNomeESForUpdate').value = escritorioNome;
    document.getElementById('EscritorioDescricaoESForUpdate').value = escritorioDescricao;
}

function deleteEscritorio(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('EscritorioIdToDelete').value = escritorioId;
}

function resetDetalhesEscritorio(){
	document.getElementById('formulario-jurisdicao-create');
	document.getElementById('formulario-jurisdicao-edit');
	document.getElementById('formulario-membro-create');
	document.getElementById('formulario-membro-edit');
}