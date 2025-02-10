import { Todo } from '../utils/types';
import instance from './interceptor';

const BASE_URI = '/todos';

const GetAllTodos = async (): Promise<Todo[]> => {
    const endpoint = `${BASE_URI}`;
    try {
        const response = await instance.axios.get<Todo[]>(endpoint);
        return response.data;
    } catch (error) {
        console.error("Error fetching todos:", error);
        throw error;
    }
};

const GetTodoByID = async (id: number): Promise<Todo> => {
    const endpoint = `${BASE_URI}/${id}`;
    try {
        const response = await instance.axios.get<Todo>(endpoint);
        return response.data;
    } catch (error) {
        console.error("Error fetching todo:", error);
        throw error;
    }
};

const CreateTodo = async (newTodo: Omit<Todo, 'id' | 'created_at' | 'updated_at'>): Promise<Todo> => {
    const endpoint = `${BASE_URI}`;
    try {
        const response = await instance.axios.post<Todo>(endpoint, newTodo);
        return response.data;
    } catch (error) {
        console.error("Error creating todo:", error);
        throw error;
    }
};

const UpdateTodoByID = async (id: number, updatedTodo: Partial<Omit<Todo, 'id' | 'created_at'>>): Promise<void> => {
    const endpoint = `${BASE_URI}/${id}`;
    try {
        await instance.axios.put<void>(endpoint, updatedTodo);
    } catch (error) {
        console.error("Error updating todo:", error);
        throw error;
    }
};

const DeleteTodoByID = async (id: number): Promise<void> => {
    const endpoint = `${BASE_URI}/${id}`;
    try {
        await instance.axios.delete(endpoint);
    } catch (error) {
        console.error("Error deleting todo:", error);
        throw error;
    }
};

export {
    GetAllTodos,
    GetTodoByID,
    CreateTodo,
    UpdateTodoByID,
    DeleteTodoByID
};