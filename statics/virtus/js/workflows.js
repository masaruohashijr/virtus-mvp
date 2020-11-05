function editWorkflow(e) { 
	var editForm = document.getElementById('edit-form');
	editForm.style.display = 'block';
	var idWF = e.parentNode.parentNode.childNodes[3].innerText;
	var name = e.parentNode.parentNode.childNodes[5].innerText;
	var description = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
	var entity = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('WorkflowIdForUpdate').value= idWF;
	document.getElementById('NameForUpdate').value= name;
	document.getElementById('DescriptionForUpdate').value= description;
	document.getElementById('EntityForUpdate').value=entity;
	loadActivitiesByWorkflowId(idWF);
}

function deleteWorkflow(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var orderId = e.parentNode.parentNode.childNodes[3].innerText; // alterado
	document.getElementById('WorkflowIdToDelete').value = orderId;
}
