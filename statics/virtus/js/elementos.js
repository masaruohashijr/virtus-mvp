function editElemento(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get elemento id to update
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
	//alert(elementoId);
    var elementoNome = e.parentNode.parentNode.childNodes[5].innerText;
	//alert(elementoNome);
    var elementoDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	//alert(elementoDescricao);
    var elementoAutor = e.parentNode.parentNode.childNodes[9].innerText;
	//alert(elementoAutor);
    var elementoCriadoEm = e.parentNode.parentNode.childNodes[11].innerText;
	//alert(elementoCriadoEm);
    var elementoStatus = e.parentNode.parentNode.childNodes[13].innerText;
	//alert(elementoStatus);
	document.getElementById('ElementoIdForUpdate').value = elementoId;
    document.getElementById('ElementoNomeForUpdate').value = elementoNome;
    document.getElementById('ElementoDescricaoForUpdate').value = elementoDescricao;
    document.getElementById('ElementoAutorForUpdate').value = elementoAutor;
    document.getElementById('ElementoCriadoEmForUpdate').value = elementoCriadoEm;
    document.getElementById('ElementoStatusForUpdate').value = elementoStatus;
	// AJAX 
	//var statusId = document.getElementById('ElementoStatusId').value;
	loadItensByElementoId(elementoId);
	//loadAllowedActions(statusId);
	//loadAvailableFeatures('edit',statusId);
}

function deleteElemento(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var elementoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ElementoIdToDelete').value = elementoId;
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
	xmlhttp.open("GET","/loadAvailableFeatures?entityType=elemento&statusId="+statusId,true);
	xmlhttp.send();
}

function loadAllowedActions(statusId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var formulario = document.getElementById('update-elemento-form');
				removeActions(formulario);
				// renderizar
				var actions = JSON.parse(xmlhttp.responseText);
				alert(xmlhttp.responseText);
				for(i = 0;i<actions.length;i++){
					var elementoId = document.getElementById('ElementoIdForUpdate').value;
					if(!document.getElementById(actions[i].id+'-'+actions[i].name)){
						var btnAction = document.createElement('input');
						btnAction.type = "button";
						btnAction.className = "w3-btn w3-cerceta w3-margin-top w3-margin-bottom w3-right";
						btnAction.style = "margin-right: 10px; background-color: #05ffb0;";
						btnAction.id = actions[i].id+'-'+actions[i].name;
						btnAction.value = actions[i].name;
						btnAction.documentId = actions[i].id;
						btnAction.isAction = true;
						btnAction.onclick = function() {executeAction(elementoId, this.documentId)};
						formulario.appendChild(btnAction);
						loadAvailableFeatures('edit',statusId);
						alert("Botão "+btnAction.value+" criado! "+btnAction.isAction);
					}
				}
			}
	}
	xmlhttp.open("GET","/loadAllowedActions?entityType=elemento&statusId="+statusId,true);
	xmlhttp.send();
}

function executeAction(elementoId, actionId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
		if (xmlhttp.readyState==4 && xmlhttp.status==200)
		{
			var status = JSON.parse(xmlhttp.responseText);
			document.getElementById("ElementoStatusForUpdate").value = status.name; 
			loadAllowedActions(status.id);
		}
	}
	xmlhttp.open("GET","/executeAction?entityType=elemento&id="+elementoId+"&actionId="+actionId,true);
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
