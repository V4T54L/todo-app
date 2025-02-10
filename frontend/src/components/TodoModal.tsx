import React, { useState } from 'react';
import { Todo } from '../utils/types';

interface TodoModalProps {
    isOpen: boolean;
    closeModal: () => void;
    editingTodo?: Todo;
    addTodo: (todo: Todo) => void
    editTodo: (id: number, updatedTodo: Partial<Todo>) => void
}

const TodoModal: React.FC<TodoModalProps> = ({ isOpen, closeModal, editingTodo, addTodo, editTodo }) => {
    const [title, setTitle] = useState(editingTodo?.title || '');
    const [content, setContent] = useState(editingTodo?.content || '');
    const [status, setStatus] = useState(editingTodo ? editingTodo.status : 'pending');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (editingTodo) {
            editTodo(editingTodo.id, { title, content, status });
        } else {
            const newTodo: Todo = {
                id: 0,
                title,
                content,
                status,
                created_at: "",
                updated_at: "",
            };
            addTodo(newTodo);
        }
        closeModal();
    };

    return (
        <div className={`fixed inset-0 flex justify-center items-center ${isOpen ? '' : 'hidden'}`}>
            <div className="bg-white p-6 rounded shadow-lg">
                <h2 className="text-lg font-semibold">{editingTodo ? 'Edit Todo' : 'Create Todo'}</h2>
                <form onSubmit={handleSubmit} className="mt-4">
                    <input
                        type="text"
                        onChange={(e) => setTitle(e.target.value)}
                        value={title}
                        placeholder="Title"
                        className="w-full p-2 border border-gray-300 rounded mb-2"
                        required
                    />
                    <textarea
                        onChange={(e) => setContent(e.target.value)}
                        value={content}
                        placeholder="Content"
                        className="w-full p-2 border border-gray-300 rounded mb-2"
                        required
                    />
                    <select
                        onChange={(e) => setStatus(e.target.value as 'pending' | 'completed')}
                        value={status}
                        className="w-full p-2 border border-gray-300 rounded mb-2"
                    >
                        <option value="pending">Pending</option>
                        <option value="completed">Completed</option>
                    </select>
                    <div className="flex justify-end">
                        <button
                            type="submit"
                            className="bg-blue-500 text-white px-4 py-2 rounded mr-2"
                        >
                            {editingTodo ? 'Update' : 'Create'}
                        </button>
                        <button
                            type="button"
                            onClick={closeModal}
                            className="border border-gray-300 px-4 py-2 rounded"
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default TodoModal;
