import { useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";

import ModernLogin from "./components/ModernLogin";
import ModernCadastro from "./components/ModernCadastro";
import RentalDashboard from "./components/RentalDashboard";
import RentClothingPage from "./components/RentClothingPage";
import AtivarConta from "./components/Sellers";

import AdminLogin from "./components/AdminLogin";
import AdminCadastro from "./components/AdminCadastro";
import AdminDashboard from "./components/AdminDashboard";

import "./App.css";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || "");
  const [role, setRole] = useState(localStorage.getItem("role") || "user");
  const [currentView, setCurrentView] = useState("login");

  const handleLogout = () => {
    setToken("");
    setRole("user");
    localStorage.removeItem("token");
    localStorage.removeItem("role");
    localStorage.removeItem("userEmail");
  };

  const switchToLogin = () => setCurrentView("login");
  const switchToCadastro = () => setCurrentView("cadastro");

  return (
    <Router>
      <Routes>
        {!token ? (
          <>
            <Route
              path="/"
              element={
                currentView === "login" ? (
                  <ModernLogin
                    setToken={setToken}
                    switchToCadastro={switchToCadastro}
                  />
                ) : (
                  <ModernCadastro
                    setToken={setToken}
                    switchToLogin={switchToLogin}
                  />
                )
              }
            />
            <Route path="/active" element={<AtivarConta />} />

            <Route
              path="/admin/login"
              element={<AdminLogin setToken={setToken} setRole={setRole} />}
            />
            <Route path="/admin/cadastro" element={<AdminCadastro />} />

            <Route path="*" element={<Navigate to="/" />} />
          </>
        ) : (
          <>
            {role === "admin" ? (
              <>
                <Route
                  path="/admin/dashboard"
                  element={
                    <AdminDashboard token={token} onLogout={handleLogout} />
                  }
                />
                <Route path="*" element={<Navigate to="/admin/dashboard" />} />
              </>
            ) : (
              <>
                <Route
                  path="/dashboard"
                  element={
                    <RentalDashboard token={token} onLogout={handleLogout} />
                  }
                />
                <Route
                  path="/alugar/:id"
                  element={
                    <RentClothingPage token={token} onLogout={handleLogout} />
                  }
                />
                <Route path="*" element={<Navigate to="/dashboard" />} />
              </>
            )}
          </>
        )}
      </Routes>
    </Router>
  );
}

export default App;
