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
	let valores = campo.name.split("_");
	let nomeCampo = valores[0];
	let entidadeId = valores[1];
	let cicloId = valores[2];
	let pilarId = valores[3];
	let componenteId = valores[4];
	let dataAnterior = valores[5];
	if(dataAnterior != "" && campo.value != dataAnterior){
		document.getElementById("AcionadoPor").value = campo.name;
		document.getElementById("motRepro_callback").value = campo.name;
		document.getElementById("motReproEntidade").value = entidadesMap.get(entidadeId);
		document.getElementById("motReproCiclo").value = ciclosMap.get(cicloId);
		document.getElementById("motReproPilar").value = pilaresMap.get(pilarId);
		document.getElementById("motReproComponente").value = componentesMap.get(componenteId);
		document.getElementById("motReproDataAnterior").value = formatarData(dataAnterior);
		document.getElementById("motReproNovaData").value = formatarData(campo.value);
		if(nomeCampo.startsWith('Inicia')){
			document.getElementById('motReproTituloDataAnterior').value = 'Início Anterior';
			document.getElementById('motReproTituloNovaData').value = 'Novo Início';
			document.getElementById("motReproDataAnterior").name = 'InicioAnterior';
			document.getElementById("motReproNovaData").name = 'Inicio';
		} else {
			document.getElementById('motReproTituloDataAnterior').value = 'Término Anterior';
			document.getElementById('motReproTituloNovaData').value = 'Novo Término';
			document.getElementById("motReproDataAnterior").name = 'TerminoAnterior';
			document.getElementById("motReproNovaData").name = 'Termino';
		}
		document.getElementById('motivar-reprogramacao-form').style.display='block';
		document.getElementById("motRepro_text").value='';
		document.getElementById("motRepro_text").focus();
	}
}

function salvarReprogramacao(){
	let motivacao = document.getElementById('motRepro_text').value;
	if(motivacao.length>3){
		document.getElementsByName('MotivacaoCronograma')[0].value=motivacao;
		document.getElementById('motivar-reprogramacao-form').style.display='none';
		let xmlhttp;
		let acionadoPor = document.getElementById('AcionadoPor').value;
		let campoData = document.getElementsByName(acionadoPor)[0];
		let valores = acionadoPor.split("_");
		xmlhttp = new XMLHttpRequest();
		xmlhttp.onreadystatechange=function()
		{
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					if(tipoData=='iniciaEm') {
						tipoData = 'início';						
					} else {
						tipoData = 'término';						
					}
					let messageText = "O início "+tipoData+" do cronograma do componente foi alterado com sucesso de "+
						formatarData(dataAnterior) +
						" para "+formatarData(novaData)+".";
					document.getElementById("messageText").innerText = messageText;
					document.getElementById("message").style.display="block";
					atualizarFieldName(campoData, novaData); 
				}
		}
		let tipoData = valores[0];
		if(tipoData.startsWith('IniciaEm')){
			tipoData = 'iniciaEm';
		} else {
			tipoData = 'terminaEm';
		}
		let entidadeId = valores[1];
		let cicloId = valores[2];
		let pilarId = valores[3];
		let componenteId = valores[4];
		let novaData = formatarData(document.getElementById('motReproNovaData').value);
		let nameAnt = campoData.name;
		let dataAnterior = nameAnt.substr(nameAnt.lastIndexOf('_')+1);
		xmlhttp.open("GET","/salvarReprogramacao?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&componenteId="+componenteId+"&motivacao="+motivacao+"&tipoData="+tipoData+"&dataAnterior="+dataAnterior+"&novaData="+novaData,true);
		xmlhttp.send();
	} else {
		let errorMsg = "Falta preencher a motivação da nota do elemento.";
		document.getElementById("Errors").innerText = errorMsg;
		document.getElementById("error-message").style.display="block";
		motivacao.focus();
		return;		
	}
}

function motivarConfigPlanos(){
	let entidade = document.getElementById('EntidadeConfigPlanos').value;
	let ciclo = document.getElementById('CicloConfigPlanos').value;
	let pilar = document.getElementById('PilarConfigPlanos').value;
	let componente = document.getElementById('ComponenteConfigPlanos').value;
	document.getElementById("motReconfEntidade").value = entidade;
	document.getElementById("motReconfCiclo").value = ciclo;
	document.getElementById("motReconfPilar").value = pilar;
	document.getElementById("motReconfComponente").value = componente;
	document.getElementById('motivar-reconfiguracao-form').style.display='block';
	document.getElementById("motReconf_text").value='';
	document.getElementById("motReconf_text").focus();
}

