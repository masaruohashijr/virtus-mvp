class NotasAtuais {
	constructor(cicloNota,pilarNota,componenteNota,planoNota,tipoNotaNota){
		this.cicloNota = cicloNota;
		this.pilarNota = pilarNota;
		this.componenteNota = componenteNota;
		this.planoNota = planoNota;
		this.tipoNotaNota = tipoNotaNota; 
	}
}

class PesosAtuais {
	constructor(cicloPeso,pilarPeso,componentePeso,planoPeso,tipoPesoPeso){
		this.cicloPeso = cicloPeso;
		this.pilarPeso = pilarPeso;
		this.componentePeso = componentePeso;
		this.planoPeso = planoPeso;
		this.tipoNotaPeso = tipoNotaPeso; 
	}
}

function resetFormAvaliarPlanos(){
	if(document.getElementById('form-avaliar-planos')==null){
		return;
	}
	let inputs = document.getElementById('form-avaliar-planos').elements;
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
	let planoId = sel.name.split("_")[5];
	let tipoNotaId = sel.name.split("_")[6];
	let elementoId = sel.name.split("_")[7];
	let notaAnterior = sel.name.split("_")[8];
	if(sel.value != notaAnterior){
		document.getElementById("AcionadoPor").value = sel.name;
		document.getElementById("motNota_callback").value = sel.name;
		document.getElementById("motNotaEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motNotaCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motNotaPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motNotaPlano").value = planosMap.get(planoId);
		document.getElementById("motNotaComponente").value = componentesMap.get(componenteId);
		document.getElementById("motNotaTipoNota").value = tiposNotasMap.get(tipoNotaId);
		document.getElementById("motNotaElemento").value = elementosMap.get(elementoId);
		document.getElementById("motNotaNotaAnterior").value = notaAnterior;
		document.getElementById("motNotaNovaNota").value = sel.value;
		document.getElementById('motivar-nota-form').style.display='block';
		document.getElementById("motNota_text").value = '';
		document.getElementById("motNota_text").focus();
	}
}

function motivarPeso(sel){
	let entidadeId = sel.name.split("_")[1];
	let cicloId = sel.name.split("_")[2];
	let pilarId = sel.name.split("_")[3];
	let componenteId = sel.name.split("_")[4];
	let planoId = sel.name.split("_")[5];
	let tipoNotaId = sel.name.split("_")[6];
	let elementoId = sel.name.split("_")[7];
	let pesoAnterior = sel.name.split("_")[8];
	if(sel.value != pesoAnterior){
		document.getElementById("AcionadoPor").value = sel.name;
		document.getElementById("motPeso_callback").value = sel.name;
		document.getElementById("motPesoEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motPesoCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motPesoPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motPesoPlano").value = planosMap.get(planoId);
		document.getElementById("motPesoComponente").value = componentesMap.get(componenteId);
		document.getElementById("motPesoTipoNota").value = tiposNotasMap.get(tipoNotaId);
		document.getElementById("motPesoElemento").value = elementosMap.get(elementoId);
		document.getElementById("motPesoPesoAnterior").value = pesoAnterior;
		document.getElementById("motPesoNovoPeso").value = sel.value;
		document.getElementById('motivar-peso-form').style.display='block';
		document.getElementById("motPeso_text").value = '';
		document.getElementById("motPeso_text").focus();
	}
}

function motivarRemocao(sel){
	let entidadeId = sel.name.split("_")[1];
	let cicloId = sel.name.split("_")[2];
	let pilarId = sel.name.split("_")[3];
	let componenteId = sel.name.split("_")[4];
	let auditorAnterior = sel.name.split("_")[5];
	if(sel.value != auditorAnterior && auditorAnterior != 0){
		document.getElementById("AcionadoPor").value = sel.name;
		document.getElementById("motRem_callback").value = sel.name;
		document.getElementById("motRemEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motRemCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motRemPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motRemComponente").value = componentesMap.get(componenteId);
		document.getElementById("motRemAuditorAnterior").value = auditoresMap.get(auditorAnterior);
		document.getElementById("motRemNovoAuditor").value = sel.options[sel.selectedIndex].text;
		document.getElementById('motivar-remocao-form').style.display='block';
		document.getElementById("motRem_text").value='';
		document.getElementById("motRem_text").focus();
	}
}

function resetAuditor(){
	let campoAuditorComponente = document.getElementById("motRem_callback").value;
	document.getElementsByName(campoAuditorComponente)[0].value = campoAuditorComponente.split("_")[5];
}

function resetNota(){
	let campoNotaElemento = document.getElementById("motNota_callback").value;
	document.getElementsByName(campoNotaElemento)[0].value = campoNotaElemento.split("_")[8];
}

function resetPeso(){
	let campoPesoElemento = document.getElementById("motPeso_callback").value;
	document.getElementsByName(campoPesoElemento)[0].value = campoPesoElemento.split("_")[8];
}

function salvarNotaElemento(){
	let motivacao = document.getElementById('motNota_text').value;
	if(motivacao.length>3){
		resetFormAvaliarPlanos();
		document.getElementsByName('MotivacaoNota')[0].value=motivacao;
		document.getElementById('motivar-nota-form').style.display='none';
		let xmlhttp;
		let acionadoPor = document.getElementById('AcionadoPor').value;
		let valores = acionadoPor.split("_");
		xmlhttp = new XMLHttpRequest();
		xmlhttp.onreadystatechange=function()
		{
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					var notasAtuaisJson = JSON.parse(xmlhttp.responseText);
					atualizarNotas(notasAtuaisJson, valores);
					let notaAnterior = document.getElementById('motNotaNotaAnterior').value;
					let novaNota = document.getElementById('motNotaNovaNota').value;
					let messageText = "A nota foi atualizada com sucesso de "+notaAnterior +" para "+novaNota+".";
					document.getElementById("messageText").innerText = messageText;
					document.getElementById("message").style.display="block";
					let sel = document.getElementsByName(acionadoPor)[0];
					atualizarSelectName(sel, novaNota); 
				}
		}
		let entidadeId = valores[1];
		let cicloId = valores[2];
		let pilarId = valores[3];
		let componenteId = valores[4];
		let planoId = valores[5];
		let tipoNotaId = valores[6];
		let elementoId = valores[7];
		let novaNota = document.getElementById('motNotaNovaNota').value;
		xmlhttp.open("GET","/salvarNotaElemento?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&planoId="+planoId+"&componenteId="+componenteId+"&tipoNotaId="+tipoNotaId+"&elementoId="+elementoId+"&motivacao="+motivacao+"&nota="+novaNota,true);
		xmlhttp.send();
	} else {
		let errorMsg = "Falta preencher a motivação da nota do elemento.";
		document.getElementById("Errors").innerText = errorMsg;
		document.getElementById("error-message").style.display="block";
		motivacao.focus();
		return;		
	}
}

function atualizarNotas(notasAtuaisJson, valores){
	let cicloNota = notasAtuaisJson.cicloNota;
	let pilarNota = notasAtuaisJson.pilarNota;
	let componenteNota = notasAtuaisJson.componenteNota;
	let planoNota = notasAtuaisJson.planoNota;
	let tipoNotaNota = notasAtuaisJson.tipoNotaNota;
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	let componenteId = valores[4];
	let planoId = valores[5];
	let tipoNotaId = valores[6];
	document.getElementById('CicloNota_'+entidadeId+'_'+cicloId).value = cicloNota;
	document.getElementById('PilarNota_'+entidadeId+'_'+cicloId+'_'+pilarId).value = pilarNota;
	document.getElementById('ComponenteNota_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId).value = componenteNota;
	document.getElementById('PlanoNota_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId+"_"+planoId).value = planoNota;
	document.getElementById('TipoNotaNota_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId+"_"+planoId+"_"+tipoNotaId).value = tipoNotaNota;
}

function salvarPesoElemento(){
	let motivacao = document.getElementById('motPeso_text').value;
	if(motivacao.length>3){
		resetFormAvaliarPlanos();
		document.getElementsByName('MotivacaoPeso')[0].value=motivacao;
		document.getElementById('motivar-peso-form').style.display='none';
		let xmlhttp;
		let acionadoPor = document.getElementById('AcionadoPor').value;
		let valores = acionadoPor.split("_");
		xmlhttp = new XMLHttpRequest();
		xmlhttp.onreadystatechange=function()
		{
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					var pesosAtuaisJson = JSON.parse(xmlhttp.responseText);
					atualizarPesos(pesosAtuaisJson, valores);
					let pesoAnterior = document.getElementById('motPesoPesoAnterior').value;
					let novoPeso = document.getElementById('motPesoNovoPeso').value;
					let messageText = "O peso foi atualizado com sucesso de "+pesoAnterior +" para "+novoPeso+".";
					document.getElementById("messageText").innerText = messageText;
					document.getElementById("message").style.display="block";
					let sel = document.getElementsByName(acionadoPor)[0];
					habilitarNotaElementoSelect(sel);
					atualizarSelectName(sel, novoPeso); 
				}
		}
		let entidadeId = valores[1];
		let cicloId = valores[2];
		let pilarId = valores[3];
		let componenteId = valores[4];
		let planoId = valores[5];
		let tipoNotaId = valores[6];
		let elementoId = valores[7];
		let pesoNovo = document.getElementById('motPesoNovoPeso').value;
		xmlhttp.open("GET","/salvarPesoElemento?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&planoId="+planoId+"&componenteId="+componenteId+"&tipoNotaId="+tipoNotaId+"&elementoId="+elementoId+"&motivacao="+motivacao+"&peso="+pesoNovo,true);
		xmlhttp.send();
	} else {
		let errorMsg = "Falta preencher a motivação do peso do elemento.";
		document.getElementById("Errors").innerText = errorMsg;
		document.getElementById("error-message").style.display="block";
		motivacao.focus();
		return;		
	}
}

function atualizarPesos(pesosAtuaisJson, valores){
	let pilarPeso = pesosAtuaisJson.pilarPeso;
	let componentePeso = pesosAtuaisJson.componentePeso;
	let planoPeso = pesosAtuaisJson.planoPeso;
	let tipoNotaPeso = pesosAtuaisJson.tipoNotaPeso;
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	let componenteId = valores[4];
	let planoId = valores[5];
	let tipoNotaId = valores[6];
	document.getElementById('PilarPeso_'+entidadeId+'_'+cicloId+'_'+pilarId).value = pilarPeso;
	document.getElementById('ComponentePeso_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId).value = componentePeso;
	document.getElementById('PlanoPeso_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId+"_"+planoId).value = planoPeso;
	document.getElementById('TipoNotaPeso_'+entidadeId+'_'+cicloId+'_'+pilarId+"_"+componenteId+"_"+planoId+"_"+tipoNotaId).value = tipoNotaPeso;
}


function habilitarNotaElementoSelect(selPeso){
	let selNota = selPeso.parentNode.parentNode.childNodes[16].childNodes[1];
	let desabilita = false;
	if(selPeso.value == 0){
		desabilita = true;
	}
	selNota.disabled = desabilita;
	selNota.readOnly = desabilita;
}

function atualizarSelectName(sel, novo){
	let nameSel = sel.name;
	let lastUnderscorePos = nameSel.lastIndexOf('_');
	let newName = nameSel.substr(0,lastUnderscorePos);
	newName = newName + "_"+novo;
	sel.name = newName;
}

function salvarRemocao(){
	let motivacao = document.getElementById('motRem_text').value;
	if(motivacao.length>3){
		resetFormAvaliarPlanos();
		document.getElementsByName('MotivacaoRemocao')[0].value=motivacao;
		document.getElementById('motivar-remocao-form').style.display='none';
		let xmlhttp;
		let acionadoPor = document.getElementById('AcionadoPor').value;
		let sel = document.getElementsByName(acionadoPor)[0];
		let valores = acionadoPor.split("_");
		xmlhttp = new XMLHttpRequest();
		xmlhttp.onreadystatechange=function()
		{
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					let messageText = "O auditor foi alterado com sucesso de "+
						auditoresMap.get(auditorAnterior) +
						" para "+auditoresMap.get(auditorNovo)+".";
					document.getElementById("messageText").innerText = messageText;
					document.getElementById("message").style.display="block";
					atualizarSelectName(sel, auditorNovo); 
				}
		}
		let entidadeId = valores[1];
		let cicloId = valores[2];
		let pilarId = valores[3];
		let componenteId = valores[4];
		let auditorAnterior = valores[5];
		let auditorNovo = sel.value;
		xmlhttp.open("GET","/salvarAuditorComponente?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&componenteId="+componenteId+"&motivacao="+motivacao+"&auditorNovo="+auditorNovo+"&auditorAnterior="+auditorAnterior,true);
		xmlhttp.send();
	} else {
		let errorMsg = "Falta preencher a motivação da remoção.";
		document.getElementById("Errors").innerText = errorMsg;
		document.getElementById("error-message").style.display="block";
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
	let tipoNotaId = btn.name.split("_")[6];
	let elementoId = btn.name.split("_")[7];
	let peso = btn.name.split("_")[8];
	let nota = btn.name.split("_")[9];	
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

function reduzirTodasAsLinhas(tabelaNome){
	let tb = document.getElementById(tabelaNome);
	let rows = tb.childNodes[1].rows;
	let nivel = '';
	for(i=1;i<rows.length;i++){
		nivel = rows[i].childNodes[3].innerText;
		//console.log(nivel)
		if(nivel == 'EFPC' || nivel == 'Ciclo'){
			continue;
		}
		rows[i].style.display = 'none';
	}
}

function expandirNivel(e, tabelaNome){
	let tb = document.getElementById(tabelaNome);
	let rows = tb.childNodes[1].rows;	
	let linhaNivelExpansao = parseInt(e.parentNode.parentNode.childNodes[1].innerText.trim());
	let nivelSuperior = e.parentNode.parentNode.childNodes[3].innerText.trim();
	if(nivelSuperior == "Ciclo"){
		nivelExpansao = "Pilar";
	} else if(nivelSuperior == "Pilar"){
		nivelExpansao = "Componente";
	} else if(nivelSuperior == "Componente"){
		nivelExpansao = "Plano";
	} else if(nivelSuperior == "Plano"){
		nivelExpansao = "Tipo de Nota";
	} else if(nivelSuperior == "Tipo de Nota"){
		nivelExpansao = "Elemento";
	} else if(nivelSuperior == "Elemento"){
		nivelExpansao = "Item";
	}
	for(i=0;rows.length;i++){
		if(rows[i] != null){
			numeroLinha = rows[i].childNodes[1].innerText.trim();
		} else {
			break;
		}
		if(!(/[a-zA-Z]/).test(numeroLinha) && numeroLinha > linhaNivelExpansao){
			if(rows[i] == null){
				break;
			}
			nivel = rows[i].childNodes[3].innerText.trim();
			if(nivel == nivelExpansao){
				rows[i].style.display = "table-row";
			}
			if(nivel == nivelSuperior){
				break;
			}
		}
	}
}

function reduzirNivel(e, tabelaNome){
	let tb = document.getElementById(tabelaNome);
	let rows = tb.childNodes[1].rows;
	let linhaNivelReducao = parseInt(e.parentNode.parentNode.childNodes[1].innerText);
	let nivelSuperior = e.parentNode.parentNode.childNodes[3].innerText;
	if(nivelSuperior == "Ciclo"){
		nivelReducao = "Pilar";
	} else if(nivelSuperior == "Pilar"){
		nivelReducao = "Componente";
	} else if(nivelSuperior == "Componente"){
		nivelReducao = "Plano";
	} else if(nivelSuperior == "Plano"){
		nivelReducao = "Tipo de Nota";
	} else if(nivelSuperior == "Tipo de Nota"){
		nivelReducao = "Elemento";
	} else if(nivelSuperior == "Elemento"){
		nivelReducao = "Item";
	}
	nivelHierarquicoReducao = parseInt(hierarquiaMap.get(nivelReducao));
	for(i=linhaNivelReducao;rows.length;i++){
		if(rows[i] != null){
			numeroLinha = rows[i].childNodes[1].innerText.trim();
		} else {
			break;
		}
		if(!(/[a-zA-Z]/).test(numeroLinha) && parseInt(numeroLinha) > linhaNivelReducao){
			if(rows[i] == null){
				break;
			}
			nivel = rows[i].childNodes[3].innerText.trim();
			if(parseInt(hierarquiaMap.get(nivel)) >= nivelHierarquicoReducao){
				if(rows[i].childNodes[5].innerHTML.includes('dropdown')){
					rows[i].childNodes[5].childNodes[1].innerHTML = '<i style="color: darkblue" class="ion-android-arrow-dropright hoverbtn"></i>';
				}
				rows[i].style.display = "none";
			}
			if(nivel == nivelSuperior){
				break;
			}
		}
	}
}

var hierarquiaMap = new Map();
{
	hierarquiaMap.set('Ciclo',1);
	hierarquiaMap.set('Pilar',2);
	hierarquiaMap.set('Componente',3);
	hierarquiaMap.set('Plano',4);
	hierarquiaMap.set('Tipo de Nota',5);
	hierarquiaMap.set('Elemento',6);
	hierarquiaMap.set('Item',7);
}

function expandir(e, tabelaNome){
	if(e.innerHTML.includes('dropright')){
		e.innerHTML = '<i style="color: orange" class="ion-android-arrow-dropdown hoverbtn"></i>';
		expandirNivel(e, tabelaNome);
	} else {
		e.innerHTML = '<i style="color: darkblue" class="ion-android-arrow-dropright hoverbtn"></i>';
		reduzirNivel(e, tabelaNome);
	}
}

function filtraTabela(input, tabelaNome, offset, colnum){
  var filter, table, tr, td, i, txtValue;
  filter = input.value.toUpperCase();
  table = document.getElementById(tabelaNome);
  tr = table.getElementsByTagName("tr");
  for (i = offset; i < tr.length; i++) {
    td = tr[i].getElementsByTagName("td")[colnum];
	console.log(td.innerText);
    if (td) {
      txtValue = td.textContent || td.innerText;
      if (txtValue.toUpperCase().indexOf(filter) > -1) {
        tr[i].style.display = "table-row";
      } else {
        tr[i].style.display = "none";
      }
    }       
  }
}