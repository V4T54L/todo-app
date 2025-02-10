import React, { useEffect, useState } from 'react';
import Header from './Header';
import TodoList from './TodoList';
import Footer from './Footer';
import TodoModal from './TodoModal';
import { Todo } from '../utils/types';
import { CreateTodo, DeleteTodoByID, GetAllTodos, UpdateTodoByID } from '../api/todo';

const Home: React.FC = () => {
    const [isModalOpen, setModalOpen] = useState(false);
    const [todos, setTodos] = useState<Todo[]>([])
    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    const addTodo = async (todo: Todo) => {
        try {
            const res = await CreateTodo(todo)
            setTodos((prev) => [...prev, res]);
        } catch (error) {
            setErrorMessage((error as Error).message || 'Failed to create todo');
        }
    };

    const editTodo = async (id: number, updatedTodo: Partial<Todo>) => {
        try {
            await UpdateTodoByID(id, updatedTodo)
            setTodos((prev) =>
                prev.map((todo) => (todo.id === id ? { ...todo, ...updatedTodo } : todo))
            );
        } catch (error) {
            setErrorMessage((error as Error).message || 'Failed to edit todo');
        }
    };

    const deleteTodo = async (id: number) => {
        try {
            await DeleteTodoByID(id)
            setTodos((prev) => prev.filter((todo) => todo.id !== id));
        } catch (error) {
            setErrorMessage((error as Error).message || 'Failed to delete todo');
        }
        setTodos((prev) => prev.filter((todo) => todo.id !== id));
    };

    useEffect(() => {
        const fetchTodos = async () => {
            try {
                const todos = await GetAllTodos();
                setTodos(todos);
            } catch (error) {
                setErrorMessage((error as Error).message || 'Failed to fetch todos');
            }
        };

        fetchTodos();
    }, [])

    return (
        <div className="flex flex-col min-h-screen">
            <Header />
            <main className="flex-grow p-6">
                <button
                    onClick={() => setModalOpen(true)}
                    className="bg-blue-500 text-white px-4 py-2 rounded"
                >
                    Create TODO
                </button>
                <TodoList addTodo={addTodo} deleteTodo={deleteTodo} editTodo={editTodo} error={errorMessage} todos={todos} />
            </main>
            <Footer />
            <TodoModal
                addTodo={addTodo}
                editTodo={editTodo}
                isOpen={isModalOpen}
                closeModal={() => setModalOpen(false)}
            />
        </div>
    );
};

export default Home;