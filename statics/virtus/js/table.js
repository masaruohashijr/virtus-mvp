function wipeRows(tableID, elems) {
	console.log('wipeRows');
	let tableRef = document.getElementById(tableID);
	console.log('tableID: '+tableID);
	console.log('tableRef: '+tableRef);
	console.log('elems.length: '+elems.length);
	for(i=0;i<elems.length;i++){
		tableRef.deleteRow(1);
	}
}
