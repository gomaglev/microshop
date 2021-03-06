#[WIP]

# microshop

A scaffolding using gRPC, Wire and Gorm. Generated by Protomicro.

## Structure [WIP]


```Kaitai Struct
├── api: *.proto files for generation
│   └── order
│       ├── v1
│       │   └── order.proto
│       └── item
│           ├── v1
│           │   └── item.proto
│           └── message
│               └── v1
│                   └── message.proto
└── internal
    └── app
        ├── service: generated from *.proto in api folder
        │   └── order
        │       ├── v1
        │       │   ├── order_grpc.pb.go
        │       │   ├── order.pb.go
        │       │   ├── order.pb.gw.go
        │       │   ├── order.pb.validate.go
        │       │   ├── order.svc.go
        │       │   └── order_test.go
        │       └── item
        │           ├── v1
        │           │   ├── item_grpc.pb.go
        │           │   ├── item.pb.go
        │           │   ├── item.pb.gw.go
        │           │   ├── item.pb.validate.go
        │           │   ├── item.svc.go
        │           │   └── item_test.go
        │           └── message
        │               └── v1
        │                   ├── message_grpc.pb.go
        │                   ├── message.pb.go
        │                   ├── message.pb.gw.go
        │                   ├── message.pb.validate.go
        │                   ├── message.svc.go
        │                   └── message_test.go
        ├── model
        │   ├── order.go
        │   ├── order_item.go
        │   ├── order_item_message.go
        │   ├── gorm: default is gorm, if a model use other databases, implement them
        │   │   │── entity
        │   │   │   ├── order.go
        │   │   │   ├── order_item.go
        │   │   │   └── order_item_message.go
        │   │   └── model
        │   │       ├── order.go
        │   │       ├── order_item.go
        │   │       └── order_item_message.go
        │   │── mongo: blank
        │   │   │── entity
        │   │   │   ├── order.go
        │   │   │   ├── order_item.go
        │   │   │   └── order_item_message.go
        │   │   └── model
        │   │       ├── order.go
        │   │       ├── order_item.go
        │   │       └── order_item_message.go
        │   │── elastic: blank
        │   │   │── entity
        │   │   │   ├── order.go
        │   │   │   ├── order_item.go
        │   │   │   └── order_item_message.go
        │   │   └── model
        │   │       ├── order.go
        │   │       ├── order_item.go
        │   │       └── order_item_message.go
        │   └── external: blank
        │       │── entity
        │       │   ├── order.go
        │       │   ├── order_item.go
        │       │   └── order_item_message.go
        │       └── model
        │           ├── order.go
        │           ├── order_item.go
        │           └── order_item_message.go
        └── dto
            ├── order.go
            ├── order_item.go
            └── order_item_message.go
 
```
