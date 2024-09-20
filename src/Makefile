readDB:
	go run cmd/readDB/main.go -f data/original_database.xml
	go run cmd/readDB/main.go -f data/stolen_database.json

compareDB:
	go run cmd/compareDB/main.go --old data/original_database.xml --new data/stolen_database.json

compareFS:
	go run cmd/compareFS/main.go --old data/snapshot1.txt --new data/snapshot2.txt
