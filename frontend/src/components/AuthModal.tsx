import React, { useState } from 'react';
import { Signup, Login } from '../api/auth';
import { LoginResponse, User } from '../utils/types';

interface AuthModalProps {
    isOpen: boolean;
    onRequestClose: () => void;
    setUser: (user: User) => void
}

const AuthModal: React.FC<AuthModalProps> = ({ isOpen, onRequestClose, setUser }) => {
    const [isSignup, setIsSignup] = useState(false);
    const [name, setName] = useState<string>('');
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [error, setError] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);

    const handleSwitch = () => {
        setIsSignup(!isSignup);
        setError('');
        setName('');
        setEmail('');
        setPassword('');
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setLoading(true);

        try {
            if (isSignup) {
                const newUser: User = await Signup(name, email, password);
                console.log("Signup successful, user:", newUser);
                onRequestClose();
            } else {
                const loginResponse: LoginResponse = await Login(email, password);
                console.log("Login successful, user:", loginResponse.user);
                setUser(loginResponse.user)
                onRequestClose();
            }
        } catch (error) {
            setError('Error: ' + (error as Error).message || 'Something went wrong');
            console.error("Authentication error:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        isOpen && (
            <div
                className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-70 transition-opacity"
                role="dialog"
                aria-modal="true"
            >
                <div className="bg-white rounded-lg shadow-lg p-8 w-full max-w-md transition-transform transform scale-95 hover:scale-100">
                    <h2 className="text-2xl font-extrabold text-center text-gray-700 mb-6">{isSignup ? 'Create Account' : 'Welcome Back'}</h2>
                    <form onSubmit={handleSubmit}>
                        {isSignup && (
                            <div className="mb-4">
                                <label className="block text-gray-600 mb-1" htmlFor="name">Name</label>
                                <input
                                    type="text"
                                    id="name"
                                    value={name}
                                    onChange={(e) => setName(e.target.value)}
                                    required
                                    className="border border-gray-300 p-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-400"
                                />
                            </div>
                        )}
                        <div className="mb-4">
                            <label className="block text-gray-600 mb-1" htmlFor="email">Email Address</label>
                            <input
                                type="email"
                                id="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                                className="border border-gray-300 p-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-400"
                            />
                        </div>
                        <div className="mb-4">
                            <label className="block text-gray-600 mb-1" htmlFor="password">Password</label>
                            <input
                                type="password"
                                id="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                                className="border border-gray-300 p-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-400"
                            />
                        </div>
                        {error && <p className="text-red-500 text-sm mb-2">{error}</p>}
                        <button
                            type="submit"
                            disabled={loading}
                            className={`w-full py-2 rounded bg-blue-500 text-white hover:bg-blue-600 transition duration-200 ${loading ? 'opacity-50 cursor-not-allowed' : ''}`}
                        >
                            {loading ? 'Loading...' : (isSignup ? 'Sign Up' : 'Log In')}
                        </button>
                    </form>
                    <p className="text-center mt-4">
                        <span onClick={handleSwitch} className="text-blue-500 cursor-pointer">
                            {isSignup ? 'Already have an account? Log In' : "Don't have an account? Sign Up"}
                        </span>
                    </p>
                    <button
                        className="mt-4 text-red-500 hover:text-red-700"
                        onClick={onRequestClose}
                    >
                        Close
                    </button>
                </div>
            </div>
        )
    );
};

export default AuthModal;