# **Thought Process**

## **Overview**
The application is a high-performance RESTful service designed to handle at least 10,000 requests per second. It supports deduplication, logging, and optional integrations with external services, ensuring scalability and extensibility. Below is a breakdown of the implementation and design considerations.

---

## **Implementation Details**

### **1. Basic Endpoint**
- **Route**: `/api/verve/accept`
- **Method**: `GET`
- **Parameters**:
    - `id` (integer, required): Used to track and deduplicate requests.
    - `endpoint` (string, optional): If provided, the service sends an HTTP POST request to given endpoint in query parameters.
- **Response**:
    - `"ok"` if processed successfully.
    - `"failed"` if there’s an error in processing.

### **2. Deduplication**
- **Requirement**: Ensure each `id` is counted only once per minute, even when multiple instances of the service are running behind a load balancer.
- **Implementation**:
    - **Single Instance**: In-memory deduplication using `sync.Map` (fast and lightweight).
    - **Multi-Instance**: Redis is used to maintain deduplication consistency across instances. The `SETNX` command is employed to atomically store keys with a TTL of 1 minute.

### **3. Logging**
- Every minute, the service logs the count of unique `id`s received in that timeframe.
- **Log Format**: `Unique requests in the last minute: <count>`.

### **4.HTTP Request to External Endpoint**
- When the `endpoint` query parameter is provided, the service sends an HTTP request to the given URL.
- **Extension 1**: Replaced the HTTP `GET` request with a `POST` request, including the unique count as a JSON payload in the body.

---

## **Extensions**

### **Extension 1: HTTP POST Request**
- **Implementation**:
    - Sends a POST request with the following JSON payload:
      ```json
      {
        "unique_request_count": <count>
      }
      ```
    - Logs the status of the HTTP response.

### **Extension 2: Multi-Instance Deduplication**
- **Challenge**: Handling deduplication when multiple instances are deployed behind a load balancer.
- **Solution**:
    - Introduced Redis as a centralized store for deduplication.
    - Used `SETNX` to ensure atomic operations.
    - Redis TTL automatically cleans up old keys after 1 minute.

### **Extension 3: Distributed Streaming**
- **Requirement**: Replace logging with a distributed streaming service (e.g., Kafka) to send unique request counts.
- **Implementation**:
    - Integrated Kafka as the streaming service.
    - Published the unique request count to a Kafka topic named `unique-request-count`.

---

## **Design Considerations**

### **1. Scalability**
- **Single Instance**: Efficient in-memory structures like `sync.Map` for minimal latency.
- **Multi-Instance**: Redis ensures consistency in deduplication across instances.

### **2. Performance**
- Redis and Kafka were chosen for their proven high-throughput capabilities, enabling the service to scale to handle large numbers of requests.

### **3. Extensibility**
- The modular design allows easy addition of features, such as integrating other distributed systems or changing the logging mechanism.

### **4. Fault Tolerance**
- Redis ensures atomic operations for deduplication, reducing the likelihood of race conditions.
- Kafka guarantees at-least-once delivery of messages.

---

## **Dependencies**

### **Redis**
- Purpose: Centralized deduplication across multiple instances.
- Commands Used: `SETNX`, `KEYS`, `DEL`.

### **Kafka**
- Purpose: Distributed streaming for unique request counts.
- Topic: `unique-request-count`.

### **Gin Framework**
- Used for routing and handling HTTP requests efficiently.

