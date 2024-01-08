import Foundation
import UpperKit

// Ensure all services that are used by Go are set up
GoServices.configure()

let input = "abc"
let output = Uppercaser.uppercase(input)

print("Input: \(input)")
print("Output: \(output)")

CallbackTest.call(number: 5) { number in
    let expectedNumber = 5
    
    guard number == expectedNumber else {
        fatalError("\(number) != \(expectedNumber)")
    }
    
    print("Callback called with number \(number)")
}
