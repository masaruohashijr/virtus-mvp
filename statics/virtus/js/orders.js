function updateorder(e) { 
	var editForm = document.getElementById('edit-form');
	editForm.style.display = 'block';
	var idOrder = e.parentNode.parentNode.childNodes[3].innerText;
	var userId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
	var orderedDT = e.parentNode.parentNode.childNodes[7].innerText.split(' ')
	var takeOutDT = e.parentNode.parentNode.childNodes[9].innerText.split(' ');
	var status = e.parentNode.parentNode.childNodes[11].innerText;
	var statusId = e.parentNode.parentNode.childNodes[11].childNodes[1].value;
	var orderedDate = orderedDT[0];
	var orderedAt = orderedDT[1];
	var takeOutDate = takeOutDT[0];
	var takeOutAt = takeOutDT[1];
	document.getElementById('OrderIdForUpdate').value= idOrder;
	document.getElementById('UserForUpdate').value= userId;
	odt = orderedDate.split("/");
	tdt = takeOutDate.split("/");
	document.getElementById('OrderDateForUpdate').value= odt[2]+'-'+odt[1]+'-'+odt[0];
	document.getElementById('OrderedAtForUpdate').value=orderedAt;
	document.getElementById('TakeOutDateForUpdate').value= tdt[2]+'-'+tdt[1]+'-'+tdt[0];
	document.getElementById('TakeOutAtForUpdate').value=takeOutAt;
	document.getElementById('StatusForUpdate').value=status;
	loadItemsByOrderId(idOrder);
	loadAllowedActions(statusId);
	loadAvailableFeatures('edit',statusId);
}

function deleteorder(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var orderId = e.parentNode.parentNode.childNodes[3].innerText; // alterado
	document.getElementById('OrderIdToDelete').value = orderId;
}

function loadAvailableFeatures(contexto,statusId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var feats = JSON.parse(xmlhttp.responseText);
				for(i = 0;feats && i<feats.length;i++){
					if(feats[i].code == 'createItem'){
						document.getElementById(contexto+"-form-item-btn").style.visibility = "visible";
						var btns = document.getElementsByTagName("input");
						for(i=0;i<btns.length;i++){
							if(btns[i].value == "Editar" || btns[i].value == "Apagar" ){
								btns[i].style.visibility = "visible";
							}
						}
						return;
					}
				}
				document.getElementById(contexto+"-form-item-btn").style.visibility = "hidden";
				var btns = document.getElementsByTagName("input");
				for(i=0;i<btns.length;i++){
					if(btns[i].value == "Editar" || btns[i].value == "Apagar" ){
						btns[i].style.visibility = "hidden";
					}
				}
			}
	}
	xmlhttp.open("GET","/loadAvailableFeatures?entityType=order&statusId="+statusId,true);
	xmlhttp.send();
}

function loadAllowedActions(statusId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var formulario = document.getElementById('update-order-form');
				removeActions(formulario);
				// renderizar
				var actions = JSON.parse(xmlhttp.responseText);
				for(i = 0;i<actions.length;i++){
					var orderId = document.getElementById('OrderIdForUpdate').value;
					if(!document.getElementById(actions[i].id+'-'+actions[i].name)){
						var btnAction = document.createElement('input');
						btnAction.type = "button";
						btnAction.className = "w3-btn w3-cerceta w3-margin-top w3-margin-bottom w3-right";
						btnAction.style = "margin-right: 10px; background-color: #05ffb0;";
						btnAction.id = actions[i].id+'-'+actions[i].name;
						btnAction.value = actions[i].name;
						btnAction.documentId = actions[i].id;
						btnAction.isAction = true;
						btnAction.onclick = function() {executeAction(orderId, this.documentId)};
						formulario.appendChild(btnAction);
						loadAvailableFeatures('edit',statusId);
						// alert("Botão "+btnAction.value+" criado! "+btnAction.isAction);
					}
				}
			}
	}
	xmlhttp.open("GET","/loadAllowedActions?entityType=order&statusId="+statusId,true);
	xmlhttp.send();
}

function executeAction(orderId, actionId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
		if (xmlhttp.readyState==4 && xmlhttp.status==200)
		{
			var status = JSON.parse(xmlhttp.responseText);
			document.getElementById("StatusForUpdate").value = status.name; 
			loadAllowedActions(status.id);
		}
	}
	xmlhttp.open("GET","/executeAction?entityType=order&id="+orderId+"&actionId="+actionId,true);
	xmlhttp.send();
}

function removeActions(formulario){
	var nodes = formulario.childNodes;
	for(i=nodes.length-1;i>=0;i--){
		if(nodes[i].isAction){
			formulario.removeChild(nodes[i]);
		}
	}
}