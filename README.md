# OTP Verification with Go
A simple OTP (One-Time Password) verification service using Go, Gin framework, and Twilio.

## Setup

1. **Configure Twilio Credentials**:
   Edit the `.env` file and replace the placeholder values with your actual Twilio credentials:
   ```
   TWILIO_ACCOUNT_SID=your_actual_account_sid
   TWILIO_AUTH_TOKEN=your_actual_auth_token  
   TWILIO_SERVICE_SID=your_actual_service_sid
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build and Run**:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Send OTP
- **POST** `/otp`
- **Body**: 
  ```json
  {
    "phone": "+1234567890"
  }
  ```

### Verify OTP
- **POST** `/verifyOTP`
- **Body**:
  ```json
  {
    "user": {
      "phone": "+1234567890"
    },
    "code": "123456"
  }
  ```

## Testing

You can test the endpoints using curl:

```bash
# Send OTP
curl -X POST http://localhost:8000/otp \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1234567890"}'

# Verify OTP
curl -X POST http://localhost:8000/verifyOTP \
  -H "Content-Type: application/json" \
  -d '{"user": {"phone": "+1234567890"}, "code": "123456"}'
```

## Notes

- Make sure to get a Twilio Verify Service SID from your Twilio Console
- Phone numbers should be in E.164 format (e.g., +1234567890)
- The application runs on port 8000 by default
