go run service/repository/api/repository.go -f service/repository/api/etc/repository-api.yaml &
go run service/share/api/share.go -f service/share/api/etc/share-api.yaml &
go run service/user/api/user.go -f service/user/api/etc/user-api.yaml &
go run service/user_repository/api/userrepository.go -f service/user_repository/api/etc/userrepository-api.yaml
