function editEntidade(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get entidade id to update
    var entidadeId = e.parentNode.parentNode.childNodes[3].innerText;
    var entidadeNome = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('EntidadeIdForUpdate').value = entidadeId;
    document.getElementById('EntidadeNomeForUpdate').value = entidadeNome;
	loadPlanosByEntidadeId(entidadeId);
	loadCiclosByEntidadeId(entidadeId);
}

function deleteEntidade(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var entidadeId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('EntidadeIdToDelete').value = entidadeId;
}

function loadPlanosByEntidadeId(entidadeId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var planosEntidade = JSON.parse(xmlhttp.responseText);
				wipeRows("table-planos-edit", planos)
				planos = [];
				for(i = 0;planosEntidade != null && i <planosEntidade.length;i++){
					planos[i]=planosEntidade[i];
					addPlanoRow("table-planos-edit");
				}
				return planos;
			}
	}
	xmlhttp.open("GET","/loadPlanosByEntidadeId?entidadeId="+entidadeId,true);
	xmlhttp.send();
}

function loadCiclosByEntidadeId(entidadeId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var ciclosEnt = JSON.parse(xmlhttp.responseText);
				wipeRows("table-ciclos-entidade-edit", ciclos)
				ciclosEntidade = [];
				for(order = 0;ciclosEnt != null && order<ciclosEnt.length;order++){
					ciclosEntidade[order]=ciclosEnt[order];
					addCicloEntidadeRow("table-ciclos-entidade-edit");
				}
				return ciclos;
			}
	}
	xmlhttp.open("GET","/loadCiclosByEntidadeId?entidadeId="+entidadeId,true);
	xmlhttp.send();
}