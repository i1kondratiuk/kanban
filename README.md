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

#### /boards/{id}/columns/{id}
* `GET`: gets a column
* `PUT` : updates a column
* `DELETE` : deletes a column

#### /boards/{id}/columns/{id}/name
* `PUT` : renames a column

#### /boards/{id}/columns/{id}/position
* `PUT` : moves a column left or right

#### /boards/{id}/columns/{id}/tasks
* `GET` : gets all tasks
* `POST` : creates a new task

#### /boards/{id}/columns/{id}/tasks/{id}
* `GET`: gets a task
* `PUT` : updates a task
* `DELETE` : deletes a task with all comments related to this task

#### /boards/{id}/columns/{id}/tasks/{id}/priority
* `PUT` : changes a task priority

#### /boards/{id}/columns/{id}/tasks/{id}/status
* `PUT` : changes a task status (moves a task across columns)

#### /boards/{id}/columns/{id}/tasks/{id}/name
* `PUT` : renames a task

#### /boards/{id}/columns/{id}/tasks/{id}/description
* `PUT` : updates a task description

#### /boards/{id}/columns/{id}/tasks/{id}/comments
* `GET` : gets all comments sorted by their creation date (from newest to oldest)
* `POST` : creates a new comment

#### /boards/{id}/columns/{id}/tasks/{id}/comments/{id}
* `PUT` : updates a comment
* `DELETE` : deletes a comment