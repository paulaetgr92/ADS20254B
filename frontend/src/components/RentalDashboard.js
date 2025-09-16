import { useState, useEffect } from "react";
import "./RentalDashboard.css";

export default function RentalDashboard({ token, onLogout }) {
  const userEmail = localStorage.getItem("userEmail") || "Usuário";
  const [activeTab, setActiveTab] = useState("catalog");
  const [products, setProducts] = useState([]);
  const [memberSince, setMemberSince] = useState("");

  const mockProducts = [
    {
      id: 1,
      name: "Vestido Elegante Preto",
      category: "Vestidos",
      size: "M",
      color: "Preto",
      brand: "Zara",
      daily_price: 89.9,
      weekly_price: 299.9,
      image_url:
        "https://images.unsplash.com/photo-1566479179817-c0b5b4b4b1b5?w=300&h=400&fit=crop",
      condition: "Excelente",
    },
    {
      id: 2,
      name: "Blazer Executivo",
      category: "Blazers",
      size: "P",
      color: "Azul Marinho",
      brand: "Hugo Boss",
      daily_price: 129.9,
      weekly_price: 449.9,
      image_url:
        "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=300&h=400&fit=crop",
      condition: "Excelente",
    },
    {
      id: 3,
      name: "Saia Midi Plissada",
      category: "Saias",
      size: "M",
      color: "Rosa",
      brand: "H&M",
      daily_price: 59.9,
      weekly_price: 199.9,
      image_url:
        "https://images.unsplash.com/photo-1583496661160-fb5886a13d77?w=300&h=400&fit=crop",
      condition: "Muito Bom",
    },
    {
      id: 4,
      name: "Calça Social Slim",
      category: "Calças",
      size: "38",
      color: "Cinza",
      brand: "Renner",
      daily_price: 69.9,
      weekly_price: 239.9,
      image_url:
        "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=300&h=400&fit=crop",
      condition: "Excelente",
    },
  ];

  useEffect(() => {
    setProducts(mockProducts);
  }, []);

  useEffect(() => {
    if (token) {
      try {
        const payloadBase64 = token.split(".")[1];
        const decodedPayload = JSON.parse(atob(payloadBase64));
        console.log("Payload do token:", decodedPayload);

        const timestamp = decodedPayload.createdAt || decodedPayload.iat;
        if (timestamp) {
          const date = new Date(timestamp * 1000);
          const formatted = date.toLocaleDateString("pt-BR", {
            year: "numeric",
            month: "long",
          });
          setMemberSince(formatted);
        }
      } catch (err) {
        console.error("Erro ao decodificar token:", err);
      }
    }
  }, [token]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("userEmail");
    onLogout();
  };

  const renderCatalog = () => (
    <div className="catalog-section">
      <div className="catalog-header">
        <h2>Catálogo de Roupas</h2>
        <div className="catalog-filters">
          <select className="filter-select">
            <option value="">Todas as categorias</option>
            <option value="vestidos">Vestidos</option>
            <option value="blazers">Blazers</option>
            <option value="saias">Saias</option>
            <option value="calcas">Calças</option>
          </select>
          <select className="filter-select">
            <option value="">Todos os tamanhos</option>
            <option value="P">P</option>
            <option value="M">M</option>
            <option value="G">G</option>
          </select>
        </div>
      </div>

      <div className="products-grid">
        {products.map((product) => (
          <div key={product.id} className="product-card">
            <div className="product-image">
              <img src={product.image_url} alt={product.name} />
              <div className="product-overlay">
                <button className="rent-btn">Alugar</button>
              </div>
            </div>
            <div className="product-info">
              <h3 className="product-name">{product.name}</h3>
              <p className="product-brand">{product.brand}</p>
              <div className="product-details">
                <span className="product-size">Tam: {product.size}</span>
                <span className="product-color">{product.color}</span>
              </div>
              <div className="product-pricing">
                <div className="price-option">
                  <span className="price-label">Diária</span>
                  <span className="price-value">R$ {product.daily_price}</span>
                </div>
                <div className="price-option">
                  <span className="price-label">Semanal</span>
                  <span className="price-value">R$ {product.weekly_price}</span>
                </div>
              </div>
              <div className="product-condition">
                <span className="condition-badge">{product.condition}</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );

  const renderMyRentals = () => (
    <div className="rentals-section">
      <h2>Meus Aluguéis</h2>
      <div className="empty-state">
        <div className="empty-icon">📦</div>
        <h3>Nenhum aluguel ainda</h3>
        <p>Explore nosso catálogo e alugue sua primeira peça!</p>
        <button className="browse-btn" onClick={() => setActiveTab("catalog")}>
          Explorar Catálogo
        </button>
      </div>
    </div>
  );

  const renderProfile = () => (
    <div className="profile-section">
      <h2>Meu Perfil</h2>
      <div className="profile-card">
        <div className="profile-avatar">
          <div className="avatar-circle">👤</div>
        </div>
        <div className="profile-info">
          <h3>Informações Pessoais</h3>
          <div className="info-item">
            <label>Email:</label>
            <span>{userEmail}</span>
          </div>
          <div className="info-item">
            <label>Status:</label>
            <span className="status-badge">Ativo</span>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div className="rental-dashboard">
      <header className="dashboard-header">
        <div className="header-content">
          <div className="brand-section">
            <h1 className="brand-title">DoutorRent</h1>
            <p className="brand-subtitle">Seu guarda-roupa virtual</p>
          </div>
          <div className="user-section">
            <div className="user-info">
              <div className="user-avatar">👤</div>
              <span className="user-name">Olá, {userEmail.split("@")[0]}!</span>
            </div>
            <button className="logout-btn" onClick={handleLogout}>
              Sair
            </button>
          </div>
        </div>
      </header>

      <nav className="dashboard-nav">
        <button
          className={`nav-btn ${activeTab === "catalog" ? "active" : ""}`}
          onClick={() => setActiveTab("catalog")}
        >
          <span className="nav-icon">👗</span>
          Catálogo
        </button>
        <button
          className={`nav-btn ${activeTab === "rentals" ? "active" : ""}`}
          onClick={() => setActiveTab("rentals")}
        >
          <span className="nav-icon">📦</span>
          Meus Aluguéis
        </button>
        <button
          className={`nav-btn ${activeTab === "profile" ? "active" : ""}`}
          onClick={() => setActiveTab("profile")}
        >
          <span className="nav-icon">👤</span>
          Perfil
        </button>
      </nav>

      <main className="dashboard-main">
        {activeTab === "catalog" && renderCatalog()}
        {activeTab === "rentals" && renderMyRentals()}
        {activeTab === "profile" && renderProfile()}
      </main>
    </div>
  );
}
