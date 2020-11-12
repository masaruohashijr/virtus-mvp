function validarDesignarEquipe(e){
	if (e.parentNode.parentNode.childNodes[7].childNodes[1].length == 0) {
		// Na tabela de Designação de Equipes
		// campo Select dos ciclos da entidade na linha da tabela
		console.log(false);
		return false;	
	} else {
		console.log(true);
		return true;
	}
}