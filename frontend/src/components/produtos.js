import { apiFetch } from "./api";

export async function criarProduto(dados) {
  try {
    const response = await apiFetch("produtos", {
      method: "POST",
      body: JSON.stringify(dados),
    });
    console.log("Produto criado:", response);
    return response;
  } catch (error) {
    console.error("Erro ao criar produto:", error.message);
    throw error;
  }
}
