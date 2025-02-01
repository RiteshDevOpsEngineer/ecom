------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
GO Modules used || DEVELOPED BY RITESH SANDILYA
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
1). GO GIN Framwork
3). GORM - go get gorm.io/gorm
2). MYSQL - go get github.com/go-sql-driver/mysql
    Redis - go get github.com/go-redis/redis/v8
    Kafka -
    MongoDb -go get go.mongodb.org/mongo-driver/mongo
	         go get go.mongodb.org/mongo-driver/mongo/options
3). JWT - go get github.com/dgrijalva/jwt-go
4). JOHO ENV reader - go get github.com/joho/godotenv
5). Log -go get github.com/sirupsen/logrus



------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Middlewares
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
1). Logging Middleware : Logs incoming requests and outgoing responses
2). Recovery Middleware : Recovers from panics and prevents the application from crashing.
3). Authentication Middleware : Verifies the identity of the requester using tokens (JWT, OAuth) or API keys.
4). Request ID Middleware : Assigns a unique ID to each request for easier tracking and correlation across services
5). CORS Middleware : Handles Cross-Origin Resource Sharing (CORS) to allow or restrict resources on a web server.
6). Rate Limiting Middleware : Limits the number of requests a user can make in a given time frame to prevent abuse.
7). Request Validation Middleware : Validates incoming request payloads to ensure they meet the expected format and requirements.
8). Content Negotiation Middleware : Ensures the API responds with the correct content type (e.g., JSON, XML) based on the Accept header.
9). Compression Middleware : Compresses response bodies to reduce payload size and improve performance.
10).Session Management Middleware : Manages user sessions for authentication state in session-based authentication scenarios.
11).Metrics/Monitoring Middleware : Collects metrics for monitoring performance, usage, and errors.

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Routes
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
1). Route Name: auth
    user registration
    auth Validation
    otp Validation
    password / pin change
    email Validation

2). Route Name : order
    place order
    Order Details
    order cancel
    order details download

3). Route Name : Product
    prodcut list
    product details

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
4). Structure

├── cmd
│   ├── main.go                    # Main entry point of the application
├── config
│   └── config.go                  # Configuration loading and parsing
├── go.mod                         # Go module definition
├── go.sum                         # Go dependencies checksum
├── assets                         # Directory for storing images (e.g., for documentation)
├── internal
│   ├── adapters                   # Adapters layer
│   │   ├── cache                  # Cache-specific adapters (e.g., Redis)
│   │   ├── database               # Database-specific adapters
│   │   │   ├── mongodb.go         # MongoDB connection and setup
│   │   │   ├── mysql.go           # MySQL connection and setup
│   │   │   └── redis.go           # Redis connection and setup
│   │   ├── handler                # HTTP handlers
│   │   ├── repository             # Repository pattern implementation
│   │   │   ├── interfaces.go      # Repository interfaces
│   │   │   ├── mongodb            # MongoDB repository implementations
│   │   │   ├── mysql              # MySQL repository implementations
│   │   │   │   ├── employee_repository.go  # MySQL repository for employee
│   │   │   │   └── user_repository.go      # MySQL repository for user
│   │   │   └── redis              # Redis repository implementations
│   │   ├── routes                 # Route definitions
│   │   │   ├── auth.go            # Authentication routes
│   │   │   ├── employee_routes.go # Employee-related routes
│   │   │   ├── order.go           # Order-related routes
│   │   │   ├── product.go         # Product-related routes
│   │   │   ├── user.go            # User-related routes
│   │   │   └── user_routes.go     # User-related routes (potentially duplicate of user.go)
│   │   └── tests                  # Tests
│   │       ├── integration        # Integration tests
│   │       └── unit               # Unit tests
│   ├── core                       # Core business logic
│   │   ├── domain                 # Domain models
│   │   │   ├── employee.go        # Employee domain model
│   │   │   ├── errors.go          # Custom error definitions
│   │   │   └── user.go            # User domain model
│   │   ├── ports                  # Port definitions (interfaces for adapters)
│   │   │   ├── repository.go      # Repository port interface
│   │   │   └── user.go            # User port interface
│   │   └── services               # Service implementations
│   │       ├── employee_service.go # Employee service implementation
│   │       └── user_service.go    # User service implementation
│   ├── middleware                 # Middleware functions
│   │   ├── authCheck.go           # JWT authentication middleware
│   │   └── maintanance.go         # Maintenance mode middleware
│   └── web                        # Web-related files
│       └── static                 # Static files (HTML, CSS, JS, etc.)
├── log
│   └── app.log                    # Application log file
└── readme.md                      # Project documentation


----------------------------------------------------------------------------------------
Status Description
1). For Account
    status 1  active
    status 2 suspended
    status 3 deleted

--------------------------------------------------------------------------------------
OTP/JWT token is binded to mobile
otp auto expires

--------------------------------------------------------------------------------------

AWS manazer to store secret keys informations and env variables
Automatically otp removed on expiration


RS256


#android deveploper or ios.... non form file submit JSON
for file upload etc.. form data format

     
