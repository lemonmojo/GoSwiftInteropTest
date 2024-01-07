import Foundation
import UpperKit

let uppercaser = Uppercaser()

let input = "abc"
let output = uppercaser.uppercase(input)

print("Input: \(input)")
print("Output: \(output)")

let callbackTest = CallbackTest()
let number: Int32 = 5

callbackTest.call(number: number)
