class Anotacao {
	constructor(id, entidadeId, entidadeSigla, assunto, risco, tendencia) {
		this.id = id;
		this.entidadeId = entidadeId;
		this.entidadeSigla = entidadeSigla;
		this.assunto = assunto;
		this.risco = risco;
		this.tendencia = tendencia;
	}
}

function carregarEFPC(campoSelect){
	let select = document.getElementById(campoSelect);
	let opt = document.createElement('option');
	opt.value = '';
	opt.innerHTML = '';
	select.appendChild(opt);

	siglasMap.forEach((value,key)=> {
			opt = document.createElement('option');
			opt.value = key;
			opt.innerHTML = value;
			select.appendChild(opt);
		});
}

function preencherAnotacoesSelect(e, contexto){
	if(contexto=='create'){
		campo = "AnotacoesForInsert";
	} else {
		campo = "AnotacoesForUpdate";
	}
	let select = document.getElementById(campo);
	let opt = document.createElement('option');
	select.length = 0;
	opt.value = '';
	opt.innerHTML = '';
	select.appendChild(opt);	
	anotacoesSiglasMap.forEach(function(value, key) {
		if(e.value == key){
			opt = document.createElement('option');
			opt.value = value[0].id;
			opt.innerHTML = value[0].assunto;
			select.appendChild(opt);	
		}
	});
}

function resetAnotacaoForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}

function openCreateAnotacao(btn){
	document.getElementById('create-form').style.display='block';
	document.getElementById('EntidadeAnotacaoForInsert').focus();
}

function deleteAnotacao(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var anotacaoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('AnotacaoIdForDelete').value = anotacaoId;
}

function editAnotacao(e){
	resetAnotacaoForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var anotacaoId = e.parentNode.parentNode.childNodes[3].innerText;
    var anotacaoEntidadeId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var anotacaoAssunto = e.parentNode.parentNode.childNodes[7].innerText;
    var anotacaoMatriz = e.parentNode.parentNode.childNodes[9].innerText;
    var anotacaoRisco = e.parentNode.parentNode.childNodes[11].innerText;
    var anotacaoTendencia = e.parentNode.parentNode.childNodes[13].innerText;
    var anotacaoDescricao = e.parentNode.parentNode.childNodes[15].innerText;
    var anotacaoAutor = e.parentNode.parentNode.childNodes[17].innerText;
    var anotacaoRelatorId = e.parentNode.parentNode.childNodes[17].childNodes[1].value;
    var anotacaoResponsavelId = e.parentNode.parentNode.childNodes[17].childNodes[3].value;
    var anotacaoCriadoEm = e.parentNode.parentNode.childNodes[19].innerText;
    var anotacaoStatus = e.parentNode.parentNode.childNodes[21].innerText;
	document.getElementById('AnotacaoIdForUpdate').value = anotacaoId;
    document.getElementById('EntidadeAnotacaoForUpdate').value = anotacaoEntidadeId;
    document.getElementById('AssuntoAnotacaoForUpdate').value = anotacaoAssunto;
	document.getElementById('MatrizAnotacaoForUpdate').value = anotacaoMatriz;
    document.getElementById('RiscoAnotacaoForUpdate').value = parseNome2Valor(anotacaoRisco);
    document.getElementById('TendenciaAnotacaoForUpdate').value = parseNome2Valor(anotacaoTendencia);
    document.getElementById('RelatorAnotacaoForUpdate').value = anotacaoRelatorId;
	document.getElementById('ResponsavelAnotacaoForUpdate').value = anotacaoResponsavelId;
	document.getElementById('DescricaoAnotacaoForUpdate').value = anotacaoDescricao;
	document.getElementById('AuthorNameAnotacaoForUpdate').value = anotacaoAutor;
    document.getElementById('CriadoEmAnotacaoForUpdate').value = anotacaoCriadoEm;
    document.getElementById('StatusAnotacaoForUpdate').value = anotacaoStatus;
    document.getElementById('NomeAnotacaoForUpdate').focus();
	// loadPilaresByCicloId(anotacaoId);	
}

