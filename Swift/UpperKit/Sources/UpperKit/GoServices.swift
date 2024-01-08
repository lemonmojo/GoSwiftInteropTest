import Foundation
import libupper

public struct GoServices { }

public extension GoServices {
    static func configure() {
        configureLogger()
    }
}

private extension GoServices {
    static func configureLogger() {
        logger_configure { // create
            let newLogger = Logger()
            let newLoggerUnmanaged = Unmanaged.passRetained(newLogger)
            let newLoggerPtr = newLoggerUnmanaged.toOpaque()
            
            return newLoggerPtr
        } _: { loggerC in // destroy
            guard let loggerC else {
                return
            }
            
            let loggerUnmanaged = Unmanaged<Logger>.fromOpaque(loggerC)
            loggerUnmanaged.release()
        } _: { loggerC, messageC, logLevelC  in // log
            guard let loggerC,
                  let messageC else {
                return
            }
            
            let loggerUnmanaged = Unmanaged<Logger>.fromOpaque(loggerC)
            let logger = loggerUnmanaged.takeUnretainedValue()
            
            let message = String(cString: messageC)
            
            guard let logLevel = Logger.Level(rawValue: logLevelC) else {
                fatalError("Invalid log level: \(logLevelC)")
            }
            
            logger.log(message, logLevel)
        }
    }
}
