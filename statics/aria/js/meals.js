function updatemeal(e) { // alterado
	var editForm = document.getElementById('edit-form');
	editForm.style.display = 'block';
	var idMeal = e.parentNode.parentNode.childNodes[3].innerText;
	var mealTypeId = e.parentNode.parentNode.childNodes[5].innerText;
	var date = e.parentNode.parentNode.childNodes[7].innerText;
	var startAt = e.parentNode.parentNode.childNodes[9].innerText;
	var endAt = e.parentNode.parentNode.childNodes[11].innerText;
	var bolus = e.parentNode.parentNode.childNodes[17].innerText;
	var sel = document.getElementById('MealTypeForUpdate');
	for(n=0;n<sel.options.length;n++){
		if(mealTypeId == sel.options[n].text){
			sel.options[n].selected = true;
			break;
		}
	}
	var dt = convertDate(date);
	document.getElementById('MealIdForUpdate').value=idMeal;
	document.getElementById('DateForUpdate').value= dt;
	document.getElementById('StartAtForUpdate').value=startAt;
	document.getElementById('EndAtForUpdate').value=endAt;
	document.getElementById('BolusForUpdate').value=bolus;
	loadItensByMealId(idMeal);
}

function deletemeal(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var mealId = e.parentNode.parentNode.childNodes[1].innerText; // alterado
	document.getElementById('MealIdToDelete').value = mealId;
}


function changeLabelMeasure(e){
	var labelMeasure = document.getElementById('labelMeasure');
	var id = e.options[e.selectedIndex].value;
	medidaUsual = ar[id].split("#")[1];
	labelMeasure.innerText = "Quantidade de "+medidaUsual;
}
	

	

