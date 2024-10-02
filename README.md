RESTful API Backend for Task Management Application

This project is a fully developed RESTful API built using Go (Golang) and GORM, focusing on efficient database management. It is designed to handle common CRUD (Create, Read, Update, Delete) operations with scalability and performance in mind. While Gin is included for handling HTTP configurations, it is not used as the primary web framework. The API is structured to be integrated into future projects, such as a Flutter application, where it will serve as the backend.
Overview

The primary goal of this project is to create a robust backend for a task management application developed in Flutter. By incorporating this REST API into the Flutter project, the application will gain enhanced functionality and backend support for features like user management, data storage, and real-time updates.
Key Features
1. User Management

    User Registration and Login: Secure user authentication processes are implemented using bcrypt for password hashing, ensuring that user credentials are stored securely.
    Account Deletion: Users have the ability to delete their accounts. The backend handles the removal of user data from the database while maintaining data integrity.
    Find User by ID: The API provides endpoints to retrieve user information based on unique user IDs.

2. Task Management

    Add New Tasks: Users can add new tasks, which are stored in the MySQL database via GORM.
    CRUD Operations on Tasks: Full control over tasks, including creating, reading, updating, and deleting.
    Task Approval, Rejection, and Fixation: The backend supports workflow management features, allowing tasks to be approved, rejected, or marked as fixed.
    Find Tasks by ID: Retrieve specific tasks using their unique IDs.
    Statistics Display: View statistics related to tasks, such as the number of tasks completed or pending.

3. Technologies and Libraries Used

    Go (Golang): Chosen for its efficiency and performance in building scalable backend services.
    GORM: An ORM library for Go that simplifies database interactions.
    MySQL Database: Used for data storage, managed locally via DBeaver.
    Bcrypt: Employed for hashing passwords, enhancing security.
    String Conversion Libraries: Utilized for data type conversions and manipulations.
    Gin Components: While Gin is not used as the main web framework, certain components or middleware are utilized for handling HTTP requests and configurations.
    HTTP Configurations: Essential for setting up routes, middleware, and managing HTTP requests and responses.

4. Architectural Pattern - MVC (Model-View-Controller)

The application follows the MVC architectural pattern for better organization and maintainability.
    Models: Define the data structures and business logic of the application.
    Controllers: Handle incoming HTTP requests, interact with models, and return responses.
