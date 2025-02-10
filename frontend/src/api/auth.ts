import instance from './interceptor';
import { LoginResponse, User } from '../utils/types';

const BASE_URI = '/auth';

const Signup = async (name: string, email: string, password: string): Promise<User> => {
    const endpoint = `${BASE_URI}/signup`;
    try {
        const response = await instance.axios.post<User>(endpoint, { name, email, password });
        return response.data;
    } catch (error) {
        console.error("Error signing up:", error);
        throw error;
    }
};

const Login = async (email: string, password: string): Promise<LoginResponse> => {
    const endpoint = `${BASE_URI}/login`;
    try {
        const response = await instance.axios.post<LoginResponse>(endpoint, { email, password });
        instance.token = response.data.token;
        return response.data;
    } catch (error) {
        console.error("Error logging in:", error);
        throw error;
    }
};

export { Signup, Login };