import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState: TodoState = {
  todos: [],
  loading: false,
  error: null,
};

const todoSlice = createSlice({
  name: "todo",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    // fetchTodos
    builder.addCase(fetchTodos.pending, (state) => {
      state.loading = true;
      state.todos = [];
    });
    builder.addCase(fetchTodos.fulfilled, (state, action) => {
      state.todos = action.payload;
      state.loading = false;
    });
    builder.addCase(fetchTodos.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
    // toggleTodo
    builder.addCase(toggleTodo.fulfilled, (state, action) => {
      const index = state.todos.findIndex(
        (todo) => todo.id === action.payload.id
      );
      state.todos[index] = action.payload;
    });
    builder.addCase(toggleTodo.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
    builder.addCase(toggleTodo.pending, (state) => {
      state.loading = true;
    });
    // createTodo
    builder.addCase(createTodo.fulfilled, (state) => {
      state.loading = false;
    });
    builder.addCase(createTodo.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
    builder.addCase(createTodo.pending, (state) => {
      state.loading = true;
    });
    // deleteTodo
    builder.addCase(deleteTodo.fulfilled, (state, action) => {
      state.todos = state.todos.filter((todo) => todo.id !== action.payload);
    });
    builder.addCase(deleteTodo.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
    builder.addCase(deleteTodo.pending, (state) => {
      state.loading = true;
    });
  },
});

export const fetchTodos = createAsyncThunk("todo/fetchTodos", async () => {
  const { data } = await axios.get("http://localhost:9090/todos");
  console.log("data", data);
  return data;
});

export const toggleTodo = createAsyncThunk(
  "todo/toggleTodo",
  async (id: number) => {
    const { data } = await axios.patch(`http://localhost:9090/todos/${id}`);
    return data;
  }
);

export const createTodo = createAsyncThunk(
  "todo/createTodo",
  async (todo: createTodo) => {
    const item = todo.item;
    await axios.post("http://localhost:9090/todos", { item });
    console.log(todo);
  }
);

export const deleteTodo = createAsyncThunk(
  "todo/deleteTodo",
  async (id: number) => {
    await axios.delete(`http://localhost:9090/todos/${id}`);
    return id;
  }
);

export default todoSlice.reducer;
