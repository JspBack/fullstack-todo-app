import { MdDone, MdDelete } from "react-icons/md";
import { IoMdClose } from "react-icons/io";
import { useAppDispatch } from "@/lib/hooks";
import { deleteTodo, toggleTodo } from "@/lib/reducers/todo";

export const TodoCard = ({ data }: { data: Todo }) => {
  const dispatch = useAppDispatch();

  const handleToggle = async () => {
    dispatch(toggleTodo(data.id));
  };

  const handleDelete = () => {
    dispatch(deleteTodo(data.id));
  };

  return (
    <div
      className="aspect-square w-60 h-60 bg-gray-400 rounded-lg shadow-lg p-6 
    justify-center items-center flex flex-col m-2 "
    >
      <div className="flex justify-center flex-col items-center h-full w-full font-medium">
        <button
          className="w-8 h-8 aspect-square bg-gray-500 rounded-full
         flex justify-center items-center self-end hover:bg-red-400 transition-colors"
          onClick={handleDelete}
        >
          <MdDelete />
        </button>
        <p className="self-start">Todo name: </p>
        <p
          className={`${
            data.completed ? "line-through self-start" : "self-start"
          }`}
        >
          {data.item}
        </p>
        <p className="pt-6">
          Todo Status: {data.completed ? "Done" : "Not Done"}
        </p>
        <button
          className={` ${
            data.completed ? "bg-red-500 " : "bg-green-500"
          } text-white p-2 rounded-lg mt-4`}
          onClick={handleToggle}
        >
          {data.completed ? <IoMdClose /> : <MdDone />}
        </button>
      </div>
    </div>
  );
};
