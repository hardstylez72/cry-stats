export $(grep -v '^#' ../cmd/.prod.env | xargs) 
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$CRYPAY_DB_TUNNELED_CONNECTION goose -table goose_cry_pay $@

