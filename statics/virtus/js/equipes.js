function updateEquipe(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get equipe id to update
    var equipeId = e.parentNode.parentNode.childNodes[3].innerText;
    var equipeTitulo = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('EquipeIdToUpdate').value = equipeId;
    document.getElementById('EquipeTitulo').value = equipeTitulo;
}

function deleteEquipe(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var equipeId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('EquipeIdToDelete').value = equipeId;
}
