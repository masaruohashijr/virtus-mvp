function openDet(btn){
	btn.disabled = true;
	document.getElementById('det-elemento-form').style.display='block';
	return false;
}
function openHist(btn){
	btn.disabled = true;
	document.getElementById('hist-elemento-form').style.display='block';
	return false;
}

function resetFormAvaliarPapeis(){
	let inputs = document.getElementById('form-avaliar-papeis').elements;
	for (i = 0; i < inputs.length; i++) {
  		inputs[i].removeAttribute("disabled");
	}
}
