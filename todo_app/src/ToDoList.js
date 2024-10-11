// src/ToDoList.js

import React, { useEffect, useState } from 'react';

const styles = {
    container: {
        padding: '20px',
        maxWidth: '600px',
        margin: 'auto',
        border: '1px solid #ddd',
        borderRadius: '5px',
        boxShadow: '0 2px 10px rgba(0, 0, 0, 0.1)',
    },
    form: {
        display: 'flex',
        alignItems: 'center',
        marginBottom: '20px',
    },
    input: {
        flex: '1',
        padding: '10px',
        marginRight: '10px',
        border: '1px solid #ddd',
        borderRadius: '4px',
    },
    checkbox: {
        marginRight: '10px',
    },
    button: {
        padding: '10px 20px',
        backgroundColor: '#28a745',
        color: 'white',
        border: 'none',
        borderRadius: '4px',
        cursor: 'pointer',
    },
    todoList: {
        listStyleType: 'none',
        padding: '0',
    },
    todoItem: {
        padding: '10px',
        borderBottom: '1px solid #ddd',
    },
};

const API_URL = process.env.REACT_APP_API_URL

const ToDoList = () => {
    const [todos, setTodos] = useState([]);
    const [title, setTitle] = useState('');
    const [done, setDone] = useState(false);

    useEffect(() => {
        // Fetch existing todos from the API on component mount
        const fetchTodos = async () => {
            const response = await fetch(API_URL);
            const data = await response.json();
            setTodos(data);
        };
        fetchTodos();
    }, []);

    const addTodo = async () => {
        if (!title) return; // Prevent adding empty to-do

        const newTodo = { title, done };

        // Send a POST request to the API
        await fetch(API_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newTodo),
        });

        // Optionally, fetch the updated list of todos or append the new one directly
        setTodos([...todos, newTodo]);
        setTitle('');
        setDone(false);
    };

    return (
        <div style={styles.container}>
            <h1>To-Do List</h1>

            {/* Form to add new todo */}
            <div style={styles.form}>
                <input
                    type="text"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    placeholder="Title"
                    style={styles.input}
                />
                <input
                    type="checkbox"
                    checked={done}
                    onChange={(e) => setDone(e.target.checked)}
                    style={styles.checkbox}
                />
                <label>Done</label>
                <button onClick={addTodo} style={styles.button}>
                    Add To-Do
                </button>
            </div>

            <h2>Current To-Dos</h2>

            {/* List all todos */}
            <ul style={styles.todoList}>
                {todos.map((todo) => (
                    <li key={todo.id} style={styles.todoItem}>
                        {todo.title} - {todo.done ? 'Done' : 'Not Done'}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ToDoList;
