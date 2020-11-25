" WITH R2 AS "+
"   (SELECT componente_id, "+
"           tipo_nota_id, "+
"           round(round(cast(peso_padrao AS numeric(10, 5))/cast(total AS numeric(10, 5)), 4)*cast(100 AS numeric), 2) AS peso_padrao "+
"    FROM "+
"      (WITH TMP AS "+
"         (SELECT componente_id, "+
"                 SUM(peso_padrao) AS peso_padrao "+
"          FROM elementos_componentes a "+
"          WHERE componente_id = a.componente_id "+
"          GROUP BY componente_id) SELECT a.componente_id, "+
"                                         a.tipo_nota_id, "+
"                                         tmp.peso_padrao AS total, "+
"                                         sum(a.peso_padrao) AS peso_padrao "+
"       FROM elementos_componentes a "+
"       LEFT JOIN TMP ON a.componente_id = TMP.componente_id "+
"       GROUP BY a.componente_id, "+
"                a.tipo_nota_id, "+
"                tmp.peso_padrao) R1 "+
"    ORDER BY 1, "+
"             2) "+
" UPDATE tipos_notas_componentes "+
" SET peso_padrao = R2.peso_padrao "+
" FROM R2 "+
" WHERE tipos_notas_componentes.componente_id = R2.componente_id "+
"   AND tipos_notas_componentes.tipo_nota_id = R2.tipo_nota_id "