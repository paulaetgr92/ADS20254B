import { apiFetch } from "../api";

// Criar produto
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

// Listar todos os produtos
export async function listarProdutos() {
  try {
    const response = await apiFetch("produtos", {
      method: "GET",
    });
    console.log("Produtos listados:", response);
    return response;
  } catch (error) {
    console.error("Erro ao listar produtos:", error.message);
    throw error;
  }
}

// Buscar produto por ID
export async function buscarProdutoPorId(id) {
  try {
    const response = await apiFetch(`produtos/${id}`, {
      method: "GET",
    });
    console.log("Produto encontrado:", response);
    return response;
  } catch (error) {
    console.error("Erro ao buscar produto:", error.message);
    throw error;
  }
}

// Atualizar produto por ID
export async function atualizarProduto(id, dados) {
  try {
    const response = await apiFetch(`produtos/${id}`, {
      method: "PUT",
      body: JSON.stringify(dados),
    });
    console.log("Produto atualizado:", response);
    return response;
  } catch (error) {
    console.error("Erro ao atualizar produto:", error.message);
    throw error;
  }
}

// Inativar produto
export async function inativarProduto(id) {
  try {
    const response = await apiFetch(`produtos/${id}/inativar`, {
      method: "PUT",
    });
    console.log("Produto inativado:", response);
    return response;
  } catch (error) {
    console.error("Erro ao inativar produto:", error.message);
    throw error;
  }
}
