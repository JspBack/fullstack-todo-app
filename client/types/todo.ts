interface Todo {
  id: number;
  item: string;
  completed: boolean;
}

interface TodoState {
  todos: Todo[];
  loading: boolean;
  error: string | null;
}

interface CreateTodoProps {
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

interface createTodo {
  item: string;
}
