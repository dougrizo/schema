#!/bin/bash
protoc --go_out=../go/schema/pb --go_opt=paths=source_relative --go-grpc_out=../go/schema/pb --go-grpc_opt=paths=source_relative ./crypto.proto
protoc --go_out=../go/schema/pb --go_opt=paths=source_relative --go-grpc_out=../go/schema/pb --go-grpc_opt=paths=source_relative ./crypto_quote.proto
protoc --go_out=../go/schema/pb --go_opt=paths=source_relative --go-grpc_out=../go/schema/pb --go-grpc_opt=paths=source_relative ./stock.proto
protoc --go_out=../go/schema/pb --go_opt=paths=source_relative --go-grpc_out=../go/schema/pb --go-grpc_opt=paths=source_relative ./stock_quote.proto
