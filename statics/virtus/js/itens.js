
var item_tobe_deleted;
	
class Item {
	constructor(order, id, elementoId, nome, descricao, referencia, autorId, autorNome, c_criadoEm, status, cStatus) {
		this.order = order;
		this.id = id;
		this.elementoId = elementoId;
		this.nome = nome;
		this.descricao = descricao;
		this.referencia = referencia;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.c_criadoEm = c_criadoEm;
		this.status = status;
		this.cStatus = cStatus;
	}
}

function criarItem(){
	var nome = document.getElementById('NomeItemForInsert').value;
	var descricao = document.getElementById('DescricaoItemForInsert').value;
	var referencia = document.getElementById('ReferenciaItemForInsert').value;
	var erros = '';
	if(nome==''){
		erros += 'Falta preencher o nome.\n';
		alert(erros);
		return;
	}
	itemId = getMaxId(itens);
	item = new Item(0, itemId, 0, nome, descricao, referencia, '', '', '', '', '');
	itens.push(item);
	addItemRow("table-itens-"+contexto);
	limparCamposItemForm();
	document.getElementById('create-item-form').style.display='none';
}

function getMaxId(col){
	let max = 0; 
	for(n=0;n<col.length;n++){
		if(col[n].id > max){
			max = col[n].id
		}
	}
	return max + 1;	
}

function updateItem(){
	var id = document.getElementById('id-edit').value;
	var order = document.getElementById('order-edit').value;
	var elementoId = document.getElementById('elementoId-edit').value;
	var titulo = document.getElementById('NomeItemForUpdate').value;
	var descricao = document.getElementById('DescricaoItemForUpdate').value;
	var referencia = document.getElementById('ReferenciaItemForUpdate').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(order, id, elementoId, titulo, descricao, referencia, '', '', '', '','');
	itens[order] = item;
	updateItemRow("table-itens-"+contexto,order);
	limparCamposItemForm();
	document.getElementById('edit-item-form').style.display='none';
}

function limparCamposItemForm(){
	document.getElementById('formulario-item-create').reset();
	document.getElementById('formulario-item-edit').reset();
}


function showDeleteItemForm(e){
	var deleteItemForm = document.getElementById('delete-item-form');
	deleteItemForm.style.display = 'block';
	item_tobe_deleted = e;
}

function deleteItem() {
	var order = item_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newItens = [];
	for(i=0;i<itens.length;i++){
		if(i != order){
			newItens.push(itens[i]);
		}
	}
	itens = newItens;
	item_tobe_deleted.parentNode.parentNode.innerHTML = '';
	var deleteItemForm = document.getElementById('delete-item-form');
	deleteItemForm.style.display = 'none';
}


function editItem(e) {
	var editItemForm = document.getElementById('edit-item-form');
	editItemForm.style.display = 'block';
	
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var elementoId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var nome = e.parentNode.parentNode.childNodes[0].innerText;
	var descricao = e.parentNode.parentNode.childNodes[1].innerText;
	var referencia = e.parentNode.parentNode.childNodes[2].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('elementoId-edit').value=elementoId;
	document.getElementById('NomeItemForUpdate').value=nome;
	document.getElementById('DescricaoItemForUpdate').value=descricao;
	document.getElementById('ReferenciaItemForUpdate').value=referencia;
	document.getElementById('NomeItemForUpdate').focus();
}	

function loadItensByElementoId(elementoId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var itensEdit = JSON.parse(xmlhttp.responseText);
				wipeRows("table-itens-edit")
				itens = [];
				for(order = 0;order<itensEdit.length;order++){
					itens[order]=itensEdit[order];
					addItemRow("table-itens-edit");
				}
				return itens;
			}
	}
	xmlhttp.open("GET","/loadItensByElementoId?elementoId="+elementoId,true);
	xmlhttp.send();
}

function addItemRow(tableID) {
	console.log(tableID);
	let tableRef = document.getElementById(tableID);
	console.log('tableRef: '+tableRef);
	let newRow = tableRef.insertRow(-1);
	console.log('newRow: '+newRow);
	order = itens.length-1;
	item = itens[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(item.nome);
	newCell.style = "text-align: left";
	var jsonItem = JSON.stringify(item);
	jsonItem = jsonItem.split(',').join('#');
	jsonItem = jsonItem.split('"').join('');
	jsonItem = jsonItem.split('{').join('');
	jsonItem = jsonItem.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="item'+item.id+'" value="'+jsonItem+'"/>'+newCell.innerHTML;
	//console.log(newCell.innerHTML);
	newCell.innerHTML = '<input type="hidden" name="elementoId" value="'+item.elementoId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+item.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// descricao
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(item.descricao);
	newCell.style = "display:none";
	newCell.appendChild(newText);
	// descricao
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(item.referencia);
	newCell.style = "display:none";
	newCell.appendChild(newText);
	// status
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(item.autorNome);
	newCell.appendChild(newText);
	// autor
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(item.c_criadoEm);
	newCell.appendChild(newText);
	// criado em
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(item.cStatus);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(6);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editItem(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteItemForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function updateItemRow(tableID, order){
	console.log('updateItemRow');
	let tbl = document.getElementById(tableID);
	console.log(tbl);
	let linhas = tbl.childNodes[1].childNodes;
	console.log(linhas);
	let row = tbl.childNodes[0];
	console.log(row);
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				if(inputOrder.value==order){
					row = linhas[y];
					break;
				}
			}
		}
	}
	let celula = row.childNodes[0];
	console.log(celula);
	console.log(itens[order].nome);
	celula.innerText = itens[order].nome;
	var json = JSON.stringify(itens[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="item'+order+'" value="'+json+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="elementoId" value="'+item.elementoId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+itens[order].id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = itens[order].descricao;
}

function openCreateItem(){
	document.getElementById('create-item-form').style.display='block';
	document.getElementById('NomeItemForInsert').focus();
}