function wipeRows(tableID) {
	console.log('wipeRows');
	let tableRef = document.getElementById(tableID);
	console.log('tableID: '+tableID);
	console.log('tableRef: '+tableRef);
	console.log('tableRef.childNodes.length: '+tableRef.rows.length);
	while(tableRef.rows.length>1){
		tableRef.deleteRow(1);
	}
}
