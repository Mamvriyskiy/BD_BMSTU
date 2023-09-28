--task_09(Инструкция SELECT, использующая простое выражение CASE)
SELECT c.first_name, c.last_name,
    CASE EXTRACT(YEAR FROM b.date_arrive)
        WHEN EXTRACT(YEAR FROM now()) THEN 'This Year'
        WHEN EXTRACT(YEAR FROM now()) - 1 THEN 'Last year'
        ELSE CAST(EXTRACT(YEAR FROM NOW()) - EXTRACT(YEAR FROM b.date_arrive) AS varchar(5)) || ' years ago'
    END AS "When"
FROM client c JOIN booking b ON c.client_id = b.client_id;