function updateMatriz(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get matriz id to update
    var matrizId = e.parentNode.parentNode.childNodes[3].innerText;
    var matrizTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('MatrizIdToUpdate').value = matrizId;
    document.getElementById('MatrizTitulo').value = matrizTitulo;
}

function deleteMatriz(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var matrizId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('MatrizIdToDelete').value = matrizId;
}
