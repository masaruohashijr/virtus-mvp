
var item_tobe_deleted;
	
class Item {
	constructor(order, id, elementoId, titulo, descricao, avaliacao, autorId, autorNome, dataCriacao, status) {
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
	}
}

function criarItem(){
	var titulo = document.getElementById('TituloElementoForInsert').value;
	var descricao = document.getElementById('DescricaoElementoForInsert').value;
	var avaliacao = document.getElementById('AvaliacaoElementoForInsert').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(0, itens.length, 0, titulo, descricao, avaliacao, '', '', '', '');
	itens.push(item);
	addRow("table-itens-create");
	limparCamposItemForm('create');
	document.getElementById('create-item-form').style.display='none';
}


function editarItem(){
	var order = document.getElementById('order-edit').value;
	var elementoId = document.getElementById('elementoId-edit').value;
	var titulo = document.getElementById('titulo-edit').value;
	var descricao = document.getElementById('descricao-edit').value;
	var avaliacao = document.getElementById('avaliacao-edit').value;
	var autorId = document.getElementById('autorId-edit').value;
	var autorNome = document.getElementById('autorNome-edit').value;
	var dataCriacao = document.getElementById('dataCriacao-edit').value;
	var status = document.getElementById('status-edit').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(order, items.length, elementoId, titulo, descricao, avaliacao, autorId, autorNome, dataCriacao, status);
	items[order] = item;
	updateRow("table-items-"+contexto,order);
	limparCamposItemForm('edit');
	document.getElementById('edit-item-form').style.display='none';
}

function limparCamposItemForm(tipoForm){
	if(tipoForm == 'create'){
		document.getElementById('TituloElementoForInsert').value;
		document.getElementById('DescricaoElementoForInsert').value;
		document.getElementById('AvaliacaoElementoForInsert').value;
	} else {
		document.getElementById('TituloElementoForUpdate').value;
		document.getElementById('DescricaoElementoForUpdate').value;
		document.getElementById('AvaliacaoElementoForUpdate').value;
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


function updateItem(e) {
	var editItemForm = document.getElementById('edit-item-form');
	editItemForm.style.display = 'block';
	
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var titulo = e.parentNode.parentNode.childNodes[0].innerText;
	var descricao = e.parentNode.parentNode.childNodes[1].innerText;
	var avaliacao = e.parentNode.parentNode.childNodes[1].childNodes[0].value;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('TituloElementoForUpdate').value=titulo;
	document.getElementById('DescricaoElementoForUpdate').value=descricao;
	document.getElementById('AvaliacaoElementoForUpdate').value=avaliacao;
}

function loadItensByElementoId(elementoId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var itensEdit = JSON.parse(xmlhttp.responseText);
				var itensEdit = xmlhttp.responseText;
				wipeRows("table-itens-edit")
				items = [];
				for(order = 0;order<itensEdit.length;order++){
					items[order]=itemsEdit[order];
					addRow("table-itens-edit");
				}
				return itens;
			}
	}
	xmlhttp.open("GET","/loadItensByElementoId?elementoId="+elementoId,true);
	xmlhttp.send();
}