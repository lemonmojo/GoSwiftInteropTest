import Foundation

public class Logger {
    public static let shared = Logger()
    
    public init() {
        log("Initializing Logger", .debug)
    }
    
    public func log(_ message: String,
                    _ level: Level) {
        let prefix = level.prefix
        let fullMessage = "\(prefix) \(message)"
        
        print(fullMessage)
    }
    
    deinit {
        log("Deinitializing Logger", .debug)
    }
}

public extension Logger {
    enum Level: UInt32 {
        case debug
        case info
        case warn
        case error
    }
}

fileprivate extension Logger.Level {
    var prefix: String {
        switch self {
            case .debug:
                "[DEBUG]"
            case .info:
                "[INFO]"
            case .warn:
                "[WARN]"
            case .error:
                "[ERROR]"
        }
    }
}
