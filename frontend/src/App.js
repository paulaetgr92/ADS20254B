import { useState } from "react";
import ModernLogin from "./components/ModernLogin";
import ModernCadastro from "./components/ModernCadastro";
import RentalDashboard from "./components/RentalDashboard";
import "./App.css";

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

  if (token) {
    return <RentalDashboard token={token} onLogout={handleLogout} />;
  }

  return (
    <div className="App">
      {currentView === "login" ? (
        <ModernLogin setToken={setToken} switchToCadastro={switchToCadastro} />
      ) : (
        <ModernCadastro setToken={setToken} switchToLogin={switchToLogin} />
      )}
    </div>
  );
}

export default App;

