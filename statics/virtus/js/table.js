
function wipeRows(tableID) {
	let tableRef = document.getElementById(tableID);
	for(i=0;i<items.length;i++){
		tableRef.deleteRow(1);
	}
}

function updateRow(tableID, order){
	let tableRef = document.getElementById(tableID);
	let rowNumber = 3+parseInt(order);
	let row = tableRef.childNodes[1].childNodes[rowNumber];
	let celula = row.childNodes[0];
	celula.innerText = items[order].beerName;
	var jsonItem = JSON.stringify(items[order]);
	jsonItem = jsonItem.split(',').join('#');
	jsonItem = jsonItem.split('"').join('');
	jsonItem = jsonItem.split('{').join('');
	jsonItem = jsonItem.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="item'+order+'" value="'+jsonItem+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="beerid" value="'+items[order].beerId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+items[order].id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = items[order].qtd;
	row.childNodes[2].innerText = items[order].price;
	row.childNodes[3].innerText = items[order].value;
}



function addRow(tableID) {
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.insertRow(-1);
	order = items.length-1;
	item = items[order];
	// beer
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(item.beerName);
	var jsonItem = JSON.stringify(item);
	jsonItem = jsonItem.split(',').join('#');
	jsonItem = jsonItem.split('"').join('');
	jsonItem = jsonItem.split('{').join('');
	jsonItem = jsonItem.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="item'+item.id+'" value="'+jsonItem+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="beerid" value="'+item.beerId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+item.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// qtd
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(item.qtd);
	newCell.appendChild(newText);
	// price
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(item.price);
	newCell.appendChild(newText);
	// value
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(item.value);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {updateitem(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteItemForm(btnApagar)};
	newCell.appendChild(btnApagar);
}
