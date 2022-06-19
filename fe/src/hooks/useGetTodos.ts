import axios from "axios";
import { useEffect, useState } from "react";

type Todo = {
  id: number;
  userid: any;
  todo: string;
  createdat: Date;
  updatedat: Date;
  isDone: boolean;
};

export const useGetTodos = () => {
  const URL = "http://127.0.0.1:8080/gettodoList";
  const [todos, setTodos] = useState<Todo[]>([]);
  const [doneTodos, setDoneTodos] = useState<Todo[]>([]);

  useEffect(() => {
    axios
      .get(URL, { params: { isdone: 0 } })
      .then((res) => {
        // console.log(res.data);
        setTodos(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  useEffect(() => {
    axios
      .get(URL, { params: { isdone: 1 } })
      .then((res) => {
        // console.log(res.data);
        setDoneTodos(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return { todos, doneTodos };
};
