interface Todo {
    id: number;
    title: string;
    content: string;
    status: string;
    created_at: string;
    updated_at: string;
}

interface User {
    id: number;
    name: string;
    email: string;
}

interface LoginResponse {
    token: string;
    user: User;
}

export type { Todo, User, LoginResponse }