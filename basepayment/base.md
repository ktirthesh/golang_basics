Patterns used in real payment systems
<br/>Chain of Responsibility → request field validation
<br/>Specification Pattern → business rule validation
<br/>Strategy Pattern → merchant / payment method specific rules
<br/>Rule Engine style validators → dynamic rule execution
<br/>
<br/>
``` go
API Request
     │
     ▼
Request Validator (schema validation)
     │
     ▼
Business Rule Validator
     │
     ▼
Merchant Config Validator
     │
     ▼
Risk/Fraud Validator
     │
     ▼
Payment Method Validator
     │
     ▼
Create Payment Link
```
```
Field validation
    ├── amount format
    ├── currency supported
    ├── email format
    └── phone format

Business rules
    ├── min amount
    ├── max amount
    ├── expiry validation
    └── merchant limit

Merchant rules
    ├── merchant active
    ├── merchant allowed payment methods
    └── merchant risk level

Fraud checks
    ├── suspicious email
    ├── phone blacklist
    └── velocity check
```
