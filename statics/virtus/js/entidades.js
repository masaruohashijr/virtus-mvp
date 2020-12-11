function editEntidade(e) {
	resetEntidadeForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var entidadeId = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
	console.log('entidadeId: '+entidadeId);
    var entidadeCodigo = e.parentNode.parentNode.childNodes[3].innerText;
	console.log('entidadeCodigo: '+entidadeCodigo);
    var entidadeSigla = e.parentNode.parentNode.childNodes[5].innerText;
	console.log('entidadeSigla: '+entidadeSigla);
    var entidadeNome = e.parentNode.parentNode.childNodes[7].innerText;
	console.log('entidadeNome: '+entidadeNome);
    var entidadeDescricao = e.parentNode.parentNode.childNodes[7].childNodes[1].value;
	console.log('entidadeDescricao: '+entidadeDescricao);
    var entidadeSituacao = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	console.log('entidadeSituacao: '+entidadeSituacao);
    var entidadeESI = e.parentNode.parentNode.childNodes[9].childNodes[3].value;
	console.log('entidadeESI: '+entidadeESI);
    var entidadeMunicipio = e.parentNode.parentNode.childNodes[9].childNodes[5].value;
	console.log('entidadeMunicipio: '+entidadeMunicipio);
    var entidadeSiglaUF = e.parentNode.parentNode.childNodes[9].childNodes[7].value;
	console.log('entidadeSiglaUF: '+entidadeSiglaUF);

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
    var entidadeId = e.parentNode.parentNode.childNodes[3].childNodes[0].value;
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
				wipeRows("table-ciclos-entidade-edit")
				ciclosEntidade = [];
				for(order = 0;ciclosEnt != null && order<ciclosEnt.length;order++){
					ciclosEntidade[order]=ciclosEnt[order];
					addCicloEntidadeRow("table-ciclos-entidade-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadCiclosByEntidadeId?entidadeId="+entidadeId,true);
	xmlhttp.send();
}