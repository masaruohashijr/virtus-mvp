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
		document.getElementById("AcionadoPor").value = sel.name;
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
		document.getElementById("AcionadoPor").value = sel.name;
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
		document.getElementById("AcionadoPor").value = sel.name;
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
	hierarquiaMap.set('Tipo de Nota',4);
	hierarquiaMap.set('Elemento',5);
	hierarquiaMap.set('Item',6);
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
var temp = [];
function filtraPapeis(input, tabelaNome){
  var filter, table, tr, td, i, txtValue;
  filter = input.value.toUpperCase();
  table = document.getElementById(tabelaNome);
  tr = table.getElementsByTagName("tr");
  for (i = 3; i < tr.length; i++) {
    td = tr[i].getElementsByTagName("td")[2];
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