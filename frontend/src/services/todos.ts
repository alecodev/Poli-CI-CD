import { type TodoList } from '../types';

// Obtenemos la URL de la API desde las variables de entorno configuradas en Vite. 
const API_URL = import.meta.env.VITE_API_URL;

// Devuelve una Promesa que resuelve en un array de tareas tipado como TodoList.
export const fetchTodos = async (): Promise<TodoList> => {
  // Hacemos una petición HTTP GET a la URL de la API usando fetch.
  const res = await fetch(API_URL);
  // Verificamos si la respuesta es correcta.
  if (!res.ok) {
    console.error('Error fetching todos');
    //Error a obtener la respuesta .
    return [];
  }

  const todos = (await res.json()) as TodoList;
  //obtener la lista de tareas desde el servidor.
  return todos;
};

// Definimos la función updateTodos para actualizar la lista de tareas en el servidor.
export const updateTodos = async ({
  todos,
}: {
  todos: TodoList;
}): Promise<boolean> => {
  // Hacemos una petición HTTP PUT a la URL de la API.
  const res = await fetch(API_URL, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(todos),
  });

  return res.ok;
};
