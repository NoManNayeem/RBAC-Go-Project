import React, { useState, createContext, useContext } from "react";
import { BrowserRouter as Router, Route, Routes, Navigate, useNavigate } from "react-router-dom";

// Create a Role Context
const RoleContext = createContext();

// Login Page
function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const { setRole } = useContext(RoleContext);
    const navigate = useNavigate();

    const handleLogin = (e) => {
        e.preventDefault();

        // Simulate login logic
        if (username === "Nayeem" && password === "Password") {
            setRole("admin");
            navigate("/admin");
        } else {
            setRole("public");
            navigate("/public");
        }
    };

    return (
        <div style={{ textAlign: "center", marginTop: "100px" }}>
            <h1>Login</h1>
            <form onSubmit={handleLogin}>
                <div>
                    <input
                        type="text"
                        placeholder="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
                <div>
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    );
}

// Public Page
function PublicPage() {
    return <h1>Welcome to the Public Page</h1>;
}

// Admin Page
function AdminPage() {
    return <h1>Welcome to the Admin Page</h1>;
}

// Protected Route Component
function ProtectedRoute({ children, allowedRoles }) {
    const { role } = useContext(RoleContext);

    if (!allowedRoles.includes(role)) {
        return <Navigate to="/" />;
    }

    return children;
}

// App Component
function App() {
    const [role, setRole] = useState(null);

    return (
        <RoleContext.Provider value={{ role, setRole }}>
            <Router>
                <Routes>
                    <Route path="/" element={<LoginPage />} />
                    <Route
                        path="/admin"
                        element={
                            <ProtectedRoute allowedRoles={["admin"]}>
                                <AdminPage />
                            </ProtectedRoute>
                        }
                    />
                    <Route
                        path="/public"
                        element={
                            <ProtectedRoute allowedRoles={["public", "admin"]}>
                                <PublicPage />
                            </ProtectedRoute>
                        }
                    />
                </Routes>
            </Router>
        </RoleContext.Provider>
    );
}

export default App;
