import { Navigate, Route, Routes, useNavigate } from "react-router-dom"
import LandingPage from "./components/LandingPage"
import React, { useState } from "react"
import { User } from "./utils/types"
import Home from "./components/Home"

function App() {
  const [user, setUser] = useState<User>()

  const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
    if (user) {
      return children
    }
    return <Navigate to={"/"} />
  }

  const navigate = useNavigate()

  const loginSuccess = (user: User) => {
    setUser(user)
    navigate("/dashboard")
  }

  return (
    <>
      <Routes>
        <Route path="/" element={<LandingPage setUser={loginSuccess} />} />
        <Route path="/dashboard" element={<PrivateRoute><Home /></PrivateRoute>} />
        <Route path="*" element={<h1>Page not found</h1>} />
      </Routes>
    </>
  )
}

export default App
