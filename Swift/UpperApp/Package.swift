// swift-tools-version: 5.9

import PackageDescription

let package = Package(
    name: "UpperApp",
    
    dependencies: [
        .package(path: "../UpperKit")
    ],
    
    targets: [
        .executableTarget(
            name: "UpperApp",
            
            dependencies: [
                .product(name: "UpperKit", package: "UpperKit")
            ]
        ),
    ]
)
