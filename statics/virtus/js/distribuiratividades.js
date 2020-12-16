function submeterDistribuirAtividadesForm(e){
	console.log(e.parentNode.parentNode.childNodes[3].childNodes[0].value);
	console.log(e.parentNode.parentNode.childNodes[7].childNodes[1].value);
	console.log("ANTES");
	console.log(document.getElementById("EntidadeId").value);
	console.log(document.getElementById("CicloId").value);
	document.getElementById("EntidadeId").value=e.parentNode.parentNode.childNodes[3].childNodes[0].value;
	document.getElementById("CicloId").value=e.parentNode.parentNode.childNodes[7].childNodes[1].value;
	console.log("DEPOIS");
	console.log(document.getElementById("EntidadeId").value);
	console.log(document.getElementById("CicloId").value);
	document.getElementById("formulario-distribuir-atividades").submit();
}

function validarDistribuirAtividades(e){
	if (e.parentNode.parentNode.childNodes[7].childNodes[1].length == 0) {
		// Na tabela de Distribuição de Atividades
		// campo Select dos ciclos da entidade na linha da tabela
		console.log(false);
		return false;	
	} else {
		console.log(true);
		return true;
	}
}

function motivarReprogramacao(campo){
	let entidadeId = campo.name.split("_")[1];
	let cicloId = campo.name.split("_")[2];
	let pilarId = campo.name.split("_")[3];
	let componenteId = campo.name.split("_")[4];
	let dataAnterior = campo.name.split("_")[5];
	if(campo.value != dataAnterior && dataAnterior != ""){
		document.getElementById("AcionadoPor").value = campo.name;
		document.getElementById("motRepro_callback").value = campo.name;
		document.getElementById("motReproEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motReproCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motReproPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motReproComponente").value = componentesMap.get(componenteId);
		document.getElementById("motReproDataAnterior").value = formatarData(dataAnterior);
		document.getElementById("motReproNovaData").value = formatarData(campo.value);
		document.getElementById('motivar-reprogramacao-form').style.display='block';
		document.getElementById("motRepro_text").value='';
		document.getElementById("motRepro_text").focus();
	}
}