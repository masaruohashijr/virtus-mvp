function updateMealType(e) {
	var editForm = document.getElementById('edit-form');
	// display update form
	editForm.style.display = 'block';
	// get mealType id to update
	var mealTypeId = e.parentNode.parentNode.childNodes[3].innerText;
	var mealTypeName = e.parentNode.parentNode.childNodes[5].innerText;
	var mealTypeStartAt = e.parentNode.parentNode.childNodes[7].innerText;
	var mealTypeEndAt = e.parentNode.parentNode.childNodes[9].innerText;
	document.getElementById('mealTypeIdToUpdate').value = mealTypeId;
	document.getElementById('mealTypeName').value = mealTypeName;
	document.getElementById('mealTypeStartAt').value = mealTypeStartAt;
	document.getElementById('mealTypeEndAt').value = mealTypeEndAt;
}

function deleteMealType(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var mealTypeId = e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById('mealTypeIdToDelete').value = mealTypeId;
}