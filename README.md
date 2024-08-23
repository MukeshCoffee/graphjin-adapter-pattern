# GraphJin Adapter Pattern

This project demonstrates how to dynamically convert GraphQL queries into SQL commands using the Adapter design pattern in GoLang. That means based on what fields or data we want to query, the system will generate the corresponding SQL command on the fly.


## Project Structure

- **cmd/main**: Contains the main entry point of the application.
- **internal/adaptee**: Contains the existing Adaptee code.
- **internal/adapter**: Contains adapter classes responsible for converting GraphQL queries to SQL commands.
- **internal/target**: Contains the interface that the adapters implement.
- **internal/target/sql_executor.go**: Defines the `SQLExecutor` interface.


## How It Works

The project dynamically generates GraphQL queries and corresponding SQL commands based on fields provided at runtime. The conversion is handled by the adapter classes (`get_user_adapter.go` and `create_user_adapter.go`).

### Dynamic Field Handling

The key feature of this project is its ability to take input fields dynamically. These fields are provided as a list or map at runtime and are used to generate both the GraphQL query and the SQL command.

### Conversion Process

1. **Adapter Pattern Implementation:**
   - The project uses the **Adapter Design Pattern** to bridge between GraphQL queries and SQL commands. Specifically, there are two adapter classes, `get_user_adapter.go` and `create_user_adapter.go`, that are responsible for handling different types of queries (like retrieving user data or creating a new user).

2. **GraphQL Query Generation:**
   - The fields provided at runtime are passed to the adapter classes. For example, in `get_user_adapter.go`, the `ExecuteSQL` method takes these fields and generates a GraphQL query string that includes only the specified fields.
   - **Example:** If you provide fields `["id", "name", "email"]`, the generated GraphQL query might look like this:
     ```graphql
     {
       user(id: 302) {
         id
         name
         email
       }
     }
     ```

3. **SQL Command Generation:**
   - After generating the GraphQL query, the same fields are used to generate the corresponding SQL command. The SQL command is structured to query or manipulate the database with the fields provided.
   - **Example:** For the same fields, the corresponding SQL command might look like this:
     ```sql
     SELECT id, name, email FROM users WHERE id = 302;
     ```

### Adapters in Action

- **Get User Adapter (`get_user_adapter.go`)**: This adapter is responsible for converting the fields related to retrieving a user (like `id`, `name`, `email`) into a GraphQL query and a corresponding SQL `SELECT` statement.
  
- **Create User Adapter (`create_user_adapter.go`)**: This adapter handles fields related to creating a new user (like `name`, `email`) and converts them into a GraphQL mutation and an SQL `INSERT` statement.

### Runtime Flexibility

The dynamic nature of the project allows us to specify different sets of fields at runtime, which gives us the flexibility. We don’t need to change the code when we want to query or manipulate different fields; we simply provide a different list of fields, and the system will handle the conversion for us.

This setup is particularly useful in scenarios where the set of fields you need to work with changes frequently or depends on user input, making the system both versatile and consistent.


### Key Files

- **main.go**: The entry point, sets up the scenario and runs the conversions.
- **get_user_adapter.go**: Adapter for fetching user details.
- **create_user_adapter.go**: Adapter for creating a new user.

### Usage

1. **Build the project**:
    ```bash
    go build ./cmd/main
    ```

2. **Run the tests**:
    ```bash
    go test ./internal/...
    ```

3. **Run the application**:
    ```bash
    go run ./cmd/main
    ```


    ### Example Output

When you run the application, you should see output that demonstrates both the old Adapter pattern and the new GraphQL to SQL Adapter pattern:

```plaintext
Demonstrating the old Adapter pattern:
Client: Adaptee's specific request

Demonstrating the new GraphQL to SQL Adapter pattern:
GraphQL Query: { user(id: 302) { id, name, email } }
SQL Command: SELECT id, name, email FROM users WHERE id = 302;
