function navGadget(){
	if(opened){
		closeNav();
		opened = false;
	} else {
		openNav();
		opened = true;
	}
}

function openNav() {
  document.getElementById("mySidenav").style.width = "260px";
  document.getElementById("main").style.marginLeft = "260px";
  document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
}

function closeNav() {
  document.getElementById("mySidenav").style.width = "0";
  document.getElementById("main").style.marginLeft= "0";
  document.body.style.backgroundColor = "white";
}

function expandirSubMenu(e, menu, inf, sup){
	let anchor = e;
	let tabela = e.parentNode.parentNode.parentNode;
	let nivel = inf;
	if(anchor.innerHTML.includes('dropright')){
		anchor.innerHTML = '<i style="font-size:1.3em;" class="ion-android-arrow-dropdown hoverbtn"></i>&nbsp;'+menu;		
		for(n=nivel;n<sup && n<tabela.rows.length;n++){
			tabela.rows[n].style = 'display:table-row';	
		}
	} else {
		anchor.innerHTML = '<i style="font-size:1.3em;" class="ion-android-arrow-dropright hoverbtn"></i>&nbsp;&nbsp;'+menu;
		for(n=nivel;n<sup && n<tabela.rows.length;n++){
			tabela.rows[n].style = 'display:none';	
		}
	}
}

