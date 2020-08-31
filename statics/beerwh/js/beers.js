function updatebeer(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get beer id to update
    var beerId = e.parentNode.parentNode.childNodes[3].innerText;
    var beerName = e.parentNode.parentNode.childNodes[5].innerText;
    var beerQtd = e.parentNode.parentNode.childNodes[7].innerText;
    var beerPrice = e.parentNode.parentNode.childNodes[9].innerText;
	document.getElementById('beerIdToUpdate').value = beerId;
    document.getElementById('beerName').value = beerName;
    document.getElementById('beerQtd').value = beerQtd;
    document.getElementById('beerPrice').value = beerPrice;
}

function deletebeer(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var beerId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('beerIdToDelete').value = beerId;
}