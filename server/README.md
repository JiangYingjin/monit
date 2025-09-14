# 🚀 Monit Server - Enterprise IT Monitoring Backend

## 📋 Project Overview

Monit Server is a high-performance, enterprise-grade IT system monitoring backend service built on **Go + Gin framework**. Featuring a sophisticated layered architecture design, it provides comprehensive machine monitoring, data collection, alert management, and real-time analytics capabilities.

## 🛠️ Technology Stack

- **Framework**: Gin (High-performance Web Framework)
- **Database**: GORM + MySQL/PostgreSQL/SQLite
- **Authentication**: JWT (JSON Web Token)
- **Logging**: Zap (High-performance Structured Logging)
- **Configuration**: Viper (Configuration Management with Hot Reload)
- **Scheduling**: Cron (Task Scheduling & Automation)
- **Documentation**: Swagger (Interactive API Documentation)
- **Caching**: Redis (Optional High-performance Caching)
- **Notifications**: SMTP (Alert Notification System)

## 🏗️ System Architecture

### Layered Architecture Design

```
┌─────────────────┐
│   Presentation  │  ← API Layer (api/v1)
├─────────────────┤
│   Business      │  ← Service Layer (service)
├─────────────────┤
│   Data Access   │  ← Model Layer (model)
├─────────────────┤
│   Infrastructure│  ← Core Layer (core, initialize)
└─────────────────┘
```

### Core Modules

#### 🔧 Core Components (core/)

- **server.go**: Server initialization and configuration management
- **viper.go**: Configuration management with hot reload capabilities
- **zap.go**: High-performance logging system initialization

#### ⚙️ Initialization Module (initialize/)

- **router.go**: Route registration and middleware configuration
- **gorm.go**: Database connection and table schema registration
- **timer.go**: Scheduled task initialization and management

#### 🛡️ Middleware Layer (middleware/)

- **jwt.go**: JWT authentication middleware supporting dual user and machine authentication
- **cors.go**: Cross-Origin Resource Sharing (CORS) request handling

#### 📊 Business Modules

##### 🖥️ Machine Management (Customize/)

- **Machine**: Comprehensive machine information management
- **MachineService**: Machine service configuration and monitoring
- **MachineWarning**: Alert rule configuration and management
- **MachineWarningLog**: Alert log recording and tracking

##### 📈 Data Management (Customize/)

- **Data**: Monitoring data storage and retrieval
- **DataType**: Data type definitions and schemas
- **ServiceTemplate**: Service monitoring templates and configurations

## 📁 Project Structure

```shell
├── api/                    # API Interface Layer
│   ├── v1/                # API Version Control
│   │   ├── Customize/     # Custom Business APIs
│   │   └── system/        # System Management APIs
│   └── enter.go           # API Entry Point
├── config/                # Configuration Structures
│   └── config.go          # Configuration Definitions
├── core/                  # Core Components
│   ├── server.go          # Server Initialization
│   ├── viper.go           # Configuration Management
│   └── zap.go             # Logging System
├── docs/                  # Swagger Documentation
├── global/                # Global Variables
│   └── global.go          # Global Object Definitions
├── initialize/            # Initialization Module
│   ├── router.go          # Route Initialization
│   ├── gorm.go            # Database Initialization
│   └── timer.go           # Scheduled Task Initialization
├── middleware/            # Middleware Layer
│   ├── jwt.go             # JWT Authentication
│   └── cors.go            # CORS Handling
├── model/                 # Data Models
│   ├── Customize/         # Custom Business Models
│   ├── system/            # System Models
│   ├── request/           # Request Structures
│   └── response/          # Response Structures
├── router/                # Route Definitions
│   ├── Customize/         # Custom Business Routes
│   └── system/            # System Routes
├── service/               # Business Logic Layer
│   ├── Customize/         # Custom Business Services
│   └── system/            # System Services
├── utils/                 # Utility Functions
│   ├── jwt.go             # JWT Utilities
│   ├── validator.go       # Parameter Validation
│   └── timer/             # Scheduled Task Utilities
├── task/                  # Scheduled Tasks
│   └── clearTable.go      # Data Cleanup Tasks
├── plugin/                # Plugin System
│   ├── email/             # Email Plugin
│   └── ws/                # WebSocket Plugin
├── resource/              # Static Resources
├── source/                # Initialization Data
├── main.go                # Application Entry Point
├── config.yaml            # Configuration File
└── go.mod                 # Go Module Definition
```

## ⭐ Core Features

### 🖥️ Machine Management

- **Machine Provisioning**: Automated SSH-based monitoring agent deployment
- **Real-time Status**: Continuous machine online status monitoring
- **Service Configuration**: Flexible service type and parameter configuration

### 📊 Data Collection

- **Real-time Data**: High-frequency monitoring data ingestion from agents
- **Multi-format Storage**: Support for diverse data type storage
- **Historical Queries**: Advanced time-range data querying capabilities

### 🚨 Alert System

- **Rule Configuration**: Sophisticated threshold-based alert configuration
- **Real-time Detection**: Instant alert triggering on data anomalies
- **Multi-channel Notifications**: Email notifications and comprehensive alert logging

### 🔐 Authentication & Authorization

- **User Authentication**: Robust JWT-based user authentication
- **Machine Authentication**: Independent machine authentication mechanism
- **Role-based Access Control**: Granular permission management system

## 🔌 API Endpoints

### 🖥️ Machine Management APIs

- `POST /machine/createMachine` - Create new machine
- `GET /machine/getMachineList` - Retrieve machine list
- `PUT /machine/updateMachine` - Update machine information
- `DELETE /machine/deleteMachine` - Remove machine

### 📊 Data APIs

- `POST /data/createData` - Create monitoring data
- `GET /data/getDataList` - Retrieve data list
- `POST /machine/uploadDataMulti` - Bulk data upload

### 🚨 Alert APIs

- `POST /machineWarning/createMachineWarning` - Create alert rule
- `GET /machineWarning/getMachineWarningList` - Retrieve alert list
- `GET /machineWarningLog/getMachineWarningLogList` - Retrieve alert logs

### 🔐 Machine Authentication APIs

- `POST /machine/machineLogin` - Machine authentication
- `POST /machine/setMachineService` - Configure machine services

## ⚙️ Configuration Guide

### 🗄️ Database Configuration

```yaml
mysql:
  path: "your-mysql-host"
  port: "3306"
  db-name: "monit"
  username: "root"
  password: "your-password"
```

### 🔑 JWT Configuration

```yaml
jwt:
  signing-key: "your-secret-key"
  expires-time: "7d"
  buffer-time: "1d"
  issuer: "monit"
```

### 📧 Email Configuration

```yaml
email:
  to: "admin@example.com"
  from: "noreply@example.com"
  host: "smtp.example.com"
  secret: "your-email-password"
  port: 587
  is-ssl: true
```

## 🚀 Deployment Guide

### 📋 Prerequisites

- Go 1.20+
- MySQL 5.7+ / PostgreSQL 12+ / SQLite 3
- Redis (Optional)

### 🛠️ Installation Steps

1. **Install Dependencies**

```bash
go mod tidy
```

2. **Configure Database**

```bash
# Modify database configuration in config.yaml
```

3. **Launch Service**

```bash
go run main.go
```

4. **Access Services**

- API Documentation: http://localhost:8888/swagger/index.html
- Health Check: http://localhost:8888/health

### 🐳 Docker Deployment

```bash
# Build Docker image
docker build -t monit-server .

# Run container
docker run -d -p 8888:8888 monit-server
```

## 👨‍💻 Development Guide

### ➕ Adding New API Endpoints

1. **Define Model** (model/)

```go
type NewModel struct {
    global.GVA_MODEL
    Name string `json:"name" gorm:"column:name"`
}
```

2. **Create Service** (service/)

```go
type NewService struct{}

func (s *NewService) CreateNewModel(model *NewModel) error {
    return global.GVA_DB.Create(model).Error
}
```

3. **Implement API** (api/v1/)

```go
type NewApi struct{}

func (api *NewApi) CreateNewModel(c *gin.Context) {
    // API implementation logic
}
```

4. **Register Routes** (router/)

```go
func (s *NewRouter) InitNewRouter(Router *gin.RouterGroup) {
    newRouter := Router.Group("new")
    var newApi = v1.ApiGroupApp.NewApi
    {
        newRouter.POST("create", newApi.CreateNewModel)
    }
}
```

### 🔧 Adding Middleware

1. **Create Middleware** (middleware/)

```go
func NewMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Middleware logic
        c.Next()
    }
}
```

2. **Register Middleware** (initialize/router.go)

```go
PrivateGroup.Use(middleware.NewMiddleware())
```

## 📊 Monitoring Metrics

### 🖥️ System Metrics

- CPU utilization rate
- Memory usage percentage
- Disk usage statistics
- Network traffic analysis

### 🔧 Service Metrics

- **MySQL**: QPS, TPS, connection count
- **Redis**: Memory usage, hit ratio
- **Nginx**: Request count, response time
- **PHP-FPM**: Process count, request handling

## 🔒 Security Features

- **🔐 JWT Authentication**: Secure user and machine authentication
- **🔒 Password Encryption**: bcrypt-based password hashing
- **🌐 CORS Configuration**: Flexible cross-origin request control
- **✅ Parameter Validation**: Strict input parameter validation
- **🛡️ SQL Injection Protection**: GORM-based SQL injection prevention

## ⚡ Performance Optimization

- **🔗 Connection Pooling**: Advanced database connection pool management
- **💾 Caching Mechanism**: Redis-based high-performance caching
- **🔄 Concurrency Control**: Single sign-on concurrency control
- **🧹 Automated Cleanup**: Scheduled cleanup of expired data

## 🔧 Troubleshooting

### ❓ Common Issues

1. **Database Connection Failure**
   - Verify database configuration
   - Confirm database service status
   - Check network connectivity

2. **JWT Authentication Failure**
   - Review JWT configuration
   - Validate token authenticity
   - Confirm signing key

3. **Agent Connection Failure**
   - Check network connectivity
   - Verify SSH configuration
   - Confirm agent status

### 📋 Log Analysis

```bash
# View service logs
tail -f server.log

# View error logs
grep "ERROR" server.log
```

---

## ⚠️ Important Notice

This service serves as the core backend for the Monit monitoring platform, responsible for data collection, storage, processing, and alerting functionalities. We strongly recommend comprehensive testing and configuration optimization before deploying in production environments.

## 🌟 Key Highlights

- **🏗️ Enterprise Architecture**: Scalable, maintainable, and robust design
- **⚡ High Performance**: Optimized for high-throughput data processing
- **🔒 Security First**: Comprehensive security measures and best practices
- **📈 Real-time Monitoring**: Sub-second data collection and alerting
- **🛠️ Developer Friendly**: Extensive documentation and easy extensibility
- **🌐 Production Ready**: Battle-tested in enterprise environments
