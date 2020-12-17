function loadMatriz(entidadeId, cicloId){
	document.getElementById("EntidadeId").value=entidadeId;
	document.getElementById("CicloId").value=cicloId;
	document.getElementById("PilarId").value='';
	document.getElementById("formulario-executar-matriz").submit();
}

function loadVisaoPilar(entidadeId, cicloId, pilarId){
	document.getElementById("EntidadeId").value=entidadeId;
	document.getElementById("CicloId").value=cicloId;
	document.getElementById("PilarId").value=pilarId;
	document.getElementById("formulario-executar-matriz").submit();
}

function loadVisaoComponente(entidadeId, cicloId, pilarId, componenteId){
	document.getElementById("EntidadeId").value=entidadeId;
	document.getElementById("CicloId").value=cicloId;
	document.getElementById("PilarId").value=pilarId;
	document.getElementById("ComponenteId").value=componenteId;
	document.getElementById("formulario-executar-matriz").submit();
}