import Foundation

public class Logger {
    public static let shared = Logger()
    
    public init() { }
    
    public func log(_ message: String) {
        print(message)
    }
}
