source venv/bin/activate
arch -x86_64 python python/main.py

go run services/manager/main.go
go run services/stream/main.go