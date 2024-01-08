import Foundation
import UpperKit

// Ensure all services that are used by Go are set up
GoServices.configure()

let logger = Logger()

let input = "abc"
let output = Uppercaser.uppercase(input)

logger.log("Input: \(input)", .info)
logger.log("Output: \(output)", .info)

CallbackTest.call(number: 5) { number in
    let expectedNumber = 5
    
    guard number == expectedNumber else {
        fatalError("\(number) != \(expectedNumber)")
    }
    
    logger.log("Callback called with number \(number)", .info)
}
