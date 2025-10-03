## MCP‑сервер (Go) для Finam PlaceOrder

Проект предоставляет MCP‑совместимый WebSocket JSON‑RPC сервер с одним инструментом: `place_order`, который отправляет заявку через Finam TradeAPI.

### Возможности

- WebSocket‑эндпоинт: `/ws` (JSON‑RPC 2.0)
- Инструменты:
  - `place_order` — повторяет параметры Finam PlaceOrder
- Эндпоинт проверки работоспособности: `GET /healthz`

### Конфигурация

Переменные окружения:

- `SERVER_ADDR` (по умолчанию `:8080`)
- `FINAM_API_BASE_URL` (по умолчанию `https://api.finam.ru`)
- `FINAM_ACCESS_TOKEN` (обязательно)

### Локальный запуск

```bash
cd mcp-server
go mod tidy
go run ./cmd/server
# WebSocket: ws://localhost:8080/ws
```

### Docker

```bash
docker build -t mcp-finam:latest .
docker run -e FINAM_ACCESS_TOKEN=... -p 8080:8080 mcp-finam:latest
```

### Интеграция с n8n

Используйте узел MCP Client, указав адрес `ws://<host>:8080/ws`.

Пример потока:

1. Telegram Trigger → AI Agent (модель с MCP‑клиентом) → `tools/list` для обнаружения `place_order`.
2. AI Agent валидирует аргументы и вызывает `tools/call` с:

```json
{
  "name": "place_order",
  "arguments": {
    "account_id": "12345",
    "symbol": "AFLT@MISX",
    "quantity": "10",
    "side": "SIDE_BUY",
    "type": "ORDER_TYPE_LIMIT",
    "time_in_force": "TIME_IN_FORCE_DAY",
    "limit_price": "50",
    "client_order_id": "tg-20251003-001"
  }
}
```

3. Сервер пересылает запрос в Finam и возвращает идентификаторы заявки, которые модель использует для финального ответа.

### Замечания по безопасности

- Храните `FINAM_ACCESS_TOKEN` в секрете (секрет‑менеджер/переменные окружения)
- В проде размещайте за TLS‑терминацией (реверс‑прокси)

### Ссылки

- Finam PlaceOrder: https://tradeapi.finam.ru/docs/guides/grpc/orders_service/PlaceOrder
