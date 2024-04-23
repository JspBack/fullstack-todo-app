"use client";
import { useFormik } from "formik";
import validateTodoSchema from "@/types/schema/validateTodo";
import { createTodo, fetchTodos } from "@/lib/reducers/todo";
import { useAppDispatch } from "@/lib/hooks";

export const CreateTodo = ({ setOpen }: CreateTodoProps) => {
  const dispatch = useAppDispatch();
  const formik = useFormik({
    initialValues: {
      todo: "",
    },
    validationSchema: validateTodoSchema,
    onSubmit: async (values) => {
      await dispatch(createTodo({ item: values.todo }));
      dispatch(fetchTodos());
      setOpen(false);
    },
  });
  return (
    <div className="fixed w-screen h-screen top-0 left-0 flex justify-center items-center bg-black bg-opacity-50 ">
      <form
        onSubmit={formik.handleSubmit}
        className="w-60 h-1/4 bg-white flex flex-col justify-around items-center rounded-lg"
      >
        <label htmlFor="todo" className="font-bold text-2xl">
          Create Todo
        </label>
        <input
          id="todo"
          name="todo"
          type="text"
          placeholder="Enter your todo"
          className="w-3/4 h-8 border-2 border-gray-400 rounded-lg px-2"
          onChange={formik.handleChange}
          value={formik.values.todo}
        />
        {formik.errors.todo ? (
          <p className="text-red-500 text-sm">{formik.errors.todo}</p>
        ) : null}
        <div className="w-1/2 flex justify-around">
          <button
            className="w-auto h-8 bg-red-500 text-white font-bold rounded-lg px-2 mr-1"
            onClick={() => setOpen(false)}
          >
            Cancel
          </button>
          <button
            type="submit"
            className="w-auto h-8 bg-green-500 text-white font-bold rounded-lg px-2 ml-1"
          >
            Create
          </button>
        </div>
      </form>
    </div>
  );
};
