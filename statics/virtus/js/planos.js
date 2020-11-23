var plano_tobe_deleted;
	
class Plano {
	constructor(order, id, entidadeId, nome, descricao, cnpb, recursoGarantidor, modalidade, autorId, autorNome, criadoEm, status, cStatus) {
		this.order = order;
		this.id = id;
		this.entidadeId = entidadeId;
		this.nome = nome;
		this.descricao = descricao;
		this.cnpb = cnpb;
		this.recursoGarantidor = recursoGarantidor;
		this.modalidade = modalidade;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.status = status;
		this.cStatus = cStatus;
	}
}

function criarPlano(){
	console.log('criarPlano');
	var nome = document.getElementById('NomePlanoForInsert').value;
	var descricao = document.getElementById('DescricaoPlanoForInsert').value;
	var erros = '';
	if(nome==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	planoId = getMaxId(planos);
	plano = new Plano(0, planoId, 0, nome, descricao, '', '', '', '', '');
	console.log(plano)
	planos.push(plano);
	console.log(planos);
	console.log('contexto: '+contexto);
	//console.log("table-planos-"+contexto);
	addPlanoRow("table-planos-"+contexto);
	limparCamposPlanoForm();
	document.getElementById('create-plano-form').style.display='none';
}

function addPlanoRow(tableID) {
	console.log('addPlanoRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = planos.length-1;
	plano = planos[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(plano.cnpb);
	let json = JSON.stringify(plano);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="plano'+plano.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+plano.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+plano.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// modalidade
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(plano.modalidade);
	newCell.appendChild(newText);
	// recurso garantidor
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(plano.recursoGarantidor);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(3);
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editPlano(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeletePlanoForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function editPlano(e) {
	console.log('editPlano');
	var editPlanoForm = document.getElementById('edit-plano-form');
	editPlanoForm.style.display = 'block';
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var entidadeId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var nome = e.parentNode.parentNode.childNodes[0].innerText;
	var descricao = e.parentNode.parentNode.childNodes[1].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('entidadeId-edit').value=entidadeId;
	document.getElementById('NomePlanoForUpdate').value=nome;
	document.getElementById('DescricaoPlanoForUpdate').value=descricao;
}

function updatePlano() {
	console.log('updatePlano');
	var id = document.getElementById('id-edit').value;
	var order = document.getElementById('order-edit').value;
	var entidadeId = document.getElementById('entidadeId-edit').value;
	var nome = document.getElementById('NomePlanoForUpdate').value;
	var descricao = document.getElementById('DescricaoPlanoForUpdate').value;
	var erros = '';
	if(nome==''){
		erros += 'Falta preencher o título.\n';
		alert(erros);
		return;
	}
	plano = new Plano(order, id, entidadeId, nome, descricao, '', '', '', '','');
	planos[order] = plano;
	console.log("table-planos-"+contexto);
	console.log('order: '+order);
	updatePlanoRow("table-planos-"+contexto,order);
	limparCamposPlanoForm();
	document.getElementById('edit-plano-form').style.display='none';
}

function updatePlanoRow(tableID, order){
	console.log('updatePlanoRow');
	let tbl = document.getElementById(tableID);
	let linhas = tbl.childNodes[1].childNodes;
	let row = tbl.childNodes[0];
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
	celula.innerText = planos[order].nome;
	var json = JSON.stringify(planos[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="plano'+order+'" value="'+json+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="entidadeId" value="'+plano.entidadeId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+planos[order].id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = planos[order].descricao;
}

function showDeletePlanoForm(e){
	console.log('showDeletePlanoForm');
	var deletePlanoForm = document.getElementById('delete-plano-form');
	deletePlanoForm.style.display = 'block';
	plano_tobe_deleted = e;
}

function deletePlano() {
	console.log('deletePlano');
	var order = plano_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newPlanos = [];
	let tbl = plano_tobe_deleted.parentNode.parentNode.parentNode;
	let linhas = tbl.childNodes;
	contadorLinha = 1;
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				if(inputOrder.value==order){
					if(inputOrder.value == order){
						tbl.deleteRow(contadorLinha);
						break;
					}
				}
				contadorLinha ++;
			}
		}
	}
	for(i=0;i<planos.length;i++){
		if(i != order){
			newPlanos.push(planos[i]);
		}
	}
	planos = newPlanos;
	var deletePlanoForm = document.getElementById('delete-plano-form');
	deletePlanoForm.style.display = 'none';
}


function limparCamposPlanoForm(){
	console.log('limparCamposPlanoForm');
	document.getElementById('formulario-plano-create').reset()
	document.getElementById('formulario-plano-edit').reset()
}

