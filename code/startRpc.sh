go run service/repository/rpc/repository.go -f service/repository/rpc/etc/repository.yaml &
go run service/user/rpc/user.go -f service/user/rpc/etc/user.yaml &
go run service/user_repository/rpc/userrepository.go -f service/user_repository/rpc/etc/userrepository.yaml
