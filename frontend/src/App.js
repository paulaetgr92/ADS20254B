import { useState } from "react";
import ModernLogin from "./components/ModernLogin";
import ModernCadastro from "./components/ModernCadastro";
import AdminLogin from "./components/AdminLogin";
import AdminCadastro from "./components/AdminCadastro";
import RentalDashboard from "./components/RentalDashboard";
import AdminDashboard from "./components/AdminDashboard";
import "./App.css";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || "");
  const [adminToken, setAdminToken] = useState(
    localStorage.getItem("adminToken") || ""
  );
  const [currentView, setCurrentView] = useState("login");
  const [isAdminMode, setIsAdminMode] = useState(false);

  const handleLogout = () => {
    setToken("");
    localStorage.removeItem("token");
    localStorage.removeItem("userEmail");
  };

  const handleAdminLogout = () => {
    setAdminToken("");
    localStorage.removeItem("adminToken");
    localStorage.removeItem("adminEmail");
  };

  const switchToLogin = () => {
    setCurrentView("login");
  };

  const switchToCadastro = () => {
    setCurrentView("cadastro");
  };

  const switchToAdminLogin = () => {
    setCurrentView("adminLogin");
  };

  const switchToAdminCadastro = () => {
    setCurrentView("adminCadastro");
  };

  const toggleMode = () => {
    setIsAdminMode(!isAdminMode);
    setCurrentView(isAdminMode ? "login" : "adminLogin");
  };

  if (adminToken) {
    return (
      <AdminDashboard adminToken={adminToken} onLogout={handleAdminLogout} />
    );
  }

  if (token) {
    return <RentalDashboard token={token} onLogout={handleLogout} />;
  }

  return (
    <div className="App">
      <div className="mode-toggle">
        <button onClick={toggleMode} className="toggle-btn">
          {isAdminMode ? "👤 Área do Cliente" : "🔧 Área Admin"}
        </button>
      </div>

      {isAdminMode ? (
        currentView === "adminLogin" ? (
          <AdminLogin
            setAdminToken={setAdminToken}
            switchToAdminCadastro={switchToAdminCadastro}
          />
        ) : (
          <AdminCadastro
            setAdminToken={setAdminToken}
            switchToAdminLogin={switchToAdminLogin}
          />
        )
      ) : currentView === "login" ? (
        <ModernLogin setToken={setToken} switchToCadastro={switchToCadastro} />
      ) : (
        <ModernCadastro setToken={setToken} switchToLogin={switchToLogin} />
      )}
    </div>
  );
}

export default App;
