function wipeRows(tableID) {
	console.log('wipeRows');
	let tableRef = document.getElementById(tableID);
	console.log('tableID: '+tableID);
	console.log('tableRef: '+tableRef);
	console.log('tableRef.childNodes.length: '+tableRef.rows.length);
	for(i=1;i<tableRef.rows.length;i++){
		tableRef.deleteRow(1);
	}
}
