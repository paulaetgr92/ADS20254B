import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { apiFetch } from "../api";
import "./RentalAuth.css";

export default function ModernCadastro({ switchToLogin }) {
    const [formData, setFormData] = useState({
        name: "",
        email: "",
        password: "",
        confirmPassword: "",
        cpf: "",
        cnpj: "",
        celular: "",
    });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");
    const navigate = useNavigate();

    const handleChange = (e) => {
        setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
        if (error) setError("");
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError("");
        setSuccess("");

        const { password, confirmPassword, cpf, cnpj } = formData;

        if (password !== confirmPassword) return setError("As senhas não coincidem");
        if (password.length < 6) return setError("A senha deve ter pelo menos 6 caracteres");
        if (!cpf && !cnpj) return setError("Preencha CPF ou CNPJ");
        if (cpf && cpf.length !== 11) return setError("CPF deve ter 11 dígitos");
        if (cnpj && cnpj.length !== 14) return setError("CNPJ deve ter 14 dígitos");

        setLoading(true);

        try {
            const response = await apiFetch("cadastros", {
                method: "POST",
                body: JSON.stringify(formData),
            });

            console.log("Resposta completa do backend:", response);


            const cadastroId = response?.cadastro?.ID;

            if (!cadastroId) throw new Error("Não foi possível identificar o ID do cadastro retornado.");

            setSuccess("Cadastro criado com sucesso! Verifique seu SMS para o código de ativação.");

            setTimeout(() => {
                navigate("/active", {
                    state: { cadastroId },
                });
            }, 1500);
        } catch (err) {
            console.error("Erro no cadastro:", err);
            setError(err.message || "Erro ao cadastrar");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="rental-auth-container">
            <div className="rental-auth-card">
                <div className="rental-logo">
                    <h1>DoutorRent</h1>
                    <p className="rental-tagline">Alugue o look perfeito</p>
                </div>

                <div className="rental-auth-content">
                    <h2>Criar sua conta</h2>
                    <p>Comece a alugar looks incríveis</p>

                    {error && <div className="rental-error">{error}</div>}
                    {success && <div className="rental-success">{success}</div>}

                    <form onSubmit={handleSubmit} className="rental-form">
                        <input type="text" name="name" placeholder="Nome completo" value={formData.name} onChange={handleChange} required className="rental-input" />
                        <input type="email" name="email" placeholder="Seu email" value={formData.email} onChange={handleChange} required className="rental-input" />
                        <input type="text" name="cpf" placeholder="CPF (somente números)" value={formData.cpf} onChange={handleChange} className="rental-input" />
                        <input type="text" name="cnpj" placeholder="CNPJ (somente números)" value={formData.cnpj} onChange={handleChange} className="rental-input" />
                        <input type="text" name="celular" placeholder="Celular (somente números)" value={formData.celular} onChange={handleChange} className="rental-input" />
                        <input type="password" name="password" placeholder="Criar senha" value={formData.password} onChange={handleChange} required className="rental-input" />
                        <input type="password" name="confirmPassword" placeholder="Confirmar senha" value={formData.confirmPassword} onChange={handleChange} required className="rental-input" />

                        <button type="submit" className="continue-btn" disabled={loading}>
                            {loading ? "Criando conta..." : "Criar conta gratuita"}
                        </button>
                    </form>

                    <div className="rental-footer">
                        <Link to="/politica" className="footer-link">Política de Privacidade</Link>
                        <Link to="/termos" className="footer-link">Termos de Uso</Link>
                    </div>

                    <div className="auth-switch-rental">
                        <p>Já tem uma conta?</p>
                        <button type="button" className="switch-button-rental" onClick={switchToLogin}>
                            Fazer login
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
