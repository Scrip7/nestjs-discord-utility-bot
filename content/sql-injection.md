**Always use prepared statements when querying a database**

With a one line change your current queries are guaranteed to be safe against dynamic data regardless of future changes in your program.
> It's somewhat shameful that there are so many successful SQL Injection attacks occurring, because it is EXTREMELY simple to avoid SQL Injection vulnerabilities in your code.
https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html