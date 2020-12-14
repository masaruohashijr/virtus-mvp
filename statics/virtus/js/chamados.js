function resetChamadoForms(){
	document.getElementById('formulario-create').reset();
	document.getElementById('formulario-edit').reset();
}


function openCreateChamado(btn){
	document.getElementById('create-form').style.display='block';
	document.getElementById('TipoChamadoForInsert').focus();
}

function deleteChamado(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var chamadoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ChamadoIdForDelete').value = chamadoId;
}

function editChamado(e){
	resetChamadoForms();
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var chamadoId = e.parentNode.parentNode.childNodes[3].innerText;
    var chamadoTipo = e.parentNode.parentNode.childNodes[5].innerText;
    var chamadoPrioridade = e.parentNode.parentNode.childNodes[7].innerText;
    var chamadoTitulo = e.parentNode.parentNode.childNodes[9].innerText;
    var chamadoDescricao = e.parentNode.parentNode.childNodes[11].innerText;
    var chamadoResponsavel = e.parentNode.parentNode.childNodes[13].childNodes[1].value;
    var chamadoRelator = e.parentNode.parentNode.childNodes[15].childNodes[1].value;
    var chamadoIniciaEm = e.parentNode.parentNode.childNodes[17].innerText;
    var chamadoProntoEm = e.parentNode.parentNode.childNodes[19].innerText;
    var chamadoEstimativa = e.parentNode.parentNode.childNodes[21].innerText;
    var chamadoAutor = e.parentNode.parentNode.childNodes[23].innerText;
    var chamadoCriadoEm = e.parentNode.parentNode.childNodes[25].innerText;
    var chamadoStatus = e.parentNode.parentNode.childNodes[27].innerText;
	document.getElementById('ChamadoIdForUpdate').value = chamadoId;
	chamadoTipo = parseNome2Valor(chamadoTipo);
    chamadoPrioridade = parseNome2Valor(chamadoPrioridade);
    document.getElementById('TipoChamadoForUpdate').value = chamadoTipo;
    document.getElementById('TituloChamadoForUpdate').value = chamadoTitulo;
    document.getElementById('DescricaoChamadoForUpdate').value = chamadoDescricao;
    document.getElementById('PrioridadeChamadoForUpdate').value = chamadoPrioridade;
    document.getElementById('RelatorChamadoForUpdate').value = chamadoRelator;
    document.getElementById('ResponsavelChamadoForUpdate').value = chamadoResponsavel;
    document.getElementById('EstimativaChamadoForUpdate').value = chamadoEstimativa;
    document.getElementById('IniciaEmChamadoForUpdate').value = formatarData(chamadoIniciaEm);
    document.getElementById('ProntoEmChamadoForUpdate').value = formatarData(chamadoProntoEm);
	document.getElementById('AuthorNameChamadoForUpdate').value = chamadoAutor;
    document.getElementById('CriadoEmChamadoForUpdate').value = chamadoCriadoEm;
    document.getElementById('StatusChamadoForUpdate').value = chamadoStatus;
    document.getElementById('TituloChamadoForUpdate').focus();
	//loadPilaresByCicloId(chamadoId);	
}
