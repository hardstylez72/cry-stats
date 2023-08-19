export $(grep -v '^#' ../cmd/.env | xargs) 
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$CRYPAY_DB_CONNECTION goose -table goose_cry_pay $@

