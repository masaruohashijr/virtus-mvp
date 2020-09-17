function updateorder(e) { 
	var editForm = document.getElementById('edit-form');
	editForm.style.display = 'block';
	var idOrder = e.parentNode.parentNode.childNodes[3].innerText;
	var clientId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
	var orderedDT = e.parentNode.parentNode.childNodes[7].innerText.split(' ')
	var takeOutDT = e.parentNode.parentNode.childNodes[9].innerText.split(' ');
	var orderedDate = orderedDT[0];
	var orderedAt = orderedDT[1];
	var takeOutDate = takeOutDT[0];
	var takeOutAt = takeOutDT[1];
	document.getElementById('OrderIdForUpdate').value= idOrder;
	document.getElementById('ClientForUpdate').value= clientId;
	let odt = orderedDate.split("/");
	let tdt = takeOutDate.split("/");
	document.getElementById('OrderDateForUpdate').value= odt[2]+'-'+odt[1]+'-'+odt[0];
	document.getElementById('OrderedAtForUpdate').value=orderedAt;
	document.getElementById('TakeOutDateForUpdate').value= tdt[2]+'-'+tdt[1]+'-'+tdt[0];
	document.getElementById('TakeOutAtForUpdate').value=takeOutAt;
	loadItemsByOrderId(idOrder);
}

function deleteorder(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var orderId = e.parentNode.parentNode.childNodes[3].innerText; // alterado
	document.getElementById('OrderIdToDelete').value = orderId;
}
