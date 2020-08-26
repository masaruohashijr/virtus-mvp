function updatefood(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get food id to update
    var foodId = e.parentNode.parentNode.childNodes[3].innerText;
    var foodName = e.parentNode.parentNode.childNodes[5].innerText;
    var foodMeasure = e.parentNode.parentNode.childNodes[7].innerText;
    var foodQtd = e.parentNode.parentNode.childNodes[9].innerText;
    var foodCho = e.parentNode.parentNode.childNodes[11].innerText;
    var foodKcal = e.parentNode.parentNode.childNodes[13].innerText;
	document.getElementById('foodIdToUpdate').value = foodId;
    document.getElementById('foodName').value = foodName;
    document.getElementById('foodMeasure').value = foodMeasure;
    document.getElementById('foodQtd').value = foodQtd;
    document.getElementById('foodCho').value = foodCho;
    document.getElementById('foodKcal').value = foodKcal;
}

function deletefood(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var foodId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('foodIdToDelete').value = foodId;
}