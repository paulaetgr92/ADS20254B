import { useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { apiFetch } from "../api";
import "./RentalAuth.css";

export default function ActivationForm() {
    const [code, setCode] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");
    const navigate = useNavigate();
    const location = useLocation();

    const cadastroId = location.state?.cadastroId;

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError("");
        setSuccess("");

        if (!code) return setError("Digite o código de ativação");

        setLoading(true);

        try {
        
            const verifyResponse = await apiFetch("activation/verify", {
                method: "POST",
                body: JSON.stringify({ id: cadastroId, code }),
            });
            console.log("Resposta da verificação:", verifyResponse);

            if (!verifyResponse?.success) {
                throw new Error(verifyResponse?.message || "Código inválido");
            }


             const getResponse = await apiFetch(`activation/get?id=${cadastroId}&code=${code}`, {
                 method: "GET",
             });
             console.log("Resposta da busca:", getResponse);

            setSuccess("Cadastro ativado com sucesso!");
            setTimeout(() => {
                navigate("/login");
            }, 1500);

        } catch (err) {
            console.error("Erro na requisição:", err);
            setError(err.message || "Erro ao ativar cadastro");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="rental-auth-container">
            <div className="rental-auth-card">
                <div className="rental-logo">
                    <h1>DoutorRent</h1>
                    <p className="rental-tagline">Ative sua conta</p>
                </div>

                <div className="rental-auth-content">
                    <h2>Ativação</h2>
                    <p>Digite o código que você recebeu via SMS</p>

                    {error && <div className="rental-error">{error}</div>}
                    {success && <div className="rental-success">{success}</div>}

                    <form onSubmit={handleSubmit} className="rental-form">
                        <input
                            type="text"
                            name="code"
                            placeholder="Código de ativação"
                            value={code}
                            onChange={(e) => setCode(e.target.value)}
                            required
                            className="rental-input"
                        />
                        <button type="submit" className="continue-btn" disabled={loading}>
                            {loading ? "Ativando..." : "Ativar cadastro"}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}
