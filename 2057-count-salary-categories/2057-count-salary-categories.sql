# Write your MySQL query statement below
# low
select
    "Low Salary" category,
    count(*) accounts_count
from Accounts
where income < 20000
# average
union
select
    "Average Salary" category,
    count(*) accounts_count
from Accounts
where income between 20000 and 50000
# high
union
select
    "High Salary" category,
    count(*) accounts_count
from Accounts
where income > 50000