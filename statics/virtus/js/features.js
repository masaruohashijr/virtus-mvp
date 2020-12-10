function editFeature(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var featureId = e.parentNode.parentNode.childNodes[3].innerText;
    var featureName = e.parentNode.parentNode.childNodes[5].innerText;
    var featureDescription = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var featureCode = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('FeatureIdForUpdate').value = featureId;
    document.getElementById('FeatureNameForUpdate').value = featureName;
    document.getElementById('FeatureDescriptionForUpdate').value = featureDescription;
    document.getElementById('FeatureCodeForUpdate').value = featureCode;
    document.getElementById('FeatureNameForUpdate').focus();
}

function deleteFeature(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var featureId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('featureIdToDelete').value = featureId;
}

function loadFeaturesByRoleId(roleId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var featuresEdit = JSON.parse(xmlhttp.responseText);
				selectOptionsFeaturesForUpdate(featuresEdit);
			}
	}
	xmlhttp.open("GET","/loadFeaturesByRoleId?roleId="+roleId,true);
	xmlhttp.send();
}

function selectOptionsFeaturesForUpdate(featuresEdit){
	let s = document.getElementById("FeaturesForUpdate");
	for(n=0;n<featuresEdit.length;n++){
		for(m=0;m<s.options.length;m++){
			if(s.options[m].value == featuresEdit[n].id){
				s.options[m].selected = 'selected';
				break;
			}
		}
	}
}

function openCreateFeature(){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NameFeatureForInsert').focus();
}
