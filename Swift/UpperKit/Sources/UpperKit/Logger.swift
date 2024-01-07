import Foundation

public class Logger {
    public static let shared = Logger()
    
    public init() {
        log("[DEBUG] Initializing Logger")
    }
    
    public func log(_ message: String) {
        print(message)
    }
    
    deinit {
        log("[DEBUG] Deinitializing Logger")
    }
}
