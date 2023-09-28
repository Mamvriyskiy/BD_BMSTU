--task_15(Инструкция SELECT, консолидирующая данные с помощью предложения GROUP BY и предложения HAVING)
select r.categorie, count(r.categorie) from room r
group by r.categorie 
having AVG(r.price) > (select AVG(price) from room)
