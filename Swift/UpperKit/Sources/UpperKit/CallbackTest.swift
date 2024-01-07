import libupper

public struct CallbackTest {
    public init() { }
    
    public func call(number: Int32) {
        libupper.test_callback(number) { number in
            print("Callback called with number \(number)")
        }
    }
}
