import libupper

public struct CallbackTest {
    public static func call(number: Int32, 
                            callback: callbackFunc) {
        libupper.test_callback(number, callback)
    }
}
