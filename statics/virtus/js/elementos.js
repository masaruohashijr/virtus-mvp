function updateElemento(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get elemento id to update
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
    var elementoTitulo = e.parentNode.parentNode.childNodes[5].innerText;
    var elementoDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var elementoAutor = e.parentNode.parentNode.childNodes[9].innerText;
    var elementoDataCriacao = e.parentNode.parentNode.childNodes[11].innerText;
    var elementoStatus = e.parentNode.parentNode.childNodes[13].innerText;
	/*document.getElementById('ElementoIdForUpdate').value = elementoId;
    document.getElementById('ElementoTituloForUpdate').value = elementoTitulo;
    document.getElementById('ElementoDescricaoForUpdate').value = elementoDescricao;
    document.getElementById('ElementoAutorForUpdate').value = elementoAutor;
    document.getElementById('ElementoDataCriacaoForUpdate').value = elementoDataCriacao;
    document.getElementById('ElementoStatusForUpdate').value = elementoStatus;*/
	// AJAX 
	loadItensByElementoId(elementoId);
}

function deleteElemento(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ElementoIdToDelete').value = elementoId;
}

