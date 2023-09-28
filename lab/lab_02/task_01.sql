--task_01 (предикат сравнения)
select distinct c.first_name, c.last_name, i.total_price from client c 
join booking as b on b.client_id = c.client_id
join invoiceforpayment as i on b.booking_id  = i.booking_id 
where c.sex = 'm' and 
i.total_price = (select MIN(total_price) from invoiceforpayment)
