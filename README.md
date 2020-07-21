Deployment 

1. Download the repo in any folder, be sure to have installed the go sdk in the computer. 
2. In an instance of my sql, local or remote, execute the file mutants_candidates.sql. 
3. Open a terminal and for the current session, configure the next environment variables:
   DB_USER = the user for the data base previously created
   DB_PASSWORD = The password for the data base previously created 
   DB_HOST = The host ip where the mysql instance is running
   DB_NAME = mutants
4. In the project root directory, execute the command go run main.go
   