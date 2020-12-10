package handlers

import (
	"log"
	mdl "virtus/models"
)

func ListTiposNotaByComponenteId(componenteId string) []mdl.TipoNota {
	sql := " SELECT " +
		" id, tipo_nota_id, " +
		" peso_padrao " +
		" FROM tipos_notas_componentes WHERE componente_id = $1 "
	log.Println(sql)
	rows, _ := Db.Query(sql, componenteId)
	defer rows.Close()
	var tipos []mdl.TipoNota
	var tipo mdl.TipoNota
	for rows.Next() {
		rows.Scan(
			&tipo.Id,
			&tipo.TipoNotaId,
			&tipo.PesoPadrao)
		tipos = append(tipos, tipo)
	}
	return tipos
}
