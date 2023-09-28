--task_07(Инструкция SELECT, использующая агрегатные функции в выражениях столбцов.)
SELECT AVG(cost_of_stay) AS "Average Cost of Stay"
FROM (SELECT i.total_price / (date_departure - date_arrive) AS cost_of_stay
        FROM booking b 
        JOIN invoiceforpayment i ON b.invoice_id = i.invoice_id
) as TotTotal
