<!-- !!!IMPORTANT  Before Doing Other Command Run This (ONLY DO THIS IF YOU HAVENT INSTALL SCOOP AND MIGRATE) -->
1. Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
2. Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
3. scoop install migrate



<!-- Below is How to Make Migration Up And DOwn File -->
migrate create -ext sql -dir migrations/sql/ -seq name_migration_table

<!-- Below is How to Up Migrate -->
migrate -path ./migrations/sql/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

<!-- Below is How To Down Migrate -->
migrate -path ./migrations/sql/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down
