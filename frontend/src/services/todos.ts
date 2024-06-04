import { type TodoList } from '../types';

interface controlledError {
  error: string;
}

export class FetchError extends Error {
  data: controlledError;
  code: number;

  constructor(data: controlledError, code: number) {
    super();
    this.data = data;
    this.code = code;
  }
}

// Obtenemos la URL de la API desde las variables de entorno configuradas en Vite
const API_URL = import.meta.env.VITE_API_URL;

// Devuelve una Promesa que resuelve en un array de tareas tipado como TodoList
export const fetchTodos = async ({
  uuid,
}: {
  uuid: string;
}): Promise<TodoList> => {
  // Hacemos una petición HTTP GET a la URL de la API usando fetch
  const res = await fetch(API_URL + '/' + uuid);
  // Verificamos si la respuesta es correcta
  if (!res.ok) {
    console.error('Error fetching todos');
    // Error a obtener la respuesta
    const message = (await res.json()) as controlledError;
    throw new FetchError(message, res.status);
  }

  const todos = (await res.json()) as TodoList;
  // obtener la lista de tareas desde el servidor
  return todos;
};

// Definimos la función updateTodos para actualizar la lista de tareas en el servidor
export const updateTodos = async ({
  uuid,
  todos,
}: {
  uuid: string;
  todos: TodoList;
}): Promise<boolean> => {
  // Hacemos una petición HTTP PUT a la URL de la API
  const res = await fetch(API_URL + '/' + uuid, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(todos),
  });

  if (!res.ok) {
    console.error('Error updating todos');
    const message = (await res.json()) as controlledError;
    throw new FetchError(message, res.status);
  }

  return res.ok;
};
