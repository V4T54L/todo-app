import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const Header: React.FC = () => {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);

    return (
        <header className="bg-gray-800 p-4 flex justify-between items-center">
            <span className="text-white text-lg font-bold">wa-TODO</span>
            <div className="relative">
                <button
                    className="text-white"
                    onClick={() => setIsDropdownOpen(!isDropdownOpen)}
                >
                    <i className="fas fa-user-circle fa-2x"></i>
                </button>

                {isDropdownOpen && (
                    <div className="absolute right-0 mt-2 w-48 bg-white shadow-lg rounded-lg">
                        <Link to="/profile" className="block px-4 py-2 text-gray-800 hover:bg-gray-200">Profile</Link>
                        <Link to="/settings" className="block px-4 py-2 text-gray-800 hover:bg-gray-200">Settings</Link>
                        <button className="block w-full text-left px-4 py-2 text-gray-800 hover:bg-gray-200">Logout</button>
                    </div>
                )}
            </div>
        </header>
    );
};

export default Header;
