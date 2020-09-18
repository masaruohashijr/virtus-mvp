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


function updateaction(e) {
	resetActionsEditForm();
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get action id to update
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    var actionName = e.parentNode.parentNode.childNodes[5].innerText;
    var except = e.parentNode.parentNode.childNodes[7].innerText;
    var originId = e.parentNode.parentNode.childNodes[9].childNodes[0].value;
    var destinationId = e.parentNode.parentNode.childNodes[11].childNodes[0].value;
	document.getElementById('actionIdToUpdate').value = actionId;
    document.getElementById('actionName').value = actionName;
	if(except=='false'){
		document.getElementById('ExceptForUpdate').checked = false;
	} else {
		document.getElementById('ExceptForUpdate').checked = true;
	}
    document.getElementById('OriginStatusForUpdate').value = originId;
    document.getElementById('DestinationStatusForUpdate').value = destinationId;
}

function deleteaction(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var actionId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('actionIdToDelete').value = actionId;
}

function resetActionsEditForm(){
	document.getElementById('actionIdToUpdate').value = '';
	document.getElementById('actionName').value='';
	document.getElementById('ExceptForUpdate').checked = false;
    document.getElementById('OriginStatusForUpdate').value = '';
    document.getElementById('DestinationStatusForUpdate').value = '';
}