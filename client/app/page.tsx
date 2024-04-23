"use client";
import { useEffect, useState } from "react";
import { CreateTodo, TodoCard } from "@/components";
import { fetchTodos } from "@/lib/reducers/todo";
import { useAppSelector, useAppDispatch } from "@/lib/hooks";

export default function Home() {
  const [open, setOpen] = useState(false);
  const dispatch = useAppDispatch();
  const todos = useAppSelector((state) => state.todo.todos);

  const handleOpen = () => {
    setOpen(true);
  };

  useEffect(() => {
    dispatch(fetchTodos());
  }, [dispatch]);
  return (
    <div className="w-screen min-h-screen flex flex-col justify-around items-center bg-gradient-to-t from-indigo-500 via-purple-500 to-pink-500 overflow-hidden">
      <div className="w-1/2 flex items-center justify-around text-nowrap">
        <h1 className="font-bold text-4xl">TODOS😄</h1>
        <button
          className="w-auto h-8 bg-white text-black font-bold rounded-lg px-2"
          onClick={handleOpen}
        >
          Add Todo
        </button>
      </div>
      {open && <CreateTodo setOpen={setOpen} />}
      <div className="flex flex-wrap w-full justify-center items-center">
        {todos.map((todo) => (
          <TodoCard key={todo.id} data={todo} />
        ))}
        {todos.length === 0 && (
          <p className="text-2xl font-bold text-white">No todos found</p>
        )}
      </div>
    </div>
  );
}
