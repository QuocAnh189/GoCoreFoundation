# GoCoreFoundation

## 1.Migrations

This guide outlines the process for creating and managing database schema changes using the project's migration system. All migration files are located in the `migrations/up` and `migrations/down` directories.

---

### How to Create a New Migration

1.  Generate new migration files using the `make create-migration` command. You must provide a descriptive `NAME`.

    ```sh
    # This will create a pair of files in the migrations/up and migrations/down directories
    make create-migration NAME=add_email_to_users_table
    ```

2.  Add SQL statements to the newly created files.
    - In `migrations/up/YYYYMMDDHHMMSS_add_email_to_users_table.sql`, add the SQL to apply your changes.
      ```sql
      ALTER TABLE users ADD COLUMN email VARCHAR(255) UNIQUE;
      ```
    - In `migrations/down/YYYYMMDDHHMMSS_add_email_to_users_table.sql`, add the SQL to revert your changes.
      ```sql
      ALTER TABLE users DROP COLUMN email;
      ```

---

### How to Run Migrations

1.  Apply all pending migrations to the database.

    This command executes all scripts in `migrations/up/` in chronological order.

    ```sh
    make migrate-up
    ```

2.  Revert (rollback) migrations.

    This command executes all scripts in `migrations/down/` in reverse chronological order.

    ```sh
    make migrate-down
    ```
