import Foundation
import libupper

public struct GoServices {
    public static let shared = GoServices()
    
    init() {
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
        } _: { loggerC, messageC in // log
            guard let loggerC,
                  let messageC else {
                return
            }
            
            let loggerUnmanaged = Unmanaged<Logger>.fromOpaque(loggerC)
            let logger = loggerUnmanaged.takeUnretainedValue()
            
            let message = String(cString: messageC)
            
            logger.log(message)
        }
    }
}