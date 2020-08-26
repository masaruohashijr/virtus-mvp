
var item_tobe_deleted;
	
class Detail {
	constructor(id, mealId, foodId, foodName, qtdMeasure, qtd, cho, kcal) {
	this.id = id;
	this.mealId = mealId;
	this.foodId = foodId;
	this.foodName = foodName;
	this.qtdMeasure = qtdMeasure;
	this.qtd = qtd;
	this.cho = cho;
	this.kcal = kcal;
	}
}
	function criarItem(){
		var a = document.getElementById('alimento-create');
		var foodId = a.options[a.selectedIndex].value;
		var foodName = a.options[a.selectedIndex].text;
		var qtdMedida = document.getElementById('qtdMedida-create').value;
		var qtd = document.getElementById('qtd-create').value;
		var cho = document.getElementById('cho-create').value;
		var kcal = document.getElementById('kcal-create').value;
		var erros = '';
		if(foodId=='' || qtd == ''){
			if(foodId==''){
				erros += 'Falta preencher o alimento.\n';
			}
			if(qtd==''){
				erros += 'Falta preencher a quantidade.\n';
			}
			alert(erros);
			return;
		}
		item = new Item(items.length, "",foodId, foodName, qtdMedida, qtd, cho, kcal);
		items.push(item);
		addRow("table-items-"+contexto);
		limparCamposItemForm('create');
		document.getElementById('create-item-form').style.display='none';
	}


	function editarItem(){
		var a = document.getElementById('alimento-edit');
		var id = document.getElementById('id-edit').value;
		var mealid = document.getElementById('meal-id-edit').value;
		var foodid = a.options[a.selectedIndex].value;
		var alimento = a.options[a.selectedIndex].text;
		var qtdMedida = document.getElementById('qtdMedida-edit').value;
		var qtd = document.getElementById('qtd-edit').value;
		var cho = document.getElementById('cho-edit').value;
		var kcal = document.getElementById('kcal-edit').value;
		var order = document.getElementById('order-edit').value;
		var erros = '';
		if(foodid=='' || qtd == ''){
			if(foodid==''){
				erros += 'Falta preencher o alimento.\n';
			}
			if(qtd==''){
				erros += 'Falta preencher a quantidade.\n';
			}
			alert(erros);
			return;
		}
		item = new Item(id, mealid, foodid, alimento, qtdMedida, qtd, cho, kcal);
		items[order]=item;
		updateRow("table-items-"+contexto,order);
		limparCamposItemForm('edit');
		document.getElementById('edit-item-form').style.display='none';
	}
	
	

function limparCamposItemForm(form){
	var a = document.getElementById('alimento-'+form);
	a.options[a.selectedIndex].selected = false;
	document.getElementById('qtdMedida-'+form).value = "";
	document.getElementById('qtd-'+form).value = "";
	document.getElementById('cho-'+form).value = "";
	document.getElementById('kcal-'+form).value = "";
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
	var mealId = document.getElementById('MealIdForUpdate').value;
	var foodId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var qtd = e.parentNode.parentNode.childNodes[1].innerText;
	var cho = e.parentNode.parentNode.childNodes[2].innerText;
	var kcal = e.parentNode.parentNode.childNodes[3].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('meal-id-edit').value=mealId;
	var options = document.getElementById('alimento-edit').options;
	for(i=0;i<options.length;i++){
		if(options[i].value == foodId){
			options[i].selected = true;
			break;
		}	
	}
	document.getElementById('qtd-edit').value=qtd;
	document.getElementById('cho-edit').value=cho;
	document.getElementById('kcal-edit').value=kcal;
	document.getElementById('order-edit').value=order;
}


function loadItensByMealId(idMeal){
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
	xmlhttp.open("GET","/updateMeal?idMeal="+idMeal,true);
	xmlhttp.send();
}