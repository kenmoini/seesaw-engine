# API Endpoints

## GET /health

Returns the health of the server.

### Response

```json
{
  "status": "ok"
}
```

---

## GET /api/v1/register/{uuid}

Returns the registration status of the given Agent UUID.

### Response

```json
{
  "status": "pending",
  "message": "Waiting for registration approval"
}
```

---

## POST /api/v1/register

Registers a new Agent.

### Request

```json
{
  "agent_version": "1.0.0",
  "token": "my-registration-token",
  "hostname": "my-hostname",
  "host_inventory": {
    ...
  }
}
```

### Response

```json
{
  "status": "ok",
  "message": "Registration successful",
  "uuid": "agent-uuid"
}
```