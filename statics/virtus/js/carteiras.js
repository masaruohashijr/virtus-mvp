function updateCarteira(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get carteira id to update
    var carteiraId = e.parentNode.parentNode.childNodes[3].innerText;
    var carteiraTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('CarteiraIdToUpdate').value = carteiraId;
    document.getElementById('CarteiraTitulo').value = carteiraTitulo;
}

function deleteCarteira(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var carteiraId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('CarteiraIdToDelete').value = carteiraId;
}
