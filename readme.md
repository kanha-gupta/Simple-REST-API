## Endpoints

### Execute a command
- **Method**: `POST`
- **Endpoint**: `/api/cmd`
- **Description**: Execute command using query parameters

- **Example Command**:
  ```bash
  curl -X POST http://localhost:8080/api/cmd?command=ls
  ```

- **Example output**:
- ```
  {"output":"go.mod\nmain.go\nreadme.md\ntest.txt\n"}
  ```
  
- **Example Command**:
  ```bash
  curl -X POST http://localhost:8080/api/cmd?command=cat%20test.txt
  ```
  
- **Example output**:
  ```
  {"output":"This is a test file\n"}
  ```
