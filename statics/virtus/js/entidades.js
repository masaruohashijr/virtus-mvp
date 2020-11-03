function editEntidade(e) {
	resetEntidadeForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var entidadeId = e.parentNode.parentNode.childNodes[3].innerText;
    var entidadeNome = e.parentNode.parentNode.childNodes[5].innerText;
    var entidadeDescricao = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var entidadeSigla = e.parentNode.parentNode.childNodes[11].childNodes[1].value;
    var entidadeCodigo = e.parentNode.parentNode.childNodes[11].childNodes[3].value;
    var entidadeSituacao = e.parentNode.parentNode.childNodes[11].childNodes[5].value;
    var entidadeESI = e.parentNode.parentNode.childNodes[11].childNodes[7].value;
    var entidadeMunicipio = e.parentNode.parentNode.childNodes[11].childNodes[9].value;
    var entidadeSiglaUF = e.parentNode.parentNode.childNodes[11].childNodes[11].value;

	document.getElementById('EntidadeIdForUpdate').value = entidadeId;
    document.getElementById('EntidadeNomeForUpdate').value = entidadeNome;
    document.getElementById('EntidadeDescricaoForUpdate').value = entidadeDescricao;
    document.getElementById('EntidadeSiglaForUpdate').value = entidadeSigla;
    document.getElementById('EntidadeCodigoForUpdate').value = entidadeCodigo;
    document.getElementById('EntidadeSituacaoForUpdate').value = entidadeSituacao;
    document.getElementById('EntidadeESIForUpdate').value = entidadeESI;
    document.getElementById('EntidadeMunicipioForUpdate').value = entidadeMunicipio;
    document.getElementById('EntidadeSiglaUFForUpdate').value = entidadeSiglaUF;
	loadPlanosByEntidadeId(entidadeId);
	loadCiclosByEntidadeId(entidadeId);
}

function resetEntidadeForms(){
	console.log("resetEntidadeForms()");
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
	document.getElementById('EntidadeDescricaoForInsert').value="";
	console.log(document.getElementById('EntidadeDescricaoForInsert').value);
	document.getElementById('EntidadeDescricaoForUpdate').value="";
	console.log(document.getElementById('EntidadeDescricaoForUpdate').value);
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
				wipeRows("table-planos-edit")
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