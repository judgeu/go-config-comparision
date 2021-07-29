# Comparision of golang config reader libs


#### Run viper
`./run-viper.sh`

#### Run configor
`./run-configor.sh`

#### Generate the results  
`./parse-result.sh  >> README.md`


### Results
| | environment.env |  environment.loglevel |  environment.loglevel is warning |  environment.loglevel is error |  database.port |  database.username |  database.database |  database.host from default value |  games |  until |  since | 
|---|---|---|---|---|---|---|---|---|---|---|---|
| go run -mod vendor viper/base1.go |  production |  warn |  true |  false |  3306 |  %s3cR3t$- |  project1 |  default-localhost |  [game1 game2 game3] |  0001-01-01T00:00:00Z |  24h0m0s | 
| VP_SINCE=600s VP_IDK=yes VP_DATABASE_PORT=5555 VP_DB_PORT=6666 VP_GAMES_4=game4 go run -mod vendor viper/base1.go |  production |  warn |  true |  false |  5555 |  %s3cR3t$- |  project1 |  default-localhost |  [game1 game2 game3] |  0001-01-01T00:00:00Z |  10m0s | 
| go run -mod vendor configor/base1.go |  production |  warn |  true |  false |  3306 |  %s3cR3t$- |  project1 |  default-localhost |  [game1 game2 game3] |  2021-01-01T12:13:14Z |  24h0m0s | 
| CONFIGOR_SINCE=600s CONFIGOR_IDK=yes CONFIGOR_DATABASE_PORT=5555 CONFIGOR_DB_PORT=6666 CONFIGOR_GAMES_4=game4 go run -mod vendor configor/base1.go |  production |  warn |  true |  false |  6666 |  %s3cR3t$- |  project1 |  default-localhost |  [game1 game2 game3] |  2021-01-01T12:13:14Z |  10m0s | 
