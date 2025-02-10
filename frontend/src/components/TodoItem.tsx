import React, { useState } from 'react';
import TodoModal from './TodoModal';
import { Todo } from '../utils/types';

interface TodoItemProps {
    todo: Todo;
    addTodo: (todo: Todo) => void
    editTodo: (id: number, updatedTodo: Partial<Todo>) => void
    deleteTodo: (id: number) => void
}

const TodoItem: React.FC<TodoItemProps> = ({ todo, addTodo, deleteTodo, editTodo }) => {
    const [isEditModalOpen, setEditModalOpen] = useState(false);

    return (
        <li className="flex justify-between items-center bg-gray-200 p-4 mb-2 rounded shadow">
            <div>
                <h3 className="font-bold">{todo.title}</h3>
                <p>{todo.content}</p>
                <p className="text-sm text-gray-500">Status: {todo.status}</p>
            </div>
            <div>
                <button
                    onClick={() => setEditModalOpen(true)}
                    className="bg-yellow-500 text-white px-2 py-1 rounded mr-2"
                >
                    Edit
                </button>
                <button
                    onClick={() => deleteTodo(todo.id)}
                    className="bg-red-500 text-white px-2 py-1 rounded"
                >
                    Delete
                </button>
            </div>
            <TodoModal
                addTodo={addTodo}
                editTodo={editTodo}
                isOpen={isEditModalOpen}
                closeModal={() => setEditModalOpen(false)}
                editingTodo={todo}
            />
        </li>
    );
};

export default TodoItem;
