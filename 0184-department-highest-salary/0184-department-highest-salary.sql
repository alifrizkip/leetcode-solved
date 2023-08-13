# Write your MySQL query statement below
select
    d.name Department,
    e.name Employee,
    e.salary Salary
from (select
    name,
    departmentId,
    salary,
    rank() over (
        partition by departmentId
        order by salary desc
    ) 'rank'
from Employee) e
inner join Department d
    on d.id = e.departmentId
where e.rank = 1