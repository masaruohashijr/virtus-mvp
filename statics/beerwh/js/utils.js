function regraDeTresQtd(e, form){
	var a = document.getElementById('alimento-'+form);
	var id = a.options[a.selectedIndex].value;
	var array = ar[id].split("#")
	var qtd = array[2];
	var cho = array[3];
	var kcal = array[4];
	var qtdInformada = e.value;
	var x = cho*qtdInformada/qtd;
	x = Math.round((x + Number.EPSILON) * 100) / 100
	var y = kcal*qtdInformada/qtd;
	y = Math.round((y + Number.EPSILON) * 100) / 100
	var cho = document.getElementById('cho-'+form);
	cho.value = x;
	var kcal = document.getElementById('kcal-'+form);
	kcal.value = y;
}

function regraDeTresMedida(e,form){
	var a = document.getElementById('alimento'+form);
	var id = a.options[a.selectedIndex].value;
	var array = ar[id].split("#")
	var qtd = 1;
	var cho = array[3];
	var kcal = array[4];
	var qtdInformada = e.value;
	var x = cho*qtdInformada/qtd;
	var y = kcal*qtdInformada/qtd;
	var choInput = document.getElementById('cho-'+form);
	x = Math.round((x + Number.EPSILON) * 100) / 100
	choInput.value = x;
	var kcalInput = document.getElementById('kcal-'+form);
	y = Math.round((y + Number.EPSILON) * 100) / 100
	kcalInput.value = y;
}

function convertDate(dt){
	var parts = dt.split("/");
	var dia = parts[0];
	var mes = parts[1];
	var ano = parts[2];
	return ano+"-"+mes+"-"+dia;
}

function resetFields(form){
	document.getElementById('qtdMedida-'+form).value='';
	document.getElementById('qtd-'+form).value='';
	document.getElementById('cho-'+form).value='';
	document.getElementById('kcal-'+form).value='';
}

function setTime(){
	if(document.getElementById("OrderDateForInsert")){
		var now = new Date();
		var tomorrow = new Date();
		tomorrow.setDate(tomorrow.getDate() + 1);
		document.getElementById("OrderDateForInsert").value= now.format('yyyy-mm-dd');
		document.getElementById("TakeOutDateForInsert").value= tomorrow.format('yyyy-mm-dd');
		dateFormat.masks.hammerTime = 'HH:MM';
		var vl = now.format("hammerTime");
		document.getElementById("OrderedAtForInsert").value= vl;
		document.getElementById("TakeOutAtForInsert").value= vl;
	}
}