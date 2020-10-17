
var item_tobe_deleted;
	
class Item {
	constructor(order, id, elementoId, titulo, descricao, avaliacao, autorId, autorNome, dataCriacao, status, cStatus) {
		this.order = order;
		this.id = id;
		this.elementoId = elementoId;
		this.titulo = titulo;
		this.descricao = descricao;
		this.avaliacao = avaliacao;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.dataCriacao = dataCriacao;
		this.status = status;
		this.cStatus = cStatus;
	}
}

function criarItem(){
	var titulo = document.getElementById('NomeItemForInsert').value;
	var descricao = document.getElementById('DescricaoItemForInsert').value;
	var avaliacao = document.getElementById('AvaliacaoItemForInsert').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(0, itens.length, 0, titulo, descricao, avaliacao, '', '', '', '', '');
	itens.push(item);
	addRow("table-itens-"+contexto);
	limparCamposItemForm(contexto);
	document.getElementById('create-item-form').style.display='none';
}


function updateItem(){
	var id = document.getElementById('id-edit').value;
	var order = document.getElementById('order-edit').value;
	var elementoId = document.getElementById('elementoId-edit').value;
	var titulo = document.getElementById('NomeItemForUpdate').value;
	var descricao = document.getElementById('DescricaoItemForUpdate').value;
	var avaliacao = document.getElementById('AvaliacaoItemForUpdate').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(order, id, elementoId, titulo, descricao, avaliacao, '', '', '', '','');
	itens[order] = item;
	updateRow("table-itens-edit",order);
	limparCamposItemForm('edit');
	document.getElementById('edit-item-form').style.display='none';
}

function limparCamposItemForm(tipoForm){
	if(tipoForm == 'create'){
		document.getElementById('NomeItemForInsert').value = '';
		document.getElementById('DescricaoItemForInsert').value = '';
		document.getElementById('AvaliacaoItemForInsert').value = '';
	} else {
		document.getElementById('NomeItemForUpdate').value = '';
		document.getElementById('DescricaoItemForUpdate').value = '';
		document.getElementById('AvaliacaoItemForUpdate').value = '';
	}
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
	var titulo = e.parentNode.parentNode.childNodes[0].innerText;
	var descricao = e.parentNode.parentNode.childNodes[1].innerText;
	var avaliacao = e.parentNode.parentNode.childNodes[1].childNodes[0].value;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('elementoId-edit').value=elementoId;
	document.getElementById('NomeItemForUpdate').value=titulo;
	document.getElementById('DescricaoItemForUpdate').value=descricao;
	document.getElementById('AvaliacaoItemForUpdate').value=avaliacao;
}	

function loadItensByElementoId(elementoId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var itensEdit = JSON.parse(xmlhttp.responseText);
				wipeRows("table-itens-edit", itens)
				itens = [];
				for(order = 0;order<itensEdit.length;order++){
					itens[order]=itensEdit[order];
					addRow("table-itens-edit");
				}
				return itens;
			}
	}
	xmlhttp.open("GET","/loadItensByElementoId?elementoId="+elementoId,true);
	xmlhttp.send();
}