import React, { useState } from 'react';
import AuthModal from './AuthModal';
import { User } from '../utils/types';

interface LandingPageProps {
    setUser: (user: User) => void;
}


const LandingPage: React.FC<LandingPageProps> = ({ setUser }) => {
    const [modalOpen, setModalOpen] = useState(false);

    return (
        <>
            <div className="flex flex-col items-center justify-center h-screen bg-gradient-to-b from-sky-950 to-indigo-700">
                <h1 className="text-white text-5xl font-extrabold">Welcome to wa-TODO</h1>
                <p className="text-white text-xl mt-4">Simplify Your Life, One Task at a Time!</p>
                <p className="text-white w-[60%] mt-4">Welcome to TodoMaster, the ultimate tool to help you manage your tasks,
                    boost productivity, and achieve your goals effortlessly. Whether you're managing personal projects,
                    work tasks, or daily errands, TodoMaster has you covered!</p>
                <button
                    onClick={() => setModalOpen(true)}
                    className="mt-8 bg-white text-blue-500 px-6 py-2 rounded shadow-md hover:bg-gray-200"
                >
                    Get Started
                </button>
            </div>
            <AuthModal setUser={setUser} isOpen={modalOpen} onRequestClose={() => setModalOpen(false)} />
        </>
    );
};

export default LandingPage;