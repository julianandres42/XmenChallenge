Deployment 

1. Clone this repository in any folder in your computer, be sure to have installed the go sdk. 
2. In an instance of my sql, local or remote, execute the file mutants_candidates.sql. 
3. Open a terminal and, for the current session, configure the next environment variables:
   DB_USER = the user for the data base previously created
   DB_PASSWORD = The password for the data base previously created 
   DB_HOST = The host ip where the mysql instance is running
   DB_NAME = mutants
   DB_MACHINE= local
4. In the project root directory, execute the command go run main.go, and in a broser or any http client (postman, curl), 
   type the url http://localhost:8080/ , you must see the message Hello, Magneto!. 
5. Use the other endpoints as the requirement says.  
   
