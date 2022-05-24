TL;DR: MySQL is a bad database. Don't use it unless you have literally no other option.

**Why shouldn't I use MySQL?**

 - MySQL has no transactional DDL. This means that if you screw up creating a table in the middle of a migration, you cannot rollback easily.
 - MySQL has a lot of security issues. New zero-day vulnerabilities are found regularly.
 - MySQL is owned by Oracle. Oracle's primary interest is money, at the expense of everything else (like developer experience and freedom).
 - MySQL is slow. It beats Postgres etc at some things, but once you move beyond basic queries it begins to slog.
 - MySQL disconnects you randomly. Unless you setup your connection in a very specific way, it will break regularly.
 - MySQL allows all kinds of insane data to be inserted. '0000-00-00' is valid in MySQL.
 - MySQL is not ANSI compliant by default. Even when turned on, it isn't fully compliant.
 - MySQL has no feature advantage over other databases; PostgreSQL has many more useful features (such as RETURNs).

There is no reason to use MySQL over PostgreSQL, or even sqlite3. See also:
https://blog.ionelmc.ro/2014/12/28/terrible-choices-mysql
https://grimoire.ca/mysql/choose-something-else