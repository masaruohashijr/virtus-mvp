
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
	var elementoId = document.getElementById('elementoId-create').value;
	var titulo = document.getElementById('titulo-create').value;
	var descricao = document.getElementById('descricao-create').value;
	var avaliacao = document.getElementById('avaliacao-create').value;
	var autorId = document.getElementById('autorId-create').value;
	var autorNome = document.getElementById('autorNome-create').value;
	var dataCriacao = document.getElementById('dataCriacao-create').value;
	var status = document.getElementById('status-create').value;
	var erros = '';
	if(titulo==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	item = new Item(0, items.length, elementoId, titulo, descricao, avaliacao, autorId, autorNome, dataCriacao, status);
	items.push(item);
	addRow("table-items-create");
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

function limparCamposItemForm(form){
	document.getElementById('order-'+form).value;
	document.getElementById('elementoId-'+form).value;
	document.getElementById('titulo-'+form).value;
	document.getElementById('descricao-'+form).value;
	document.getElementById('avaliacao-'+form).value;
	document.getElementById('autorId-'+form).value;
	document.getElementById('autorNome-'+form).value;
	document.getElementById('dataCriacao-'+form).value;
	document.getElementById('status-'+form).value;
}


function showDeleteItemForm(e){
	var deleteItemForm = document.getElementById('delete-item-form');
	deleteItemForm.style.display = 'block';
	item_tobe_deleted = e;
}

function deleteitem() {
	var order = item_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newItems = [];
	for(i=0;i<items.length;i++){
		if(i != order){
			newItems.push(items[i]);
		}
	}
	items = newItems;
	item_tobe_deleted.parentNode.parentNode.innerHTML = '';
	var deleteItemForm = document.getElementById('delete-item-form');
	deleteItemForm.style.display = 'none';
}


function updateitem(e) {
	var editItemForm = document.getElementById('edit-item-form');
	editItemForm.style.display = 'block';
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var beerId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var qtd = e.parentNode.parentNode.childNodes[1].innerText;
	var price = e.parentNode.parentNode.childNodes[2].innerText;
	var value = e.parentNode.parentNode.childNodes[3].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	var options = document.getElementById('beer-edit').options;
	for(i=0;i<options.length;i++){
		if(options[i].value == beerId){
			options[i].selected = true;
			break;
		}	
	}
	for(i=0;i<ar.length;i++){
		var reg = ar[i].split('#');
		if(reg[0]==beerId){
			stockValue = parseInt(reg[1]);
			document.getElementById('stock-edit').value=stockValue - qtd;
			break;
		}
	}
	document.getElementById('qtd-edit').value=qtd;
	document.getElementById('price-edit').value=price;
	document.getElementById('value-edit').value=value;
	document.getElementById('order-edit').value=order;
}

function loadItensByElementoId(elementoId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				// var itensEdit = JSON.parse(xmlhttp.responseText);
				var itensEdit = xmlhttp.responseText;
				alert(itensEdit);
				/*wipeRows("table-itens-edit")
				items = [];
				for(order = 0;order<itensEdit.length;order++){
					items[order]=itemsEdit[order];
					addRow("table-itens-edit");
				}
				return itens;*/
			}
	}
	alert('loadItensByElementoId');
	xmlhttp.open("GET","/loadItensByElementoId?elementoId="+elementoId,true);
	xmlhttp.send();
}