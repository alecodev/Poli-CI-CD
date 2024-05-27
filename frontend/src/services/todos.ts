import { type TodoList } from '../types';

const API_URL = import.meta.env.VITE_API_URL;
const API_BIN_KEY = import.meta.env.VITE_API_BIN_KEY;
console.info({ API_URL, API_BIN_KEY });

interface Todo {
  id: string;
  title: string;
  completed: boolean;
  order: number;
}

export const fetchTodos = async (): Promise<Todo[]> => {
  const res = await fetch(API_URL);
  if (!res.ok) {
    console.error('Error fetching todos');
    return [];
  }

  const { record: todos } = (await res.json()) as { record: Todo[] };
  return todos;
};

export const updateTodos = async ({
  todos,
}: {
  todos: TodoList;
}): Promise<boolean> => {
  const res = await fetch(API_URL, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'X-Master-Key': API_BIN_KEY,
    },
    body: JSON.stringify(todos),
  });

  return res.ok;
};
