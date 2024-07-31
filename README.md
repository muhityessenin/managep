Todo List API 
=============

API Endpoints
-------------

### /users

*   GET /users: Retrieve a list of all users.
*   POST /users: Create a new user.
*   GET /users/{id}: Retrieve details of a specific user.
*   PUT /users/{id}: Update details of a specific user.
*   DELETE /users/{id}: Delete a specific user.
*   GET /users/{id}/tasks: Retrieve the list of tasks for a specific user.
*   GET /users/search?name={name}: Search users by name.
*   GET /users/search?email={email}: Search users by email.

### /tasks

*   GET /tasks: Retrieve a list of all tasks.
*   POST /tasks: Create a new task.
*   GET /tasks/{id}: Retrieve details of a specific task.
*   PUT /tasks/{id}: Update details of a specific task.
*   DELETE /tasks/{id}: Delete a specific task.
*   GET /tasks/search?title={title}: Search tasks by title.
*   GET /tasks/search?status={status}: Search tasks by status.
*   GET /tasks/search?priority={priority}: Search tasks by priority.
*   GET /tasks/search?assignee={userId}: Search tasks by assignee ID.
*   GET /tasks/search?project={projectId}: Search tasks by project ID.

### /projects

*   GET /projects: Retrieve a list of all projects.
*   POST /projects: Create a new project.
*   GET /projects/{id}: Retrieve details of a specific project.
*   PUT /projects/{id}: Update details of a specific project.
*   DELETE /projects/{id}: Delete a specific project.
*   GET /projects/{id}/tasks: Retrieve the list of tasks in a project.
*   GET /projects/search?title={title}: Search projects by title.
*   GET /projects/search?manager={userId}: Search projects by manager ID.

HTTP Methods
------------

*   GET: Retrieve a list or a single entity.
*   POST: Create a new entity.
*   PUT: Update an existing entity (by ID).
*   DELETE: Delete an entity (by ID).

HTTP Response Codes
-------------------

*   **200 OK:** Successful GET, PUT, DELETE request.
*   **201 Created:** Successful POST request.
*   **400 Bad Request:** Incorrect request.
*   **404 Not Found:** Resource not found.
*   **405 Method Not Allowed:** Unsupported HTTP method.

How to run it locally
------------------------
**Required Tools: Docker**

**git clone https://github.com/muhityessenin/managep.git**

**make build**

**make up**

**Those commands will create docker containers and run the program**
