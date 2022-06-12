# Tugas 4 -- Layanan Aplikasi Web

## How to run
1. Clone this repository
    ```
    $ git clone https://github.com/rafimuhammad01/tugas-4-law.git
    ```
2. If you have specific configuration, you can fill .env file from .env.example, if you want to use default value just go to the next step 
    ```
    $ cp .env.example .env
    $ vim .env
    ```
3. Run all applications via docker compose
    ```
    $ docker compose up -d
    ```
4. Run migrations manually by login to postgres and run this command
    ```
    CREATE TABLE IF NOT EXISTS student (
        "id" serial primary key,
        "npm" varchar(64) not null unique,
        "name" varchar(128) not null
    );
    ``` 
    -  Step to do migration manually :
        -   Execute postgres docker container
            ```
            $ docker exec -it postgres-tugas4 bin/sh
            ```
        - Go to your postgres account and database that created by docker compose
            ```
            psql -U <DB_USERNAME> -d <DB_NAME>
            ```
            note: the default value for `DB_USERNAME` is `postgres` and the default value for `DB_NAME` is `tugas4db`
        - Run this command for creating table
            ```
            CREATE TABLE IF NOT EXISTS student (
                "id" serial primary key,
                "npm" varchar(64) not null unique,
                "name" varchar(128) not null
            );
            ```
        - Exit postgres container
5. Applications is ready to use.

## API Documentation
1. GET /read/<npm>
    - Response Body
        - 200 Success   
            ```
            {
                "status": "OK",
                "data": {
                    "transaction_id": 1655034013,
                    "student": {
                        "id": 1,
                        "npm": "1906398603",
                        "name": "Rafi Muhammad"
                    }
                }
            }
            ```
        - 404 Not Found
            ```
            {
                "status": "student not found",
                "error": "student with npm 19063986031 is not exist"
            }
            ```
        - 500 Internal Server Error
            ```
            {
                "status": "internal server error",
            }
            ```
2. GET /read/<npm>/<trxID>
    - Response Body
        - 200 Success
            ```
            {
                "status": "OK",
                "data": {
                    "transaction_id": 1655030270,
                    "student": {
                        "id": 1,
                        "npm": "1906398603",
                        "name": "Rafi Muhammad"
                    }
                }
            }
            ```
        - 404 Not Found
            ```
            {
                "status": "student not found",
                "error": "student with npm 1906398603 and transaction id 165503027 is not found in cache, please use /read endpoint"
            }
            ```
        - 400 Bad Request
            ```
            {
                "status": "invalid transaction id",
                "error": "transaction id must be positive integer"
            }
            ```
         - 500 Internal Server Error
            ```
            {
                "status": "internal server error",
            }
            ```
3. POST /update
    - Request Body
         ```
            {
                "npm" : "1906398603",
                "name" : "Rafi Muhammad"
            }
         ```
    - Response Body
        - 200 Success
            ```
            {
                "status": "OK"
            }
            ```
        - 400 Bad Request
            ```
            {
                "status": "invalid npm",
                "error": "there exist student with npm 1906398603"
            }
            ```
            ```
            {
                "status": "invalid npm",
                "error": "npm should only contains number"
            }
            ```
        - 500 Internal Server Error
            ```
            {
                "status": "internal server error",
            }
            ```
            
You can try the API using this url `http://34.67.40.34`