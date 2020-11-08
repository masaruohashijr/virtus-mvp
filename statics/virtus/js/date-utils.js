function formatarData(date){
	console.log('formatarData');
	if(date=='') return date;
	let tokenBr = "/";
	let tokenEn = "-";
	let snippet = [];
	if(date.indexOf(tokenBr)>-1){
		newToken = tokenEn;
		snippet = date.split(tokenBr);
		return snippet[2]+newToken+snippet[1]+newToken+snippet[0];
	} 
	snippet = date.split(tokenEn);
	console.log(snippet.length);
	console.log(snippet);
	if(snippet.length>0){
		newToken = tokenBr
		console.log('retorno aqui');
		return snippet[2]+newToken+snippet[1]+newToken+snippet[0];
	}
	console.log('retorno date');
	return date;
}