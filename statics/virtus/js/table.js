
function wipeRows(tableID) {
	let tableRef = document.getElementById(tableID);
	for(i=0;i<itens.length;i++){
		tableRef.deleteRow(1);
	}
}

function updateRow(tableID, order){
	let tableRef = document.getElementById(tableID);
	let rowNumber = 3+parseInt(order);
	let row = tableRef.childNodes[1].childNodes[rowNumber];
	let celula = row.childNodes[0];
	celula.innerText = itens[order].beerName;
	var jsonItem = JSON.stringify(itens[order]);
	jsonItem = jsonItem.split(',').join('#');
	jsonItem = jsonItem.split('"').join('');
	jsonItem = jsonItem.split('{').join('');
	jsonItem = jsonItem.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="item'+order+'" value="'+jsonItem+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="beerid" value="'+itens[order].beerId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+itens[order].id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = itens[order].qtd;
	row.childNodes[2].innerText = itens[order].price;
	row.childNodes[3].innerText = itens[order].value;
}



function addRow(tableID) {
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.insertRow(-1);
	order = itens.length-1;
	item = itens[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(item.titulo);
	var jsonItem = JSON.stringify(item);
	jsonItem = jsonItem.split(',').join('#');
	jsonItem = jsonItem.split('"').join('');
	jsonItem = jsonItem.split('{').join('');
	jsonItem = jsonItem.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="item'+item.id+'" value="'+jsonItem+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+item.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// descricao
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(item.descricao);
	newCell.innerHTML = '<input type="hidden" name="avaliacao" value="'+item.avaliacao+'"/>'+newCell.innerHTML;
	newCell.appendChild(newText);
	// status
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(item.status);
	newCell.appendChild(newText);
	// data de criação
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(item.dataCriacao);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {updateItem(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteItemForm(btnApagar)};
	newCell.appendChild(btnApagar);
}
