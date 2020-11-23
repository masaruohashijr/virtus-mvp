package db

import (
//	"log"
)

func createMembros() {
	sql := "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'alessandro' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'alexandre' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'alfredo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'alvaro' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'andre.silva' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'andre.goncalves' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'angelica' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'annette' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'antonio.portes' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'antonio.garcia' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'antonio.frainer' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'arnaldo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'carlos.silveira' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'carlos.neves' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'carlos' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'charles.silva' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'charles.dantas' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'chow' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'clovis' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'dagomar' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'dauto' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'david' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'delma' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'diogo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'douglas' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'eduardo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'eliane' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'elyson' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'enaide' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'estevam' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'evelyn' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'felipe' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'francisco.coelho' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'francisco.junior' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'germano' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'giselle' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'hamilton' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'helvio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'hilton' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'humberto' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'isabel' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'izabel' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'james' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'jorge' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'jose.chedeak' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'jose.pires' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'jose.cestari' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'jose.fernanes' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'jucinea' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'juliana' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'leonardo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'luciano.draghetti' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'luciano.pinheiro' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'luis.pugliese' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'luis.barbosa' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'luis.angoti' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'luiz' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'maique' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'marcelo.melo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'marcelo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'marcia' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'marcio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'marcus' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'maria' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'maria.silva' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'maria.pimenta' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'marina' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'mauricio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'mauricio.lundgren' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'maury' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'nercilia' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'nivea' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'otavio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'patricia' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'paulo.matsumoto' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'paulo.diniz' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'paulo.vitorino' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'pedro' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'pedro.eugenio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'peterson' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'rafael' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'raquel' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'rita' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'roberto' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'rodrigo.oliveira' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'rodrigo.abreu' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'rodrigo.andrade' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'romulo' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'sergio' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRJ' AND a.username = 'simone' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'vandeisa' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERRS' AND a.username = 'vanessa' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'veronica' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'vitor' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERPE' AND a.username = 'walter' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERSP' AND a.username = 'wander' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'wania' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERDF' AND a.username = 'wellington.marques' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
	sql = "INSERT INTO membros (usuario_id, escritorio_id) SELECT a.id, b.id FROM users a, escritorios b WHERE b.abreviatura = 'ERMG' AND a.username = 'wellington.pereira' AND NOT EXISTS (SELECT 1 FROM membros c WHERE c.usuario_id = a.id and c.escritorio_id = b.id )"
	db.Exec(sql)
}
