var activity_tobe_deleted;

class Activity {
	constructor(order, id, wid, actionId, actionName, expActionId, expActionName, expTime, startAt, endAt, roles, roleNames, roles_array, features) {
		this.order = order;
		this.id = id;
		this.wid = wid;
		this.actionId = actionId;
		this.actionName = actionName;
		this.expActionId = expActionId;
		this.expActionName = expActionName;
		this.expTime = expTime;
		this.startAt = startAt;
		this.endAt = endAt;
		this.roles = roles;
		this.roleNames = roleNames;
		this.roles_array = roles_array;
		this.features = features;
	}
}

function limparCamposActivityForm(form) {
	document.getElementById('action-' + form).value = "";
	document.getElementById('start-at-' + form).value = "";
	document.getElementById('end-at-' + form).value = "";
	document.getElementById('exp-time-' + form).value = "";
	document.getElementById('exp-action-' + form).value = "";
	document.getElementById('roles-' + form).value = "";
	document.getElementById('features-' + form).value = "";
}

function criarActivity() {
	var wId = 0;
	if (contexto != 'insert') {
		var wId = document.getElementById('WorkflowIdForUpdate').value;
	}
	var a = document.getElementById('action-create');
	var actionId = a.options[a.selectedIndex].value;
	var erros = '';
	if (actionId == '') {
		erros += 'Falta preencher a ação.\n';
		alert(erros);
		return;
	}
	var actionName = a.options[a.selectedIndex].text;
	var startAt = document.getElementById('start-at-create').value;
	var endAt = document.getElementById('end-at-create').value;
	var expTime = document.getElementById('exp-time-create').value;
	var ea = document.getElementById('exp-action-create');
	var expActionId = ea.options[ea.selectedIndex].value;
	var expActionName = ea.options[ea.selectedIndex].text;
	var roles = getSelectedItems(document.getElementById('roles-create'));
	var roleNames = getSelectedRoleNames(document.getElementById('roles-create'));
	var features = getSelectedItems(document.getElementById('features-create'));
	activity = new Activity(activities.length, 0, wId, actionId, actionName, expActionId, expActionName, expTime, startAt, endAt, roles, roleNames, '', features);
	activities.push(activity);
	addActRow("table-activities-" + contexto);
	limparCamposActivityForm('create');
	document.getElementById('create-activity-form').style.display = 'none';
}


function editActivity(e) {
	limparCamposActivityForm('edit');
	var editActivityForm = document.getElementById('edit-activity-form');
	editActivityForm.style.display = 'block';
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	//alert("id: "+id);
	var actionId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	//alert("actionId: "+actionId);
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	//alert("order: "+order);
	var origin = e.parentNode.parentNode.childNodes[1].innerText;
	//alert("origin: "+origin);
	var destination = e.parentNode.parentNode.childNodes[2].innerText;
	//alert("destination: "+destination);
	var startAt = e.parentNode.parentNode.childNodes[3].innerText;
	//alert("startAt: "+startAt);
	var expTime = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
	//alert("expTime: "+expTime);
	var endAt = e.parentNode.parentNode.childNodes[4].innerText;
	//alert("endAt: "+endAt);
	var expActionId = e.parentNode.parentNode.childNodes[4].childNodes[1].value;
	//alert("expActionId: "+expActionId);
	var expActionName = e.parentNode.parentNode.childNodes[4].childNodes[0].value;
	//alert("expActionName: "+expActionName);
	var rolesIds = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
	//alert("rolesIds: "+rolesIds);
	var featuresIds = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
	// alert("featuresIds: "+featuresIds);
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value = id;
	document.getElementById('action-edit').value = actionId;
	document.getElementById('start-at-edit').value = startAt;
	document.getElementById('end-at-edit').value = endAt;
	document.getElementById('exp-time-edit').value = expTime;
	document.getElementById('exp-action-edit').value = expActionId;
	setSelected(document.getElementById('roles-edit'), rolesIds);
	setSelected(document.getElementById('features-edit'), featuresIds);
	document.getElementById('order-edit').value = order;
}

function updateActivity() {
	var wId = document.getElementById('WorkflowIdForUpdate').value;
	var id = document.getElementById('id-edit').value;
	var a = document.getElementById('action-edit');
	var actionId = a.options[a.selectedIndex].value;
	var erros = '';
	if (actionId == '') {
		erros += 'Falta preencher a ação.\n';
		alert(erros);
		return;
	}
	var actionName = a.options[a.selectedIndex].text;
	var startAt = document.getElementById('start-at-edit').value;
	var endAt = document.getElementById('end-at-edit').value;
	var expTime = document.getElementById('exp-time-edit').value;
	var ea = document.getElementById('exp-action-edit');
	if(ea.selectedIndex>-1){
		var expActionId = ea.options[ea.selectedIndex].value;
		var expActionName = ea.options[ea.selectedIndex].text;
	} else {
		var expActionId = "";
		var expActionName = "";
	}
	var roles = getSelectedItems(document.getElementById('roles-edit'));
	var roleNames = getSelectedRoleNames(document.getElementById('roles-edit'));
	var features = getSelectedItems(document.getElementById('features-edit'));
	var order = document.getElementById('order-edit').value;
	activity = new Activity(order, id, wId, actionId, actionName, expActionId, expActionName, expTime, startAt, endAt, roles, roleNames, '', features);
	activities[order] = activity;
	updateActRow("table-activities-" + contexto, order);
	//limparCamposActivityForm('edit');
	var editActivityForm = document.getElementById('edit-activity-form');
	editActivityForm.style.display = 'none';
}

function showDeleteActivityForm(e) {
	var deleteActivityForm = document.getElementById('delete-activity-form');
	deleteActivityForm.style.display = 'block';
	activity_tobe_deleted = e;
}


function deleteActivity() {
	var order = activity_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newActivities = [];
	for (i = 0; i < activities.length; i++) {
		if (i != order) {
			newActivities.push(activities[i]);
		}
	}
	activities = newActivities;
	activity_tobe_deleted.parentNode.parentNode.innerHTML = '';
	var deleteActivityForm = document.getElementById('delete-activity-form');
	deleteActivityForm.style.display = 'none';
}

function addActRow(tableID) {
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.insertRow(-1);
	order = activities.length - 1;
	activity = activities[order];
	// actvt
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(activity.actionName);
	var jsonActvt = JSON.stringify(activity);
	jsonActvt = jsonActvt.split(',').join('#');
	jsonActvt = jsonActvt.split('"').join('');
	jsonActvt = jsonActvt.split('{').join('');
	jsonActvt = jsonActvt.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="activity' + activity.actionId + '" value="' + jsonActvt + '"/>' + newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="actionId" value="' + activity.actionId + '"/>' + newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="' + activity.id + '"/>' + newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="' + order + '"/>' + newCell.innerHTML;
	// Origin
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(getActionStatus('origin',activity.actionId));
	newCell.appendChild(newText);
	// Destination
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(getActionStatus('destination',activity.actionId));
	newCell.appendChild(newText);
	// Inicia Em - Start At
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(activity.startAt);
	newCell.appendChild(newText);
	// Expiration Time in Days
	newCell.innerHTML = '<input type="hidden" name="expTime" value="' + activity.expTime + '"/>' + newCell.innerHTML;
	// Termina Em - End At
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(activity.endAt);
	newCell.appendChild(newText);
	// Expiration Action (Id,Name)
	newCell.innerHTML = '<input type="hidden" name="expActionId" value="' + activity.expActionId + '"/>' + newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="expActionName" value="' + activity.expActionName + '"/>' + newCell.innerHTML;
	// Roles Allowed
	newCell = newRow.insertCell(5);
	newText = document.createTextNode('');
	newCell.appendChild(newText);
	str = activity.roleNames.split(".").join(", ");
	newCell.innerHTML = str;
	newCell.innerHTML = '<input type="hidden" name="features" value="' + activity.features + '"/>' + newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="roles" value="' + activity.roles + '"/>' + newCell.innerHTML;
	newCell = newRow.insertCell(6);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() { editActivity(btnEditar) };
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnEditar.style = "margin-right: 10px";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() { showDeleteActivityForm(btnApagar) };
	newCell.appendChild(btnApagar);
}

function updateActRow(tableID, order) {
	let tableRef = document.getElementById(tableID);
	let rowNumber = 3 + parseInt(order);
	let row = tableRef.childNodes[1].childNodes[rowNumber];
	let celula = row.childNodes[0];
	celula.innerText = activities[order].actionName;
	var jsonActivity = JSON.stringify(activities[order]);
	jsonActivity = jsonActivity.split(',').join('#');
	jsonActivity = jsonActivity.split('"').join('');
	jsonActivity = jsonActivity.split('{').join('');
	jsonActivity = jsonActivity.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="activity' + activity.actionId + '" value="' + jsonActivity + '"/>' + celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="actionId" value="' + activity.actionId + '"/>' + celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="' + activity.id + '"/>' + celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="' + order + '"/>' + celula.innerHTML;
	// Origin
	row.childNodes[1].innerText = getActionStatus('origin',activities[order].actionId);
	// Destination
	row.childNodes[2].innerText = getActionStatus('destination',activities[order].actionId);
	// Inicia Em - Start At
	row.childNodes[3].innerText = activities[order].startAt;
	// Expiration Time in Days
	row.childNodes[3].innerHTML = '<input type="hidden" name="expTime" value="' + activities[order].expTime + '"/>' + row.childNodes[3].innerHTML;
	// Termina Em - End At
	row.childNodes[4].innerText = activities[order].endAt;
	// Expiration Action (Id,Name)
	row.childNodes[4].innerHTML = '<input type="hidden" name="expActionId" value="' + activities[order].expActionId + '"/>' + row.childNodes[4].innerHTML;
	row.childNodes[4].innerHTML = '<input type="hidden" name="expActionName" value="' + activities[order].expActionName + '"/>' + row.childNodes[4].innerHTML;
	// Roles Allowed
	row.childNodes[5].innerHTML = activities[order].roleNames;
	row.childNodes[5].innerHTML = '<input type="hidden" name="features" value="' + activities[order].features + '"/>' + row.childNodes[5].innerHTML;
	row.childNodes[5].innerHTML = '<input type="hidden" name="roles" value="' + activities[order].roles + '"/>' + row.childNodes[5].innerHTML;
}

function getSelectedFeatures(select) {
	var result = '';
	var options = select && select.options;
	var opt;
	for (var i = 0, iLen = options.length; i < iLen; i++) {
		opt = options[i];
		if (opt.selected) {
			result += opt.value + ".";
		}
	}
	if (result.length > 0) {
		result = result.substring(0, result.length - 1);
	}
	return result;
}

function getSelectedItems(select) {
	var result = '';
	var options = select && select.options;
	var opt;
	for (var i = 0, iLen = options.length; i < iLen; i++) {
		opt = options[i];
		if (opt.selected) {
			result += opt.value + ".";
		}
	}
	if (result.length > 0) {
		result = result.substring(0, result.length - 1);
	}
	return result;
}

function getSelectedRoleNames(select) {
	var result = '';
	var options = select && select.options;
	var opt;
	for (var i = 0, iLen = options.length; i < iLen; i++) {
		opt = options[i];
		if (opt.selected) {
			result += opt.text + ". ";
		}
	}
	if (result.length > 0) {
		result = result.substring(0, result.length - 2);
	}
	return result;
}


function setSelected(select, ids) {
	var sels = ids.split('.');
	for (var n = 0; n < sels.length; n++) {
		var options = select.options;
		var opt;
		for (var i = 0; i < options.length; i++) {
			opt = options[i];
			let sel = sels[n].trim();
			if (opt.value == sel) {
				opt.selected = 'selected';
			}
		}
	}
}

function loadActivitiesByWorkflowId(idWF) {
	var xmlhttp;
	xmlhttp = new XMLHttpRequest();
	xmlhttp.onreadystatechange = function() {
		if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
			var activitiesEdit = JSON.parse(xmlhttp.responseText);
			let tableRef = document.getElementById("table-activities-" + contexto);
			for (i = 0; i < activities.length; i++) {
				tableRef.deleteRow(1);
			}
			activities = [];
			for (order = 0; order < activitiesEdit.length; order++) {
				activities[order] = activitiesEdit[order];
				addActRow("table-activities-" + contexto);
			}
			return activities;
		}
	}
	xmlhttp.open("GET", "/loadActivitiesByWorkflowId?idWF=" + idWF, true);
	xmlhttp.send();
}