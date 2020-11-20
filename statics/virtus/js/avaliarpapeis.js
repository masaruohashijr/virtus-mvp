function resetFormAvaliarPapeis(){
	let inputs = document.getElementById('form-avaliar-papeis').elements;
	for (i = 0; i < inputs.length; i++) {
		console.log(inputs[i].name + " - " + inputs[i].type);
		if(inputs[i].type == "submit"){
			inputs[i].removeAttribute("disabled");
		}
	}
}

function motivarNota(sel){
	let entidadeId = sel.name.split("_")[1];
	let cicloId = sel.name.split("_")[2];
	let pilarId = sel.name.split("_")[3];
	let componenteId = sel.name.split("_")[4];
	let tipoNotaId = sel.name.split("_")[5];
	let elementoId = sel.name.split("_")[6];
	let notaAnterior = sel.name.split("_")[7];
	if(sel.value != notaAnterior){
		document.getElementById("motNota_callback").value = sel.name;
		document.getElementById("motNotaEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motNotaCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motNotaPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motNotaComponente").value = componentesMap.get(componenteId);
		document.getElementById("motNotaTipoNota").value = tiposNotasMap.get(tipoNotaId);
		document.getElementById("motNotaElemento").value = elementosMap.get(elementoId);
		document.getElementById("motNotaNotaAnterior").value = notaAnterior;
		document.getElementById("motNotaNovaNota").value = sel.value;
		document.getElementById('motivar-nota-form').style.display='block';
		document.getElementById("motNota_text").focus();
	}
}

function motivarPeso(sel){
	let entidadeId = sel.name.split("_")[1];
	let cicloId = sel.name.split("_")[2];
	let pilarId = sel.name.split("_")[3];
	let componenteId = sel.name.split("_")[4];
	let tipoNotaId = sel.name.split("_")[5];
	let elementoId = sel.name.split("_")[6];
	let pesoAnterior = sel.name.split("_")[7];
	if(sel.value != pesoAnterior){
		document.getElementById("motPeso_callback").value = sel.name;
		document.getElementById("motPesoEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motPesoCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motPesoPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motPesoComponente").value = componentesMap.get(componenteId);
		document.getElementById("motPesoTipoNota").value = tiposNotasMap.get(tipoNotaId);
		document.getElementById("motPesoElemento").value = elementosMap.get(elementoId);
		document.getElementById("motPesoPesoAnterior").value = pesoAnterior;
		document.getElementById("motPesoNovoPeso").value = sel.value;
		document.getElementById('motivar-peso-form').style.display='block';
		document.getElementById("motPeso_text").focus();
	}
}

function motivarRemocao(sel){
	let entidadeId = sel.name.split("_")[1];
	let cicloId = sel.name.split("_")[2];
	let pilarId = sel.name.split("_")[3];
	let componenteId = sel.name.split("_")[4];
	let auditorAnterior = sel.name.split("_")[5];
	if(sel.value != auditorAnterior){
		document.getElementById("motRem_callback").value = sel.name;
		document.getElementById("motRemEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motRemCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motRemPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motRemComponente").value = componentesMap.get(componenteId);
		document.getElementById("motRemAuditorAnterior").value = auditoresMap.get(auditorAnterior);
		document.getElementById("motRemNovoAuditor").value = sel.options[sel.selectedIndex].text;
		document.getElementById('motivar-remocao-form').style.display='block';
		document.getElementById("motRem_text").focus();
	}
}

function resetAuditor(){
	let campoAuditorComponente = document.getElementById("motRem_callback").value;
	document.getElementsByName(campoAuditorComponente)[0].value = campoAuditorComponente.split("_")[5];
}

function resetNota(){
	let campoNotaElemento = document.getElementById("motNota_callback").value;
	document.getElementsByName(campoNotaElemento)[0].value = campoNotaElemento.split("_")[7];
}

function resetPeso(){
	let campoPesoElemento = document.getElementById("motPeso_callback").value;
	document.getElementsByName(campoPesoElemento)[0].value = campoPesoElemento.split("_")[7];
}

function salvarNota(){
	let motivacao = document.getElementById('motNota_text');
	if(motivacao.value.length>3){
		resetFormAvaliarPapeis();
		document.getElementsByName('MotivacaoNota')[0].value=motivacao.value;
		document.getElementById('motivar-nota-form').style.display='none';
		document.getElementById('form-avaliar-papeis').submit();
	} else {
		alert("Falta preencher a motivação da nota.");
		motivacao.focus();
		return;		
	}
}

function salvarPeso(){
	let motivacao = document.getElementById('motPeso_text');
	if(motivacao.value.length>3){
		resetFormAvaliarPapeis();
		document.getElementsByName('MotivacaoPeso')[0].value=motivacao.value;
		document.getElementById('motivar-peso-form').style.display='none';
		document.getElementById('form-avaliar-papeis').submit();
	} else {
		alert("Falta preencher a motivação do peso.");
		motivacao.focus();
		return;		
	}
}

function salvarRemocao(){
	let motivacao = document.getElementById('motRem_text');
	if(motivacao.value.length>3){
		resetFormAvaliarPapeis();
		document.getElementsByName('MotivacaoRemocao')[0].value=motivacao.value;
		document.getElementById('motivar-remocao-form').style.display='none';
		document.getElementById('form-avaliar-papeis').submit();
	} else {
		alert("Falta preencher a motivação da remoção.");
		motivacao.focus();
		return;		
	}
}

function openDet(btn){
	btn.disabled = true;
	document.getElementById('det-elemento-form').style.display='block';
	let entidadeId = btn.name.split("_")[1];
	let cicloId = btn.name.split("_")[2];
	let pilarId = btn.name.split("_")[3];
	let componenteId = btn.name.split("_")[4];
	let tipoNotaId = btn.name.split("_")[5];
	let elementoId = btn.name.split("_")[6];
	let peso = btn.name.split("_")[7];
	let nota = btn.name.split("_")[8];	
	document.getElementById('det-elemento-form').style.display='block';
	document.getElementById("detEntidade").value = entidadesMap.get(entidadeId);
	document.getElementById("detCiclo").value = ciclosMap.get(cicloId);
	document.getElementById("detPilar").value = pilaresMap.get(pilarId);
	document.getElementById("detComponente").value = componentesMap.get(componenteId);
	document.getElementById("detTipoNota").value = tiposNotasMap.get(tipoNotaId);
	document.getElementById("detElemento").value = elementosMap.get(elementoId);
	document.getElementById("detPeso").value = peso;
	document.getElementById("detNota").value = nota;	
	return false;
}