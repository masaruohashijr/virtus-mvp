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

function motivarReprogramacao(campo, titulo){
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
		if(titulo == 'início'){
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
		document.getElementsByName('MotivacaoNota')[0].value=motivacao;
		document.getElementById('motivar-reprogramacao-form').style.display='none';
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
					atualizarFieldName(sel, novaNota); 
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
		let nameAnt = document.getElementsByName(acionadoPor)[0].name;
		let newName = nameAnt.substr(0,nameAnt.lastIndexOf('_'))+'_'+novaNota;
		document.getElementsByName(acionadoPor)[0].name = newName;
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
