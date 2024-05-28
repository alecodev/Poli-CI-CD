import { type TodoList } from '../types';

const API_URL = import.meta.env.VITE_API_URL;

export const fetchTodos = async (): Promise<TodoList> => {
  const res = await fetch(API_URL);
  if (!res.ok) {
    console.error('Error fetching todos');
    return [];
  }

  const todos = (await res.json()) as TodoList;
  return todos;
};

export const updateTodos = async ({
  todos,
}: {
  todos: TodoList;
}): Promise<boolean> => {
  const res = await fetch(API_URL, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(todos),
  });

  return res.ok;
};
