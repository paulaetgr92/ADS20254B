import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import "./AdminAuth.css";

const AdminCadastro = () => {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    cnpj: "",
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleChange = (e) =>
    setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));

  const validateForm = () => {
    if (formData.password !== formData.confirmPassword) {
      setError("As senhas não coincidem");
      return false;
    }
    if (formData.password.length < 6) {
      setError("A senha deve ter pelo menos 6 caracteres");
      return false;
    }
    if (formData.cnpj.length !== 14) {
      setError("CNPJ deve ter 14 dígitos");
      return false;
    }
    return true;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setSuccess("");
    setLoading(true);

    if (!validateForm()) {
      setLoading(false);
      return;
    }

    try {
      const response = await fetch(
        "http://localhost:8080/api/v1/admin/create",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            cnpj: formData.cnpj,
            name: formData.name,
            email: formData.email,
            password: formData.password,
          }),
        }
      );

      const data = await response.json();

      if (!response.ok || !data.success) {
        throw new Error(data.error || "Erro ao cadastrar admin");
      }

      setSuccess("Admin cadastrado com sucesso! Redirecionando...");
      setTimeout(() => navigate("/admin/login"), 2000);
    } catch (err) {
      console.error("Erro ao cadastrar admin:", err);
      setError(err.message || "Erro ao conectar com o servidor");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="admin-auth-container">
      <Link to="/admin/login" className="admin-link">
        Voltar ao Login
      </Link>

      <div className="admin-auth-card">
        <div className="admin-auth-header">
          <h2>Cadastro Admin</h2>
          <p>Crie sua conta administrativa</p>
        </div>

        <form onSubmit={handleSubmit} className="admin-auth-form">
          {error && <div className="error-message">{error}</div>}
          {success && <div className="success-message">{success}</div>}

          <div className="form-group">
            <label>CNPJ</label>
            <input
              name="cnpj"
              value={formData.cnpj}
              onChange={handleChange}
              required
              placeholder="Digite o CNPJ (14 dígitos)"
              maxLength="14"
            />
          </div>

          <div className="form-group">
            <label>Nome da Empresa</label>
            <input
              name="name"
              value={formData.name}
              onChange={handleChange}
              required
              placeholder="Digite o nome da empresa"
            />
          </div>

          <div className="form-group">
            <label>Email</label>
            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              required
              placeholder="Digite seu email"
            />
          </div>

          <div className="form-group">
            <label>Senha</label>
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              required
              placeholder="Digite sua senha (mín. 6 caracteres)"
            />
          </div>

          <div className="form-group">
            <label>Confirmar Senha</label>
            <input
              type="password"
              name="confirmPassword"
              value={formData.confirmPassword}
              onChange={handleChange}
              required
              placeholder="Confirme sua senha"
            />
          </div>

          <button
            type="submit"
            className="admin-auth-button"
            disabled={loading}
          >
            {loading ? "Cadastrando..." : "Cadastrar"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default AdminCadastro;
