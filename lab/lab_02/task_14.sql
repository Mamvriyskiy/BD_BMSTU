--task_14(Инструкция SELECT, консолидирующая данные с помощью предложения GROUP BY, но без предложения HAVING)
select r.categorie, count(r.categorie) from room r
group by r.categorie
