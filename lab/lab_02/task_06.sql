--task_06 (использование предикат сравнения с квантором. 
select distinct c.first_name, c.last_name, i.total_price, i.date_of_issue  from client c 
join booking as b on b.client_id = c.client_id
join invoiceforpayment as i on b.booking_id  = i.booking_id 
where i.total_price > ALL(select total_price from invoiceforpayment where 
date_of_issue between '2002-01-01' and '2003-05-01')