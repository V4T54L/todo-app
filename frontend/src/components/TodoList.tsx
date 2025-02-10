import React from 'react';
import { Todo } from '../utils/types';
import TodoItem from './TodoItem';

interface TodoListProps {
    todos: Todo[],
    error: string|null,
    addTodo: (todo: Todo) => void
    editTodo: (id: number, updatedTodo: Partial<Todo>) => void
    deleteTodo: (id: number) => void
}

const TodoList: React.FC<TodoListProps> = ({ error, todos, addTodo, deleteTodo, editTodo }) => {
    if (error) {
        return <div>Error loading todos: {error}</div>;
    }


    return (
        <ul className="mt-4">
            {todos.map(todo => (
                <TodoItem addTodo={addTodo} deleteTodo={deleteTodo} editTodo={editTodo} key={todo.id} todo={todo} />
            ))}
        </ul>
    );
};

export default TodoList;