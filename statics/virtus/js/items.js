
var item_tobe_deleted;
	
class Item {
	constructor(id, beerId, beerName, qtd, price, value) {
		this.id = id;
		this.beerId = beerId;
		this.beerName = beerName;
		this.qtd = qtd;
		this.price = price;
		this.value = value;
	}
}

function validaQtd(e, form){
	var qtdValue = parseInt(e.value);
	var beerId = document.getElementById('beer-'+form).value;
	for(i=0;i<ar.length;i++){
		var reg = ar[i].split('#');
		if(reg[0]==beerId){
			stockValue = reg[1];
		}
	}
	if(qtdValue > stockValue){
		document.getElementById('qtd-'+form).focus();
		return false;
	}
	return true;
}

function calcularValorItem(e, form){
	var qtdValue = parseInt(e.value);
	var beerId = document.getElementById('beer-'+form).value;
	for(i=0;i<ar.length;i++){
		var reg = ar[i].split('#');
		if(reg[0]==beerId){
			stockValue = reg[1];
			price = reg[2];
			document.getElementById('stock-'+form).value = stockValue - qtdValue;
			document.getElementById('value-'+form).value = qtdValue*price;
			break;
		}
	}
}

function preencheCamposItem(e, form){
	for(i=0;i<ar.length;i++){
		var reg = ar[i].split('#');
		if(reg[0]==e.value){
			document.getElementById('qtd-'+form).value='';
			document.getElementById('stock-'+form).value=reg[1];
			document.getElementById('price-'+form).value=reg[2];
			document.getElementById('value-'+form).value='';
			break;
		}
	}
}

function criarItem(){
	var a = document.getElementById('beer-create');
	var beerId = a.options[a.selectedIndex].value;
	var beerName = a.options[a.selectedIndex].text;
	var qtd = document.getElementById('qtd-create').value;
	var value = document.getElementById('value-create').value;
	var erros = '';
	if(beerId=='' || qtd == ''){
		if(beerId==''){
			erros += 'Falta preencher a cerveja.\n';
		}
		if(qtd==''){
			erros += 'Falta preencher a quantidade.\n';
		}
		alert(erros);
		return;
	}
	item = new Item(items.length, beerId, beerName, qtd, price, value);
	items.push(item);
	addRow("table-items-"+contexto);
	limparCamposItemForm('create');
	document.getElementById('create-item-form').style.display='none';
}


function editarItem(){
	var a = document.getElementById('beer-edit');
	var id = document.getElementById('id-edit').value;
	var beerId = a.options[a.selectedIndex].value;
	var beerName = a.options[a.selectedIndex].text;
	var qtd = document.getElementById('qtd-edit').value;
	var price = document.getElementById('price-edit').value;
	var value = document.getElementById('value-edit').value;
	var order = document.getElementById('order-edit').value;
	var erros = '';
	if(beerId=='' || qtd == ''){
		if(beerId==''){
			erros += 'Falta preencher o alimento.\n';
		}
		if(qtd==''){
			erros += 'Falta preencher a quantidade.\n';
		}
		alert(erros);
		return;
	}
	item = new Item(id, beerId, beerName, qtd, price, value);
	items[order]=item;
	updateRow("table-items-"+contexto,order);
	limparCamposItemForm('edit');
	document.getElementById('edit-item-form').style.display='none';
}

function limparCamposItemForm(form){
	var a = document.getElementById('beer-'+form);
	a.options[a.selectedIndex].selected = false;
	document.getElementById('qtd-'+form).value = "";
	document.getElementById('stock-'+form).value = "";
	document.getElementById('price-'+form).value = "";
	document.getElementById('value-'+form).value = "";
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

function loadItemsByOrderId(idOrder){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var itemsEdit = JSON.parse(xmlhttp.responseText);
				wipeRows("table-items-"+contexto)
				items = [];
				for(order = 0;order<itemsEdit.length;order++){
					items[order]=itemsEdit[order];
					addRow("table-items-"+contexto);
				}
			return items;
			}
	}
	xmlhttp.open("GET","/loadItemsByOrderId?idOrder="+idOrder,true);
	xmlhttp.send();
}