# kanban

## API

#### /boards
* `GET` : gets all boards sorted by their position
* `POST` : creates a new board

#### /boards/{id}
* `GET`: gets a board
* `PUT` : updates a board
* `DELETE` : deletes a board

#### /boards/{id}/columns
* `GET` : gets all columns
* `POST` : creates a new column

#### /columns/{id}
* `GET`: gets a column
* `PUT` : updates a column
* `DELETE` : deletes a column

#### /columns/{id}/name
* `PUT` : renames a column

#### /columns/{id}/position
* `PUT` : moves a column left or right

#### /columns/{id}/tasks
* `GET` : gets all tasks
* `POST` : creates a new task

#### /tasks/{id}
* `GET`: gets a task
* `PUT` : updates a task
* `DELETE` : deletes a task with all comments related to this task

#### /tasks/{id}/priority
* `PUT` : changes a task priority

#### /tasks/{id}/status
* `PUT` : changes a task status (moves a task across columns)

#### /tasks/{id}/name
* `PUT` : renames a task

#### /tasks/{id}/description
* `PUT` : updates a task description

#### /tasks/{id}/comments
* `GET` : gets all comments sorted by their creation date (from newest to oldest)
* `POST` : creates a new comment

#### /comments/{id}
* `PUT` : updates a comment
* `DELETE` : deletes a comment