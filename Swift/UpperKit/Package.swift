// swift-tools-version: 5.9

import PackageDescription
import Foundation

let libupperPath = URL(fileURLWithPath: #file)
    .deletingLastPathComponent()
    .appending(path: "Sources")
    .appending(path: "libupper")
    .path

let package = Package(
    name: "UpperKit",
    
    platforms: [
        .macOS(.v14)
    ],
    
    products: [
        .library(
            name: "UpperKit",
            targets: [ "UpperKit" ]),
    ],
    
    targets: [
        .target(
            name: "UpperKit",
            
            dependencies: [ "libupper" ],
            
            linkerSettings: [
                .unsafeFlags([
                    "-L\(libupperPath)"
                ])
            ]
        ),
        
        .systemLibrary(name: "libupper")
    ]
)
