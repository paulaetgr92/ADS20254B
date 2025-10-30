import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { createSale } from "../api"; // envia para o backend
import "./RentClothingPage.css";

export default function RentClothingPage({ token }) {
  const { id } = useParams();
  const navigate = useNavigate();
  const [product, setProduct] = useState(null);
  const [rentalType, setRentalType] = useState("daily");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [confirming, setConfirming] = useState(false);

  const mockProducts = [
    {
      id: 1,
      name: "Vestido Elegante Preto",
      daily_price: 89.9,
      weekly_price: 299.9,
      image_url:
        "https://images.unsplash.com/photo-1566479179817-c0b5b4b4b1b5?w=300&h=400&fit=crop",
    },
    {
      id: 2,
      name: "Blazer Executivo",
      daily_price: 129.9,
      weekly_price: 449.9,
      image_url:
        "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=300&h=400&fit=crop",
    },
  ];

  useEffect(() => {
    const found = mockProducts.find((p) => p.id === parseInt(id));
    if (found) setProduct(found);
  }, [id]);

  const handleConfirmRental = async () => {
    if (!startDate || !endDate) {
      alert("Preencha as datas de início e fim!");
      return;
    }
    setConfirming(true);
    try {
      const saleData = {
        produtoId: product.id,
        quantidade: 1,
        tempoValor: rentalType === "daily" ? 1 : 7, // 1 dia ou 7 dias
      };
      console.log("Enviando para backend:", saleData);
      const res = await createSale(saleData, token);
      console.log("Resposta backend:", res);
      alert(`Aluguel confirmado para "${product.name}"!`);
      navigate("/dashboard");
    } catch (err) {
      console.error("Erro ao enviar aluguel:", err);
      alert("Falha ao enviar aluguel. Veja o console.");
    } finally {
      setConfirming(false);
    }
  };

  if (!product) return <div>Carregando...</div>;

  return (
    <div className="rent-page">
      <button onClick={() => navigate(-1)}>← Voltar</button>
      <div className="rent-container">
        <img src={product.image_url} alt={product.name} />
        <h2>{product.name}</h2>
        <label>Tipo de Aluguel:</label>
        <select value={rentalType} onChange={(e) => setRentalType(e.target.value)}>
          <option value="daily">Diária — R$ {product.daily_price}</option>
          <option value="weekly">Semanal — R$ {product.weekly_price}</option>
        </select>
        <label>Data de Início:</label>
        <input type="date" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
        <label>Data de Devolução:</label>
        <input type="date" value={endDate} onChange={(e) => setEndDate(e.target.value)} />
        <button onClick={handleConfirmRental} disabled={confirming}>
          {confirming ? "Confirmando..." : "Confirmar Aluguel"}
        </button>
      </div>
    </div>
  );
}
