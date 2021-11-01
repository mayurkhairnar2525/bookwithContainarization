#bookwithContainarization

#PORT - 0.0.0.0:9999
#Table Name - bookcount
#DB_Name - Bookmanagemet
#PORT - 9099
#Container_Name - book_container
 - docker run --name book-container -p 9999:3306 -e MYSQL_ROOT_PASSWORD=12345678 -d mysql
 - docker container start book-container

#Schema - CREATE TABLE `bookcount`.`bookmanagement` (
`id` INT NOT NULL,
`name` VARCHAR(45) NOT NULL,
`author` VARCHAR(45) NOT NULL,
`prices` INT NULL,
`available` VARCHAR(45) NULL,
`pagequality` VARCHAR(45) NULL,
`lauchedyear` VARCHAR(45) NULL,
`isbn` VARCHAR(45) NOT NULL,
`stock` INT NULL,
PRIMARY KEY (`id`, `isbn`));

#  docker build -t  book-containerization .
#  (Docker imageName - book-containerization)



# docker run -p 8590:9099 book-containerization
# docker run -p 9999:3306 -d -e MYSQL_ROOT_PASSWORD=12345678 mysql

#  docker-compose -f docker-compose.yaml up

CONTAINER ID   IMAGE                                       COMMAND                  CREATED          STATUS          PORTS                                                  NAMES
707f0fa37b9b   mysql                                       "docker-entrypoint.s…"   26 seconds ago   Up 25 seconds   33060/tcp, 0.0.0.0:9999->3306/tcp, :::9999->3306/tcp   bookwithcontainarization_database_1
39057e6b6ce0   bookwithcontainarization_book-application   "./containerization"     26 seconds ago   Up 25 seconds   0.0.0.0:8080->8081/tcp, :::8080->8081/tcp              bookwithcontainarization_book-application_1
