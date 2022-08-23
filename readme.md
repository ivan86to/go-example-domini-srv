# Inizializzazione Progetto

go mod init <project-name>

# Installazione Dipendenze

## Mongo Db

go get -u go.mongodb.org/mongo-driver/mongo
go get -u go.mongodb.org/mongo-driver/bson

## Gin Framework

go get -u github.com/gin-gonic/gin

## Local MongoDB instance for develop

docker pull mongo
docker run -d --name mongodb -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin mongo


# Entity Domain Model to map

-- bank domain
abi, codice associazione bancaria italiana
cab, codice avviamento bancario 
bic, Banc Identifier Code


-- currency domain
cambi_valuta,
divise, 


-- geo domain
cap,
comuni,
province,
regioni, 
stati, 

-- attività economiche 
ateco,
sae,
rae,
tae,

-- commons 
tipi_intermediari
voci_type