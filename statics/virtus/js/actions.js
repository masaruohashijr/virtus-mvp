class Action {
	constructor(id, origin, destination) {
		this.id = id;
		this.origin = origin;
		this.destination = destination;
	}
}

function getActionStatus(type, actionId){
	let acts = actions_array;
	for(n=0;n<acts.length;n++){
		if(acts[n].id == actionId){
			if('origin' == type){
				return acts[n].origin;
			} else {
				return acts[n].destination;
			}
		}
	}
}


function editAction(e) {
	resetActionsEditForm();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    var actionName = e.parentNode.parentNode.childNodes[5].innerText;
    var actionDescription = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
    var except = e.parentNode.parentNode.childNodes[7].innerText;
    var originId = e.parentNode.parentNode.childNodes[9].childNodes[0].value;
    var destinationId = e.parentNode.parentNode.childNodes[11].childNodes[0].value;
	document.getElementById('ActionIdForUpdate').value = actionId;
    document.getElementById('ActionNameForUpdate').value = actionName;
    document.getElementById('ActionDescriptionForUpdate').value = actionDescription;
	if(except=='false'){
		document.getElementById('ExceptForUpdate').checked = false;
	} else {
		document.getElementById('ExceptForUpdate').checked = true;
	}
    document.getElementById('OriginStatusForUpdate').value = originId;
    document.getElementById('DestinationStatusForUpdate').value = destinationId;
}

function deleteAction(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('actionIdToDelete').value = actionId;
}

function resetActionsEditForm(){
	document.getElementById('formulario-action-create').reset();
	document.getElementById('formulario-action-edit').reset();
}