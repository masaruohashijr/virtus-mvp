function editPilar(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var pilarId = e.parentNode.parentNode.childNodes[3].innerText;
    var pilarNome = e.parentNode.parentNode.childNodes[5].innerText;
    var pilarDescricao = e.parentNode.parentNode.childNodes[7].innerText;
    var pilarReferencia = e.parentNode.parentNode.childNodes[9].innerText;
	document.getElementById('PilarIdForUpdate').value = pilarId;
    document.getElementById('PilarNomeForUpdate').value = pilarNome;
    document.getElementById('PilarDescricaoForUpdate').value = pilarDescricao;
    document.getElementById('PilarReferenciaForUpdate').value = pilarReferencia;
document.getElementById('PilarNomeForUpdate').focus();
	loadComponentesByPilarId(pilarId);
}

function deletePilar(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var pilarId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('PilarIdToDelete').value = pilarId;
}

function loadComponentesByPilarId(pilarId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var componentesJson = JSON.parse(xmlhttp.responseText);
				wipeRows("table-componentes-pilar-edit")
				componentesPilar = [];
				for(order = 0;componentesJson != null && order<componentesJson.length;order++){
					componentesPilar[order]=componentesJson[order];
					addComponentePilarRow("table-componentes-pilar-edit");
				}
			}
	}
	xmlhttp.open("GET","/loadComponentesByPilarId?pilarId="+pilarId,true);
	xmlhttp.send(); 
}

function openCreatePilar(){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NomePilarForInsert').focus();
}