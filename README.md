# Go Book Store Application

## A simple book store application written in GO programming language


### Tools, Libraries and Frameworks
---

- #### [GORM](https://gorm.io/) - An ORM library for GO
- #### [Gorilla MUX](https://github.com/gorilla/mux) - A request router and dispatcher for GO
- #### [Docker](https://www.docker.com/) - An open source containerization platform
- #### [MySQL](https://www.mysql.com/) - An open source relational database


### Application APIs
---
    [GET] /book/
    
    [GET] /book/<<userId>>
    
    [POST] /book/
    
    [UPDATE] /book/<<userId>>
    
    [DELETE] /book/<<userId>>
    

### Run Application
---
- **Clone the git repositiory**
```
git clone https://github.com/hasnatsaeed/go-bookstore.git
```
- **Build the docker image**
```
docker build --no-cache --tag go-bookstore .
```
- **Run the application using docker compose**
```
docker compose up
```
- **Access application APIs**
  - Get a book by ID  
  ```
  curl --location --request GET 'http://localhost:9010/book/1'
  ```
  
  - Get all books
  ```
  curl --location --request GET 'http://localhost:9010/book/'
  ```
  
  - Create a book
  ```
  curl --location --request POST 'http://localhost:9010/book/' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name":"Atomic Habits",
      "author":"john griffen",
      "publication":"penguins"
  }'
  ```
  
  - Update a book
  ```
  curl --location --request PUT 'http://localhost:9010/book/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "publication":"jordan"

  }'
  ```
  - Delete a book
  ```
  curl --location --request DELETE 'http://localhost:9010/book/1'
  ```
