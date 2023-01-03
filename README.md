# Worlder Test

## Folder Architecture

![Alt text](assets/hexagonal.png?raw=true "Hexagonal")

1. Interface Layer --> from outside to inside
2. Business Layer --> usecase layer
3. Infrastructure Layer --> from inside to outside
   

       ├── ...
       ├── app                  # this folder for all application logic
            └── core            # business layer
                └── models
                └── ports
                └── usecase
            └── infrastructure  # infrastructure layer
                └── database
                └── mqtt
            └── interface       # interface layer
                └── api
            └── utils           # utility function
       ├── assets               # all image 
       ├── docker-compose.yml
       ├── Dockerfile
       ├── internal
            └── config          # global and init config
       └── main.go              # main file to run

## Software Architecture

![Alt text](assets/software-architecture.png?raw=true "Software")

![Alt text](assets/api-architecture.png?raw=true "API")

## How to Run
run: ***docker compose up -d***

# Example API
curl: 
    ***curl --location --request GET 'localhost:9000/engine/api/sensors/list'***