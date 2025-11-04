import { useState } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import ModernLogin from "./components/ModernLogin";
import ModernCadastro from "./components/ModernCadastro";
import RentalDashboard from "./components/RentalDashboard";
import RentClothingPage from "./components/RentClothingPage"; // ← nova página importada
import "./App.css";
import AtivarConta from "./components/Sellers";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || "");
  const [currentView, setCurrentView] = useState("login");

  const handleLogout = () => {
    setToken("");
    localStorage.removeItem("token");
    localStorage.removeItem("userEmail");
  };

  const switchToLogin = () => {
    setCurrentView("login");
  };

  const switchToCadastro = () => {
    setCurrentView("cadastro");
  };

  return (
    <Router>
      <Routes>
        {/* 🔐 Rota de login/cadastro */}
        {!token ? (
          <>
            <Route
              path="/"
              element={
                currentView === "login" ? (
                  <ModernLogin setToken={setToken} switchToCadastro={switchToCadastro} />
                ) : (
                  <ModernCadastro setToken={setToken} switchToLogin={switchToLogin} />
                )
              }
            />
              <Route
                  path="/active"
                  element={<AtivarConta/>}
              />
            {/* Redireciona qualquer outra rota para login se não tiver token */}
            <Route path="*" element={<Navigate to="/" />} />
          </>
        ) : (
          <>
            {/* 🏠 Dashboard principal */}
            <Route
              path="/dashboard"
              element={<RentalDashboard token={token} onLogout={handleLogout} />}
            />

            {/* 👕 Nova página de aluguel de roupa */}
            <Route
              path="/alugar/:id"
              element={<RentClothingPage token={token} onLogout={handleLogout} />}
            />

            {/* Se o usuário tentar acessar "/" redireciona pro dashboard */}
            <Route path="*" element={<Navigate to="/dashboard" />} />
          </>
        )}
      </Routes>
    </Router>
  );
}

export default App;
