## Endpoints

### 1. Add a New Item
- **Method**: `POST`
- **Endpoint**: `/api/cmd`
- **Description**: Adds a new item to the list.
- **Request Body**:
  ```json
  {
    "name": "ItemName"
  }
  ```
- **Example Command**:
  ```bash
  curl -X POST http://localhost:8080/api/cmd -H "Content-Type: application/json" -d '{"name": "Item1"}'
  ```

### 2. Get All Items
- **Method**: `GET`
- **Endpoint**: `/api/cmd`
- **Description**: Retrieves all items from the list.
- **Example Command**:
  ```bash
  curl -X GET http://localhost:8080/api/cmd
  ```

### 3. Delete an Item
- **Method**: `DELETE`
- **Endpoint**: `/api/cmd`
- **Description**: Deletes an item by ID.
- **Query Parameter**: `id` (the ID of the item to delete)
- **Example Command**:
  ```bash
  curl -X DELETE "http://localhost:8080/api/cmd?id=1"
  ```

## Internal Details

- **Language**: Go
- **Dependencies**: None (standard library only)
- **Concurrency**: Synchronization is handled using `sync.Mutex`
- **Data Storage**: In-memory storage using a map