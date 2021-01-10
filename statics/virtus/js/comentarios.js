var comentario_tobe_deleted;
	
class Comentario {
	constructor(order, id, chamadoId, texto, autor, c_criadoEm) {
		this.order = order;
		this.id = id;
		this.chamadoId = chamadoId;
		this.texto = texto;
		this.autor = autor;
		this.c_criadoEm = c_criadoEm;
	}
}

function openCreateComentario(){
	document.getElementById('create-comentario-anotacao-form').style.display='block';
	document.getElementById('ComentarioForInsert').focus();
}

function criarComentario(){
	console.log('criarComentario');
	var comentario = document.getElementById('ComentarioForInsert').value;
	comentarioId = getMaxId(comentarios);
	comentario = new Comentario(0, comentarioId, comentario);
	comentarios.push(comentario);
	addComentarioRow("table-comentarios-anotacao-"+contexto);	
	document.getElementById('ComentarioForInsert').value = '';
	document.getElementById('create-comentario-chamado-form').style.display='none';
}

function addComentarioRow(tableID) {
	console.log(tableID);
	let tableRef = document.getElementById(tableID);
	console.log('tableRef: '+tableRef);
	let newRow = tableRef.insertRow(-1);
	console.log('newRow: '+newRow);
	order = comentarios.length-1;
	comentario = comentarios[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(comentario.texto);
	newCell.style = "text-align: left";
	var jsonComentario = JSON.stringify(comentario);
	jsonComentario = jsonComentario.split(',').join('#');
	jsonComentario = jsonComentario.split('"').join('');
	jsonComentario = jsonComentario.split('{').join('');
	jsonComentario = jsonComentario.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="comentario'+comentario.id+'" value="'+jsonItem+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="chamadoId" value="'+comentario.chamadoId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+comentario.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// comentario
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(comentario.texto);
	newCell.appendChild(newText);
	// autor
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(comentario.autor);
	newCell.appendChild(newText);
	// criado em
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(comentario.c_criadoEm);
	newCell.appendChild(newText);
	// Botões
	newCell = newRow.insertCell(4);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editComentario(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteComentarioForm(btnApagar)};
	newCell.appendChild(btnApagar);
}
