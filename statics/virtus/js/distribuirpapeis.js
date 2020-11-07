function submeterDistribuirPapeisForm(e){
	console.log(e.parentNode.parentNode.childNodes[3].innerText);
	console.log(e.parentNode.parentNode.childNodes[7].childNodes[1].value);
	console.log(document.getElementById("EntidadeId").value);
	console.log(document.getElementById("CicloId").value);
	document.getElementById("EntidadeId").value=e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById("CicloId").value=e.parentNode.parentNode.childNodes[7].childNodes[1].value;
	document.getElementById("formulario-distribuir-papeis").submit();
}
