export async function apiFetch(endpoint, options = {}) {
  const token = localStorage.getItem("token");
  const headers = {
    "Content-Type": "application/json",
    ...(token && { Authorization: `Bearer ${token}` }),
    ...options.headers,
  };

  let response, data;

  try {
    response = await fetch(`http://localhost:8080/api/v1/${endpoint}`, {
      ...options,
      headers,
    });

    const text = await response.text();
    try {
      data = text ? JSON.parse(text) : {};
    } catch {
      data = { message: text };
    }
  } catch (error) {
    throw new Error("Erro de conexão com o servidor");
  }

  if (!response.ok) {
    console.error("Erro do backend:", data);
    const errorMessage = data.error || data.message || "Erro na API";
    throw new Error(errorMessage);
  }

  return data;
}
