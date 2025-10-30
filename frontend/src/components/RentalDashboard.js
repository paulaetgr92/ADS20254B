import { useState, useEffect } from "react";
import { createSale } from "../api";
import "./RentalDashboard.css";

export default function RentalDashboard({ token, onLogout }) {
  const userEmail = localStorage.getItem("userEmail") || "Usuário";
  const [activeTab, setActiveTab] = useState("catalog");
  const [products, setProducts] = useState([]);
  const [cart, setCart] = useState([]);
  const [rentals, setRentals] = useState([]);
  const [toast, setToast] = useState({ message: "", type: "success" });

const mockProducts = [
  { id:1, name:"Vestido Elegante Preto", daily_price:89.9, weekly_price:299.9, image_url:"https://images.unsplash.com/photo-1566479179817-c0b5b4b4b1b5?w=300&h=400&fit=crop" },
  { id:2, name:"Blazer Executivo", daily_price:129.9, weekly_price:449.9, image_url:"https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=300&h=400&fit=crop" },
  { id:3, name:"Casaco de Inverno Masculino", daily_price:99.9, weekly_price:349.9, image_url:"https://images.unsplash.com/photo-1600180758895-798f2b7c4c32?w=300&h=400&fit=crop" },
  { id:4, name:"Calça Jeans Slim", daily_price:69.9, weekly_price:239.9, image_url:"https://images.unsplash.com/photo-1593032457868-bc0f2f5a7f39?w=300&h=400&fit=crop" },
  { id:5, name:"Camisa Social Branca", daily_price:59.9, weekly_price:199.9, image_url:"https://images.unsplash.com/photo-1593032466438-8fc45a75aeb5?w=300&h=400&fit=crop" },
  { id:6, name:"Saia Midi Estampada", daily_price:79.9, weekly_price:269.9, image_url:"https://images.unsplash.com/photo-1593032457459-3b6d75a9e7f1?w=300&h=400&fit=crop" }
];


  useEffect(() => setProducts(mockProducts), []);

  const showToast = (message, type = "success") => {
    setToast({ message, type });
    setTimeout(() => setToast({ message: "", type: "success" }), 3000);
  };

  const addToCart = (product) => {
    if (cart.find((p) => p.id === product.id)) {
      showToast("⚠️ Já está no carrinho!", "warning");
      return;
    }
    setCart([...cart, product]);
    showToast(`✨ ${product.name} adicionado ao carrinho!`);
  };

  const removeFromCart = (id) => {
    setCart(cart.filter((p) => p.id !== id));
    showToast("❌ Item removido do carrinho.");
  };

  const confirmRental = async () => {
    if (cart.length === 0) {
      showToast("⚠️ Carrinho vazio!", "warning");
      return;
    }

    try {
      const promises = cart.map((item) => {
        const saleData = {
          produtoId: item.id,
          quantidade: 1,
          tempoValor: 1,
        };
        console.log("Enviando para backend:", saleData);
        return createSale(saleData, token).then((res) => ({
          ...item,
          backendId: res.id || null,
        }));
      });

      const results = await Promise.all(promises);
      setRentals([...rentals, ...results]);
      setCart([]);
      setActiveTab("rentals");
      showToast("🎉 Aluguel confirmado e salvo no backend!");
    } catch (err) {
      console.error("Erro ao enviar aluguel:", err);
      showToast("❌ Falha ao enviar aluguel!", "error");
    }
  };

  return (
    <div className="rental-dashboard">
      <header>
        <h1>DoutorRent</h1>
        <button onClick={onLogout}>Sair</button>
      </header>

      <nav>
        <button onClick={() => setActiveTab("catalog")}>👗 Catálogo</button>
        <button onClick={() => setActiveTab("cart")}>🛒 Carrinho ({cart.length})</button>
        <button onClick={() => setActiveTab("rentals")}>📦 Meus Aluguéis ({rentals.length})</button>
      </nav>

      <main>
        {activeTab === "catalog" && (
          <div className="catalog-section">
            {products.map((p) => (
              <div key={p.id} className="product-card">
                <img src={p.image_url} alt={p.name} />
                <h3>{p.name}</h3>
                <button onClick={() => addToCart(p)}>Alugar</button>
              </div>
            ))}
          </div>
        )}

        {activeTab === "cart" && (
          <div className="cart-section">
            {cart.length === 0 ? <p>Carrinho vazio</p> : cart.map((p) => (
              <div key={p.id}>
                <span>{p.name}</span>
                <button onClick={() => removeFromCart(p.id)}>❌</button>
              </div>
            ))}
            {cart.length > 0 && <button onClick={confirmRental}>Confirmar Aluguel</button>}
          </div>
        )}

        {activeTab === "rentals" && (
          <div className="rentals-section">
            {rentals.length === 0 ? <p>Nenhum aluguel ainda</p> : rentals.map((r) => (
              <div key={r.id || r.backendId}>
                <span>{r.name}</span>
                <span>ID backend: {r.backendId}</span>
              </div>
            ))}
          </div>
        )}
      </main>

      {toast.message && <div className={`toast-message ${toast.type}`}>{toast.message}</div>}
    </div>
  );
}
