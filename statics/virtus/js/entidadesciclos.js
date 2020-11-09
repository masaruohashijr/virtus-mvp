function submeterEntidadeCicloForm(e, formId){
	console.log(e.parentNode.parentNode.childNodes[3].innerText);
	//alert(e.parentNode.parentNode.childNodes[3].innerText);
	console.log(e.parentNode.parentNode.childNodes[7].childNodes[1].value);
	//alert(e.parentNode.parentNode.childNodes[7].innerText);
	console.log(document.getElementById("EntidadeId").value);
	//alert(document.getElementById("EntidadeId").value);
	console.log(document.getElementById("CicloId").value);
	//alert(document.getElementById("CicloId").value);
	document.getElementById("EntidadeId").value=e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById("CicloId").value=e.parentNode.parentNode.childNodes[7].childNodes[1].value;
	document.getElementById(formId).submit();
}

function submeterDistribuirPapeisForm(e){
	console.log(e.parentNode.parentNode.childNodes[3].innerText);
	console.log(e.parentNode.parentNode.childNodes[7].childNodes[1].value);
	console.log(document.getElementById("EntidadeId").value);
	console.log(document.getElementById("CicloId").value);
	document.getElementById("EntidadeId").value=e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById("CicloId").value=e.parentNode.parentNode.childNodes[7].childNodes[1].value;
	document.getElementById("formulario-distribuir-papeis").submit();
}

function editIntegrantes(e) {
//	resetDetalhesEscritorio();
    var editForm = document.getElementById('edit-equipe-form');
    editForm.style.display = 'block';
    var escritorioId = e.parentNode.parentNode.childNodes[3].innerText;
    var escritorioNome = e.parentNode.parentNode.childNodes[5].innerText;
    var escritorioDescricao = e.parentNode.parentNode.childNodes[7].innerText;
	var escritorioAbreviatura = e.parentNode.parentNode.childNodes[9].innerText;
	if(e.parentNode.parentNode.childNodes[11].childNodes.length > 1){
	    var escritorioChefeId = e.parentNode.parentNode.childNodes[11].childNodes[1].value;
	    document.getElementById('EscritorioChefeESForUpdate').value = escritorioChefeId;
	}
	document.getElementById('EscritorioIdESForUpdate').value = escritorioId;
    document.getElementById('EscritorioNomeESForUpdate').value = escritorioNome;
    document.getElementById('EscritorioDescricaoESForUpdate').value = escritorioDescricao;
    document.getElementById('EscritorioAbreviaturaESForUpdate').value = escritorioAbreviatura;
	loadMembrosByEscritorioId(escritorioId);
}