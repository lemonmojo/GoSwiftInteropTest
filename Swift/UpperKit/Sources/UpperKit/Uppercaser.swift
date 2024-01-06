import libupper

public struct Uppercaser {
    public init() { }
    
    public func uppercase(_ input: String) -> String {
        let output = input.withCString { cStringPtr in
            let cStringPtrMutable = UnsafeMutablePointer(mutating: cStringPtr)
            
            guard let outputC = uppercase_string(cStringPtrMutable) else {
                fatalError("uppercase_string returned nil")
            }
            
            let innerOutput = String(cString: outputC)
            
            return innerOutput
        }
        
        return output
    }
}
