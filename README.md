# Quiz Test RESTful API in Go

This project is a Quiz Test RESTful API implemented in Go programming language. It allows users to create quizzes, take quizzes, and view quiz results. The API is designed to be simple, efficient, and easy to use.

## Getting Started

To get started with the Quiz Test RESTful API, follow the instructions below:

- Go programming language (version 1.15 or higher)
- MongoDB (Make sure it's installed and running on your system)

## Installation

Clone this repository to your local machine.

```bash
git clone https://github.com/BaseMax/QuizTestAPIGo.git
```

Change into the project directory.

```bash
cd QuizTestAPIGo
```

Install the required dependencies.
```bash
go get ./...
```

Rename the `.env.example` file to `.env` and update the configuration with your MongoDB connection details and any other required settings.

## Running the API

To run the API server, execute the following command:

```go
go run main.go
```

By default, the server will run on localhost:8080.

## API Endpoints

The following endpoints are available in the Quiz Test API:

- **POST** `/api/quizzes` - Create a new quiz.
- **GET** `/api/quizzes/{id}` - Retrieve a specific quiz by ID.
- **GET** `/api/quizzes` - Retrieve all quizzes.
- **PUT** `/api/quizzes/{id}` - Update an existing quiz by ID.
- **DELETE** `/api/quizzes/{id}` - Delete a quiz by ID.
- **POST** `/api/quizzes/{id}/questions` - Add a new question to a specific quiz by ID.
- **PUT** `/api/quizzes/{id}/questions/{questionID}` - Update a question within a specific quiz by ID and question ID.
- **DELETE** `/api/quizzes/{id}/questions/{questionID}` - Delete a question within a specific quiz by ID and question ID.
- **GET** `/api/quizzes/{id}/questions` - Retrieve all questions for a specific quiz by ID.
- **POST** `/api/quizzes/{id}/submit` - Submit a quiz with answers.
- **GET** `/api/results/{id}` - Retrieve the result of a submitted quiz by ID.
- **GET** `/api/results` - Retrieve all submitted quiz results.
- **GET** `/api/quizzes/{id}/results` - Retrieve all submitted results for a specific quiz by ID.
- **POST** `/api/questions` - Create a new question.
- **GET** `/api/questions/{id}` - Retrieve a specific question by ID.
- **GET** `/api/questions` - Retrieve all questions.
- **PUT** `/api/questions/{id}` - Update an existing question by ID.
- **DELETE** `/api/questions/{id}` - Delete a question by ID.
- **POST** `/api/questions/{id}/options` - Add a new option to a specific question by ID.
- **PUT** `/api/questions/{id}/options/{optionID}` - Update an option within a specific question by ID and option ID.
- **DELETE** `/api/questions/{id}/options/{optionID}` - Delete an option within a specific question by ID and option ID.
- **GET** `/api/quizzes/{id}/questions/{questionID}` - Retrieve a specific question within a quiz by its ID.
- **GET** `/api/quizzes/{id}/questions/{questionID}/options` - Retrieve all options for a specific question within a quiz by its ID.
- **GET** `/api/quizzes/{id}/questions/{questionID}/options/{optionID}` - Retrieve a specific option for a specific question within a quiz by its ID and option ID.
- **PUT** `/api/quizzes/{id}/questions/{questionID}/options/{optionID}` - Update an option within a specific question within a quiz by its ID and option ID.
- **DELETE** `/api/quizzes/{id}/questions/{questionID}/options/{optionID}` - Delete an option within a specific question within a quiz by its ID and option ID.
- **POST** `/api/results` - Submit quiz answers and get the result instantly (without storing the submission).
- **GET** `/api/quizzes/{id}/results/statistics` - Retrieve statistics for a specific quiz, such as the number of times it has been taken and the average score.
- **GET** `/api/quizzes/{id}/questions/random` - Retrieve a random question from a specific quiz.

## API Documentation

For detailed information on each endpoint and their usage, you can access the API documentation through the following URL after running the server:

```bash
http://localhost:8080/swagger/index.html
```

## Database Structure

```sql
CREATE TABLE quizzes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE questions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    quiz_id INT NOT NULL,
    question_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (quiz_id) REFERENCES quizzes (id) ON DELETE CASCADE
);

CREATE TABLE options (
    id INT PRIMARY KEY AUTO_INCREMENT,
    question_id INT NOT NULL,
    option_text TEXT NOT NULL,
    is_correct BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES questions (id) ON DELETE CASCADE
);

CREATE TABLE participants (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE quiz_results (
    id INT PRIMARY KEY AUTO_INCREMENT,
    quiz_id INT NOT NULL,
    participant_id INT NOT NULL,
    score INT NOT NULL,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (quiz_id) REFERENCES quizzes (id) ON DELETE CASCADE,
    FOREIGN KEY (participant_id) REFERENCES participants (id) ON DELETE CASCADE
);

CREATE TABLE quiz_result_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    quiz_result_id INT NOT NULL,
    question_id INT NOT NULL,
    selected_option_id INT NOT NULL,
    FOREIGN KEY (quiz_result_id) REFERENCES quiz_results (id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions (id) ON DELETE CASCADE,
    FOREIGN KEY (selected_option_id) REFERENCES options (id) ON DELETE CASCADE
);
```

## Contributing

Contributions to this project are welcome. If you find any bugs or have any new features to add, feel free to submit a pull request. Please ensure that your changes follow the coding conventions and include appropriate tests.

## License

This project is licensed under the GPL-3.0 License.

Copyright 2023, Max Base
