#ifndef DISABLE_LOGGING
#define LOG(format, ...) \
    if LoggingEnabled { \
        fmt.Printf("LOG(macro): " + format +"\n", __VA_ARGS__); \
    }
#else
#define LOG(...)
#endif