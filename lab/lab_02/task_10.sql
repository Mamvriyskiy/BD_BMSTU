--task_10(Инструкция SELECT, использующая поисковое выражение CASE.)
SELECT c.first_name, c.last_name,
    CASE 
        WHEN b.date_arrive < now() and b.date_departure > now() THEN 'lives'
        WHEN b.date_arrive > now() THEN 'will visit'
        ELSE 'visited'
    END AS "STATUS"
from client c join booking b on c.client_id = b.client_id;
