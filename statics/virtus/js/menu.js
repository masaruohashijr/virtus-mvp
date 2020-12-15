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

function expandirSubMenu(row, menu){
	let tabela = row.parentNode;
	console.log(tabela.rows.length);		
	let nivelSubmenu = row.rowIndex+1;
	console.log(nivelSubmenu);
	let anchor = row.childNodes[0].childNodes[1];
	console.log(anchor.innerHTML);		
	if(anchor.innerHTML.includes('dropright')){
		console.log('entrei dropright');		
		anchor.innerHTML = '<i style="font-size:1.3em;" class="ion-android-arrow-dropdown hoverbtn"></i>&nbsp;'+menu;
		console.log('entrei dropright');		
		for(n=nivelSubmenu;n<tabela.rows.length && tabela.rows[n] && tabela.rows[n].classList.contains("submenu");n++){
			tabela.rows[n].style = 'display:table-row';	
		}
	} else {
		console.log('entrei dropdown');		
		anchor.innerHTML = '<i style="font-size:1.3em;" class="ion-android-arrow-dropright hoverbtn"></i>&nbsp;&nbsp;'+menu;
		for(n=nivelSubmenu;tabela.rows.length && tabela.rows[n] && tabela.rows[n].classList.contains("submenu");n++){
			tabela.rows[n].style = 'display:none';	
		}
	}
}

